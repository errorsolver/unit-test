package repository

import (
	"go-unit-test/model"
	"math"
)

type CubeRepository interface {
	Volume() float64
	Area() float64
	Perimeter() float64
}

type cubeRepository struct {
	cube model.Cube
}

func (c *cubeRepository) Volume() float64 {
	return math.Pow(c.cube.Side, 3)
}

func (c *cubeRepository) Area() float64 {
	return math.Pow(c.cube.Side, 2) * 6
}

func (c *cubeRepository) Perimeter() float64 {
	return c.cube.Side * 12
}

func NewCubeRepository(cube model.Cube) CubeRepository {
	return &cubeRepository{cube: cube}
}
