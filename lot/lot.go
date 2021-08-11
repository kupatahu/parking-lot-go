package lot

import (
	"errors"
	"github.com/kupatahu/parking-lot-go/spot"
)

type Lot struct {
	spots []*spot.Spot
}

func New(size int) *Lot {
	l := new(Lot)

	for i := 0; i < size; i++ {
		l.spots = append(l.spots, spot.New())
	}

	return l
}

func (l *Lot) Park(plate string) error {
	var found bool

	for _, spot := range l.spots {
		if spot.IsAvailable() {
			spot.Park(plate)
			found = true
			break
		}
	}

	if found {
		return nil
	}

	return errors.New("parking lot full")
}
