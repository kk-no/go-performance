package goroutine

import (
	"io"
	"os"
	"testing"
)

var image []byte

func init() {
	body, err := os.Open("../assets/image.jpg")
	if err != nil {
		panic(err)
	}
	defer body.Close()
	image, _ = io.ReadAll(body)
}

func BenchmarkBasicAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicAppend(image)
	}
}

func BenchmarkUseGoroutineAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useGoroutineAppend(image)
	}
}
