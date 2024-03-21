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

type BinaryObject struct {
	Type string `json:"type"`

	Value Bytes `json:"value"`
}

const BinaryObjectAvroCRC64Fingerprint = "b\xc0\x12\xb4s|}\xe2"

func NewBinaryObject() BinaryObject {
	r := BinaryObject{}
	return r
}

func DeserializeBinaryObject(r io.Reader) (BinaryObject, error) {
	t := NewBinaryObject()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBinaryObjectFromSchema(r io.Reader, schema string) (BinaryObject, error) {
	t := NewBinaryObject()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBinaryObject(r BinaryObject, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r BinaryObject) Serialize(w io.Writer) error {
	return writeBinaryObject(r, w)
}

func (r BinaryObject) Schema() string {
	return "{\"fields\":[{\"name\":\"type\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"com.bytedance.dts.subscribe.avro.BinaryObject\",\"type\":\"record\"}"
}

func (r BinaryObject) SchemaName() string {
	return "com.bytedance.dts.subscribe.avro.BinaryObject"
}

func (_ BinaryObject) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BinaryObject) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BinaryObject) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BinaryObject) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BinaryObject) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BinaryObject) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BinaryObject) SetString(v string)   { panic("Unsupported operation") }
func (_ BinaryObject) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BinaryObject) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Type}

		return w

	case 1:
		w := BytesWrapper{Target: &r.Value}

		return w

	}
	panic("Unknown field index")
}

func (r *BinaryObject) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *BinaryObject) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BinaryObject) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BinaryObject) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BinaryObject) HintSize(int)                     { panic("Unsupported operation") }
func (_ BinaryObject) Finalize()                        {}

func (_ BinaryObject) AvroCRC64Fingerprint() []byte {
	return []byte(BinaryObjectAvroCRC64Fingerprint)
}

func (r BinaryObject) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BinaryObject) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for value")
	}
	return nil
}
