package serve

import (
	"fmt"
	"github.com/andygeiss/tinygo/internal/pkg/check"
	"github.com/andygeiss/tinygo/internal/pkg/log"
	"net/http"
)

var handlers = map[string]http.Handler{}

// SafeListenAndServe ...
func SafeListenAndServe(addr string) {
	log.Info(fmt.Sprintf("Start listening at [%-20s] ...", addr))
	err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for p, h := range handlers {
			if r.RequestURI == p {
				h.ServeHTTP(w, r)
				return
			}
		}
		log.Warn(fmt.Sprintf("Receiving [%-20s] [%d] [%s] [%-40s]", r.RemoteAddr, 404, r.Method, r.RequestURI))
		w.WriteHeader(http.StatusNotFound)
	}))
	check.Fatal(err)
}
