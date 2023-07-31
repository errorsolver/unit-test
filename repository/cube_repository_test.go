package repository

import (
	"go-unit-test/model"
	"math"
	"testing"
)

func TestNewCubeRepository(t *testing.T) {
	t.Run("It should return cube repository", func(t *testing.T) {
		cubeModelMockup := model.Cube{Side: 2}
		newRepoCube := NewCubeRepository(cubeModelMockup)

		if newRepoCube == nil {
			t.Error("Cannot be null")
		}
	})
}

func TestCubeRepository_Volume_Success(t *testing.T) {
	t.Run("It should return volume", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 2,
		}
		expected := float64(math.Pow(cubeModelMockup.Side, 3))
		cubeRepo := NewCubeRepository(cubeModelMockup)
		actual := cubeRepo.Volume()

		if actual != expected {
			t.Errorf("Volume result expected %2.f, but %2.f", expected, actual)
		}
	})
}

func TestCubeRepository_Area_Success(t *testing.T) {
	t.Run("It should return area", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 4,
		}
		expected := float64(math.Pow(cubeModelMockup.Side, 2) * 6)
		cubeRepo := NewCubeRepository(cubeModelMockup)
		actual := cubeRepo.Area()

		if actual != expected {
			t.Errorf("Area result %2.f, expected %2.f", actual, expected)
		}
	})
}

func TestCubeRepository_Perimeter_Success(t *testing.T) {
	t.Run("It should return perimeter", func(t *testing.T) {
		cubeModelMockup := model.Cube{
			Side: 4,
		}
		expected := float64(cubeModelMockup.Side * 12)
		cubeRepo := NewCubeRepository(cubeModelMockup)
		actual := cubeRepo.Perimeter()

		if actual != expected {
			t.Errorf("Perimeter result %2.f, expected %2.f", actual, expected)
		}
	})
}
