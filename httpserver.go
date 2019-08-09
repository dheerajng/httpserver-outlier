// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "log"
    "net/http"
    "net"
    "os"
)

func getMyIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        os.Stderr.WriteString("Oops: " + err.Error() + "\n")
        os.Exit(1)
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, You asked for %s! I am %s", r.URL.Path[1:], getMyIP())
}

func misbehaveHandler(w http.ResponseWriter, r *http.Request) {
     // see http://golang.org/pkg/net/http/#pkg-constants
     w.WriteHeader(http.StatusServiceUnavailable)
     fmt.Fprintf(w, "MISBEHAVING! HTTP status code returned from %s!", getMyIP())
     //w.Write([]byte("MISBEHAVING! HTTP status code returned!"))
}

func main() {
    http.HandleFunc("/", viewHandler)
    http.HandleFunc("/misbehave", misbehaveHandler)
    log.Fatal(http.ListenAndServe(":4040", nil))
}


