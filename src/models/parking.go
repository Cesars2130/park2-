package models

import "sync"

type Parking struct {
	Subject
	spaces        []bool
	entranceMutex *sync.Mutex
	spacesChannel chan int
	totalSpaces   int
}

func NewParking(totalSpaces int) *Parking {
	return &Parking{
		Subject:       Subject{},
		spaces:        make([]bool, totalSpaces),
		entranceMutex: &sync.Mutex{},
		spacesChannel: make(chan int, totalSpaces),
		totalSpaces:   totalSpaces,
	}
}

func (p *Parking) OccupySpace(spaceIndex int) {
	if spaceIndex >= 0 && spaceIndex < len(p.spaces) {
		p.spaces[spaceIndex] = true
		p.NotifyObservers("CarEntered", spaceIndex)
	}
}

func (p *Parking) FreeSpace(spaceIndex int) {
	if spaceIndex >= 0 && spaceIndex < len(p.spaces) {
		p.spaces[spaceIndex] = false
		p.NotifyObservers("CarExited", spaceIndex)
	}
}

func (p *Parking) FindAvailableSpace() int {
	for i, occupied := range p.spaces {
		if !occupied {
			return i
		}
	}
	return -1
}

func (p *Parking) GetSpaces() chan int {
	return p.spacesChannel
}

func (p *Parking) GetEntrance() *sync.Mutex {
	return p.entranceMutex
}

func (p *Parking) NotifyObservers(event string, data interface{}) {
	for _, observer := range p.observers {
		observer.Update(event, data)
	}
}
