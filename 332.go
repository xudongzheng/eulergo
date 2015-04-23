package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y, z int
}

func cross(a, b point) (int, int, int) {
	return (a.y*b.z - b.y*a.z), (a.x*b.z - b.x*a.z), (a.x*b.y - b.x*a.y)
}

// Return true if three points are circular.
func cxtest(a, b, c point) bool {
	cx1x, cx1y, cx1z := cross(a, b)
	cx2x, cx2y, cx2z := cross(a, c)

	if cx1x*cx2y == cx2x*cx1y && cx1x*cx2z == cx2x*cx1z {
		return true
	}

	return false
}

func rad(f float64, a, b point) float64 {
	x1, y1, z1 := float64(a.x)/f, float64(a.y)/f, float64(a.z)/f
	x2, y2, z2 := float64(b.x)/f, float64(b.y)/f, float64(b.z)/f
	dproduct := x1*x2 + y1*y2 + z1*z2

	return math.Acos(dproduct)
}

func area(i int, a, b, c point) float64 {
	// Do cross product test to make sure the three points are not circular.
	if cxtest(a, b, c) {
		return math.MaxFloat64
	}

	f := float64(i)
	arc1 := rad(f, a, b)
	arc2 := rad(f, a, c)
	arc3 := rad(f, b, c)
	s := (arc1 + arc2 + arc3) / 2

	calc := 1.0
	calc *= math.Tan(s / 2)
	calc *= math.Tan((s - arc1) / 2)
	calc *= math.Tan((s - arc2) / 2)
	calc *= math.Tan((s - arc3) / 2)
	calc = math.Sqrt(calc)
	calc = math.Atan(calc)
	calc *= 4

	return f * f * calc
}

func handle(i int) float64 {
	target := i * i
	var points []point
	for a := 0; a <= i; a++ {
		for b := -i; b <= i; b++ {
			for c := -i; c <= i; c++ {
				sum := a*a + b*b + c*c
				if sum == target {
					points = append(points, point{a, b, c})
				}
			}
		}
	}

	count := len(points)
	min := math.MaxFloat64
	for a := 0; a < count; a++ {
		for b := a + 1; b < count; b++ {
			for c := b + 1; c < count; c++ {
				t := area(i, points[a], points[b], points[c])
				if t < min {
					min = t
				}
			}
		}
	}

	return min
}

func main() {
	var sum float64
	for i := 1; i <= 50; i++ {
		sum += handle(i)
	}
	fmt.Printf("%.6f\n", sum)
}
