// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     record.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Record struct {
	// version infomation
	Version int32 `json:"version"`
	// unique id of this record in the whole stream
	Id int64 `json:"id"`
	// record log timestamp
	SourceTimestamp int64 `json:"sourceTimestamp"`
	// record source location information
	SourcePosition string `json:"sourcePosition"`
	// safe record source location information, use to recovery.
	SafeSourcePosition string `json:"safeSourcePosition"`
	// record transation id
	SourceTxid string `json:"sourceTxid"`
	// source dataource
	Source Source `json:"source"`

	Operation Operation `json:"operation"`

	ObjectName *UnionNullString `json:"objectName"`
	// time when this record is processed along the stream dataflow
	ProcessTimestamps *UnionNullArrayLong `json:"processTimestamps"`
	// tags to identify properties of this record
	Tags map[string]string `json:"tags"`

	Fields *UnionNullStringArrayField `json:"fields"`

	BeforeImages *UnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject `json:"beforeImages"`

	AfterImages *UnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject `json:"afterImages"`
	// the timestamp in unit of millisecond that record is born in source
	BornTimestamp int64 `json:"bornTimestamp"`
}

const RecordAvroCRC64Fingerprint = "\x0f9q.9\x9d\xea\xfc"

func NewRecord() Record {
	r := Record{}
	r.SafeSourcePosition = ""
	r.SourceTxid = ""
	r.Source = NewSource()

	r.ObjectName = nil
	r.ProcessTimestamps = nil
	r.Tags = make(map[string]string)

	r.Fields = nil
	r.BeforeImages = nil
	r.AfterImages = nil
	r.BornTimestamp = 0
	return r
}

