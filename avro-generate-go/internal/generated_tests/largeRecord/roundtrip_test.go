// Code generated by generatetestcode.go; DO NOT EDIT.

package largeRecord

import (
	"testing"

	"github.com/rogpeppe/avro/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "sample",
                "type": "record",
                "fields": [
                    {
                        "name": "header",
                        "type": [
                            "null",
                            {
                                "name": "Data0",
                                "type": "record",
                                "fields": [
                                    {
                                        "name": "uuid",
                                        "type": [
                                            "null",
                                            {
                                                "name": "UUID0",
                                                "type": "record",
                                                "fields": [
                                                    {
                                                        "name": "uuid",
                                                        "type": "string",
                                                        "default": ""
                                                    }
                                                ],
                                                "namespace": "headerworks.datatype",
                                                "doc": "A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014"
                                            }
                                        ],
                                        "default": null,
                                        "doc": "Unique identifier for the event used for de-duplication and tracing."
                                    },
                                    {
                                        "name": "hostname",
                                        "type": [
                                            "null",
                                            "string"
                                        ],
                                        "default": null,
                                        "doc": "Fully qualified name of the host that generated the event that generated the data."
                                    },
                                    {
                                        "name": "trace",
                                        "type": [
                                            "null",
                                            {
                                                "name": "Trace0",
                                                "type": "record",
                                                "fields": [
                                                    {
                                                        "name": "traceId",
                                                        "type": [
                                                            "null",
                                                            "headerworks.datatype.UUID0"
                                                        ],
                                                        "default": null,
                                                        "doc": "Trace Identifier"
                                                    }
                                                ],
                                                "doc": "Trace0"
                                            }
                                        ],
                                        "default": null,
                                        "doc": "Trace information not redundant with this object"
                                    }
                                ],
                                "namespace": "headerworks",
                                "doc": "Common information related to the event which must be included in any clean event"
                            }
                        ],
                        "default": null,
                        "doc": "Core data information required for any event"
                    },
                    {
                        "name": "body",
                        "type": [
                            "null",
                            {
                                "name": "Data1",
                                "type": "record",
                                "fields": [
                                    {
                                        "name": "uuid",
                                        "type": [
                                            "null",
                                            {
                                                "name": "UUID1",
                                                "type": "record",
                                                "fields": [
                                                    {
                                                        "name": "uuid",
                                                        "type": "string",
                                                        "default": ""
                                                    }
                                                ],
                                                "namespace": "bodyworks.datatype",
                                                "doc": "A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014"
                                            }
                                        ],
                                        "default": null,
                                        "doc": "Unique identifier for the event used for de-duplication and tracing."
                                    },
                                    {
                                        "name": "hostname",
                                        "type": [
                                            "null",
                                            "string"
                                        ],
                                        "default": null,
                                        "doc": "Fully qualified name of the host that generated the event that generated the data."
                                    },
                                    {
                                        "name": "trace",
                                        "type": [
                                            "null",
                                            {
                                                "name": "Trace1",
                                                "type": "record",
                                                "fields": [
                                                    {
                                                        "name": "traceId",
                                                        "type": [
                                                            "null",
                                                            "headerworks.datatype.UUID0"
                                                        ],
                                                        "default": null,
                                                        "doc": "Trace Identifier"
                                                    }
                                                ],
                                                "doc": "Trace1"
                                            }
                                        ],
                                        "default": null,
                                        "doc": "Trace information not redundant with this object"
                                    }
                                ],
                                "namespace": "bodyworks",
                                "doc": "Common information related to the event which must be included in any clean event"
                            }
                        ],
                        "default": null,
                        "doc": "Core data information required for any event"
                    }
                ],
                "namespace": "com.avro.test",
                "doc": "GoGen test"
            }`,
	GoType: new(Sample),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "header": {
                            "headerworks.Data0": {
                                "hostname": {
                                    "string": "myhost.com"
                                }
                            }
                        }
                    }`,
		OutDataJSON: `{
                        "body": null,
                        "header": {
                            "headerworks.Data0": {
                                "hostname": {
                                    "string": "myhost.com"
                                },
                                "trace": null,
                                "uuid": null
                            }
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
