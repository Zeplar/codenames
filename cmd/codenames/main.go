package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/zeplar/codenames"
)

const defaultListenAddr = ":9091"

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Printf("[STARTUP] Listening on addr %s\n", defaultListenAddr)
	server := &codenames.Server{
		Server: http.Server{
			Addr: defaultListenAddr,
		},
	}
	if err := server.Start(nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
