package request

import (
  "io"
  "log"
  "strings"
)

type Request struct {
  RequestLine RequestLine
}

type RequestLine struct {
  HttpVersion   string
  RequestTarget string
  Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
  data, err := io.ReadAll(reader)
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  req := strings.Split(string(data), "\r\n")[0]
  reqLine := strings.Split(req, " ")

  log.Println(req)
  return &Request{RequestLine{
    HttpVersion:    strings.Split(reqLine[2], "HTTP/")[1],
    RequestTarget:  reqLine[1],
    Method:         reqLine[0],
  }}, nil
}
