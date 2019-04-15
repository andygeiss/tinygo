package serve

import (
	"bytes"
	"fmt"
	"github.com/andygeiss/tinygo/internal/pkg/check"
	"github.com/andygeiss/tinygo/internal/pkg/log"
	"html/template"
	"net/http"
)

type templatesHandler struct {
	content     []byte
	contentType string
	data        interface{}
}

func (h *templatesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", h.contentType)
	log.Info(fmt.Sprintf("Receiving [%-20s] [%d] [%s] [%-40s]", r.RemoteAddr, 200, r.Method, r.RequestURI))
	_, err := w.Write(h.content)
	check.Error(err)
}

// AddTemplatesHandler ...
func AddTemplatesHandler(pattern, contentType, filename string, data interface{}) {
	log.Info(fmt.Sprintf("Adding [%-40s] from file [%-40s]", pattern, filename))
	content := readAll(filename)
	merged := execute(content, data)
	h := &templatesHandler{content: merged, contentType: contentType, data: data}
	http.Handle(pattern, h)
	handlers[pattern] = h
}

func execute(content []byte, data interface{}) (merged []byte) {
	tmpl, err := template.New("t").Parse(string(content))
	check.Fatal(err)
	var m bytes.Buffer
	err = tmpl.Execute(&m, data)
	check.Fatal(err)
	merged = m.Bytes()
	return
}
