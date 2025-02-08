# Go: Silent Panic from Division by Zero

This repository demonstrates a common error in Go: a potential runtime panic caused by division by zero that is not handled gracefully.

The `bug.go` file shows the erroneous code that does not comprehensively check for zero divisors. The `bugSolution.go` file provides a solution that improves error handling.

## Bug
The `myFunc` function correctly handles the case where `a` is zero, but fails to check if `b` is zero.  This leads to a runtime panic if `b` is zero.

## Solution
The solution includes comprehensive checks for both `a` and `b` to prevent a panic.