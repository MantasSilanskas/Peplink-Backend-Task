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

	fmt.Println("Stop application by pressing ctrl + C buttons at the same time.")

	resultMap := make(map[int]bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ticker := time.NewTicker(30 * time.Second)
	done := make(chan bool)

	peplink.Parse(resultMap)

	go func() {
		for {
			select {
			case <-done:
				continue
			case <-ticker.C:
				peplink.Parse(resultMap)
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()

}
