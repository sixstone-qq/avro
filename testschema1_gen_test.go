// Code generated by avrogen. DO NOT EDIT.

package avro_test

import "github.com/heetch/avro/avrotypegen"

type TestRecord struct {
	A int
	B int
}

// AvroRecord implements the avro.AvroRecord interface.
func (TestRecord) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"default":42,"name":"A","type":{"type":"int"}},{"name":"B","type":{"type":"int"}}],"name":"TestRecord","type":"record"}`,
		Required: []bool{
			1: true,
		},
		Defaults: []func() interface{}{
			0: func() interface{} {
				return 42
			},
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?
