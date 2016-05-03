package pingpong

import (
	"net"
	"time"
	"github.com/tatsushid/go-fastping"
)

func Ping(address string) (rs int64) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", address)
	if err != nil {
		Log(err)
		rs = 0
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		rs = rtt.Nanoseconds()
	}

	err = p.Run()
	if err != nil {
		Log(err)
		rs = 0
	}
	return
}