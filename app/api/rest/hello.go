package rest

import (
	"net/http"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

type HelloRest struct {
	mux     *ctx.CtxMux
	service *service.Service
}

func NewHelloRest(mux *ctx.CtxMux) *HelloRest {
	hello := &HelloRest{mux: mux}
	hello.route()
	return hello
}

func (r *HelloRest) route() {
	r.mux.Get("/", r.Index)
}

func (r *HelloRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *HelloRest) Index(c *ctx.Context) error {
	message := r.getService().Hello.GetMessage()
	return c.JSON(http.StatusOK, ngamux.Map{
		"message": message,
	})
}
