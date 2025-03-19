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
  
  lc := getLinesChannel(file)
  for line := range lc {
    fmt.Printf("read: %s\n", line)
  }

  file.Close()
}

func getLinesChannel(f io.ReadCloser) <-chan string {
  lc := make(chan string)

  data := make([]byte, 8)
  var line string
  go func() {
    for {
      _, err := f.Read(data)
      if err == io.EOF {
        if len(line) > 0 {
          lc <- line
          close(lc)
          break
        }
        os.Exit(0)
      }

      parts := strings.Split(string(data), "\n")
      line += parts[0]
      if len(parts) == 2 {
        lc <- line
        line = parts[1]
      }
    }
  }()

  return lc
}
