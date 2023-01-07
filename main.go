package main

import (
	"flag"
	"log"
	"net"

	"github.com/anthdm/ggcahce/cache"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("SET Foo Bar 40000"))
	if err != nil {
		log.Fatal(err)
	}

	select {}

	return
	var (
		listenAddr = flag.String("listenaddr", ":3000", "listen address of the server")
		leaderAddr = flag.String("leaderaddr", "", "listen address of the leader")
	)
	flag.Parse()

	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   len(*leaderAddr) == 0,
		LeaderAddr: *leaderAddr,
	}

	server := NewServer(opts, cache.New())
	server.Start()
}
