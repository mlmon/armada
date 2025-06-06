package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: exmembers <port> [join_address:port]")
		fmt.Println("Example: exmembers 7946")
		fmt.Println("Example: exmembers 7947 127.0.0.1:7946")
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	config := memberlist.DefaultLocalConfig()
	config.Name = fmt.Sprintf("node-%d", port)
	config.BindPort = port
	config.AdvertisePort = port

	list, err := memberlist.Create(config)
	if err != nil {
		log.Fatalf("Failed to create memberlist: %v", err)
	}

	if len(os.Args) > 2 {
		joinAddr := os.Args[2]
		fmt.Printf("Attempting to join cluster at %s\n", joinAddr)
		_, err := list.Join([]string{joinAddr})
		if err != nil {
			log.Fatalf("Failed to join cluster: %v", err)
		}
		fmt.Printf("Successfully joined cluster\n")
	}

	fmt.Printf("Node %s started on port %d\n", config.Name, port)
	fmt.Printf("Local node: %s\n", list.LocalNode().Address())

	go func() {
		for {
			time.Sleep(5 * time.Second)
			members := list.Members()
			fmt.Printf("\n--- Cluster Members (%d) ---\n", len(members))
			for _, member := range members {
				fmt.Printf("  %s (%s:%d)\n", member.Name, member.Addr, member.Port)
			}
			fmt.Println("---")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Printf("\nShutting down node %s\n", config.Name)
	if err := list.Shutdown(); err != nil {
		log.Printf("Error shutting down: %v", err)
	}
}