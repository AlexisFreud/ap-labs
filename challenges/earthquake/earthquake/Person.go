package main

import (
	"fmt"
	"sync"
	"time"
)

// Persona is defined
type Persona struct {
	speed   float32
	id      int
	isDeath bool
	pos     *Node
}

func (p *Persona) setPos(MyNode *Node) {
	p.pos = MyNode
	p.pos.mux.Lock()
	//fmt.Printf("Person %d is on node \n", p.id)
}

func (p *Persona) walk(wg *sync.WaitGroup, exit chan int) {
	fmt.Printf("Person %d starts walking\n", p.id)
	for {
		select {
		case <-exit:
			fmt.Printf("Person %d has died\n", p.id)
			defer wg.Done()
			return

		default:
			if p.pos.isExit != true {
				//fmt.Printf("Person %d is in node\n", p.id)
				//fmt.Printf("Person %d has started to move to %d\n", p.id, p.pos.nextHop.name)

				p.pos.nextHop.mux.Lock() // move right foot
				time.Sleep(time.Duration(p.speed) * time.Second)
				p.pos.mux.Unlock() // set free the origin MyNode
				p.pos = p.pos.nextHop
				//fmt.Printf("Person %d just moved to MyNode %d\n", p.id, p.pos.name)
			} else {
				fmt.Printf("Person %d is out of danger\n", p.id)
				p.pos.mux.Unlock()
				p.isDeath = false
				defer wg.Done()
				return
			}
		}
	}
}