func DeserializeRecord(r io.Reader) (Record, error) {
	t := NewRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecordFromSchema(r io.Reader, schema string) (Record, error) {
	t := NewRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecord(r Record, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Version, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.SourceTimestamp, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SourcePosition, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SafeSourcePosition, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SourceTxid, w)
	if err != nil {
		return err
	}
	err = writeSource(r.Source, w)
	if err != nil {
		return err
	}
	err = writeOperation(r.Operation, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ObjectName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayLong(r.ProcessTimestamps, w)
	if err != nil {
		return err
	}
	err = writeMapString(r.Tags, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStringArrayField(r.Fields, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject(r.BeforeImages, w)
	if err != nil {
		return err
	}
	err = writeUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject(r.AfterImages, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.BornTimestamp, w)
	if err != nil {
		return err
	}
	return err
}

func (r Record) Serialize(w io.Writer) error {
	return writeRecord(r, w)
}

func (r Record) Schema() string {
	return "{\"fields\":[{\"doc\":\"version infomation\",\"name\":\"version\",\"type\":\"int\"},{\"doc\":\"unique id of this record in the whole stream\",\"name\":\"id\",\"type\":\"long\"},{\"doc\":\"record log timestamp\",\"name\":\"sourceTimestamp\",\"type\":\"long\"},{\"doc\":\"record source location information\",\"name\":\"sourcePosition\",\"type\":\"string\"},{\"default\":\"\",\"doc\":\"safe record source location information, use to recovery.\",\"name\":\"safeSourcePosition\",\"type\":\"string\"},{\"default\":\"\",\"doc\":\"record transation id\",\"name\":\"sourceTxid\",\"type\":\"string\"},{\"doc\":\"source dataource\",\"name\":\"source\",\"type\":{\"fields\":[{\"name\":\"sourceType\",\"type\":{\"name\":\"SourceType\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"symbols\":[\"MySQL\",\"Oracle\",\"SQLServer\",\"PostgreSQL\",\"MongoDB\",\"Redis\",\"DB2\",\"PPAS\",\"DRDS\",\"HBASE\",\"HDFS\",\"FILE\",\"OTHER\"],\"type\":\"enum\"}},{\"doc\":\"source datasource version information\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"Source\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"}},{\"name\":\"operation\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":{\"name\":\"Operation\",\"symbols\":[\"INSERT\",\"UPDATE\",\"DELETE\",\"DDL\",\"BEGIN\",\"COMMIT\",\"ROLLBACK\",\"ABORT\",\"HEARTBEAT\",\"CHECKPOINT\",\"COMMAND\",\"FILL\",\"FINISH\",\"CONTROL\",\"RDB\",\"NOOP\",\"INIT\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"objectName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"time when this record is processed along the stream dataflow\",\"name\":\"processTimestamps\",\"type\":[\"null\",{\"items\":\"long\",\"type\":\"array\"}]},{\"default\":{},\"doc\":\"tags to identify properties of this record\",\"name\":\"tags\",\"type\":{\"type\":\"map\",\"values\":\"string\"}},{\"default\":null,\"name\":\"fields\",\"type\":[\"null\",\"string\",{\"items\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataTypeNumber\",\"type\":\"int\"}],\"name\":\"Field\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"beforeImages\",\"type\":[\"null\",\"string\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Integer\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"charset\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"Character\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"value\",\"type\":\"string\"},{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"scale\",\"type\":\"int\"}],\"name\":\"Decimal\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"value\",\"type\":\"double\"},{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"scale\",\"type\":\"int\"}],\"name\":\"Float\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"timestamp\",\"type\":\"long\"},{\"name\":\"millis\",\"type\":\"int\"}],\"name\":\"Timestamp\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"default\":null,\"name\":\"year\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"month\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"day\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"hour\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"minute\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"second\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"millis\",\"type\":[\"null\",\"int\"]}],\"name\":\"DateTime\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"value\",\"type\":\"com.bytedance.dts.subscribe.avro.DateTime\"},{\"name\":\"timezone\",\"type\":\"string\"}],\"name\":\"TimestampWithTimeZone\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"BinaryGeometry\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextGeometry\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"BinaryObject\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextObject\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"type\":\"record\"},{\"name\":\"EmptyObject\",\"namespace\":\"com.bytedance.dts.subscribe.avro\",\"symbols\":[\"NULL\",\"NONE\"],\"type\":\"enum\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"afterImages\",\"type\":[\"null\",\"string\",{\"items\":[\"null\",\"com.bytedance.dts.subscribe.avro.Integer\",\"com.bytedance.dts.subscribe.avro.Character\",\"com.bytedance.dts.subscribe.avro.Decimal\",\"com.bytedance.dts.subscribe.avro.Float\",\"com.bytedance.dts.subscribe.avro.Timestamp\",\"com.bytedance.dts.subscribe.avro.DateTime\",\"com.bytedance.dts.subscribe.avro.TimestampWithTimeZone\",\"com.bytedance.dts.subscribe.avro.BinaryGeometry\",\"com.bytedance.dts.subscribe.avro.TextGeometry\",\"com.bytedance.dts.subscribe.avro.BinaryObject\",\"com.bytedance.dts.subscribe.avro.TextObject\",\"com.bytedance.dts.subscribe.avro.EmptyObject\"],\"type\":\"array\"}]},{\"default\":0,\"doc\":\"the timestamp in unit of millisecond that record is born in source\",\"name\":\"bornTimestamp\",\"type\":\"long\"}],\"name\":\"com.bytedance.dts.subscribe.avro.Record\",\"type\":\"record\"}"
}

func (r Record) SchemaName() string {
	return "com.bytedance.dts.subscribe.avro.Record"
}

func (_ Record) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Record) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Record) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Record) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Record) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Record) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Record) SetString(v string)   { panic("Unsupported operation") }
func (_ Record) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Record) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Version}

		return w

	case 1:
		w := types.Long{Target: &r.Id}

		return w

	case 2:
		w := types.Long{Target: &r.SourceTimestamp}

		return w

	case 3:
		w := types.String{Target: &r.SourcePosition}

		return w

	case 4:
		w := types.String{Target: &r.SafeSourcePosition}

		return w

	case 5:
		w := types.String{Target: &r.SourceTxid}

		return w

	case 6:
		r.Source = NewSource()

		w := types.Record{Target: &r.Source}

		return w

	case 7:
		w := OperationWrapper{Target: &r.Operation}

		return w

	case 8:
		r.ObjectName = NewUnionNullString()

		return r.ObjectName
	case 9:
		r.ProcessTimestamps = NewUnionNullArrayLong()

		return r.ProcessTimestamps
	case 10:
		r.Tags = make(map[string]string)

		w := MapStringWrapper{Target: &r.Tags}

		return &w

	case 11:
		r.Fields = NewUnionNullStringArrayField()

		return r.Fields
	case 12:
		r.BeforeImages = NewUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

		return r.BeforeImages
	case 13:
		r.AfterImages = NewUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

		return r.AfterImages
	case 14:
		w := types.Long{Target: &r.BornTimestamp}

		return w

	}
	panic("Unknown field index")
}

