package vector_test

import (
	"testing"

	"github.com/axamon/q/gate"
	"github.com/axamon/q/matrix"
	"github.com/axamon/q/vector"
)

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