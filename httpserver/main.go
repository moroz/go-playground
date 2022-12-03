package main

import (
	"context"
	"errors"
	"fmt"
	"io"
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

	log.Printf("%s: got / request. first(%t)=%s, second(%t)=%s\n", ctx.Value(keyServerAddr), hasFirst, first, hasSecond, second)
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

	ctx, cancelCtx := context.WithCancel(context.Background())

	baseContext := func(l net.Listener) context.Context {
		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
		return ctx
	}

	serverOne := &http.Server{
		Addr:        getPort(),
		Handler:     mux,
		BaseContext: baseContext,
	}

	serverTwo := &http.Server{
		Addr:        ":4444",
		Handler:     mux,
		BaseContext: baseContext,
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
