# Gaussian Elimination

Gaussian elimination is a algorithm to solve systems of linear equations by transforming a matrix into a simpler, upper-triangular form. It uses basic row operations to eliminate variables step by step until the solution becomes clear through back-substitution.

## How to Run

1. **Navigate to the project directory**:
   ```sh
   cd numerical-analysis/program3
   ```

3. **Run the program**:
   ```sh
   go run main.go
   ```

## Implementation Details
### Gaussian Elimination
- **Input**:
  - `input_matrix`: Augmented matrix (A|b)

- **Output**:
  - `solution`: Vector of solutions
  - `matrix`: Final state of matrix when finished
  - `error`: If the method fails (Last element is Zero)

## Test Cases
## Function Used
At the current moment the matrixes are hardcoded into the program. To use another matrixes please add/modify matrixes.

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