package models

import (
	"main/src/scenes"
	"sync"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"golang.org/x/exp/rand"
)

type Car struct {
	id            int
	timeInParking time.Duration
	image         *canvas.Image
	parkLot       int
}

func NewCar(id int) *Car {
	rand.Seed(uint64(time.Now().UnixNano()))
	timeInParking := time.Duration(rand.Intn(3)+15) * time.Second

	images := []string{
		"assets/carro2.png",
	}
	selectedImage := images[rand.Intn(len(images))]

	return &Car{
		id:            id,
		timeInParking: timeInParking,
		image:         canvas.NewImageFromURI(storage.NewFileURI(selectedImage)),
	}
}

func (c *Car) GetCarImage() *canvas.Image {
	return c.image
}

func (c *Car) TryPark(p *Parking, wg *sync.WaitGroup, s *scenes.Scene, parkingSpaces map[int][2]float32) {
	c.JoinPark(p, parkingSpaces)
	time.Sleep(c.timeInParking)
	c.LeavePark(p, s)
	wg.Done()
}

func (c *Car) JoinPark(p *Parking, parkingSpaces map[int][2]float32) {
	p.GetSpaces() <- c.id
	c.parkLot = p.FindAvailableSpace()
	p.OccupySpace(c.parkLot)

	// Notificar al observador (Scene) que el carro se movió
	pos := parkingSpaces[c.parkLot]
	p.NotifyObservers("CarMoved", map[string]interface{}{
		"image":    c.image,
		"position": [2]float32{pos[0], pos[1]},
	})
}

func (c *Car) LeavePark(p *Parking, s *scenes.Scene) {
	p.GetEntrance().Lock()
	<-p.GetSpaces()

	p.FreeSpace(c.parkLot)
	p.NotifyObservers("CarRemoved", c.image) // Notificar que el carro salió

	p.GetEntrance().Unlock()
}
