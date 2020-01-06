// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var testCodeTemplate = template.Must(
	template.New("").
		Delims("«", "»").
		Funcs(template.FuncMap{
			"quote":   quote,
			"mapKeys": mapKeys,
		}).
		Parse(`
// Code generated by generatetestcode.go; DO NOT EDIT.

package «.TestName»

import (
	"testing"

	"github.com/heetch/avro/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: «quote .InSchema»,
	GoType: new(«.GoType»),
	Subtests: []testutil.RoundTripSubtest{
		«- range $k := mapKeys .Subtests»
		«- $t := index $.Subtests $k»{
			TestName: «printf "%q" $t.TestName»,
			InDataJSON: «quote $t.InData»,
			OutDataJSON: «quote $t.OutData»,
		}, «end»},
}

«if .GoTypeBody -»
type «.GoType» «.GoTypeBody»
«end»

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
`[1:]))

const generateDir = "internal/generated_tests"

type testData struct {
	TestName   string                 `json:"testName"`
	InSchema   json.RawMessage        `json:"inSchema"`
	OutSchema  json.RawMessage        `json:"outSchema"`
	GoType     string                 `json:"goType"`
	GoTypeBody string                 `json:"goTypeBody"`
	Subtests   map[string]subtestData `json:"subtests"`
}

type subtestData struct {
	TestName string          `json:"testName"`
	InData   json.RawMessage `json:"inData"`
	OutData  json.RawMessage `json:"outData"`
}

func mapKeys(m map[string]subtestData) (ks []string) {
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

type subtest struct {
	Name    string
	InData  json.RawMessage `json:"inData"`
	OutData json.RawMessage `json:"outData"`
}

func main() {
	var exported struct {
		Tests map[string]testData `json:"tests"`
	}
	var buf bytes.Buffer
	cmd := exec.Command("cue", "export")
	cmd.Dir = "./testdata"
	cmd.Stderr = os.Stderr
	cmd.Stdout = &buf
	err := cmd.Run()
	check("run cue export", err)
	err = json.Unmarshal(buf.Bytes(), &exported)
	if err != nil {
		log.Printf("bad export data: %q", buf.Bytes())
	}
	check("unmarshal test data", err)
	os.RemoveAll(generateDir)
	for _, test := range exported.Tests {
		dir := filepath.Join(generateDir, test.TestName)
		err := os.MkdirAll(dir, 0777)
		check("mkdir", err)
		if test.OutSchema != nil {
			file := filepath.Join(dir, "schema.avsc")
			err = ioutil.WriteFile(file, []byte(test.OutSchema), 0666)
			check("create schema file", err)

			cmd := exec.Command("avro-generate-go", "-p", test.TestName, "schema.avsc")
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Dir = dir
			err = cmd.Run()
			check("avro-generate-go "+file, err)
		}
		var buf bytes.Buffer
		err = testCodeTemplate.Execute(&buf, test)
		check("generate test code", err)
		goSource, err := format.Source(buf.Bytes())
		check("gofmt test code", err)
		err = ioutil.WriteFile(filepath.Join(dir, "roundtrip_test.go"), goSource, 0666)
		check("write test go file", err)
	}
}

func check(what string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "generatetestcode: %s: %v\n", what, err)
		os.Exit(1)
	}
}

func quote(b []byte) string {
	s := string(b)
	if !strings.Contains(s, "`") {
		return "`" + s + "`"
	}
	return strconv.Quote(s)
}
