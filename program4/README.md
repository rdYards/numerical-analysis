# Interpolation Programming Assignment
## Newton’s Forward Divided Difference


## Natural Cubic Spline



**[Assignment Instructions](interpolation-program.pdf)**

## How to Run

### Newton's Divided-Difference Method
- **Input**:
  - `x`: slice of float64 values representing x-coordinates
  - `y`: slice of float64 values representing corresponding y-coordinates
- **Output**:
  - `coefficients`: slice of float64 values representing polynomial coefficients
  - `error`: If the method fails (Last element is Zero)

### Natural Cubic Spline
- **Input**:
  - `x`: slice of float64
  - `y`: slice of float64
- **Output**:
  - `coefficients`: 2D slice each row represents [a,b,c,d] for an interval
  - `error`: If the method fails (Last element is Zero)

## Test Cases
### Newton's Method Test Cases

1. **Exponential Function**:
   - Points: x=[0,1,2,3], y=[1,e,e²,e³]
   - Evaluation at x=1.5

2. **Decay Function**:
   - Points: x=[1.0,1.3,1.6,1.9,2.2], y=[0.7651977,0.6200860,0.4554022,0.2818186,0.1103623]
   - Evaluation at x=1.5

### Cubic Spline Test Cases

1. **Exponential Function**:
   - Same points as Newton's Method Test Case 1
   - Evaluations at x=0.5, 1.5, 2.5