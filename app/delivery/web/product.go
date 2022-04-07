package web

import (
	"net/http"

	"github.com/hadihammurabi/dummy-online-shop/app/helper"
	"github.com/ngamux/ngamux"
)

type ProductWeb struct {
	mux *ngamux.Ngamux
}

func NewProductWeb(mux *ngamux.Ngamux) *ProductWeb {
	r := &ProductWeb{mux: mux}
	r.route()
	return r
}

func (w *ProductWeb) route() {
	w.mux.Get("/", w.Index)
	w.mux.Get("/produk/tambah", w.Create)
}

func (w *ProductWeb) Index(rw http.ResponseWriter, r *http.Request) error {
	return helper.RenderPage(rw, "index.html", nil)
}

func (w *ProductWeb) Create(rw http.ResponseWriter, r *http.Request) error {
	return helper.RenderPage(rw, "produk/tambah.html", nil)
}
