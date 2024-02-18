package shapes

import "testing"

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{ 10, 12 }
		expectedArea := float64(120)
		receivedArea := ret.Area()

		if expectedArea != receivedArea {
			t.Errorf("The received area %f is different from the expected %f")
		}
	})

	t.Run("Circle", func(t t*testing.T){

	})
}
