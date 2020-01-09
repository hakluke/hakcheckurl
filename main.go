package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		log.Fatal(err)
	}
	maxFileDescriptors, err := strconv.Atoi(strings.TrimSuffix(string(out), "\n"))
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	lock := make(chan struct{}, maxFileDescriptors)
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		Dial:                (&net.Dialer{Timeout: 0, KeepAlive: 0}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}}

	for scanner.Scan() {
		wg.Add(1)
		lock <- struct{}{}
		go func(url string) {
			defer wg.Done()
			defer func() { <-lock }()

			req, err := http.NewRequest("GET", url, nil)
			req.Header.Set("Connection", "close")
			if err != nil {
				fmt.Println(-1, err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(-1, err, url)
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
