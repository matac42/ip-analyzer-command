package main

import (
	"flag"

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
	analyzer.Analyze(duration, iface, networkIP)
}
