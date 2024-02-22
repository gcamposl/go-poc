package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := rect.Area()

		if expectedArea != receivedArea {
			t.Fatalf("The received area %f is different from the expected %f", receivedArea, expectedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if expectedArea != receivedArea {
			t.Fatalf("The received area %f is different from the expected %f", receivedArea, expectedArea)
		}
	})
}
