// simple benchmaker, in go
// moritz haarmann hats verbrochen.

package main

import "fmt"
import "flag"
import "./benchmark"

var  omitNewline = flag.Bool("n", false, "don't print final newline")
var url = flag.String("h","","the host to check")
var numRequests *int = flag.Int("r",100,"the number of requests to perform")
var numWorkers = flag.Int("c",1,"the number of workers")

func main() {
    flag.Parse();
    b := []int{1,2,3,4,56,6}
    fmt.Printf("%d", len(b[0:4]))
    var exists  = benchmark.HostExists(*url)
    if ( exists ){
    fmt.Printf("Firing %d requests using %d parallel workers.", *numRequests, *numWorkers)
        results := benchmark.PerformTests(url, numRequests, numWorkers)
        fmt.Printf("Average response time: %f", results.AverageResponseTime())
        total, error := results.ErrorRate()
        fmt.Printf("\nTotal requests: %d\nFailed requests: %d",total, error )
        
    } else {
        fmt.Printf("Can't perform any tests on down host.")
    }
}
