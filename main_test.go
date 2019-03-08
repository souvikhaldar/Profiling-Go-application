package main

import "testing"

// TestPrime tests for prime
func TestPrime(t *testing.T) {
	x := prime(10000)
	t.Log(x)
}
