package main

import (
  "fmt"
  "io"
  "os"
  "strings"
)

func main() {
  file, err := os.Open("messages.txt")
  if err != nil {
    os.Exit(1)
  }

  data := make([]byte, 8)
  var line string
  for {
    _, err := file.Read(data)
    if err == io.EOF {
      if len(line) > 0 {
        fmt.Printf("read: %s\n", line)
      }
      os.Exit(0)
    }

    parts := strings.Split(string(data), "\n")
    line += parts[0]
    if len(parts) == 2 {
      fmt.Printf("read: %s\n", line)
      line = parts[1]
    }
  }
}
