package lot

import (
	"fmt"
	"testing"
)

func TestLot(t *testing.T) {
	t.Run("#Park", func(t *testing.T) {
		t.Run("return error if no available spot", func(t *testing.T) {
			l := New(0)
			want := "parking lot full"

			err := l.Park("abc123")

			if err.Error() != want {
				t.Errorf("want %v, got %v", want, err.Error())
			}
		})
	})

	t.Run("#New", func(t *testing.T) {
		t.Run("provide parking spots", func(t *testing.T) {
			want := "parking lot full"
			
			l := New(1)

			err1 := l.Park("abc123")
			err2 := l.Park("def456")
			if err1 != nil {
				t.Errorf("want %v, got %v", nil, err1)
			}
			if err2 != nil && err2.Error() != want {
				t.Errorf("want %v, got %v", want, err2.Error())
			}
		})
	})

	t.Run("#Unpark", func(t *testing.T) {
		t.Run("return error if car not found", func(t *testing.T) {
			l := New(1)
			want := "car not found"

			_, err := l.Unpark("abc")
			
			if err == nil {
				t.Errorf("want %v, got %v", want, nil)
			}
		})

		t.Run("return plate if car found", func(t *testing.T) {
			l := New(1)
			plate := "abc"
			l.Park(plate)

			got, err := l.Unpark(plate)
			
			if got != plate {
				t.Errorf("want %v, got %v", plate, got)
			}
			if err != nil {
				t.Errorf("want %v, got %v", nil, err.Error())
			}
		})

		t.Run("release occupied spot", func(t *testing.T) {
			l := New(1)
			plate := "abc"
			
			l.Park(plate)
			l.Unpark(plate)
			err := l.Park(plate)
			
			if err != nil {
				t.Errorf("want %v, got %v", nil, err.Error())
			}
		})
	})

	t.Run("#Status", func(t *testing.T) {
		t.Run("return current parking lot status", func(t *testing.T) {
			l := New(2)
			plate:= "abc"
			l.Park(plate)
			want := fmt.Sprintf("1. available: false, plate number: %v\n2. available: true, plate number: \n", plate)
			
			got := l.Status()

			if got != want {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	})
}
