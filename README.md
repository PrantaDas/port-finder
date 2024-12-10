# Port Finder

A fast and efficient **Port Finder** built with Go to identify open ports on a host. This project demonstrates Go's concurrency capabilities.

---

## Features

- **Concurrency**: Utilizes Goroutines and channels for fast scanning.
- **Customizable Scans**: Specify the target host, port range, and number of workers.
- **Performance**: Efficient use of resources, with a default timeout for port checks.
- **User-Friendly**: Clear output and input validation for a smooth experience.

---

## Getting Started

### Prerequisites

- Go 1.20 or later installed on your machine.
- Basic understanding of running Go programs.

### Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:PrantaDas/port-finder.git
   cd port-finder
   ```
2. Initialize go modules
   ```bash
   go mod tidy
   ```

### Usage

Run the Application

```bash
go run main.go -host <target-host> -start <start-port> -end <end-port> -workers <number-of-workers>
```

### Example

```bash
go run main.go -host example.com -start 1 -end 100 -workers 50
```

Output

```bash
Scanning example.com from ports 1 to 100 with 50 workers...

Scan Results:
Port 22 is open
Port 80 is open

Scan completed in 2.34 seconds.
```

This project is licensed under the MIT License. See the [LICENSE](https://github.com/PrantaDas/port-finder/blob/main/LICENSE) file for details.
