package main

import (
	"errors"
	"fmt"
	"math"
)

// Function used with bisection_method
func f(x float64) float64 {
	return math.Pow(x, 3) - (9 * x)
}

// Input: Interval[a, b], tolerance, max Iteration
// Output: Statement, Covergence, iteration
func bisection_method(
	interval [2]float64,
	tol uint,
	max_iter int,
) (float64, float64, int, error) {
	if f(interval[0]) * f(interval[1]) >= 0 {
		return 0.0, 0.0, 0, errors.New("No root guaranteed in the interval")
	}

	var iteration int = 0
	var tolerance float64 = math.Pow(10, -float64(tol))
	var convergance float64 = 0.0
	if max_iter == 0 {
		max_iter = math.MaxInt
	}

	for (interval[1]-interval[0])/2 > tolerance && iteration < max_iter {
		iteration++
		convergance = (interval[0] + interval[1]) / 2
		if f(convergance) == 0 {
			break
		} else if f(interval[0]) * f(convergance) < 0 {
			interval[1] = convergance
		} else {
			interval[0] = convergance
		}
	}
	if iteration == max_iter {
		return 0.0, 0.0, 0, errors.New("Max iterations reached without convergence")
	}
	return convergance, f(convergance), iteration, nil
}

func main() {
	// Verification of failure cases:
	// The function f (x) does not have a root on the interval [1, 2].
	var convergance, convergance_f, iteration, error = bisection_method([2]float64{1, 2}, 3, 0)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf(
			"Root found: convergance ≈ %f, f(convergance) ≈ %f, with %d iterations.\n",
			convergance, convergance_f, iteration,
		)
	}

	// The function f (x) on the interval [−1, 4] is not guaranteed to have a root by the IVT
	convergance, convergance_f, iteration, error = bisection_method([2]float64{-1, 4}, 3, 0)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf(
			"Root found: convergance ≈ %f, f(convergance) ≈ %f, with %d iterations.\n",
			convergance, convergance_f, iteration,
		)
	}

	// The function f (x) has a root at x = 0.
	// However, starting with the interval [−0.5, 1] and seeking this root with a tolerance of 10−3
	// and a max of 8 iterations should fail.
	convergance, convergance_f, iteration, error = bisection_method([2]float64{-0.5, 1}, 3, 8)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf(
			"Root found: convergance ≈ %f, f(convergance) ≈ %f, with %d iterations.\n",
			convergance, convergance_f, iteration,
		)
	}

	// Verification of successful implementation
	// If the root happens at precisely the midpoint of some iteration, we should terminate.
	// This should happen in 3 iterations on f (x) on the interval [2.7, 3.5].
	convergance, convergance_f, iteration, error = bisection_method([2]float64{2.7, 3.5}, 3, 0)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf(
			"Root found: convergance ≈ %f, f(convergance) ≈ %f, with %d iterations.\n",
			convergance, convergance_f, iteration,
		)
	}

	// The function f (x) on the interval [−4, −1] with a tolerance of 10−3 and max iterations
	// 20 should terminate in 12 iterations, with p ≈ −3.00024, f (p) ≈ −0.00439507.
	convergance, convergance_f, iteration, error = bisection_method([2]float64{-4, -1}, 3, 20)
	if error != nil {
		fmt.Printf("Error: %v.\n", error)
	} else {
		fmt.Printf(
			"Root found: convergance ≈ %f, f(convergance) ≈ %f, with %d iterations.\n",
			convergance, convergance_f, iteration,
		)
	}
}
