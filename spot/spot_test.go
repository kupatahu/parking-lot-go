package spot

import (
	"fmt"
	"testing"
)

func TestSpot(t *testing.T) {
	t.Run("#Unpark", func(t *testing.T) {
		t.Run("return plate number of parked car", func(t *testing.T) {
			s := New()
			plate := "abc123"
			s.Park(plate)

			got, _ := s.Unpark()

			if got != plate {
				t.Errorf("want %v, got %v", plate, got)
				t.Fail()
			}
		})
		
		t.Run("return error if no parked car", func(t *testing.T) {
			s := New()
			want := "parking spot empty"

			_, err := s.Unpark()

			if err.Error() != want {
				t.Errorf("want %v, got %v", want, err.Error())
				t.Fail()
			}
		})
	})
	
	t.Run("#Park", func(t *testing.T) {
		t.Run("return error if space occupied", func(t *testing.T) {
			s := New()
			want := "parking spot occupied"
			s.Park("abc123")

			err := s.Park("abc123")

			if err.Error() != want {
				t.Errorf("want %v, got %v", want, err.Error())
				t.Fail()
			}
		})
	})
	
	t.Run("#IsAvailable", func(t *testing.T) {
		t.Run("return true if no parked car", func(t *testing.T) {
			s := New()

			got := s.IsAvailable()

			if got != true {
				t.Errorf("want %v, got %v", true, got)
				t.Fail()
			}
		})
		
		t.Run("return false if occupied", func(t *testing.T) {
			s := New()
			s.Park("abc123")

			got := s.IsAvailable()

			if got != false {
				t.Errorf("want %v, got %v", false, got)
				t.Fail()
			}
		})
	})
	
	t.Run("#Status", func(t *testing.T) {
		t.Run("return available true if no parked car", func(t *testing.T) {
			s := New()
			want := "available: true, plate number: "

			got := s.Status()

			if got != want {
				t.Errorf("want %v, got %v", want, got)
				t.Fail()
			}
		})

		t.Run("return available false with plate number if occupied", func(t *testing.T) {
			s := New()
			plate := "abc"
			s.Park(plate)
			want := fmt.Sprintf("available: false, plate number: %v", plate)

			got := s.Status()

			if got != want {
				t.Errorf("want %v, got %v", want, got)
				t.Fail()
			}
		})
	})
	
	t.Run("#PlateEquals", func(t *testing.T) {
		t.Run("return false if parked car plate not equals to given plate", func(t *testing.T) {
			s := New()
			want := false

			got := s.PlateEquals("abc")

			if got != want {
				t.Errorf("want %v, got %v", want, got)
				t.Fail()
			}
		})

		t.Run("return true if parked car plate equals to given plate", func(t *testing.T) {
			s := New()
			plate := "abc"
			s.Park(plate)
			want := true

			got := s.PlateEquals("abc")

			if got != want {
				t.Errorf("want %v, got %v", want, got)
				t.Fail()
			}
		})
	})
}
