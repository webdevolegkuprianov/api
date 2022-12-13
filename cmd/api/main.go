package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/pkg/controller/entity_controller"
	"test/pkg/domains/entity"
	"test/pkg/router"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var ctx context.Context
var ctxCancel context.CancelFunc
var r *mux.Router

func init() {
	ctx, ctxCancel = context.WithCancel(context.Background())
	{
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		multi := zerolog.MultiLevelWriter(consoleWriter)
		ll := zerolog.New(multi).
			With().
			Timestamp().
			Str("Project", "api").
			Logger()
		zerolog.DefaultContextLogger = &ll
	}
	router := router.NewRouter(
		&entity_controller.EntityController{
			Domain: entity.NewDomain(ctx),
		})
	router.RouterEntityInit()
	r = router.Router

}

func main() {
	l := zerolog.Ctx(ctx)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)
	signal.Notify(sigCh, syscall.SIGTERM)
	defer func() {
		l.Info().Msgf("останавливаю HTTP сервер на порту: %d", 8888)
		ctxCancel()
		signal.Stop(sigCh)
		close(sigCh)
		shutdown()
	}()
	go func() {
		defer ctxCancel()
		l.Info().Msgf("запускаю HTTP сервер на порту: %d", 8888)
		err := http.ListenAndServe(fmt.Sprintf(":%d", 8888), r)
		if err != nil {
			l.Err(err).Msg("не могу запустить HTTP сервер")
		}
	}()

	for {
		select {
		case <-sigCh:
			return
		case <-ctx.Done():
			return
		}
	}
}

func shutdown() {}
