// Package webserver provides the necessary functionality to run a server.
package webserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ListenAndServe runs the server with the configured address and handler.
func ListenAndServe(port string, handler http.Handler) {
	s := &http.Server{
		Addr:           port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Curio QA Server is running on port %s\n", port)

	log.Fatal(s.ListenAndServe())
}
