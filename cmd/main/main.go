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

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	var printed [4]bool
	fmt.Println(printed)
	fmt.Println(time.Now())

	printed, err := peplink.Parse(printed)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case <-done:
				continue
			case <-ticker.C:
				fmt.Println(time.Now())
				printed, err = peplink.Parse(printed)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()

}
