package main

import (
	"flag"
	"fmt"
	"log"
	"port-finder/scanner"
	"time"
)

func main() {
	host := flag.String("host", "localhost", "Host to scan")
	startPort := flag.Int("start", 1, "Start of port range")
	endPort := flag.Int("end", 1024, "End of port range")
	workers := flag.Int("workers", 100, "Number of concurrent workers")
	flag.Parse()

	if *startPort > *endPort || *startPort < 1 || *endPort > 65535 {
		log.Fatalf("Invalid port range: %d-%d", *startPort, *endPort)
	}

	fmt.Printf("Scanning %s from ports %d to %d with %d workers...\n", *host, *startPort, *endPort, *workers)

	startTime := time.Now()

	results := scanner.Scan(*host, *startPort, *endPort, *workers)

	fmt.Println("\nScan Results:")
	for _, port := range results {
		fmt.Printf("Port %d is open\n", port)
	}

	fmt.Printf("\nScan completed in %s\n", time.Since(startTime))
}
