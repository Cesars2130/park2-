package models

type Observer interface {
	Update(event string, data interface{})
}

type Subject struct {
	observers []Observer
}

func (s *Subject) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) RemoveObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) NotifyObservers(event string, data interface{}) {
	for _, observer := range s.observers {
		observer.Update(event, data)
	}
}
