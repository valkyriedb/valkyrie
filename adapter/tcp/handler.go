package tcp

import (
	"errors"
	"log/slog"
	"net"
	"sync"

	"github.com/valkyriedb/valkyrie/adapter/message"
	"github.com/valkyriedb/valkyrie/adapter/message/status"
	decode "github.com/valkyriedb/valkyrie/internal/decoder"
	"github.com/valkyriedb/valkyrie/internal/logger"
	"github.com/valkyriedb/valkyrie/service"
)

var ErrWrongPassword = errors.New("wrong password")

type Handler struct {
	srv      service.Service
	password string

	lstn     net.Listener
	shutdown struct {
		ch chan struct{}
		wg sync.WaitGroup
	}
	log *slog.Logger
}

func NewHandler(srv service.Service, password, address string, log *slog.Logger) (*Handler, error) {
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

	for {
		err := h.auth(conn)
		if err == nil {
			_, err = conn.Write(message.Response{Status: status.OK}.ToBytes())
			if err != nil {
				l.Error("can't write to connection", logger.Err(err))
				return
			}
			break
		}

		_, err = conn.Write(message.Response{
			Status: status.Unauth,
		}.ToBytes())
		if err != nil {
			l.Error("can't write to connection", logger.Err(err))
			return
		}
	}

	for {
		select {
		default:
			req, err := message.ReadRequest(conn)
			if err != nil {
				if netErr, ok := err.(net.Error); ok {
					l.Error("can't read connection", logger.Err(netErr))
					return
				}

				_, err = conn.Write(message.Response{Status: status.InvalidReq}.ToBytes())
				if err != nil {
					l.Error("can't write to connection", logger.Err(err))
					return
				}
			}

			res := h.srv.Do(req)
			_, err = conn.Write(res.ToBytes())
			if err != nil {
				l.Error("can't write to connection", logger.Err(err))
				return
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

func (h *Handler) auth(conn net.Conn) error {
	reqPass, err := decode.String(conn)
	if err != nil {
		return err
	}

	if reqPass != h.password {
		return ErrWrongPassword
	}

	return nil
}
