package jsonany

import (
	"github.com/mailru/easyjson/buffer"
	"github.com/mailru/easyjson/jwriter"
	"google.golang.org/protobuf/types/known/anypb"
)

type Any anypb.Any

func (m *Any) MarshalEasyJSON(w *jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}
	comma := false
	w.RawByte('{')
	if m.TypeUrl != "" {
		w.RawString(`"type_url":`)
		w.String(m.TypeUrl)
		comma = true
	}
	if comma {
		w.RawByte(',')
	}
	if len(m.Value) > 0 {
		w.RawString(`"value":`)
		w.Base64Bytes(m.Value)
	}
	w.RawByte('}')
}

func (m *Any) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{Buffer: buffer.Buffer{Buf: make([]byte, 2048)}}
	m.MarshalEasyJSON(&w)
	return w.Buffer.Buf, nil
}
