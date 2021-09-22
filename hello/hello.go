package main

import (
    "fmt"
    "github.com/krzysbaranski/golang-experiments/greetings"
    "log"
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

    names := []string{"Krzys", "Jane"}
    greets, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    for _, g := range greets {
        fmt.Println(g)
    }
}
