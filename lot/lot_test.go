package lot

import "testing"

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
}
