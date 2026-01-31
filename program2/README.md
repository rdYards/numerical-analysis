# Newton’s Method

Newton's Method (also known as the Newton-Raphson method) is an iterative numerical technique for finding roots of a real-valued function. The method uses the first few terms of the function's Taylor series expansion around a given point to approximate the root.

# Secant Method

The Secant Method is another iterative root-finding technique that approximates the root of a function. It's similar to Newton's Method but doesn't require the calculation of the derivative. Instead, it uses the secant line (a line connecting two points on the function) to approximate the root.

## How to Run

1. **Navigate to the project directory**:
   ```sh
   cd numerical-analysis/program2
   ```

3. **Run the program**:
   ```sh
   go run main.go
   ```

## Implementation Details
### Newton's Method
- **Input**:
  - `f`: The function whose root is to be found
  - `df`: The derivative of the function
  - `initial_guess`: Starting point for the iteration
  - `tol`: Tolerance (precision) for convergence (e.g., `3` means `1e-3`)
  - `max_iter`: Maximum allowed iterations (set to `0` for unlimited)

- **Output**:
  - `p`: The approximate root found
  - `p_1` (for Newton's Method): The previous approximation
  - `f(p)`: The function value at the approximate root
  - `iterations`: Number of iterations performed
  - `error`: If the method fails (e.g., no root in interval or max iterations reached)

### Secant Method
- **Input**:
  - `f`: The function whose root is to be found
  - `guess0`: First initial guess
  - `guess1`: Second initial guess
  - `tol`: Tolerance (precision) for convergence (e.g., `3` means `1e-3`)
  - `max_iter`: Maximum allowed iterations (set to `0` for unlimited)

- **Output**:
  - `p`: The approximate root found
  - `f(p)`: The function value at the approximate root
  - `iterations`: Number of iterations performed
  - `error`: If the method fails (e.g., max iterations reached)

## Test Cases
## Function Used
At the current moment the functions are hardcoded into the program. To use another function please add new function.

### Newton's Method Test Cases

1. **Failure Case (x² + 1)**:
   - Expected: Error ("Max iterations reached")

2. **Failure Case (arctan(x))**:
   - Expected: Error ("Max iterations reached")

3. **Unexpected Root (j(x))**:
   - Expected: Find root near -1

4. **Linear Function (5x + 7)**:
   - Expected: Find exact root in 2 iterations

5. **Cubic Function (x³ + 4x² - 10)**:
   - Expected: Find root ≈ 1.36523 in 4 iterations

### Secant Method Test Cases

1. **Cubic Function (x³ + 4x² - 10)**:
   - Expected: Find root ≈ 1.36523 in 6 iterations

2. **Trigonometric Function (cos(x) - x)**:
   - Expected: Find root ≈ 0.73909 in 4 iterations
```