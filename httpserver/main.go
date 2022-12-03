package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("could not read body: %s\n", err)
	}

	log.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n", ctx.Value(keyServerAddr), hasFirst, first, hasSecond, second, body)
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello, HTTP!\n")
}

const DEFAULT_PORT = ":4000"

func getPort() string {
	port := os.Getenv("PORT")
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return DEFAULT_PORT
	}
	return fmt.Sprintf(":%d", portNumber)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()

	server := &http.Server{
		Addr:    getPort(),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
}
