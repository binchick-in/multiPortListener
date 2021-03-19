package main

import (
    "fmt"
    "net/http"
    "sync"
)

func soleRoute(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    fmt.Fprintf(w, ":)")
}

// Give a start and end port to
// generate a slice of port numbers
func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func main() {
    fmt.Println("Staring Main Thread...")
    portRange := makeRange(1000, 2000)
    var wg sync.WaitGroup
    wg.Add(len(portRange))
    http.HandleFunc("/", soleRoute)
    for _, port := range portRange {
        fmt.Printf("Starting concurrent server on... %d\n", port)
        go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    }
    wg.Wait()
}
