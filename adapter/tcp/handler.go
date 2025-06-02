package tcp

import (
	"log/slog"
	"net"
	"sync"

	"github.com/valkyriedb/valkyrie/internal/logger"
)

type Handler struct {
	lstn     net.Listener
	shutdown struct {
		ch chan struct{}
		wg sync.WaitGroup
	}
	log *slog.Logger
}

func NewHandler(address string, log *slog.Logger) (*Handler, error) {
	lstn, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Handler{
		lstn: lstn,
		shutdown: struct {
			ch chan struct{}
			wg sync.WaitGroup
		}{
			ch: make(chan struct{}),
		},
		log: log,
	}, nil
}

func (h *Handler) ListenAndServe() {
	h.log.Info("listening on address", slog.String("addr", h.lstn.Addr().String()))

	for {
		conn, err := h.lstn.Accept()
		if err != nil {
			return
		}
		h.shutdown.wg.Add(1)

		go h.handleConn(conn)
	}
}

func (h *Handler) Shutdown() error {
	close(h.shutdown.ch)
	err := h.lstn.Close()
	h.shutdown.wg.Wait()
	return err
}

func (h *Handler) handleConn(conn net.Conn) {
	addr := conn.RemoteAddr()
	if addr == nil {
		addr = conn.LocalAddr()
	}
	l := h.log.With(slog.String("addr", addr.String()))

	defer func() {
		h.shutdown.wg.Done()
		if err := conn.Close(); err != nil {
			l.Error("can't close connection", logger.Err(err))
		}
	}()

	l.Info("accepted new connection")

	buf := make([]byte, 1024)
	for {
		select {
		default:
			n, err := conn.Read(buf)
			if err != nil {
				l.Error("can't read data", logger.Err(err))
				break
			}

			_, err = conn.Write(buf[:n])
			if err != nil {
				l.Error("can't write data", logger.Err(err))
				break
			}

		case <-h.shutdown.ch:
			_, err := conn.Write([]byte("close"))
			if err != nil {
				l.Info("can't write close message in connection", logger.Err(err))
			}
			return
		}
	}
}
