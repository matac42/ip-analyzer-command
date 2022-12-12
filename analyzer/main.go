package analyzer

import (
	"fmt"
	"log"
	"net"
	"net/netip"
	"regexp"
	"time"

	"github.com/matac42/ip-analyzer/address"
	"github.com/mdlayher/arp"
)

func Analyze(timeout int, iface string, networkIP string) {
	duration := time.Duration(timeout) * time.Microsecond

	// Request hardware address for IP address
	re := regexp.MustCompile(`\/\d*`)
	networkAddress := re.ReplaceAllString(networkIP, "")
	fmt.Println("networkAddress: ", networkAddress)

	re = regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\/`)
	prefix := re.ReplaceAllString(networkIP, "")
	fmt.Println("prefix: ", prefix)

	ip, err := netip.ParseAddr(networkAddress)
	if err != nil {
		log.Fatal(err)
	}

	broadcast, err := netip.ParseAddr(address.AddrArray2String(address.CalcBroadCastAddr(address.AddrString2AddrArray(networkIP), address.AddrString2AddrArray("255.255.255.0"))))
	fmt.Println("broadcast: ", broadcast)
	if err != nil {
		log.Fatal(err)
	}
	ip = ip.Next()
	for ip.Less(broadcast) {
		ifi, err := net.InterfaceByName(iface)
		if err != nil {
			log.Fatal(err)
		}

		c, err := arp.Dial(ifi)
		if err != nil {
			log.Fatal(err)
		}

		if err := c.SetDeadline(time.Now().Add(duration)); err != nil {
			log.Fatal(err)
		}

		mac, err := c.Resolve(ip)
		if err != nil {
			// fmt.Println("error: ", err)
		} else {
			fmt.Printf("%s -> %s\n", ip, mac)
		}
		err = c.Close()
		if err != nil {
			log.Fatal(err)
		}
		ip = ip.Next()
	}
}
