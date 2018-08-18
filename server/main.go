package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Hello world")

	mux := http.NewServeMux()
	mux.Handle("recommendations", recommendations.Handler{})
}