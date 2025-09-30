package main

import (
	"fmt"
	"strings"
)

// Vector 向量结构体
type Vector struct {
	values []float64
}

// NewVector 构造函数，相当于Python的__init__
func NewVector(values []float64) *Vector {
	// 创建副本，避免外部修改
	valuesCopy := make([]float64, len(values))
	copy(valuesCopy, values)

	return &Vector{values: valuesCopy}
}

// Add 向量加法，相当于Python的__add__
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

// Sub 向量减法，相当于Python的__sub__
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

// Mul 向量数乘，相当于Python的__mul__
func (v *Vector) Mul(k float64) *Vector {
	result := make([]float64, len(v.values))
	for i := 0; i < len(v.values); i++ {
		result[i] = k * v.values[i]
	}
	return NewVector(result)
}

// Pos 向量取正，相当于Python的__pos__
func (v *Vector) Pos() *Vector {
	return v.Mul(1)
}

// Neg 向量取负，相当于Python的__neg__
func (v *Vector) Neg() *Vector {
	return v.Mul(-1)
}

// Get 获取指定索引的元素，相当于Python的__getitem__
func (v *Vector) Get(index int) (float64, error) {
	if index < 0 || index >= len(v.values) {
		return 0, fmt.Errorf("index out of range")
	}
	return v.values[index], nil
}

// Len 获取向量长度，相当于Python的__len__
func (v *Vector) Len() int {
	return len(v.values)
}

// String 字符串表示，相当于Python的__str__
func (v *Vector) String() string {
	strValues := make([]string, len(v.values))
	for i, val := range v.values {
		strValues[i] = fmt.Sprintf("%.2f", val)
	}
	return fmt.Sprintf("(%s)", strings.Join(strValues, ", "))
}

// GoString 详细字符串表示，相当于Python的__repr__
func (v *Vector) GoString() string {
	return fmt.Sprintf("Vector%v", v.values)
}

// Values 获取向量的值切片（用于迭代），相当于Python的__iter__
func (v *Vector) Values() []float64 {
	// 返回副本，避免外部修改
	valuesCopy := make([]float64, len(v.values))
	copy(valuesCopy, v.values)
	return valuesCopy
}

// 使用示例
func main() {
	// 创建向量
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{4, 5, 6})

	fmt.Printf("v1: %s\n", v1) // (1.00, 2.00, 3.00)
	fmt.Printf("v2: %s\n", v2) // (4.00, 5.00, 6.00)

	// 向量加法
	sum, err := v1.Add(v2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("v1 + v2 = %s\n", sum) // (5.00, 7.00, 9.00)
	}

	// 向量减法
	diff, err := v1.Sub(v2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("v1 - v2 = %s\n", diff) // (-3.00, -3.00, -3.00)
	}

	// 数乘
	scaled := v1.Mul(2.5)
	fmt.Printf("v1 * 2.5 = %s\n", scaled) // (2.50, 5.00, 7.50)

	// 取负
	neg := v1.Neg()
	fmt.Printf("-v1 = %s\n", neg) // (-1.00, -2.00, -3.00)

	// 获取元素
	val, err := v1.Get(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("v1[1] = %.2f\n", val) // 2.00
	}

	// 获取长度
	fmt.Printf("len(v1) = %d\n", v1.Len()) // 3

	// 迭代向量
	fmt.Print("v1 elements: ")
	for i, val := range v1.Values() {
		fmt.Printf("[%d]=%.2f ", i, val)
	}
	fmt.Println()

	// 详细表示
	fmt.Printf("GoString: %#v\n", v1) // Vector[1 2 3]
}
