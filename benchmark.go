package benchmark

import "fmt"
import "http"
import "time"

type BenchmarkResultItem struct {
    timeTaken int64 // miliseconds
    status int    
}

type BenchmarkResults struct {
    timeTaken float
    requestsPerformed int
    successfulRequests int
    failedRequests int
    url string
    items []BenchmarkResultItem
}

func HostExists( url string) bool {
    var r *http.Response
//    var finalUrl string
 //   var Error err
    r,_,_ = http.Get(url)
    if ( r.StatusCode == 200){
        fmt.Printf("Host " + url + " is reachable.")
        return true
    }else{
        fmt.Printf("Host " + url + " is down.")
        return false
    }
    return false;
}

func PerformTests(url *string, numRequests *int, numWorkers *int) (results *BenchmarkResults){
    results = new (BenchmarkResults)
    results.items = make([]BenchmarkResultItem, *numRequests)
    var requestsPerWorker = *numRequests/ *numWorkers
    
    channel := make(chan int, *numWorkers)
    
    for i:=0;i<*numWorkers;i++ {
        slice  :=results.items[i*requestsPerWorker:(i+1)*requestsPerWorker]
    go worker(url, &requestsPerWorker ,slice,channel)

    }
    val := 0
    
    for {
        val += <-channel
        if (val%(*numRequests/10)==0){
            fmt.Printf("Completed %d requests.\n",val)
            
        }
        if ( val == *numRequests ){
            break;
        }
    }
    results.requestsPerformed = *numRequests
    return results
}

// note beginning of function name: small caps means not exported! 
func worker(url *string, numRequests *int,slice []BenchmarkResultItem, c chan int){
    var startTime int64
    z := 4
    fmt.Printf("foo %d", z)
    for i:=0; i < *numRequests; i++ {
        startTime = time.Nanoseconds()
        var r *http.Response
        r,_,_ = http.Get(*url)
        var b = new(BenchmarkResultItem)
        b.timeTaken = (time.Nanoseconds() - startTime) / 1000000
        b.status = r.StatusCode
       slice[i] = *b
       
        c <- 1
    }
    
    return
}

func (b*BenchmarkResults) AverageResponseTime() ( avg float){
    timeTaken  := int64(0)
    forRequests := 0
    for _,item := range b.items {
        timeTaken += item.timeTaken
        forRequests += 1
    }
    avg = float(timeTaken)/float(forRequests)
    return
}

func (b*BenchmarkResults) ErrorRate() (total int, errors int){
    errors = 0
    total = 0
    for _,item := range b.items {
        if (item.status - 200 > 100){
            errors += 1
        }
        total += 1
    }

    return
}
