// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     record.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayField(r []Field, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeField(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayFieldWrapper struct {
	Target *[]Field
}

func (_ ArrayFieldWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayFieldWrapper) Finalize()                        {}
func (_ ArrayFieldWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayFieldWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Field, 0, s)
	}
}
func (r ArrayFieldWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayFieldWrapper) AppendArray() types.Field {
	var v Field
	v = NewField()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
