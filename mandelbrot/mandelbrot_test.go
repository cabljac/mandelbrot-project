package mandelbrot

import (
	"math/cmplx"
	"math/rand"
	"testing"
)

func randomComplexNumber(t *testing.T) complex128 {
	t.Helper()
	real := rand.Float64()*4 - 2
	imag := rand.Float64()*4 - 2

	return complex(real, imag)
}

func TestMandelbrotIterationSimple(t *testing.T) {
	result := mandelbrotIteration(1i, 1)
	expected := complex128(0)
	tolerance := 1e-10

	if cmplx.Abs(result-expected) > tolerance {
		t.Errorf("mandelbrotIteration(1i, 1) = %v; want %v", result, expected)
	}
}

func TestMandelbrotIterationRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		z, c := randomComplexNumber(t), randomComplexNumber(t)

		result := mandelbrotIteration(z, c)
		expected := z*z + c
		tolerance := 1e-10

		if cmplx.Abs(result-expected) > tolerance {
			t.Errorf("mandelbrotIteration(%v, %v) = %v; want %v", z, c, result, expected)
		}
	}
}

func TestHasEscaped(t *testing.T) {
	tests := []struct {
		name     string
		input    complex128
		expected bool
	}{
		{"inside radius", 1 + 0i, false},
		{"exactly on radius", 2 + 0i, false},
		{"just outside", 2.1 + 0i, true},
		{"far outside", 10 + 0i, true},
		{"3-4-5 triangle outside", 3 + 4i, true}, // |3+4i| = 5 > 2
		{"imaginary only inside", 0 + 1.5i, false},
		{"zero", 0 + 0i, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasEscaped(tt.input)
			if result != tt.expected {
				t.Errorf("hasEscaped(%v) = %t; want %t",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestIterateUntilEscape(t *testing.T) {
	tests := []struct {
		name          string
		z, c          complex128
		maxIterations int
		expectedIter  int
		shouldEscape  bool
	}{
		{
			name:          "escapes immediately",
			z:             10 + 0i, // Already outside escape radius
			c:             0 + 0i,
			maxIterations: 100,
			expectedIter:  0,
			shouldEscape:  true,
		},
		{
			name:          "escapes quickly",
			z:             0 + 0i,
			c:             2.1 + 0i, // Will escape fast
			maxIterations: 100,
			expectedIter:  1, // Should escape after first iteration
			shouldEscape:  true,
		},
		{
			name:          "never escapes - origin",
			z:             0 + 0i,
			c:             0 + 0i, // Classic point in set
			maxIterations: 50,
			expectedIter:  50, // Should hit max iterations
			shouldEscape:  false,
		},
		{
			name:          "never escapes - limited iterations",
			z:             0 + 0i,
			c:             0.25 + 0i, // Point in set
			maxIterations: 10,
			expectedIter:  10,
			shouldEscape:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			escaped, iterations := iterateUntilEscape(tt.z, tt.c, tt.maxIterations)

			// Check iteration count
			if iterations != tt.expectedIter {
				t.Errorf("iterateUntilEscape(%v, %v, %d) iterations = %d; want %d",
					tt.z, tt.c, tt.maxIterations, iterations, tt.expectedIter)
			}

			// Check if escape status matches expectation
			if escaped != tt.shouldEscape {
				t.Errorf("iterateUntilEscape(%v, %v, %d) escaped = %t; want %t",
					tt.z, tt.c, tt.maxIterations, escaped, tt.shouldEscape)
			}
		})
	}
}

// Test the Config.Calculate method
func TestConfigCalculate(t *testing.T) {
	config := NewConfig()

	tests := []struct {
		name         string
		c            complex128
		shouldEscape bool
	}{
		{"origin - in set", 0 + 0i, false},
		{"point in set", 0.25 + 0i, false},
		{"point outside set", 2 + 0i, true},
		{"far outside", 10 + 0i, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			escaped, iterations := config.Calculate(tt.c)

			if escaped != tt.shouldEscape {
				t.Errorf("Calculate(%v) escaped = %t; want %t", tt.c, escaped, tt.shouldEscape)
			}

			// Iterations should be reasonable
			if iterations < 0 || iterations > config.MaxIterations {
				t.Errorf("Calculate(%v) iterations = %d; should be between 0 and %d",
					tt.c, iterations, config.MaxIterations)
			}
		})
	}
}
