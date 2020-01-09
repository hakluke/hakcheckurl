package main

import (
	"fmt"
	"net/http"
	"sync"
	"bufio"
	"os"
)

func worker(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil{
		return
	}
	fmt.Println(resp.StatusCode, url)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	for scanner.Scan() {
		go worker(scanner.Text(), &wg)
		wg.Add(1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	wg.Wait()
}
