package io

import (
	"bytes"
	"io"
	"net/http/httptest"
	"os"
	"testing"
)

var image []byte

func init() {
	body, err := os.Open("../assets/image.jpg")
	if err != nil {
		panic(err)
	}
	image, _ = io.ReadAll(body)
	body.Close()
}

func BenchmarkIOReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		r := bytes.NewReader(image)
		ioReadAll(r, w)
	}
}

func BenchmarkIOCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		r := bytes.NewReader(image)
		ioCopy(r, w)
	}
}