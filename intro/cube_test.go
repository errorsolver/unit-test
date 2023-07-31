package intro

import "testing"

// Positive test
func TestCube_Volume_Success(t *testing.T) {
	cubeMockup := Cube{Side: 2}
	expected := float64(8)

	actual, err := cubeMockup.Volume()
	if expected != actual {
		t.Errorf("Result should be %2.f, but return %2.f", expected, actual)
	}

	if err != nil {
		t.Error("Volume result error: ", err)
	}
}

// Negative Test

func TestCube_Volume_Failed(t *testing.T) {
	cubeMockup := Cube{Side: -1}
	expected := float64(0.0)

	actual, err := cubeMockup.Volume()

	if err == nil {
		t.Error("Should return error when value is negetive")
	}

	if actual != expected {
		t.Errorf("expected %2.f, but return %2.f", expected, actual)
	}
}
