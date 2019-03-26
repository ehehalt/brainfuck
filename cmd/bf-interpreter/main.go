package main

import (
  "fmt"
  "io/ioutil"
  "os"

  "github.com/ehehalt/brainfuck/interpreter"
)

func main() {
  fmt.Println("Brainf*ck")
  fileName := os.Args[1]
  code, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Fprintf(os.Stderr, "error: %s\n", err)
    os.Exit(-1)
  }

  m := interpreter.NewMachine(string(code), os.Stdin, os.Stdout)
  m.Execute()
}
