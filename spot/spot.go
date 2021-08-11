package spot

import "errors"

type Spot struct {
	plate string
}

func New() *Spot {
	return new(Spot)
}

func (s *Spot) Park(plate string) error {
	if !s.IsAvailable() {
		return errors.New("parking spot occupied")
	} 
	s.plate = plate

	return nil
}

func (s *Spot) Unpark() (string, error) {
	plate := s.plate

	if s.IsAvailable() {
		return s.plate, errors.New("parking spot empty")
	}

	s.plate = ""

	return plate, nil
}

func (s *Spot) IsAvailable() bool {
	return s.plate == ""
}
