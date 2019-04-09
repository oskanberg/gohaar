package gohaar_test

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/oskanberg/gohaar"
)

const tolerance = 0.000000000000001

func TestTransform(t *testing.T) {
	data := []float64{1, 4, 2, 3}
	app, det := gohaar.Transform(data)

	a := []float64{3.5355339059327373, 3.5355339059327373}
	if !cmp.Equal(a, app) {
		t.Fatal(cmp.Diff(a, app))
	}

	d := []float64{-2.1213203435596424, -0.7071067811865475}
	if !cmp.Equal(d, det) {
		t.Fatal(cmp.Diff(d, det))
	}
}
func TestTransformOddSlice(t *testing.T) {
	data := []float64{1, 4, 2, 3, 4}
	app, det := gohaar.Transform(data)

	a := []float64{3.5355339059327373, 3.5355339059327373}
	if !cmp.Equal(a, app) {
		t.Fatal(cmp.Diff(a, app))
	}

	d := []float64{-2.1213203435596424, -0.7071067811865475}
	if !cmp.Equal(d, det) {
		t.Fatal(cmp.Diff(d, det))
	}
}
func TestConverseTransform(t *testing.T) {
	data := []float64{1, 4, 2, 3}
	app, det := gohaar.Transform(data)

	orig, err := gohaar.ConverseTransform(app, det)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(orig, data, cmpopts.EquateApprox(0, tolerance)) {
		t.Fatal(cmp.Diff(orig, data))
	}
}

func TestConverseTransformDifferentLengthArgs(t *testing.T) {
	_, err := gohaar.ConverseTransform([]float64{}, []float64{1})
	if err == nil {
		t.Fatal("expected err")
	}
}

/*
 Benchmarks
*/

// these exist to stop the compiler optimising out the call to Transform
var r1, r2 []float64

func benchmarkTransform(len int, b *testing.B) {
	d := make([]float64, len)
	for i := range d {
		d[i] = rand.NormFloat64() * 100
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r1, r2 = gohaar.Transform(d)
	}
}

func BenchmarkTransform2(b *testing.B)     { benchmarkTransform(2, b) }
func BenchmarkTransform20(b *testing.B)    { benchmarkTransform(20, b) }
func BenchmarkTransform200(b *testing.B)   { benchmarkTransform(200, b) }
func BenchmarkTransform2000(b *testing.B)  { benchmarkTransform(2000, b) }
func BenchmarkTransform20000(b *testing.B) { benchmarkTransform(20000, b) }
