// Code generated by generatetestcode.go; DO NOT EDIT.

package simpleEnum

import (
	"testing"

	"github.com/heetch/avro/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "E",
                        "type": {
                            "name": "MyEnum",
                            "type": "enum",
                            "symbols": [
                                "a",
                                "b",
                                "c"
                            ]
                        }
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "E": "b"
                    }`,
		OutDataJSON: `{
                        "E": "b"
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
