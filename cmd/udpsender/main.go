package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    addr, err := net.ResolveUDPAddr("udp", ":42069")
    if err != nil {
        log.Fatalf("Error resolving: %v", err)
    }

    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        log.Fatalf("Error dialing: %v", err)
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Fatalf("Error reading: %v", err)
        }

        _, err = conn.Write([]byte(line))
        if err != nil {
            log.Fatalf("Error writing: %v", err)
        }
    }
}
