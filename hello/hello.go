package main

import (
    "fmt"

    "go/greetings"
)

func main() {
  message := greetings.Hello("Gladys");
  fmt.Println(message)
}
