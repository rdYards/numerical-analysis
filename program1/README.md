# Bisection Method Implementation

This program implements the **Bisection Method**, a root-finding algorithm for continuous functions. The implementation includes error handling and additional outputs to verify convergence.

## How to Run

1. **Navigate to the project directory**:
   ```sh
   cd numerical-analysis/program1
   ```

3. **Run the program**:
   ```sh
   go run main.go
   ```

## Implementation Details

### Bisection Method
- **Input**:
  - `interval [a, b]`: The interval to search for a root.
  - `tol`: Tolerance (precision) for convergence (e.g., `3` means `1e-3`).
  - `max_iter`: Maximum allowed iterations (set to `0` for unlimited).

- **Output**:
  - `convergence`: Approximate root.
  - `f(convergence)`: Function value at the root (should be close to 0).
  - `iterations`: Number of iterations performed.
  - `error`: If the method fails (e.g., no root in interval or max iterations reached).

## Test Cases

## Function Used
At the current moment the function is hardcoded into the program as a `func`. To use another function please change `func f` to desired function.
```go
f(x) = math.Pow(x, 3) - (9 * x)
```

### Test 1: Failure Cases
- **1. Interval [1, 2]**:
  - Expected: Error ("No root guaranteed in the interval").
- **2. Interval [-1, 4]**:
  - Expected: Error (IVT not satisfied at endpoints).
- **3. Interval [-0.5, 1] with `tol=3` and `max_iter=8`**:
  - Expected: Error ("Max iterations reached").

### Test 2: Successful Cases
- **1. Interval [2.7, 3.5]**:
  - Expected: Converges in **3 iterations** (midpoint root).
- **2. Interval [-4, -1] with `tol=3` and `max_iter=20`**:
  - Expected: Converges in **12 iterations**, root ≈ `-3.00024`, `f(root) ≈ -0.00439507`.
