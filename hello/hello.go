package main

import (
    "fmt"
    "log"
    "github.com/krzysbaranski/golang-experiments/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    // Get a greeting message and print it.
    message, err := greetings.Hello("Gladys")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}
