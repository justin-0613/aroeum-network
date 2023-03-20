// Code generated by rlpgen. DO NOT EDIT.

//go:build !norlpgen
// +build !norlpgen

package types

import "github.com/justin-0613/go-aroeum/rlp"
import "io"

func (obj *rlpLog) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteBytes(obj.Address[:])
	_tmp1 := w.List()
	for _, _tmp2 := range obj.Topics {
		w.WriteBytes(_tmp2[:])
	}
	w.ListEnd(_tmp1)
	w.WriteBytes(obj.Data)
	w.ListEnd(_tmp0)
	return w.Flush()
}
