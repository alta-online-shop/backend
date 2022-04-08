package web

import (
	"net/http"
	"path"

	"github.com/hadihammurabi/dummy-online-shop/app/config"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/ngamux/middleware/static"
	"github.com/ngamux/ngamux"
)

type Web struct {
	mux *ngamux.Ngamux
}

func (w *Web) buildRoute() {
	NewProductWeb(w.mux)
	NewCategoryWeb(w.mux)
}

func New() *Web {
	mux := ngamux.New()
	rest := &Web{mux}
	cfg := ioc.Use(config.Config{}).(*config.Config)

	mux.Use(static.New(static.Config{
		Root: path.Join(cfg.ResourceDir, "public"),
	}))

	rest.buildRoute()

	return rest
}

func (r *Web) Start() error {
	rest.NewWith(r.mux)
	return http.ListenAndServe(":8080", r.mux)
}
