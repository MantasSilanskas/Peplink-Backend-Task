package main

import (
	"context"
	"fmt"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg"
	"os"
	"os/signal"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ticker := time.NewTicker(30 * time.Second)
	done := make(chan bool)

	go func() {
		for {

			select {
			case <-done:
				continue
			case <-ticker.C:
				fmt.Println(time.Now())
				pkg.Parse()
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()

}
