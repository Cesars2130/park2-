package scenes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Scene struct {
	scene     fyne.Window
	container *fyne.Container
}

func NewScene(scene fyne.Window) *Scene {
	return &Scene{scene: scene, container: nil}
}

func (s *Scene) Init() {
	// Cargar el fondo
	bgImage := canvas.NewImageFromFile("assets/Parking.png")
	bgImage.FillMode = canvas.ImageFillOriginal
	bgImage.Resize(fyne.NewSize(1280, 720)) // Ajustar al tamaño de la ventana

	// Contenedor principal
	s.container = container.NewWithoutLayout(bgImage)
	s.scene.SetContent(s.container)

	// Dibujar carriles
	s.DrawLanes()
}

func (s *Scene) DrawLanes() map[int][2]float32 {
	const (
		startX       = 300 // Coordenada X inicial para ambos carriles
		startYUp     = 250 // Coordenada Y inicial para carriles hacia arriba
		startYDown   = 350 // Coordenada Y inicial para carriles hacia abajo (menos separación)
		carrilWidth  = 60  // Ancho de cada carril
		carrilHeight = 100 // Altura de cada carril
		spacing      = 5   // Espaciado entre carriles horizontales
	)

	positions := make(map[int][2]float32)

	for i := 0; i < 10; i++ {
		carril := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
		carril.Resize(fyne.NewSize(carrilWidth, carrilHeight))
		x := float32(startX) + float32(i)*(carrilWidth+spacing)
		y := float32(startYUp)
		carril.Move(fyne.NewPos(x, y))
		s.container.Add(carril)
		positions[i] = [2]float32{x, y}
	}

	for i := 10; i < 20; i++ {
		carril := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
		carril.Resize(fyne.NewSize(carrilWidth, carrilHeight))
		x := float32(startX) + float32(i-10)*(carrilWidth+spacing)
		y := float32(startYDown)
		carril.Move(fyne.NewPos(x, y))
		s.container.Add(carril)
		positions[i] = [2]float32{x, y}
	}

	return positions
}

func (s *Scene) AddImage(image *canvas.Image, posX, posY float32) {
	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(60, 60))
	image.Move(fyne.NewPos(posX, posY))
	s.container.Add(image)
	s.container.Refresh()
}

func (s *Scene) DeleteImage(image *canvas.Image) {
	s.container.Remove(image)
	s.container.Refresh()
}

func (s *Scene) GetContainer() *fyne.Container {
	return s.container
}

func (s *Scene) Update(event string, data interface{}) {
	switch event {
	case "CarMoved":
		// Mover un carro a su nueva posición
		info := data.(map[string]interface{})
		carImage := info["image"].(*canvas.Image)
		pos := info["position"].([2]float32)
		carImage.Move(fyne.NewPos(pos[0], pos[1]))
		s.container.Refresh()
	case "CarRemoved":
		// Eliminar la imagen de un carro
		carImage := data.(*canvas.Image)
		s.DeleteImage(carImage)
	}
}
