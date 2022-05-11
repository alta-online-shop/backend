package rest

import (
	"fmt"
	"net/http"

	"github.com/ngamux/ctx"
	"github.com/ngamux/middleware/cors"
	"github.com/ngamux/ngamux"
)

type Rest struct {
	mux *ngamux.Ngamux
}

func (r *Rest) buildRoute() {
	NewHelloRest(ctx.Mux(r.mux.Group("/api/hello")))
	NewAuthRest(ctx.Mux(r.mux.Group("/api/auth")))
	NewProductRest(ctx.Mux(r.mux.Group("/api/products")))
	NewCategoryRest(ctx.Mux(r.mux.Group("/api/categories")))
	NewOrderRest(ctx.Mux(r.mux.Group("/api/orders")))
}

func New() *Rest {
	mux := ngamux.New()
	mux.Use(cors.New(cors.Config{
		AllowHeaders: "content-type,authorization",
	}))
	// mux.Use(recover.New())
	rest := &Rest{mux}

	rest.buildRoute()

	return rest
}

func (r *Rest) Start() error {
	addr := "0.0.0.0:8081"
	fmt.Printf("App run at %s\n", addr)

	return http.ListenAndServe(addr, r.mux)
}
