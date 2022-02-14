package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"workspace/product-api/handlers"
)

var serverStarting = flag.Bool("server", true, "Started to server")

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	flag.Parse()
	//****** create the handlers ******
	ph := handlers.NewProducts(l)
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)

	//****** create a new serve mux and register the handlers ******
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	fmt.Println("server:", *serverStarting)
	//sm.Handle("/", hh)
	//sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan) // For all catch signals
	//	signal.Notify(sigChan, os.Interrupt)
	//	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("Got signal:", sig)
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
