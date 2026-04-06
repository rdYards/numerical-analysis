package main

import (
	"errors"
	"fmt"
	"math"
)

type Function func(float64) float64

// Functions used
func f1(x float64) float64 {
	return math.Pow(x, 3) - 3*math.Pow(x, 2) - 8
}
func f2(x float64) float64 {
	return math.Exp(x)
}
func f3(x float64) float64 {
	return 100 / (math.Pow(x, 2)) * math.Sin(10/x)
}

// Input: Function, Interval [a, b], Number of subintervals
// Output: Integral approximation, Error
func CompositeSimpson(f Function, interval [2]float64, n int) (float64, error) {
	if n%2 != 0 {
		return 0, errors.New("n must be even for Simpson's rule")
	}
	if n < 2 {
		return 0, errors.New("n must be at least 2")
	}

	h := (interval[1] - interval[0]) / float64(n)
	sum := f(interval[0]) + f(interval[1])

	var sumOdd, sumEven float64
	for i := 1; i < n; i++ {
		x := interval[0] + h*float64(i)
		if i%2 == 1 {
			sumOdd += f(x)
		} else {
			sumEven += f(x)
		}
	}

	return (h / 3) * (sum + 4*sumOdd + 2*sumEven), nil
}

// Input: Function, Interval [a, b], epsilon, Number of subintervals
// Output: Integral approximation, Error
func AdaptiveQuadrature(f Function, interval [2]float64, epsilon float64, N int) (float64, int, error) {
	maxDepth := 3
	depth := 0

	if N <= 0 {
		return 0, 0, errors.New("N must be positive")
	}

	approx, err := CompositeSimpson(f, interval, N)
	if err != nil {
		return 0, 0, err
	}
	NUsed := N

	for depth < maxDepth {
		newApprox, err := CompositeSimpson(f, interval, 2*N)
		if err != nil {
			return 0, NUsed, err
		}

		if math.Abs(newApprox-approx) < epsilon {
			return newApprox, NUsed, nil
		}

		approx = newApprox
		NUsed = 2 * N
		N *= 2
		depth++
	}

	return approx, NUsed, errors.New("maximum recursion depth reached without convergence")
}

func main() {
	// Composite Simpson’s Rule.
	// This should work exactly for any polynomial of degree at most 3.
	for _, n := range []int{2, 8, 64} {
		result, err := CompositeSimpson(f1, [2]float64{0, 4}, n)
		if err != nil {
			fmt.Printf("Error with n=%d: %v\n", n, err)
		} else {
			fmt.Printf("n=%d: %f\n", n, result)
		}
	}

	// approximating $\int_{0}^{4} e^x \, dx$ using $n = 2$, $n = 4$, $n = 8$,
	for _, n := range []int{2, 4, 8} {
		result, err := CompositeSimpson(f2, [2]float64{0, 4}, n)
		if err != nil {
			fmt.Printf("Error with n=%d: %v\n", n, err)
		} else {
			fmt.Printf("n=%d: %f\n", n, result)
		}
	}

	// Approximating $\int_{1}^{3} \frac{100}{x^2} \sin\left(\frac{10}{x}\right) \, dx$
	// with $n = 176$ yields $-1.42601386$.
	result, err := CompositeSimpson(f3, [2]float64{1, 3}, 176)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}

	// Adaptive Quadrature
	// This should work exactly and with no recursion for any polynomial of degree at most 3.
	result, nUsed, err := AdaptiveQuadrature(f1, [2]float64{0, 4}, 1e-7, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %f (used N=%d)\n", result, nUsed)
	}

	// The algorithm should fail
	result, nUsed, err = AdaptiveQuadrature(f2, [2]float64{0, 4}, 1e-4, 4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %f (used N=%d)\n", result, nUsed)
	}

	// Approximating $\int_{1}^{3} \frac{100}{x^2} \sin\left(\frac{10}{x}\right) \, dx$
	// with $\epsilon = 10^{-4}$ and $N = 100$
	result, nUsed, err = AdaptiveQuadrature(f3, [2]float64{1, 3}, 1e-4, 100)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %f (used N=%d)\n", result, nUsed)
	}
}
