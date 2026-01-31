package main

import (
	"errors"
	"fmt"
	"math"
)

// Newton's Method implementation
// Input: function needed, derivative to function, initial_guess, tolerance, max_iterations
// Output: p, p_(-1), f(p), iteration, error
func newtons_method(
	f func(float64) float64,
	df func(float64) float64,
	initial_guess float64,
	tol uint,
	max_iter int,
) (float64, float64, float64, int, error) {
	var p float64 = initial_guess - (f(initial_guess) / df(initial_guess))
	var p_1 = initial_guess
	var iteration int = 1
	var tolerance float64 = math.Pow(10, -float64(tol))
	
	if max_iter == 0 {
		max_iter = math.MaxInt
	}

	for iteration < max_iter {
		if math.Abs(p - p_1) < tolerance {
			return p, p_1, f(p), iteration, nil
		}
		iteration++

		p_1 = p
		p = p - (f(p) / df(p))
	}
	return p, p_1, f(p), iteration, errors.New("Max iterations reached")
}

// Secant Method implementation
// Input: function, guess0, guess1, tolerance, max_iterations
// Output: p, f(p) iteration, error
func secant_method(
	f func(float64) float64,
	guess0 float64,
	guess1 float64,
	tol uint,
	max_iter int,
) (float64, float64, int, error) {
	var p = 0.0
	var p0 = guess0
	var p1 = guess1
	var iteration int = 1
	var tolerance float64 = math.Pow(10, -float64(tol))

	if max_iter == 0 {
		max_iter = math.MaxInt
	}
	
	for iteration <= max_iter {
		p = p1 - ((f(p1) * (p1 - p0)) / (f(p1) - f(p0)))
		
		if math.Abs(p - p1) < tolerance {
			return p, f(p), iteration, nil
		}
		
		p0 = p1
		p1 = p
		iteration++
	}
	return p, f(p), iteration, errors.New("Max iterations reached")
}

// Functions for testing
// Used for Test 1 (a)
func f1(x float64) float64  { return math.Pow(x, 2) + 1 }
func df1(x float64) float64 { return 2 * x }

// Used for Test 1 (b)
func f2(x float64) float64  { return math.Atan(x) }
func df2(x float64) float64 { return 1 / (1 + x*x) }

// Used for Test 2 (a)
func j(x float64) float64  { return 1/(1 + math.Pow(x, 2)) - .5 }
func dj(x float64) float64 { return -2 * x / math.Pow(1+math.Pow(x, 2), 2) }

// Used for Test 3 (a)
func f3(x float64) float64  { return (5 * x) + 7 }
func df3(x float64) float64 { return 5 }

// Used for Test 3 (b) && Test 4 (a)
func f4(x float64) float64  { return math.Pow(x, 3) + (4 * math.Pow(x, 2)) - 10 }
func df4(x float64) float64 { return 3*x*x + 8*x }

// Used for Test 4 (b)
func f7(x float64) float64 { return math.Cos(x) - x }

