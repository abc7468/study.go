package main

import (
	"net/http"

	"github.com/abc7468/study.go/decoapp"
)

func main() {
	mux := decoapp.NewDecoHandler()
	http.ListenAndServe(":3000", mux)
	//	http.ListenAndServe(":3000", myapp.NewHandler())
}
