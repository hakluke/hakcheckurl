// Receives lines from stdin, processes them with a user-defined number of threads

package main

import (
    "bufio"
    "flag"
    "fmt"
    "sync"
    "os"
    "net/http"
    "net"
    "time"
    "crypto/tls"
)


func main() {
        concurrencyPtr := flag.Int("t", 8, "Number of threads to utilise. Default is 8.")
        flag.Parse()
        client := &http.Client{Transport: &http.Transport{
                TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
                Dial:                (&net.Dialer{Timeout: 0, KeepAlive: 0}).Dial,
                TLSHandshakeTimeout: 5 * time.Second,
        }}

        numWorkers := *concurrencyPtr
        work := make(chan string)
        go func() {
            s := bufio.NewScanner(os.Stdin)
            for s.Scan() {
                work <- s.Text()
            }
            close(work)
        }()

        wg := &sync.WaitGroup{}

        for i := 0; i < numWorkers; i++ {
            wg.Add(1)
            go doWork(work, client, wg)
        }
        wg.Wait()
}

func doWork(work chan string, client *http.Client, wg *sync.WaitGroup) {
    defer wg.Done()
    for url := range work {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                fmt.Println(999, err)
                continue
        }
        req.Header.Set("Connection", "close")

        resp, err := client.Do(req)
        if err != nil {
                fmt.Println(999, err, url)
                continue
        }
        resp.Body.Close()
        fmt.Println(resp.StatusCode, url)
    }
}
