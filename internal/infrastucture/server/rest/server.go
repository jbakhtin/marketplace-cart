package rest

import (
	"context"
	"fmt"
	router "github.com/jbakhtin/marketplace-cart/internal/infrastucture/server/rest/router/chi"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart"
	"github.com/jbakhtin/marketplace-cart/internal/modules/cart/ports"
	"net/http"
)

type Server struct {
	logger ports.Logger
	http.Server
}

type Config interface {
	GetServerHTTPAddress() string
}

func NewWebServer(
	cfg Config,
	logger ports.Logger,
	cartModule cart.Module,
) (Server, error) {
	handler, err := router.NewRouter(&cfg, logger, cartModule.GetCartUseCase())
	if err != nil {
		return Server{}, err
	}

	return Server{
		logger: logger,
		Server: http.Server{
			Addr:    cfg.GetServerHTTPAddress(),
			Handler: handler,
		},
	}, nil
}

func (s *Server) Start(ctx context.Context) (err error) {
	go func() {
		err = s.ListenAndServe()
		s.logger.Info(err.Error())
	}()

	s.logger.Info(fmt.Sprintf("webserver available on: %v", s.Server.Addr))

	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
