package web

import (
	"net/http"

	"github.com/hadihammurabi/dummy-online-shop/app/helper"
	"github.com/ngamux/ngamux"
)

type CategoryWeb struct {
	mux *ngamux.Ngamux
}

func NewCategoryWeb(mux *ngamux.Ngamux) *CategoryWeb {
	r := &CategoryWeb{mux: mux}
	r.route()
	return r
}

func (w *CategoryWeb) route() {
	w.mux.Get("/kategori/tambah", w.Create)
}

func (w *CategoryWeb) Create(rw http.ResponseWriter, r *http.Request) error {
	return helper.RenderPage(rw, "kategori/tambah.html", nil)
}
