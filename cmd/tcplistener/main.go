package main

import (
  "fmt"
  "io"
  "net"
  "os"
  "strings"
)

func main() {
  //file, err := os.Open("messages.txt")
  //if err != nil {
  //  os.Exit(1)
  //}
  
  srv, err := net.Listen("tcp", "localhost:42069")
    if err != nil {
        os.Exit(1)
    }
    defer srv.Close()

    for {
        conn, err := srv.Accept()
        if err != nil {
            os.Exit(2)
        }

        fmt.Println("connection accepted...")
        lc := getLinesChannel(conn)
        for line := range lc {
            fmt.Printf("%s\n", line)
        }
        fmt.Println("connection closed...")
        conn.Close()
    }
  
  

  //file.Close()
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
