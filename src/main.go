package main

import (
  "fmt"
  "os"
  "os/user"
  "repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hello %s! This is a toy C compiler!\n",
    user.Username)
  fmt.Printf("Feel free to type in codes\n")
  repl.Start(os.Stdin, os.Stdout)
}
