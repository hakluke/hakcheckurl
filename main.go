package main

import (
    "bufio"
    "crypto/tls"
    "flag"
    "fmt"
    "net"
    "net/http"
    "os"
    "sync"
    "time"
)

func main() {
    concurrencyPtr := flag.Int("t", 8, "Number of threads to utilise. Default is 8.")
    timeoutPtr := flag.Int("timeout", 8, "Timeout for each request in seconds. Default is 8 seconds.")
    retryPtr := flag.Int("retry", 3, "Number of retries for failed requests. Default is 3.")
    retrySleepPtr := flag.Int("retry-sleep", 1, "Sleep duration between retries in seconds. Default is 1 second.")
    flag.Parse()

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
            Dial:                (&net.Dialer{Timeout: time.Duration(*timeoutPtr) * time.Second, KeepAlive: 0}).Dial,
            TLSHandshakeTimeout: time.Duration(*timeoutPtr) * time.Second,
        },
    }

    numWorkers := *concurrencyPtr
    work := make(chan string)

    go func() {
        s := bufio.NewScanner(os.Stdin)
        for s.Scan() {
            work <- s.Text()
        }
        if s.Err() != nil {
            fmt.Println("Error reading from stdin:", s.Err())
        }
        close(work)
    }()

    wg := &sync.WaitGroup{}

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go doWork(work, client, wg, *retryPtr, *retrySleepPtr)
    }
    wg.Wait()
}

func doWork(work chan string, client *http.Client, wg *sync.WaitGroup, retries int, retrySleep int) {
    defer wg.Done()
    for url := range work {
        var resp *http.Response
        var err error
        for attempts := 0; attempts <= retries; attempts++ {
            req, reqErr := http.NewRequest("GET", url, nil)
            if reqErr != nil {
                fmt.Println(999, reqErr, url)
                break
            }
            req.Header.Set("Connection", "close")

            resp, err = client.Do(req)
            if err == nil || attempts == retries {
                if err != nil {
                    fmt.Println(999, err, url)
                }
                if resp != nil {
                    resp.Body.Close()
                    if err == nil {
                        fmt.Println(resp.StatusCode, url)
                    }
                }
                break
            }
            time.Sleep(time.Duration(retrySleep) * time.Second)
        }
    }
}
