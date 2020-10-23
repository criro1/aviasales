package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buaazp/fasthttprouter"
	fasthttpprometheus "github.com/flf2ko/fasthttp-prometheus"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"

	"github.com/criro1/aviasales/pkg/api/v1"
	"github.com/criro1/aviasales/pkg/service"
	"github.com/criro1/aviasales/pkg/service/httperrors"
	"github.com/criro1/aviasales/pkg/service/httpserver"
)

type configuration struct {
	Port               string        `envconfig:"PORT" required:"true" default:"8080"`
	ReadTimeout        time.Duration `envconfig:"READ_TIMEOUT" required:"true" default:"1s"`
	ReadBufferSize     int           `envconfig:"READ_BUFFER_SIZE" required:"true" default:"16384"`
	MaxRequestBodySize int           `envconfig:"MAX_REQUEST_BODY_SIZE" required:"true" default:"10485760"`
}

func main() {
	// logger
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", "for aviasales")

	// configuration
	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}

	// error processor for UI
	errorProcessor := httperrors.NewErrorProcessor(
		v1.ErrMap,
	)

	// service
	svc := service.NewService(
		make(map[string]bool),
	)

	router := fasthttprouter.New()
	httpserver.New(
		router,
		svc,
		httperrors.DecodeJSONErrorCreator,
		httperrors.EncodeJSONErrorCreator,
		errorProcessor,
	)

	p := fasthttpprometheus.NewPrometheus("")
	fasthttpServer := &fasthttp.Server{
		Handler:            p.WrapHandler(router),
		MaxRequestBodySize: cfg.MaxRequestBodySize,
		ReadBufferSize:     cfg.ReadBufferSize,
		ReadTimeout:        cfg.ReadTimeout,
	}

	go func() {
		_ = level.Info(logger).Log("msg", "starting http server", "port", cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + cfg.Port); err != nil {
			_ = level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)
		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "goodbye")
	}(<-c)
}
