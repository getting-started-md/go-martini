package martini_amber

import (
	"github.com/eknkc/amber"
	"github.com/go-martini/martini"
	"github.com/oxtoacart/bpool"
	"html/template"
	"io"
	"net/http"
)

const (
	ContentType       = "Content-Type"
	ContentHTML       = "text/html"
	TemplateDirectory = "directory"
)

var bufpool *bpool.BufferPool

type Render interface {
	Amber(status int, template string, variables interface{})
	AmberOK(template string, variables interface{})
	AmberNotFound(template string, variables interface{})
}

func Renderer(options map[string]string) martini.Handler {

	templateDir := "templates/"
	if value, containKey := options[TemplateDirectory]; containKey {
		templateDir = value
	}

	cache, _ := amber.CompileDir(templateDir, amber.DefaultDirOptions, amber.DefaultOptions)
	bufpool = bpool.NewBufferPool(64)
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		var templates map[string]*template.Template
		if martini.Env == martini.Dev {
			templates, _ = amber.CompileDir(templateDir, amber.DefaultDirOptions, amber.DefaultOptions)
		} else {
			templates = cache
		}
		c.MapTo(&AmberRenderer{res, req, templates}, (*Render)(nil))
	}
}

type AmberRenderer struct {
	http.ResponseWriter
	req       *http.Request
	templates map[string]*template.Template
}

func (r *AmberRenderer) AmberOK(template string, variables interface{}) {
	r.Amber(200, template, variables)
}

func (r *AmberRenderer) AmberNotFound(template string, variables interface{}) {
	r.Amber(404, template, variables)
}

func (r *AmberRenderer) Amber(status int, template string, variables interface{}) {
	buf := bufpool.Get()
	err := r.templates[template].Execute(buf, variables)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Header().Set(ContentType, ContentHTML)
	r.WriteHeader(status)
	io.Copy(r, buf)
	bufpool.Put(buf)

}
