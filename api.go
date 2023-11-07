package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
	"github.com/imnulhaqueruman/docker-and-go/handlers"
)

func main() {
	mux := bone.New()

	mux.Get("/", http.HandlerFunc(handlers.NewSession))
	mux.Get("/sessions/:sessionId", http.HandlerFunc(handlers.GetSession))
	mux.Post("/sessions/:sessionId/instances", http.HandlerFunc(handlers.NewInstance))
	// mux.Delete("/sessions/:sessionId/instances/:instanceId", http.HandlerFunc(handlers.DeleteInstance))

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./www/index.html")
	})
	mux.Get("/p/:sessionId", h)
	mux.Get("/assets/*", http.FileServer(http.Dir("./www")))

	mux.Get("/exec/:id", websocket.Handler(handlers.Exec))

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe(":3000", n))

}