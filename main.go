package main

import (
	"flag"
	"fmt"

	"github.com/matac42/ip-analyzer/analyzer"
)

var (
	duration  int
	iface     string
	networkIP string
)

func main() {
	flag.IntVar(&duration, "t", 100, "timeout in milliseconds")
	flag.StringVar(&iface, "i", "eth0", "interface")
	flag.StringVar(&networkIP, "n", "", "network IP address")

	flag.Parse()
	fmt.Println(duration, iface, networkIP)
	analyzer.Analyze(duration, iface, networkIP)
}
