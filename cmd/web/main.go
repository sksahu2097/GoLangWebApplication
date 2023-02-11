package main

import (
	"fmt"
	"net/http"

	"github.com/sksahu2097/go-project/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting the application on 8080")
	_ = http.ListenAndServe(":8080", nil)

}
