package main

import (
    "bufio"
    "log"
    "net"
)


func main() {
    ln, err := net.Listen("tcp", ":9999")

    if err != nil {
        log.Fatalf("listen error: %v", err)
    }

    accepted := 0

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatalf("accept error: %v", err)
        }
        accepted++
        go serve(conn)
        log.Printf("Accepted %d", accepted)
    }
}


func serve(conn net.Conn) {
    bufr := bufio.NewReader(conn)

    for {
        line, err := bufr.ReadString('\n')
        if err != nil {
            return
        }
        conn.Write([]byte(line))
    }
}
