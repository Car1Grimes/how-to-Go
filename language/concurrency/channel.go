package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	eatLimit        = 3
	maxConcurrent   = 2
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id                int
	leftChopstick     *Chopstick
	rightChopstick    *Chopstick
	eatCount          int
	hostPermissionReq chan int
	hostDone          chan int
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for p.eatCount < eatLimit {
		p.hostPermissionReq <- p.id

		<-p.hostDone

		if rand.Intn(2) == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		p.eatCount++

		p.hostPermissionReq <- -p.id
	}
}

func host(permission chan int, done chan int) {
	currentlyEating := 0
	waitQueue := []int{}

	for {
		id := <-permission
		if id > 0 {
			if currentlyEating < maxConcurrent {
				currentlyEating++
				done <- id
			} else {
				waitQueue = append(waitQueue, id)
			}
		} else {
			currentlyEating--
			if len(waitQueue) > 0 && currentlyEating < maxConcurrent {
				nextID := waitQueue[0]
				waitQueue = waitQueue[1:]
				currentlyEating++
				done <- nextID
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}

	hostPermission := make(chan int)
	hostDone := make(chan int)

	var wg sync.WaitGroup

	go host(hostPermission, hostDone)

	for i := 0; i < numPhilosophers; i++ {
		p := &Philosopher{
			id:                i + 1,
			leftChopstick:     chopsticks[i],
			rightChopstick:    chopsticks[(i+1)%numPhilosophers],
			eatCount:          0,
			hostPermissionReq: hostPermission,
			hostDone:          hostDone,
		}
		wg.Add(1)
		go p.dine(&wg)
	}

	wg.Wait()
}
