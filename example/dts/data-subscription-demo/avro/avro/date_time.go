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

type DateTime struct {
	Year *UnionNullInt `json:"year"`

	Month *UnionNullInt `json:"month"`

	Day *UnionNullInt `json:"day"`

	Hour *UnionNullInt `json:"hour"`

	Minute *UnionNullInt `json:"minute"`

	Second *UnionNullInt `json:"second"`

	Millis *UnionNullInt `json:"millis"`
}

const DateTimeAvroCRC64Fingerprint = "\x05\x8b\x81\xd7fe\xd6%"

func NewDateTime() DateTime {
	r := DateTime{}
	r.Year = nil
	r.Month = nil
	r.Day = nil
	r.Hour = nil
	r.Minute = nil
	r.Second = nil
	r.Millis = nil
	return r
}

func DeserializeDateTime(r io.Reader) (DateTime, error) {
	t := NewDateTime()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDateTimeFromSchema(r io.Reader, schema string) (DateTime, error) {
	t := NewDateTime()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDateTime(r DateTime, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.Year, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Month, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Day, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Hour, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Minute, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Second, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Millis, w)
	if err != nil {
		return err
	}
	return err
}

func (r DateTime) Serialize(w io.Writer) error {
	return writeDateTime(r, w)
}

func (r DateTime) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"year\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"month\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"day\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"hour\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"minute\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"second\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"millis\",\"type\":[\"null\",\"int\"]}],\"name\":\"com.bytedance.dts.subscribe.avro.DateTime\",\"type\":\"record\"}"
}

func (r DateTime) SchemaName() string {
	return "com.bytedance.dts.subscribe.avro.DateTime"
}

func (_ DateTime) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DateTime) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DateTime) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DateTime) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DateTime) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DateTime) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DateTime) SetString(v string)   { panic("Unsupported operation") }
func (_ DateTime) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DateTime) Get(i int) types.Field {
	switch i {
	case 0:
		r.Year = NewUnionNullInt()

		return r.Year
	case 1:
		r.Month = NewUnionNullInt()

		return r.Month
	case 2:
		r.Day = NewUnionNullInt()

		return r.Day
	case 3:
		r.Hour = NewUnionNullInt()

		return r.Hour
	case 4:
		r.Minute = NewUnionNullInt()

		return r.Minute
	case 5:
		r.Second = NewUnionNullInt()

		return r.Second
	case 6:
		r.Millis = NewUnionNullInt()

		return r.Millis
	}
	panic("Unknown field index")
}

func (r *DateTime) SetDefault(i int) {
	switch i {
	case 0:
		r.Year = nil
		return
	case 1:
		r.Month = nil
		return
	case 2:
		r.Day = nil
		return
	case 3:
		r.Hour = nil
		return
	case 4:
		r.Minute = nil
		return
	case 5:
		r.Second = nil
		return
	case 6:
		r.Millis = nil
		return
	}
	panic("Unknown field index")
}

func (r *DateTime) NullField(i int) {
	switch i {
	case 0:
		r.Year = nil
		return
	case 1:
		r.Month = nil
		return
	case 2:
		r.Day = nil
		return
	case 3:
		r.Hour = nil
		return
	case 4:
		r.Minute = nil
		return
	case 5:
		r.Second = nil
		return
	case 6:
		r.Millis = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DateTime) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DateTime) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DateTime) HintSize(int)                     { panic("Unsupported operation") }
func (_ DateTime) Finalize()                        {}

func (_ DateTime) AvroCRC64Fingerprint() []byte {
	return []byte(DateTimeAvroCRC64Fingerprint)
}

func (r DateTime) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["year"], err = json.Marshal(r.Year)
	if err != nil {
		return nil, err
	}
	output["month"], err = json.Marshal(r.Month)
	if err != nil {
		return nil, err
	}
	output["day"], err = json.Marshal(r.Day)
	if err != nil {
		return nil, err
	}
	output["hour"], err = json.Marshal(r.Hour)
	if err != nil {
		return nil, err
	}
	output["minute"], err = json.Marshal(r.Minute)
	if err != nil {
		return nil, err
	}
	output["second"], err = json.Marshal(r.Second)
	if err != nil {
		return nil, err
	}
	output["millis"], err = json.Marshal(r.Millis)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DateTime) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["year"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Year); err != nil {
			return err
		}
	} else {
		r.Year = NewUnionNullInt()

		r.Year = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["month"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Month); err != nil {
			return err
		}
	} else {
		r.Month = NewUnionNullInt()

		r.Month = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["day"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Day); err != nil {
			return err
		}
	} else {
		r.Day = NewUnionNullInt()

		r.Day = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["hour"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Hour); err != nil {
			return err
		}
	} else {
		r.Hour = NewUnionNullInt()

		r.Hour = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["minute"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Minute); err != nil {
			return err
		}
	} else {
		r.Minute = NewUnionNullInt()

		r.Minute = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["second"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Second); err != nil {
			return err
		}
	} else {
		r.Second = NewUnionNullInt()

		r.Second = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["millis"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Millis); err != nil {
			return err
		}
	} else {
		r.Millis = NewUnionNullInt()

		r.Millis = nil
	}
	return nil
}
