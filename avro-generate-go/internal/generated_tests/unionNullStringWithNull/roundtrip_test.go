// Code generated by generatetestcode.go; DO NOT EDIT.

package unionNullStringWithNull

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
                            "null",
                            "string"
                        ]
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "OptionalString": null
                    }`,
		OutDataJSON: `{
                        "OptionalString": null
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
