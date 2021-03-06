package vector_test

import (
	"fmt"
	"testing"

	"github.com/axamon/q/gate"
	"github.com/axamon/q/matrix"
	"github.com/axamon/q/vector"
)

func ExampleNewZero() {
	v0 := vector.NewZero(4)
	fmt.Printf("%v\n", v0)
	// Output:
	// [(0+0i) (0+0i) (0+0i) (0+0i)]
}

var v = vector.New(
	complex128(1+2i),
	complex128(3+4i),
)

func ExampleVector_Clone() {
	vCloned := v.Clone()
	fmt.Printf("%v\n", vCloned)
	// Output:
	// [(1+2i) (3+4i)]
}

func ExampleVector_Dual() {
	vDual := v.Dual()
	fmt.Printf("%v\n", vDual)
	// Output:
	// [(1-2i) (3-4i)]
}

func ExampleVector_Add() {
	result := v.Add(v)
	fmt.Printf("%v\n", result)
	// Output:
	// [(2+4i) (6+8i)]
}

func ExampleVector_Mul() {
	result := v.Mul(complex128(1 - 1i))
	fmt.Printf("%v\n", result)
	result = v.Mul(complex128(5 + 2i))
	fmt.Printf("%v\n", result)
	result = result.Mul(complex128(0 + 0i))
	fmt.Printf("%v\n", result)
	// Output:
	// [(3+1i) (7+1i)]
	// [(1+12i) (7+26i)]
	// [(0+0i) (0+0i)]
}

func ExampleVector_Dimension() {
	numOdDimensions := v.Dimension()
	fmt.Println(numOdDimensions)
	// Output:
	// 2
}

func ExampleVector_Equals() {
	v3 := vector.NewZero(3)
	v2 := vector.NewZero(2)
	vCloned := v.Clone()
	fmt.Println(v.Equals(v3))
	fmt.Println(v.Equals(v2))
	fmt.Println(v.Equals(vCloned))
	// Output:
	// false
	// false
	// true

}

func ExampleTensorProductN() {
	fmt.Println(vector.TensorProductN(v))
	fmt.Println(vector.TensorProductN(v, 0))
	fmt.Println(vector.TensorProductN(v, 1))
	fmt.Println(vector.TensorProductN(v, 2))
	fmt.Println(vector.TensorProductN(v, 3))
	// Output:
	// [(1+2i) (3+4i)]
	// [(1+2i) (3+4i)]
	// [(1+2i) (3+4i)]
	// [(-3+4i) (-5+10i) (-5+10i) (-7+24i)]
	// [(-11-2i) (-25+0i) (-25+0i) (-55+10i) (-25+0i) (-55+10i) (-55+10i) (-117+44i)]
}

func TestOuterProduct(t *testing.T) {
	v0 := vector.New(1, 0)
	out := v0.OuterProduct(v0)

	if out[0][0] != complex(1, 0) {
		t.Fail()
	}

	if out[1][0] != complex(0, 0) {
		t.Fail()
	}

	if out[0][1] != complex(0, 0) {
		t.Fail()
	}

	if out[1][1] != complex(0, 0) {
		t.Fail()
	}
}

func TestVector(t *testing.T) {
	v0 := vector.New(1, 1)
	v1 := vector.New(1, -1)

	if v0.InnerProduct(v1) != complex(0, 0) {
		t.Error(v0.InnerProduct(v1))
	}

	if !v0.IsOrthogonal(v1) {
		t.Error(v0.InnerProduct(v1))
	}

	v3 := vector.New(1, 0)
	if v3.InnerProduct(v3) != complex(1, 0) {
		t.Error(v3.InnerProduct(v3))
	}

	if v3.IsOrthogonal(v3) {
		t.Error(v3.InnerProduct(v3))
	}

	if !v3.IsUnit() {
		t.Error(v3.IsUnit())
	}

	if v3.Norm() != complex(1, 0) {
		t.Error(v3.Norm())
	}

	v4 := vector.New(0, 1)
	if v3.InnerProduct(v4) != complex(0, 0) {
		t.Error(v3.InnerProduct(v4))
	}
	if !v3.IsOrthogonal(v4) {
		t.Error(v3.InnerProduct(v4))
	}

}

func TestTensorProduct(t *testing.T) {
	v := vector.New(1, 0)

	v4 := vector.TensorProduct(v, v)
	x4 := matrix.TensorProduct(gate.X(), gate.X())

	xv4 := v4.Apply(x4)
	expected := vector.TensorProduct(vector.New(0, 1), vector.New(0, 1))
	if !xv4.Equals(expected) {
		t.Error(xv4)
	}

	v16 := vector.TensorProduct(v4, v4)
	x16 := matrix.TensorProduct(x4, x4)
	xv16 := v16.Apply(x16)

	expected16 := vector.TensorProduct(vector.New(0, 0, 0, 1), vector.New(0, 0, 0, 1))
	if !xv16.Equals(expected16) {
		t.Error(xv16)
	}

}
