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

type loadRequest struct {
	Words []string `json:"words"`
}

// LoadTransport transport interface
type LoadTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (words []string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type loadTransport struct {
	decodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *loadTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (words []string, err error) {
	tmp := r.Body()
	if tmp[0] != '[' || tmp[len(tmp)-1] != ']' || bytes.Count(tmp, []byte(`"`))%2 == 1 || bytes.Count(tmp, []byte(`"`))/2-1 != bytes.Count(tmp, []byte(",")) {
		err = t.decodeJSONErrorCreator(fmt.Errorf("Error: format of the body"))
		return
	}

	str := string(tmp[2 : len(tmp)-2])
	words = strings.Split(str, `", "`)

	return
}

// EncodeResponse method for encoding response on server side
func (t *loadTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {

	r.Header.Set("Content-Type", "application/json")

	r.Header.SetStatusCode(v1.HTTPStatusOK)
	return
}

// NewLoadTransport the transport creator for http requests
func NewLoadTransport(

	decodeJSONErrorCreator errorCreator,

) LoadTransport {
	return &loadTransport{

		decodeJSONErrorCreator: decodeJSONErrorCreator,
	}
}

//easyjson:json
type getResponse struct {
	v1.GetResponse
}

// GetTransport transport interface
type GetTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (word *string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.GetResponse) (err error)
}

type getTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (word *string, err error) {

	word = ptr(ctx.QueryArgs().Peek("word"))

	return
}

// EncodeResponse method for encoding response on server side
func (t *getTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, response v1.GetResponse) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse getResponse

	theResponse.GetResponse = response

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(err)
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(v1.HTTPStatusOK)
	return
}

// NewGetTransport the transport creator for http requests
func NewGetTransport(

	encodeJSONErrorCreator errorCreator,

) GetTransport {
	return &getTransport{

		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

func ptr(in []byte) *string {
	if bytes.Equal(in, emptyBytes) {
		return nil
	}
	i := string(in)
	return &i
}
