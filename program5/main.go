package main

import (
	"errors"
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Derivatives Programming Assignment
// Input: points ([[x, y], [x, y]])
// Output: derivatives, error
func derivatives(points []Point) ([]float64, error) {
	if len(points) < 5 {
		return nil, errors.New("at least 5 points are required")
	}

	dx := points[1].X - points[0].X
	if dx == 0 {
		return nil, errors.New("x-coordinates must be equally spaced")
	}

	derivatives := make([]float64, len(points))

	// For endpoints (first and last points), use two-point formula
	derivatives[0] = (points[1].Y - points[0].Y) / dx
	derivatives[len(points)-1] = (points[len(points)-1].Y - points[len(points)-2].Y) / dx

	// For second-from-left point, use three-point midpoint formula
	if len(points) >= 3 {
		derivatives[1] = (-3*points[0].Y + 4*points[1].Y - points[2].Y) / (2 * dx)
	}

	// For second-from-right point, use three-point midpoint formula
	if len(points) >= 3 {
		derivatives[len(points)-2] = (points[len(points)-1].Y - points[len(points)-3].Y) / (2 * dx)
	}

	// For interior points, use five-point midpoint formula
	for i := 2; i < len(points)-2; i++ {
		derivatives[i] = (-points[i+2].Y + 8*points[i+1].Y - 8*points[i-1].Y + points[i-2].Y) / (12 * dx)
	}

	return derivatives, nil
}

func main() {
	// Linear function y = 2x + 1
	linearPoints := []Point{
		{X: 0, Y: 1},
		{X: 1, Y: 3},
		{X: 2, Y: 5},
		{X: 3, Y: 7},
		{X: 4, Y: 9},
	}
	derivs, err := derivatives(linearPoints)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for i, d := range derivs {
			fmt.Printf("x = %.1f: derivative ≈ %.1f\n", linearPoints[i].X, d)
		}
	}

	// Quadratic function y = x^2 - 2x + 3
	quadraticPoints := []Point{
		{X: -1.0, Y: 6.0},
		{X: -0.7, Y: 4.89},
		{X: -0.4, Y: 3.96},
		{X: -0.1, Y: 3.21},
		{X: 0.2, Y: 2.64},
		{X: 0.5, Y: 2.25},
	}
	derivs, err = derivatives(quadraticPoints)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		expected := []float64{-3.7, -3.4, -2.8, -2.2, -1.6, -1.3}
		for i, d := range derivs {
			fmt.Printf("x = %.1f: derivative ≈ %.1f (expected: %.1f)\n", quadraticPoints[i].X, d, expected[i])
		}
	}

	// Quartic function y = x^4 - x^2
	quarticPoints := []Point{
		{X: -3, Y: 72},
		{X: -2, Y: 12},
		{X: -1, Y: 0},
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 12},
		{X: 3, Y: 72},
	}
	derivs, err = derivatives(quarticPoints)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		expected := []float64{-60, -36, -2, 0, 2, 36, 60}
		for i, d := range derivs {
			fmt.Printf("x = %.1f: derivative ≈ %.1f (expected: %.1f)\n", quarticPoints[i].X, d, expected[i])
		}
	}

	// Function y = xe^x
	exponentialPoints := []Point{
		{X: 1.8, Y: 1.8 * math.Exp(1.8)},
		{X: 1.9, Y: 1.9 * math.Exp(1.9)},
		{X: 2.0, Y: 2.0 * math.Exp(2.0)},
		{X: 2.1, Y: 2.1 * math.Exp(2.1)},
		{X: 2.2, Y: 2.2 * math.Exp(2.2)},
	}
	derivs, err = derivatives(exponentialPoints)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		expected := []float64{18.13834004, 19.44373381, 22.16699562, 25.3845875, 27.06072882}
		for i, d := range derivs {
			fmt.Printf("x = %.1f: derivative ≈ %.8f (expected: %.8f)\n", exponentialPoints[i].X, d, expected[i])
		}
	}
}
