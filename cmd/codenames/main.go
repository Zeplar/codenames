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
const expiryDur = -24 * time.Hour

func main() {
	rand.Seed(time.Now().UnixNano())

	var listenAddr string
	log.Printf("[STARTUP] Listening on addr %s\n", listenAddr)
	server := &codenames.Server{
		Server: http.Server{
			Addr: listenAddr,
		},
	}
	if err := server.Start(nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
