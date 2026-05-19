package main

import (
	"fmt"
	"strings"
)

// Vector struct
type Vector struct {
	values []float64
}

// NewVector constructor, equivalent to Python __init__
func NewVector(values []float64) *Vector {
	// Create a copy to prevent external modification
	valuesCopy := make([]float64, len(values))
	copy(valuesCopy, values)

	return &Vector{values: valuesCopy}
}

// Add vector addition, equivalent to Python __add__
func (v *Vector) Add(another *Vector) (*Vector, error) {
	if len(v.values) != len(another.values) {
		return nil, fmt.Errorf("Error in adding 维度不相等")
	}

	result := make([]float64, len(v.values))
	for i := 0; i < len(v.values); i++ {
		result[i] = v.values[i] + another.values[i]
	}
	return NewVector(result), nil
}

// Sub vector subtraction, equivalent to Python __sub__
func (v *Vector) Sub(another *Vector) (*Vector, error) {
	if len(v.values) != len(another.values) {
		return nil, fmt.Errorf("Error in subtracting. Length of vectors must be same")
	}

	result := make([]float64, len(v.values))
	for i := 0; i < len(v.values); i++ {
		result[i] = v.values[i] - another.values[i]
	}
	return NewVector(result), nil
}

// Mul scalar multiplication, equivalent to Python __mul__
func (v *Vector) Mul(k float64) *Vector {
	result := make([]float64, len(v.values))
	for i := 0; i < len(v.values); i++ {
		result[i] = k * v.values[i]
	}
	return NewVector(result)
}

// Pos vector positive, equivalent to Python __pos__
func (v *Vector) Pos() *Vector {
	return v.Mul(1)
}

// Neg vector negation, equivalent to Python __neg__
func (v *Vector) Neg() *Vector {
	return v.Mul(-1)
}

// Get retrieves element at index, equivalent to Python __getitem__
func (v *Vector) Get(index int) (float64, error) {
	if index < 0 || index >= len(v.values) {
		return 0, fmt.Errorf("index out of range")
	}
	return v.values[index], nil
}

// Len returns vector length, equivalent to Python __len__
func (v *Vector) Len() int {
	return len(v.values)
}

// String representation, equivalent to Python __str__
func (v *Vector) String() string {
	strValues := make([]string, len(v.values))
	for i, val := range v.values {
		strValues[i] = fmt.Sprintf("%.2f", val)
	}
	return fmt.Sprintf("(%s)", strings.Join(strValues, ", "))
}

// GoString detailed string representation, equivalent to Python __repr__
func (v *Vector) GoString() string {
	return fmt.Sprintf("Vector%v", v.values)
}

// Values returns a copy of vector values (for iteration), equivalent to Python __iter__
func (v *Vector) Values() []float64 {
	// Return a copy to prevent external modification
	valuesCopy := make([]float64, len(v.values))
	copy(valuesCopy, v.values)
	return valuesCopy
}

// Usage example
func main() {
	// Create vector
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{4, 5, 6})

	fmt.Printf("v1: %s\n", v1) // (1.00, 2.00, 3.00)
	fmt.Printf("v2: %s\n", v2) // (4.00, 5.00, 6.00)

	// Vector addition
	sum, err := v1.Add(v2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("v1 + v2 = %s\n", sum) // (5.00, 7.00, 9.00)
	}

	// Vector subtraction
	diff, err := v1.Sub(v2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("v1 - v2 = %s\n", diff) // (-3.00, -3.00, -3.00)
	}

	// Scalar multiplication
	scaled := v1.Mul(2.5)
	fmt.Printf("v1 * 2.5 = %s\n", scaled) // (2.50, 5.00, 7.50)

	// Negation
	neg := v1.Neg()
	fmt.Printf("-v1 = %s\n", neg) // (-1.00, -2.00, -3.00)

	// Get element
	val, err := v1.Get(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("v1[1] = %.2f\n", val) // 2.00
	}

	// Get length
	fmt.Printf("len(v1) = %d\n", v1.Len()) // 3

	// Iterate vector
	fmt.Print("v1 elements: ")
	for i, val := range v1.Values() {
		fmt.Printf("[%d]=%.2f ", i, val)
	}
	fmt.Println()

	// Detailed representation
	fmt.Printf("GoString: %#v\n", v1) // Vector[1 2 3]
}
