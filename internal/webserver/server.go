package webserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ListenAndServe(port string, handler http.Handler) {
	s := &http.Server{
		Addr:           port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Curio QA Server Running on port %s\n", port)

	log.Fatal(s.ListenAndServe())
}
