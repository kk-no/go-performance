package slice

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var image []byte

func init() {
	body, _ := os.Open("../assets/image.jpg")
	defer body.Close()
	image, _ = ioutil.ReadAll(body)
	fmt.Println(len(image))
}

func BenchmarkNoMakeAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noMakeAppend(image)
	}
}

func BenchmarkUseMakeAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useMakeAppend(image)
	}
}

func BenchmarkUseMakeIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useMakeIndex(image)
	}
}
