// Code generated by avrogen. DO NOT EDIT.

package unionToScalar

import "github.com/heetch/avro/avrotypegen"

type PrimitiveUnionTestRecord struct {
	UnionField int
}

// AvroRecord implements the avro.AvroRecord interface.
func (PrimitiveUnionTestRecord) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"default":1234,"name":"UnionField","type":"int"}],"name":"PrimitiveUnionTestRecord","type":"record"}`,
		Defaults: []func() interface{}{
			0: func() interface{} {
				return 1234
			},
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?