func main() {
	// Newton’s Method
	// Verification of failure cases.
	// The function x2 + 1 does not have a root. Running with initial guess 0.5, tolerance
	// 0.0001, max iterations 10 should fail with pn ≈ −2.4, pn−1 ≈ 0.2, and f (pn) ≈ 6.76.
	p, p_1, fp, iter, err := newtons_method(f1, df1, 0.5, 4, 10)
	if err != nil {
		fmt.Printf("Error: %v, Iterations: %d | pn ≈ %.2f, pn−1 ≈ %.2f, f(pn) ≈ %.2f\n",
			err, iter, p, p_1, fp)
	} else {
		fmt.Printf("Root found, Iterations: %d | pn ≈ %.2f, pn−1 ≈ %.2f, f(pn) ≈ %.2f\n",
            iter, p, p_1, fp)
	}

	// The function arctan(x) has a root at 0. However, an initial guess of 1.4 with any tolerance
	// and 8 max iterations (or any higher, but I’ll use 8 in line with the lecture notes) should
	// fail with pn ≈ 16540.56, pn−1 ≈ −103.25, and f (pn) ≈ 1.57.
	p, p_1, fp, iter, err = newtons_method(f2, df2, 1.4, 4, 8)
    if err != nil {
        fmt.Printf("Error: %v, Iterations: %d | pn ≈ %.2f, pn−1 ≈ %.2f, f(pn) ≈ %.2f\n",
            err, iter, p, p_1, fp)
    } else {
        fmt.Printf("Root found, Iterations: %d | pn ≈ %.2f, pn−1 ≈ %.2f, f(pn) ≈ %.2f\n",
            iter, p, p_1, fp)
    }

	// Verification of unexpected root finding.
	// The function j(x) = 1/1+x2 − 1/2 has roots at ±1. Running with initial guess 2,
	// tolerance 0.001, and 10 max iterations should succeed in approximating the root at −1 as
	// −0.9999998144 (previous approximation −0.99939) in 8 iterations with f (p) = 9.2796 · 10−8
	p, p_1, fp, iter, err = newtons_method(j, dj, 2, 3, 10)
    if err != nil {
        fmt.Printf("Error: %v, Iterations: %d | pn ≈ %f, pn−1 ≈ %.2f, f(pn) ≈ %.20f\n",
            err, iter, p, p_1, fp)
    } else {
        fmt.Printf("Root found, Iterations: %d | pn ≈ %f, pn−1 ≈ %.2f, f(pn) ≈ %.20f\n",
            iter, p, p_1, fp)
    }

	// Verification of successful implementation.
	// When the function is a line, we should find the actual root on our second iteration
	// (technically, we find it on the first but don’t verify it until the second). Use 5x + 7,
	// initial guess 1000, tolerance 0.0001, max iterations 100.
	p, p_1, fp, iter, err = newtons_method(f3, df3, 1000, 4, 100)
    if err != nil {
        fmt.Printf("Error: %v, Iterations: %d | pn ≈ %f, pn−1 ≈ %f, f(pn) ≈ %.20f\n",
            err, iter, p, p_1, fp)
    } else {
        fmt.Printf("Root found, Iterations: %d | p ≈ %f, f(p) ≈ %.20f\n",iter, p, fp)
    }

	// As seen in example 2.2.1, the function x3 + 4x2 − 10 with initial guess 1, tolerance
	// 0.0001, max iterations 20 should converge in 4 iterations at 1.365230013 with f (p) =
	// 3.512354851 · 10−10.
	p, p_1, fp, iter, err = newtons_method(f4, df4, 1, 4, 20)
    if err != nil {
        fmt.Printf("Error: %v, Iterations: %d | pn ≈ %f, pn−1 ≈ %f, f(pn) ≈ %.20f\n",
            err, iter, p, p_1, fp)
    } else {
        fmt.Printf("Root found, Iterations: %d | pn ≈ %f, pn−1 ≈ %f, f(pn) ≈ %.20f\n",
            iter, p, p_1, fp)
    }

	// Secant Method
	// Verification of successful implementation.
	// As in lecture notes, the function x3 + 4x2 − 10 with initial guesses 1 and 2, tolerance
	// 0.0001, max iterations 20 should converge in 6 iterations at 1.365230001 with f (p) =
	// −2.031682733 · 10−7.
	p, fp, iter, err = secant_method(f4, 1, 2, 3, 20)
	if err != nil {
        fmt.Printf("Error: %v, Iterations: %d | p ≈ %f, f(p) ≈ %.20f\n",
            err, iter, p, fp)
    } else {
        fmt.Printf("Root found, Iterations: %d | pn ≈ %f, pn−1 ≈ %f, f(pn) ≈ %.20f\n",
            iter, p, p_1, fp)
    }

	// As in lecture notes, the function cos(x) − x with initial guesses 0.5 and π/4 , tolerance
	// 0.0001, max iterations 20 should converge in 4 iterations at 0.7390851493 with f (p) =
	// −2.698216706 · 10−8.
	p, fp, iter, err = secant_method(f7, .5, (math.Pi / 4), 3, 20)
	if err != nil {
        fmt.Printf("Error: %v, Iterations: %d | p ≈ %f, f(p) ≈ %.20f\n",
            err, iter, p, fp)
    } else {
        fmt.Printf("Root found, Iterations: %d | p ≈ %f, f(p) ≈ %.20f\n",iter, p, fp)
    }
}
