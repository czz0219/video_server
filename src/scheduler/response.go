package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, sc int, resp string) error {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
	return nil
}
