package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"syscall"
	"time"
)

func ulimit() (uint64, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return 0, err
	}
	return rLimit.Cur, nil
}

func main() {
	maxFileDescriptors, err := ulimit()
	if err != nil {
		log.Fatal(err)
	}
	if maxFileDescriptors-100 < 0 {
		log.Fatalf("maxFileDescriptors==%d is not enough", maxFileDescriptors)
	}

	var wg sync.WaitGroup
	lock := make(chan struct{}, maxFileDescriptors-100)
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		Dial:                (&net.Dialer{Timeout: 0, KeepAlive: 0}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		wg.Add(1)
		lock <- struct{}{}
		go func(url string) {
			defer wg.Done()
			defer func() { <-lock }()

			req, err := http.NewRequest("GET", url, nil)
			req.Header.Set("Connection", "close")
			if err != nil {
				fmt.Println(999, err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(999, err, url)
				return
			}
			resp.Body.Close()
			fmt.Println(resp.StatusCode, url)
		}(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	wg.Wait()
}
