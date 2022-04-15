package rest

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

type Rest struct {
	mux *ngamux.Ngamux
}

func (r *Rest) buildRoute() {
	NewHelloRest(ctx.Mux(r.mux.Group("/api/hello")))
	NewProductRest(ctx.Mux(r.mux.Group("/api/products")))
	NewCategoryRest(ctx.Mux(r.mux.Group("/api/categories")))
}

func New() *Rest {
	mux := ngamux.New()
	rest := &Rest{mux}

	rest.buildRoute()

	return rest
}

func (r *Rest) Start() error {
	addr := "0.0.0.0:8081"
	fmt.Printf("App run at %s", addr)

	return http.ListenAndServe(addr, r.mux)
}
