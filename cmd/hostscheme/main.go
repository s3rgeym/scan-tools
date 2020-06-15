// Утилита для определения протокола хоста
package main

import (
	"flag"
	"fmt"
	"http-tools/internal/utils"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var schemes = map[string]string{
	"https": "443",
	"http":  "80",
}

type cmdFlags struct {
	In          string
	Out         string
	Concurrency int
	Timeout     int
	ShowHelp    bool
}

func setupFlags(flags *cmdFlags) {
	flag.StringVar(&flags.In, "i", "-", "Input filename")
	flag.StringVar(&flags.Out, "o", "-", "Output filename")
	flag.IntVar(&flags.Concurrency, "c", 10, "Concurrent requests")
	flag.IntVar(&flags.Timeout, "t", 300, "Connect timeout in milliseconds")
	flag.BoolVar(&flags.ShowHelp, "h", false, "Show help and exit")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage %s [options...]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flags := cmdFlags{}
	setupFlags(&flags)
	flag.Parse()
	if flags.ShowHelp {
		flag.Usage()
		os.Exit(1)
	}
	in := os.Stdin
	if flags.In != "-" {
		path, err := utils.ExpandPath(flags.In)
		if err != nil {
			log.Panic(err)
		}
		file, err := os.Open(path)
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		in = file
	}
	out := os.Stdout
	if flags.Out != "-" {
		path, err := utils.ExpandPath(flags.Out)
		if err != nil {
			log.Panic(err)
		}
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		out = file
	}
	hosts, err := utils.ReadLines(in)
	if err != nil {
		log.Panic(err)
	}
	hostsLen := len(hosts)
	numberOfWorkers := utils.Min(flags.Concurrency, hostsLen)
	jobs := make(chan string, numberOfWorkers)
	go func() {
		for _, host := range hosts {
			jobs <- host
		}
		close(jobs)
	}()
	var wg sync.WaitGroup
	wg.Add(hostsLen)
	connectTimeout := time.Duration(flags.Timeout) * time.Millisecond
	for i := 0; i < numberOfWorkers; i++ {
		go worker(jobs, &wg, out, connectTimeout)
	}
	wg.Wait()
}

func worker(
	jobs <-chan string,
	wg *sync.WaitGroup,
	out *os.File,
	connectTimeout time.Duration,
) {
	for host := range jobs {
		for scheme, port := range schemes {
			conn, _ := net.DialTimeout(
				"tcp",
				net.JoinHostPort(host, port),
				connectTimeout,
			)
			if conn != nil {
				conn.Close()
				hostname := fmt.Sprintf("%s://%s", scheme, host)
				fmt.Fprintln(out, hostname)
				break
			}
		}
		wg.Done()
	}
}
