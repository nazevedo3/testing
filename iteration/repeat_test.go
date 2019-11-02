package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a")
	expected := repeated

	if repeated != expected {
		t.Errorf("got %s expected %s", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
