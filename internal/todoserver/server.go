package todoserver

import (
	"context"
	"protobuf_using_go/twitchtv/twirp"

	// Import our protobuf package, namespace it to pb. This is a convention!
	pb "github.com/coredefend/rpc-todos/rpc/todo"
)

// Create our Server type
type Server struct{}

// Seed our application with one To-Do to start with.
// The map holds the To-Do ID as the key to a To-Do value.
var todos = map[int]*pb.Todo{
	1: &pb.Todo{ID: 1, Title: "hello world", Complete: false},
}

// MakeTodo creates a To-Do
func (s *Server) MakeTodo(ctx context.Context, todo *pb.Todo) (*pb.Empty, error) {
	// Validate the given ID, return error if not > 0
	if todo.ID < 1 {
		return nil, twirp.InvalidArgumentError("ID:", "No ID provided")
	}
	// Validate Title, should not add empty To-Do.
	if todo.Title == "" {
		return nil, twirp.InvalidArgumentError("Title:", "No title given")
	}
	// Set To-Do in the Todos map, passing ID as key.
	todos[todo.ID] = todo
	// Return Empty response and nil error.
	return &pb.Empty{}, nil
}

// GetTodo returns a single To-Do for given ID.
func (s *Server) GetTodo(ctx context.Context, query *pb.TodoQuery) (*pb.Todo, error) { return nil, nil }

// AllTodos returns all To-Dos.
func (s *Server) AllTodos(ctx context.Context, e *pb.Empty) (*pb.TodoResponse, error) { return nil, nil }

// RemoveTodo deletes the To-Do with the given ID.
func (s *Server) RemoveTodo(ctx context.Context, q *pb.TodoQuery) (*pb.Empty, error) { return nil, nil }

// MarkComplete sets the To-Do with given ID to completed.
func (s *Server) MarkComplete(ctx context.Context, q *pb.TodoQuery) (*pb.Empty, error) {
	return nil, nil
}
