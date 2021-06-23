package io

import (
	"bytes"
	"net/http/httptest"
	"os"
	"testing"
)

var image []byte

func init() {
	body, err := os.ReadFile("../assets/image.jpg")
	if err != nil {
		panic(err)
	}
	image = body
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
