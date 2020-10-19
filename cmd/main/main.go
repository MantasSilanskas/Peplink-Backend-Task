package main

import (
	"Peplink-Backend-Task/pkg/peplink"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println("Stop application by pressing Ctrl + C buttons at the same time.")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	loop(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	cancel()
}

func loop(ctx context.Context) {

	resultMap := make(map[int]bool)      // šis map užtikrina, kad nebūtų rezultatų dublikatų
	rulesPrices := make(map[int]float64) // šis map leidžia, atsispausdinti atsakymą taisyklei, jei jos atsakymas jau buvo atspausdintas tačiau programos veikimo metų jos "Price" buvo pakeistas.

	ticker := time.NewTicker(2 * time.Second)

	// tickeri visada reikia sustabdyti
	defer ticker.Stop()

	loadData(resultMap, rulesPrices)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				loadData(resultMap, rulesPrices)
			}
		}
	}()
}

func loadData(rm map[int]bool, rp map[int]float64) {
	_, _, err := peplink.Parse(rm, rp)
	if err != nil {
		fmt.Println(err)
	}
}
