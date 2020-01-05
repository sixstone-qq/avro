// Code generated by generatetestcode.go; DO NOT EDIT.

package unionNullStringReverseWithString

import (
	"testing"

	"github.com/rogpeppe/avro/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "OptionalString",
                        "type": [
                            "string",
                            "null"
                        ]
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "OptionalString": {
                            "string": "hello"
                        }
                    }`,
		OutDataJSON: `{
                        "OptionalString": {
                            "string": "hello"
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
