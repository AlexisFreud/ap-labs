package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// PersonManager this struch handle the instatiation and walking
type PersonManager struct {
	wg      *sync.WaitGroup
	mapa    *Map
	people  []Persona
	seconds int
	exit    chan int
}

func (pm *PersonManager) initPeople() {
	for index := 0; index < len(pm.people); index++ {
		pm.people[index] = Persona{float32(index / 2), index, true, nil}
	}
}

func (pm *PersonManager) setPeoplePos() {
	//var rows =  make([]int, pm.mapa.height)
	//var cols = make([]int, pm.mapa.height)

	availableNode := pm.mapa.getRealNodes()

	//fmt.Println(len(availableNode))

	for index := 0; index < len(pm.people); index++ {

		rand.Seed(time.Now().UnixNano())
		pos := rand.Intn(len(availableNode)-0) + 0

		pm.people[index].setPos(availableNode[pos])

		// Remove the element at index i from availableNode.
		availableNode[pos] = availableNode[len(availableNode)-1] // Copy last element to index i.
		availableNode[len(availableNode)-1] = nil                // Erase last element (write zero value).
		availableNode = availableNode[:len(availableNode)-1]     // Truncate slice.
	}
}

func (pm *PersonManager) startTimer(timeToFinsh int) {
	pm.wg.Add(1)
	timer := time.NewTimer(time.Duration(timeToFinsh) * time.Second)
	//fmt.Println("Timer hasnt' done")
	<-timer.C
	fmt.Println("Timer done")
	close(pm.exit)
	defer pm.wg.Done()
}

func (pm *PersonManager) startWalking() {
	for index := 0; index < len(pm.people); index++ {
		go pm.people[index].walk(pm.wg, pm.exit)
		pm.wg.Add(1)
	}
}

func (pm *PersonManager) runSimulation() {
	pm.initPeople()
	pm.setPeoplePos()
	go pm.startTimer(pm.seconds)
	pm.startWalking()
	pm.wg.Wait()
}

func newPersonManager(nPeople int, seconds int, mapa *Map) *PersonManager {

	if len(mapa.getRealNodes()) < nPeople {
		return nil
	}
	return &PersonManager{
		wg:      new(sync.WaitGroup),
		mapa:    mapa,
		seconds: seconds,
		people:  make([]Persona, nPeople),
		exit:    make(chan int),
	}

}

func (pm *PersonManager) damageReport() {
	deadPeople := 0
	survivors := 0

	for index := 0; index < len(pm.people); index++ {
		if pm.people[index].isDeath {
			deadPeople++
		}
	}

	survivors = len(pm.people) - deadPeople

	fmt.Printf("\n\nOn simulation with %d people and a timeout of %d\n", len(pm.people), pm.seconds)
	fmt.Printf("%d people have survided and %d have died\n", deadPeople, survivors)

}
