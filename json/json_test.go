package json

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var body []byte

func init() {
	b, err := os.ReadFile("../assets/large.json")
	if err != nil {
		panic(err)
	}
	body = b
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	cli := &http.Client{}
	req, _ := http.NewRequestWithContext(context.TODO(), "GET", ts.URL, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ := cli.Do(req)
		jsonUnmarshal(res.Body)
		res.Body.Close()
	}
}

func BenchmarkJSONDecode(b *testing.B) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	cli := &http.Client{}
	req, _ := http.NewRequestWithContext(context.TODO(), "GET", ts.URL, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ := cli.Do(req)
		jsonDecode(res.Body)
		res.Body.Close()
	}
}