func (r *Record) SetDefault(i int) {
	switch i {
	case 4:
		r.SafeSourcePosition = ""
		return
	case 5:
		r.SourceTxid = ""
		return
	case 8:
		r.ObjectName = nil
		return
	case 9:
		r.ProcessTimestamps = nil
		return
	case 10:

		return
	case 11:
		r.Fields = nil
		return
	case 12:
		r.BeforeImages = nil
		return
	case 13:
		r.AfterImages = nil
		return
	case 14:
		r.BornTimestamp = 0
		return
	}
	panic("Unknown field index")
}

func (r *Record) NullField(i int) {
	switch i {
	case 8:
		r.ObjectName = nil
		return
	case 9:
		r.ProcessTimestamps = nil
		return
	case 11:
		r.Fields = nil
		return
	case 12:
		r.BeforeImages = nil
		return
	case 13:
		r.AfterImages = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Record) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Record) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Record) HintSize(int)                     { panic("Unsupported operation") }
func (_ Record) Finalize()                        {}

func (_ Record) AvroCRC64Fingerprint() []byte {
	return []byte(RecordAvroCRC64Fingerprint)
}

func (r Record) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["version"], err = json.Marshal(r.Version)
	if err != nil {
		return nil, err
	}
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["sourceTimestamp"], err = json.Marshal(r.SourceTimestamp)
	if err != nil {
		return nil, err
	}
	output["sourcePosition"], err = json.Marshal(r.SourcePosition)
	if err != nil {
		return nil, err
	}
	output["safeSourcePosition"], err = json.Marshal(r.SafeSourcePosition)
	if err != nil {
		return nil, err
	}
	output["sourceTxid"], err = json.Marshal(r.SourceTxid)
	if err != nil {
		return nil, err
	}
	output["source"], err = json.Marshal(r.Source)
	if err != nil {
		return nil, err
	}
	output["operation"], err = json.Marshal(r.Operation)
	if err != nil {
		return nil, err
	}
	output["objectName"], err = json.Marshal(r.ObjectName)
	if err != nil {
		return nil, err
	}
	output["processTimestamps"], err = json.Marshal(r.ProcessTimestamps)
	if err != nil {
		return nil, err
	}
	output["tags"], err = json.Marshal(r.Tags)
	if err != nil {
		return nil, err
	}
	output["fields"], err = json.Marshal(r.Fields)
	if err != nil {
		return nil, err
	}
	output["beforeImages"], err = json.Marshal(r.BeforeImages)
	if err != nil {
		return nil, err
	}
	output["afterImages"], err = json.Marshal(r.AfterImages)
	if err != nil {
		return nil, err
	}
	output["bornTimestamp"], err = json.Marshal(r.BornTimestamp)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Record) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["version"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Version); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for version")
	}
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sourceTimestamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SourceTimestamp); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sourceTimestamp")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sourcePosition"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SourcePosition); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sourcePosition")
	}
	val = func() json.RawMessage {
		if v, ok := fields["safeSourcePosition"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SafeSourcePosition); err != nil {
			return err
		}
	} else {
		r.SafeSourcePosition = ""
	}
	val = func() json.RawMessage {
		if v, ok := fields["sourceTxid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SourceTxid); err != nil {
			return err
		}
	} else {
		r.SourceTxid = ""
	}
	val = func() json.RawMessage {
		if v, ok := fields["source"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Source); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for source")
	}
	val = func() json.RawMessage {
		if v, ok := fields["operation"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Operation); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for operation")
	}
	val = func() json.RawMessage {
		if v, ok := fields["objectName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ObjectName); err != nil {
			return err
		}
	} else {
		r.ObjectName = NewUnionNullString()

		r.ObjectName = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["processTimestamps"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProcessTimestamps); err != nil {
			return err
		}
	} else {
		r.ProcessTimestamps = NewUnionNullArrayLong()

		r.ProcessTimestamps = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tags"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tags); err != nil {
			return err
		}
	} else {
		r.Tags = make(map[string]string)

	}
	val = func() json.RawMessage {
		if v, ok := fields["fields"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fields); err != nil {
			return err
		}
	} else {
		r.Fields = NewUnionNullStringArrayField()

		r.Fields = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["beforeImages"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BeforeImages); err != nil {
			return err
		}
	} else {
		r.BeforeImages = NewUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

		r.BeforeImages = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["afterImages"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AfterImages); err != nil {
			return err
		}
	} else {
		r.AfterImages = NewUnionNullStringArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject()

		r.AfterImages = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["bornTimestamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BornTimestamp); err != nil {
			return err
		}
	} else {
		r.BornTimestamp = 0
	}
	return nil
}
