// Package httpserver ...
package httpserver

import (
	"context"

	"github.com/valyala/fasthttp"

	v1 "github.com/criro1/aviasales/pkg/api/v1"
)

type service interface {
	Load(ctx context.Context, words []string) (err error)
	Get(ctx context.Context, word *string) (response v1.GetResponse, err error)
}

type load struct {
	transport      LoadTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *load) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		words []string
		err   error
	)
	words, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.Load(ctx, words)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewLoad the server creator
func NewLoad(transport LoadTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := load{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type get struct {
	transport      GetTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *get) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		word     *string
		response v1.GetResponse
		err      error
	)
	word, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err = s.service.Get(ctx, word)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGet the server creator
func NewGet(transport GetTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := get{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
