package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

var s = gocron.NewScheduler(time.UTC)
var count = 0
var count1 = 0
var wg sync.WaitGroup

func task() {
	count += 1
	fmt.Println("||==LOL=========||", count)
	if count == 8 {
		wg.Done()
		s.RemoveByTag("tag1")
	}
}

func task1() {
	count1 += 1
	fmt.Println("||========ROFL==||", count1)
	if count1 == 8 {
		wg.Done()
		s.RemoveByTag("tag2")
	}
}

func task2() {
	fmt.Println("OMG! Finally, done. ( ͡° ͜ʖ ͡°)")
}

func main() {
	s.Every(2).Seconds().Tag("tag1").Do(task)
	s.Every(2).Seconds().Tag("tag2").Do(task1)
	wg.Add(2)
	go s.StartBlocking()
	wg.Wait()
	task2()
}
