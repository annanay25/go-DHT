package go_DHT

import (
	"fmt"
	"net"
	"container/list"
	
	"github.com/annanay25/go-DHT"
	)

type contact struct {
	node NodeID
	ipaddress string
	port string
}

//make routing table a list type with elements- InfoHash -> contact.
routing_table := list.New()


//once you know there's a new peer that wants to connect with the DHT do the 'pinger' check in the appropriate bucket. And then add this guy.
//you'll find out from the httpdaemon.
