// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package httpserver

import (
	json "encoding/json"
	_v1 "github.com/criro1/aviasales/pkg/api/v1"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver(in *jlexer.Lexer, out *loadTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver(out *jwriter.Writer, in loadTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v loadTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v loadTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *loadTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *loadTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver(l, v)
}
func easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver1(in *jlexer.Lexer, out *loadRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "words":
			if in.IsNull() {
				in.Skip()
				out.Words = nil
			} else {
				in.Delim('[')
				if out.Words == nil {
					if !in.IsDelim(']') {
						out.Words = make([]string, 0, 4)
					} else {
						out.Words = []string{}
					}
				} else {
					out.Words = (out.Words)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Words = append(out.Words, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver1(out *jwriter.Writer, in loadRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"words\":"
		out.RawString(prefix[1:])
		if in.Words == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Words {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v loadRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v loadRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *loadRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *loadRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver1(l, v)
}
func easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver2(in *jlexer.Lexer, out *getTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver2(out *jwriter.Writer, in getTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver2(l, v)
}
func easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver3(in *jlexer.Lexer, out *getResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgApiV1(in, &out.Data)
		case "error":
			out.Error = bool(in.Bool())
		case "errorText":
			out.ErrorText = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver3(out *jwriter.Writer, in getResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgApiV1(out, in.Data)
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.Bool(bool(in.Error))
	}
	{
		const prefix string = ",\"errorText\":"
		out.RawString(prefix)
		out.String(string(in.ErrorText))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgServiceHttpserver3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgServiceHttpserver3(l, v)
}
func easyjsonC87d08bdDecodeGithubComCriro1AviasalesPkgApiV1(in *jlexer.Lexer, out *_v1.Anagrams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "anagram":
			if in.IsNull() {
				in.Skip()
				out.Anagram = nil
			} else {
				in.Delim('[')
				if out.Anagram == nil {
					if !in.IsDelim(']') {
						out.Anagram = make([]string, 0, 4)
					} else {
						out.Anagram = []string{}
					}
				} else {
					out.Anagram = (out.Anagram)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Anagram = append(out.Anagram, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComCriro1AviasalesPkgApiV1(out *jwriter.Writer, in _v1.Anagrams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"anagram\":"
		out.RawString(prefix[1:])
		if in.Anagram == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Anagram {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
