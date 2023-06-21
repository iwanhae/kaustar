package server

import (
	"context"
	"net/http"

	"github.com/iwanhae/kaustar/api"
	"github.com/iwanhae/kaustar/pkg/config"
	"github.com/iwanhae/kaustar/pkg/server/proxy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func NewServer(cfg config.Config) (*echo.Echo, error) {
	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	{ // Porxy to Kubernetes API server
		config, err := proxy.LoadConfig(cfg.Kubernetes)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load kubernetes config")
		}
		e.Any("/proxy/*", proxy.KubeAPIServer(config),
			middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{Skipper: middleware.DefaultSkipper, RedirectCode: 307}),
			middleware.Rewrite(map[string]string{"/proxy/*": "/$1"}),
		)
	}

	{ // Kaustar API
		const prefix = "/api/v1"
		server, err := api.NewServer(&Server{}, &Authenticator{},
			api.WithPathPrefix(prefix),
			api.WithNotFound(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "not found", http.StatusNotFound)
			}),
			api.WithErrorHandler(func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}),
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to initialize api server")
		}
		e.Any(prefix+"/*", echo.WrapHandler(server))
	}
	return e, nil
}
