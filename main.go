package main

import (
	"flag"
	"fmt"
)

func main() {
	nodeID := flag.String("id", "node1", "Unique node ID")
	httpAddr := flag.String("http", ":8080", "HTTP server address")
	raftAddr := flag.String("raft", ":9001", "Raft communication address")
	dataDir := flag.String("data", "/tmp/raft3d", "Raft data directory")
	joinAddr := flag.String("join", "", "Join address if not bootstrapping")
	flag.Parse()
	fmt.Println(*nodeID, *httpAddr, *raftAddr, *dataDir, *joinAddr)
}
