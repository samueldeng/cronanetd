package main

import (
	"fmt"
	"net"
	"time"

	ping "github.com/digineo/go-ping"
)

func main() {
	destination := "192.168.31.1"

	remoteAddr, err := net.ResolveIPAddr("ip4", destination)
	if err != nil {
		panic(err)
		return
	}

	pinger, err := ping.New("0.0.0.0", "")
	if err != nil {
		panic(err)
		return
	}
	defer pinger.Close()

	if pinger.PayloadSize() != uint16(56) {
		pinger.SetPayloadSize(uint16(56))
	}

	if remoteAddr.IP.IsLinkLocalMulticast() {
		return
	}

	rtt, err := pinger.PingAttempts(remoteAddr, time.Second, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("ping %s (%s) rtt=%v\n", destination, remoteAddr, rtt)
}
