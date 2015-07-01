//library for the health checker - ping tests and returns.

import (
	"fmt"
	"net"
	"container/list"
	
	"github.com/annanay25/go-fastping"
	"github.com/annanay25/go-DHT"
)

//to ping check and update the routing table:

p:= fastping.NewPinger()
//1. Ping the last peer in the routing table:

pingC := routing_table.last()
ra, err := net.ResolveIPAddr("ip4:icmp", JoinHostPort(pingC.ipaddress, pingC.port))
if err != nil {
	fmt.Println(err)
}
p.AddIPAddr(ra)
p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
	//update the routing_table and move the peer up/down depending on the RTT.
	fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	
}
p.OnIdle = func() {
	//delete the peer from the routing_table.
	routing_table.Remove(pingC)
	fmt.Println("finish")
}
err = p.Run()
if err != nil {
	fmt.Println(err)
}
