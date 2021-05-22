package webnew

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	})

	mux.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Haee")
	})

	mux.HandleFunc("/images/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "images")
	})

	mux.HandleFunc("/images/thumbnail", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "images/thumbnail")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RemoteAddr)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Header)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
