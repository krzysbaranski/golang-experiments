package main

import (
    "fmt"

    "github.com/krzysbaranski/golang-experiments/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
