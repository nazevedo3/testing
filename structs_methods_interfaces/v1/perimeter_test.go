package perimeter

import "testing"

var expectedOut = func(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}

}

func TestPerimeter(t *testing.T) {

	got := Perimeter(10.0, 10.0)
	want := 40.0

	expectedOut(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	expectedOut(t, got, want)
}
