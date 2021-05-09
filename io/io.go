package io

import (
	"io"
	"net/http"
)

func ioReadAll(body io.Reader, w http.ResponseWriter) {
	b, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func ioCopy(body io.Reader, w http.ResponseWriter) {
	_, err := io.Copy(w, body)
	if err != nil {
		panic(err)
	}
}