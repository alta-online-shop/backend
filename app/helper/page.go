package helper

import (
	"net/http"
	"path"
	"text/template"

	"github.com/hadihammurabi/dummy-online-shop/app/config"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
)

func Render(rw http.ResponseWriter, file string, data interface{}) error {
	cfg := ioc.Use(config.Config{}).(*config.Config)
	t, err := template.ParseFiles(path.Join(cfg.ResourceDir, file))
	if err != nil {
		return err
	}
	if err = t.Execute(rw, data); err != nil {
		return err
	}

	return nil
}

func RenderPage(rw http.ResponseWriter, file string, data interface{}) error {
	file = path.Join("pages", file)
	return Render(rw, file, data)
}
