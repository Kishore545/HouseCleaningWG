//With Goroutine channel
//-------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

type Area struct {
	name string
}

func cleanArea(area Area, wg *sync.WaitGroup, progress chan string) {
	defer wg.Done()
	progress <- fmt.Sprintf("Starting to clean %s...", area.name)
	time.Sleep(time.Duration(2) * time.Second) // Simulating cleaning time
	progress <- fmt.Sprintf("%s is now clean and shining!", area.name)
}

func main() {
	areas := []Area{{"Room 1"}, {"Room 2"}, {"Hall"}}
	var wg sync.WaitGroup
	progress := make(chan string, len(areas))

	for _, area := range areas {
		wg.Add(1)
		go cleanArea(area, &wg, progress)
	}

	// Close the channel once all cleaning tasks are done
	go func() {
		wg.Wait()
		close(progress)
	}()

	// Display progress messages
	for msg := range progress {
		fmt.Println(msg)
	}

	fmt.Println("The entire house is now clean and shining!")
}
