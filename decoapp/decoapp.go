package decoapp

import (
	"fmt"
	"log"
	"net/http"
	"time"

	decohandler "github.com/abc7468/study.go/decoHandler"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed", time.Since(start).Milliseconds())

}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed", time.Since(start).Milliseconds())

}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func NewDecoHandler() http.Handler {
	h := NewHandler()
	h = decohandler.NewDecoHandler(h, logger)
	h = decohandler.NewDecoHandler(h, logger2)
	return h
}
