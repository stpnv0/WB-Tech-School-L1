package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt((other.x-p.x)*(other.x-p.x) + (other.y-p.y)*(other.y-p.y))
}

func main() {
	var x1, x2, y1, y2 float64
	fmt.Println("Введите координаты первой точки(x1, y1): ")
	fmt.Scan(&x1)
	fmt.Scan(&y1)

	fmt.Println("Введите координаты второй точки(x2, y2): ")
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	point1 := NewPoint(x1, y1)
	point2 := NewPoint(x2, y2)

	fmt.Printf("Расстояние между точками: %f\n", point1.Distance(point2))
}
