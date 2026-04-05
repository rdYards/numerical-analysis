# Integration Programming Assignment

**[Assignment Instructions](integration-program.pdf)**

## How to Run

### Composite Simpson's Rule
- **Input**:
  - `f`: Function to integrate
  - `interval`: `[a, b]` - Integration interval
  - `n`: Number of subintervals (must be even and ≥ 2)
- **Output**:
  - `result`: Numerical approximation of the integral
  - `error`: Error

### Adaptive Quadrature
- **Input**:
  - `f`: Function to integrate
  - `interval`: `[a, b]` - Integration interval
  - `epsilon`: Error tolerance
  - `N`: Initial number of subintervals (must be > 0)
- **Output**:
  - `result`: Numerical approximation of the integral
  - `NUsed`: Number of subintervals used
  - `error`: Error

## Test Cases
### Derivatives Programming

1. **Composite Simpson's Rule**:
   - This should work exactly for any polynomial of degree at most 3. Verify this using $f(x) = x^3 - 3x^2 - 8$ on the interval $[0, 4]$ using $n = 2$, $n = 8$, and $n = 64$.
   - As in the notes, approximating $\int_{0}^{4} e^x \, dx$ using $n = 2$, $n = 4$, $n = 8$, should yield 56.76958, 53.86385, 53.61622, respectively.
   - Approximating $\int_{1}^{3} \frac{100}{x^2} \sin\left(\frac{10}{x}\right) \, dx$ with $n = 176$ yields $-1.42601386$.
2. **Adaptive Quadrature**:
   - This should work exactly and with no recursion for any polynomial of degree at most 3. Verify this using $f(x) = x^3 - 3x^2 - 8$ on the interval $[0, 4]$ using $\epsilon = 0.0000001$ and $N = 2$.
   - The algorithm should fail when approximating $\int_{0}^{4} e^x \, dx$ using $\epsilon = 0.0001$ and $N = 4$.
   - As in the notes, approximating $\int_{1}^{3} \frac{100}{x^2} \sin\left(\frac{10}{x}\right) \, dx$ with $\epsilon = 10^{-4}$ and $N = 100$ (this is overkill) should yield $-1.42601481$.
