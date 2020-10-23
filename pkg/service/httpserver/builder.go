// Package httpserver ...
package httpserver

import (
	"context"
	"net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	httpMethodLoad = "POST"
	uriPathLoad    = "/load"
	httpMethodGet  = "GET"
	uriPathGet     = "/get"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, errorProcessor errorProcessor) {

	loadTransport := NewLoadTransport(
		decodeJSONErrorCreator,
	)
	router.Handle(httpMethodLoad, uriPathLoad, NewLoad(loadTransport, svc, errorProcessor))

	getTransport := NewGetTransport(

		encodeJSONErrorCreator,
	)
	router.Handle(httpMethodGet, uriPathGet, NewGet(getTransport, svc, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
