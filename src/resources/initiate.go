package resources

import (
	"main/src/models"
	"main/src/scenes"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"golang.org/x/exp/rand"
)

func CreateApp() {
	myApp := app.New()
	stage := myApp.NewWindow("Parking Simulator")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(1280, 720))
	stage.SetFixedSize(true)

	scene := scenes.NewScene(stage)
	scene.Init()

	go Run(scene)
	stage.ShowAndRun()
}

func Run(s *scenes.Scene) {
	var wg sync.WaitGroup
	const carsNumber = 100

	parking := models.NewParking(20)
	monitor := &models.Monitor{}
	parking.AddObserver(monitor)
	parking.AddObserver(s) // Registrar el Scene como observador

	// Obtener las posiciones de los carriles desde la escena
	parkingSpaces := s.DrawLanes()

	for i := 0; i < carsNumber; i++ {
		wg.Add(1)
		go func(id int) {
			car := models.NewCar(id)
			carImage := car.GetCarImage()
			carImage.Resize(fyne.NewSize(60, 60))

			s.AddImage(carImage, 425, 450) // Posición inicial
			s.GetContainer().Refresh()

			car.TryPark(parking, &wg, s, parkingSpaces)
		}(i)
		time.Sleep((time.Duration(rand.Intn(2)+1) * time.Second))
	}

	wg.Wait()
}
