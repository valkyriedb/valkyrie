package main

import (
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/valkyriedb/valkyrie/adapter/tcp"
	"github.com/valkyriedb/valkyrie/config"
	"github.com/valkyriedb/valkyrie/internal/logger"
	"github.com/valkyriedb/valkyrie/service"
	"github.com/valkyriedb/valkyrie/storage"
)

var db storage.DB

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("can't load config: ", err)
	}

	l := logger.New(os.Stdout, cfg.Env)

	srv := service.New(&db)
	addr := net.JoinHostPort("", cfg.Port)
	tcph, err := tcp.NewHandler(srv, cfg.Password, addr, l)
	if err != nil {
		l.Error("can't listen address", slog.String("addr", addr), logger.Err(err))
		os.Exit(1)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go tcph.ListenAndServe()

	s := <-interrupt
	l.Info("signal interrupt", slog.String("error", s.String()))

	err = tcph.Shutdown()
	if err != nil {
		l.Error("can't shutdown server", logger.Err(err))
		os.Exit(1)
	}
}
