package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected: %q but got: %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 3)
	}
}

func TestClone(t *testing.T) {
	cloned := OwnClone("a")
	expected := "a"

	if cloned != expected {
		t.Errorf("expected: %q but got: %q", expected, cloned)
	}
}
