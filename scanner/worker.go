package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(host string, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range jobs {
		if scanPort(host, port) {
			results <- port
		}
	}
}

func scanPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
