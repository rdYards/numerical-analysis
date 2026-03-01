package main

import (
	"errors"
	"fmt"
	"math"
)

// Newton's Divided-Difference Function
// Input: x, y
// Output: coefficients for the polynomial
func dividedDifference(x, y []float64) ([]float64, error) {
	n := len(x)
	if n != len(y) || n < 2 {
		return nil, errors.New("x and y must have same length and at least 2 points")
	}

	// Create a 2D slice to store the divided differences
	F := make([][]float64, n)
	for i := range F {
		F[i] = make([]float64, n)
	}

	// Initialize the first column with y values
	for i := range y {
		F[i][0] = y[i]
	}

	// Compute divided differences
	for j := 1; j < n; j++ {
		for i := j; i < n; i++ {
			F[i][j] = (F[i][j-1] - F[i-1][j-1]) / (x[i] - x[i-j])
		}
	}

	// Extract the coefficients
	coefficients := make([]float64, n)
	for i := range coefficients {
		coefficients[i] = F[i][i]
	}

	return coefficients, nil
}

// Evaluate the polynomial at x using the coefficients
func evaluatePolynomial(coefficients []float64, x_value float64, x0 []float64) float64 {
	n := len(coefficients)
	if n != len(x0) {
		return math.NaN()
	}

	result := coefficients[0]
	for i := 1; i < n; i++ {
		term := coefficients[i]
		for j := 0; j < i; j++ {
			term *= (x_value - x0[j])
		}
		result += term
	}

	return result
}

// Natural Cubic Spline Function
// Input: x, y
// Output: coefficients a, b, c, d for each interval
func naturalCubicSpline(x, y []float64) ([][]float64, error) {
	n := len(x)
	if n != len(y) || n < 2 {
		return nil, errors.New("x and y must have same length and at least 2 points")
	}

	// Initialize h (interval lengths)
	h := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		h[i] = x[i+1] - x[i]
	}

	// Initialize alpha vector
	alpha := make([]float64, n-1)
	for i := 1; i < n-1; i++ {
		alpha[i] = 3 * ((y[i+1]-y[i])/h[i] - (y[i]-y[i-1])/h[i-1])
	}

	// Initialize l, mu, z vectors
	l := make([]float64, n)
	mu := make([]float64, n)
	z := make([]float64, n)

	l[0] = 1
	mu[0] = 0
	z[0] = 0

	for i := 1; i < n-1; i++ {
		l[i] = 2*(x[i+1]-x[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}

	l[n-1] = 1
	z[n-1] = 0

	// Back substitution
	c := make([]float64, n)
	c[n-1] = 0
	for j := n - 2; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
	}

	// Compute b and d coefficients
	b := make([]float64, n-1)
	d := make([]float64, n-1)

	for i := 0; i < n-1; i++ {
		b[i] = (y[i+1]-y[i])/h[i] - h[i]*(c[i+1]+2*c[i])/3
		d[i] = (c[i+1] - c[i]) / (3 * h[i])
	}

	// Prepare the result as a slice of slices [a, b, c, d] for each interval
	result := make([][]float64, n-1)
	for i := 0; i < n-1; i++ {
		result[i] = []float64{y[i], b[i], c[i], d[i]}
	}

	return result, nil
}

// Evaluate the spline at x using the coefficients
func evaluateSpline(coefficients [][]float64, x_value float64, x0 []float64) float64 {
	n := len(coefficients)
	if n != len(x0)-1 {
		return math.NaN()
	}

	// Find the interval containing x_value
	interval := 0
	for i := 1; i < len(x0); i++ {
		if x_value <= x0[i] {
			break
		}
		interval = i
	}

	if interval >= n {
		return math.NaN()
	}

	// Evaluate the spline in the found interval
	a, b, c, d := coefficients[interval][0], coefficients[interval][1], coefficients[interval][2], coefficients[interval][3]
	dx := x_value - x0[interval]
	return a + b*dx + c*dx*dx + d*dx*dx*dx
}

func main() {
	// Verification of Newton’s Forward Divided Difference
	// Using Newton’s Forward Divided Difference to approximate f(x) = e^x on the interval
	x1 := []float64{0, 1, 2, 3}
	y1 := []float64{1, math.E, math.E * math.E, math.E * math.E * math.E}
	coeffs1, err := dividedDifference(x1, y1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Coefficients:", coeffs1)
	p15 := evaluatePolynomial(coeffs1, 1.5, x1)
	fmt.Printf("P(1.5) = %.9f\n", p15)

	// Using Newton’s Forward Divided Difference on this list of 5 points
	x2 := []float64{1.0, 1.3, 1.6, 1.9, 2.2}
	y2 := []float64{0.7651977, 0.6200860, 0.4554022, 0.2818186, 0.1103623}
	coeffs2, err := dividedDifference(x2, y2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Coefficients:", coeffs2)
	p15b := evaluatePolynomial(coeffs2, 1.5, x2)
	fmt.Printf("P(1.5) = %.7f\n", p15b)

	// Verification of Natural Cubic Spline
	// Natural Cubic Spline for f(x) = e^x
	splineCoeffs, err := naturalCubicSpline(x1, y1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Spline coefficients (a, b, c, d for each interval):")
	for i, coeff := range splineCoeffs {
		fmt.Printf("Interval %d: a=%.9f, b=%.9f, c=%.9f, d=%.9f\n",
			i, coeff[0], coeff[1], coeff[2], coeff[3])
	}
	s05 := evaluateSpline(splineCoeffs, 0.5, x1)
	s15 := evaluateSpline(splineCoeffs, 1.5, x1)
	s25 := evaluateSpline(splineCoeffs, 2.5, x1)
	fmt.Printf("S(0.5) = %.9f\n", s05)
	fmt.Printf("S(1.5) = %.9f\n", s15)
	fmt.Printf("S(2.5) = %.9f\n", s25)
}
