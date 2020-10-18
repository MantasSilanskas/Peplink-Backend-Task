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

	fmt.Println("Stop application by pressing Ctrl + C buttons at the same time.")

	resultMap := make(map[int]bool)      // šis map užtikrina, kad nebūtų rezultatų dublikatų
	rulesPrices := make(map[int]float64) // šis map leidžia, atsispausdinti atsakymą taisyklei, jei jos atsakymas jau buvo atspausdintas tačiau programos veikimo metų jos "Price" buvo pakeistas.

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ticker := time.NewTicker(30 * time.Second)
	done := make(chan bool)

	_, _, err := peplink.Parse(resultMap, rulesPrices)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case <-done:
				cancel()
				return
			case <-ticker.C:
				_, _, err := peplink.Parse(resultMap, rulesPrices)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
