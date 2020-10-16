package main

import (
	"context"
	"fmt"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg/peplink"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println("Stop application by pressing ctrl and C buttons at the same time.")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ticker := time.NewTicker(30 * time.Second)
	done := make(chan bool)

	fmt.Println(time.Now())
	peplink.Parse()

	go func() {
		for {
			select {
			case <-done:
				continue
			case <-ticker.C:
				fmt.Println(time.Now())
				peplink.Parse()
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()

}
