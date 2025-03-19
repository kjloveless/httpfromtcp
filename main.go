package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  file, err := os.Open("messages.txt")
  if err != nil {
    os.Exit(1)
  }

  data := make([]byte, 8)
  for {
    _, err := file.Read(data)
    if err == io.EOF {
      os.Exit(0)
    }
    fmt.Printf("read: %s\n", data)
  }
}
