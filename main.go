package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/yanshuy/go+htmx/components"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Has("global") {
		global.Count++
	}
	templ.Handler(components.Form(global.Count, 0)).ServeHTTP(w, r)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static/"))
	router.Handle("GET /static/", http.StripPrefix("/static", fs))

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div>hile</div>"))
	})

	component := components.Home("john")
	router.Handle("GET /home", templ.Handler(component))
	router.Handle("/404", templ.Handler(components.NotFoundComponent(), templ.WithStatus(http.StatusNotFound)))
	router.Handle("GET /counter", templ.Handler(components.Page(global.Count, 0)))
	router.HandleFunc("POST /counter", postHandler)

	server := &http.Server{
		Addr:    ":4200",
		Handler: router,
	}

	logger.Info("Server listening on port " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("error", "err", err)
	}

}
