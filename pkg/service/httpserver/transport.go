// Package httpserver ...
package httpserver

import (
	"bytes"
	"fmt"
	"strings"

	v1 "github.com/criro1/aviasales/pkg/api/v1"
	"github.com/valyala/fasthttp"
)

var (
	emptyBytes = []byte("")
)

// LoadTransport transport interface
type LoadTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (words []string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type loadTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *loadTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (words []string, err error) {
	tmp := r.Body()
	if tmp[0] != '[' || tmp[len(tmp)-1] != ']' || bytes.Count(tmp, []byte(`"`))%2 == 1 || bytes.Count(tmp, []byte(`"`))/2-1 != bytes.Count(tmp, []byte(",")) {
		err = fmt.Errorf("Error: format of the body")
		return
	}
	str := string(tmp[2 : len(tmp)-2])
	words = strings.Split(str, `", "`)
	return
}

// EncodeResponse method for encoding response on server side
func (t *loadTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {
	r.Header.SetStatusCode(v1.HTTPStatusOK)
	return
}

// NewLoadTransport the transport creator for http requests
func NewLoadTransport() LoadTransport {
	return &loadTransport{}
}

// GetTransport transport interface
type GetTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (word *string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.Anagrams) (err error)
}

type getTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *getTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (word *string, err error) {
	word = ptr(ctx.QueryArgs().Peek("word"))
	return
}

// EncodeResponse method for encoding response on server side
func (t *getTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.Anagrams) (err error) {
	if len(response.Anagram) != 0 {
		body := []byte(`["` + strings.Join(response.Anagram, `", "`) + `"]`)
		r.SetBody(body)
	}
	r.Header.SetStatusCode(v1.HTTPStatusOK)
	return
}

// NewGetTransport the transport creator for http requests
func NewGetTransport() GetTransport {
	return &getTransport{}
}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
