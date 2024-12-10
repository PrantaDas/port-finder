package scanner

import (
	"sync"
)

func Scan(host string, startPort, endPort, workers int) []int {
	portJobs := make(chan int, workers)
	results := make(chan int)
	var openPorts []int
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(host, portJobs, results, &wg)
	}

	go func() {
		for port := range results {
			openPorts = append(openPorts, port)
		}
	}()

	for port := startPort; port <= endPort; port++ {
		portJobs <- port
	}
	close(portJobs)

	wg.Wait()
	close(results)

	return openPorts
}
