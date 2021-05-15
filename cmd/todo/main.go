package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/coredefend/rpc-todo/rpc/todo"
)

// Create global To-Do counter
var todoCount int32 = 1

func main() {
	// Define our interaction flags.
	action := flag.String("command", "list", "Command to interact with Todos. create, get, list, delete")
	title := flag.String("title", "", "The todo title")
	id := flag.Int("id", 1, "The ID of the todo you wish to retrieve")
	flag.Parse()

	// Create the generated Twirp Protobuf Client for our Todo service
	client := todo.NewTodoProtobufClient("http://localhost:8080", &http.Client{})

	switch *action {
	case "create":
		_, err := client.MakeTodo(context.Background(), &todo.Todo{
			ID:       todoCount + 1,
			Title:    *title,
			Complete: false,
		})
		if err != nil {
			fmt.Println("could not create todo:", err)
		}
	case "list":
		res, err := client.AllTodos(context.Background(), &todo.Empty{})
		if err != nil {
			fmt.Println("could not fetch todos:", err)
		}
		for i, t := range res.Todos {
			fmt.Printf("%d. %s [%v]\n", i+1, t.Title, t.Complete)
		}
	case "get":
		todo, err := client.GetTodo(context.Background(), &todo.TodoQuery{
			Id: int32(*id),
		})
		if err != nil {
			fmt.Println("could not fetch todo:", err)
		}

		fmt.Printf("%d. %s [%v]\n", todo.ID, todo.Title, todo.Complete)
	case "delete":
		_, err := client.RemoveTodo(context.Background(), &todo.TodoQuery{
			Id: int32(*id),
		})
		if err != nil {
			fmt.Println("could not remove todo:", err)
		}
	case "complete":
		_, err := client.MarkComplete(context.Background(), &todo.TodoQuery{
			Id: int32(*id),
		})
		if err != nil {
			fmt.Println("could not complete todo:", err)
		}
	default:
		fmt.Println("invalid command, please try again.")
	}
}
