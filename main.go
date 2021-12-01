package main

import (
	"net/http"

	"github.com/abc7468/study.go/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
