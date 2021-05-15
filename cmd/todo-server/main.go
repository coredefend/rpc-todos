package main

import (
	"fmt"
	"net/http"

	"github.com/coredefend/rpc-todos/internal/todoserver"
	pb "github.com/coredefend/rpc-todos/rpc/todo"
)

func main() {
	// Create new Server from our server package.
	server := &todoserver.Server{}
	// Create Twirp Handler for our server from generated server stub.
	// This takes the form pb.New<SERVICE_NAME>Server()
	twirpHandler := pb.NewTodoServer(server, nil)

	// Start HTTP server on 8080, passing our twirpHandler as the servemux.
	http.ListenAndServe(":8080", twirpHandler)
	fmt.Println("To-Do server listening on :8080")
}
