package main

import (
	"github.com/andygeiss/serve"
	"github.com/andygeiss/serve/mime"
	"github.com/andygeiss/tinygo/config"
)

func main() {
	serve.AddStaticHandler("/app.wasm", mime.Wasm, "web/app/app.wasm")
	serve.AddStaticHandler("/favicon.ico", mime.Ico, "web/static/favicon.ico")
	serve.AddStaticHandler("/assets/css/styles.css", mime.Css, "web/static/css/styles.css")
	serve.AddStaticHandler("/assets/img/logo.png", mime.Png, "web/static/img/logo.png")
	serve.AddStaticHandler("/assets/js/pwa_init.js", mime.Javascript, "web/static/js/pwa_init.js")
	serve.AddStaticHandler("/assets/js/wasm_exec.js", mime.Javascript, "web/static/js/wasm_exec.js")
	serve.AddStaticHandler("/assets/js/wasm_init.js", mime.Javascript, "web/static/js/wasm_init.js")
	serve.AddTemplatesHandler("/manifest.json", mime.Json, "web/templates/manifest.json", config.DefaultApp)
	serve.AddTemplatesHandler("/service-worker.js", mime.Javascript, "web/templates/service-worker.js", config.DefaultApp)
	serve.AddTemplatesHandler("/", mime.Html, "web/templates/index.html", config.DefaultApp)
	serve.SafeListenAndServe(":80")
}
