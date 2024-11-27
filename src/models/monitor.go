package models

import "fmt"

type Monitor struct{}

func (m *Monitor) Update(event string, data interface{}) {
	switch event {
	case "CarEntered":
		fmt.Printf("Un auto ha ocupado el espacio: %d\n", data.(int))
	case "CarExited":
		fmt.Printf("Un auto ha liberado el espacio: %d\n", data.(int))
	}
}
