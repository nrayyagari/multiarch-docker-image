package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    
    hostname, _ := os.Hostname()
    
    fmt.Printf("Hello! I'm container '%s'\n", hostname)
    fmt.Printf("and I'm running on host machine with %s/%s architecture!\n", runtime.GOOS, runtime.GOARCH)
}