package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/yanshuy/go+htmx/components"
)

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
