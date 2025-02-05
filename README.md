# Uninitialized Map Access Panic in Go

This repository demonstrates a common error in Go: panicking due to accessing a key in an uninitialized map.

## The Problem

In Go, maps must be initialized before use. Attempting to access a key in an uninitialized map will result in a runtime panic.  The `bug.go` file contains code that illustrates this problem.

## The Solution

The `bugSolution.go` file shows the correct way to initialize a map. Maps can be initialized using the `make` function or using a map literal.

## How to Run

1. Clone the repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` to observe the panic.  
4. Run `go run bugSolution.go` to see the corrected version.

This example highlights the need for careful map initialization in Go to prevent unexpected program crashes.