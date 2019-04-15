package serve

import (
	"fmt"
	"github.com/andygeiss/tinygo/internal/pkg/check"
	"github.com/andygeiss/tinygo/internal/pkg/log"
	"github.com/andygeiss/tinygo/internal/pkg/serve/mime"
	"io/ioutil"
	"net/http"
	"os"
)

type staticHandler struct {
	content     []byte
	contentType string
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", h.contentType)
	if h.contentType == mime.Wasm {
		w.Header().Set("content-encoding", "gzip")
	}
	log.Info(fmt.Sprintf("Receiving [%-20s] [%d] [%s] [%-40s]", r.RemoteAddr, 200, r.Method, r.RequestURI))
	_, err := w.Write(h.content)
	check.Error(err)
}

// AddStaticHandler ...
func AddStaticHandler(pattern, contentType, filename string) {
	log.Info(fmt.Sprintf("Adding [%-40s] from file [%-40s]", pattern, filename))
	h := &staticHandler{content: readAll(filename), contentType: contentType}
	http.Handle(pattern, h)
	handlers[pattern] = h
}

func readAll(filename string) (content []byte) {
	file, err := os.Open(filename)
	check.Fatal(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	check.Fatal(err)
	content = data
	return
}
