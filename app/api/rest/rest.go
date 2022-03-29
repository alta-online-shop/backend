package rest

import (
	"net/http"

	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

type Rest struct {
	mux *ngamux.Ngamux
}

func New() *Rest {
	mux := ngamux.New()
	rest := &Rest{mux}

	NewHelloRest(ctx.Mux(mux.Group("/hello")))
	NewProductRest(ctx.Mux(mux.Group("/products")))

	return rest
}

func (r *Rest) Start() error {
	return http.ListenAndServe(":8080", r.mux)
}
