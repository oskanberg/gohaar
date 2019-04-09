package gohaar // import "github.com/oskanberg/gohaar"

import (
	"errors"
	"math"
)

// Transform takes the Haar wavelet transform on f
// It returns the result in two slices:
// the approximation coefficients, and the detail coefficients.
func Transform(f []float64) ([]float64, []float64) {
	nl := len(f) / 2
	app := make([]float64, nl)
	det := make([]float64, nl)

	for i := 0; i < nl; i++ {
		j := f[2*i]
		k := f[(2*i)+1]
		app[i] = (j + k) / math.Sqrt2
		det[i] = (j - k) / math.Sqrt2
	}

	return app, det
}

// ConverseTransform takes a slice of approximation coefficients, and
// a slice of detail coefficients reconstructs the original signal.
// It will return an error if the two supplied slices are not equal in length.
func ConverseTransform(app []float64, det []float64) ([]float64, error) {
	l := len(app)
	if len(det) != l {
		return []float64{}, errors.New("arguments should be the same length")
	}

	f := make([]float64, l*2)

	for i := 0; i < l; i++ {
		j := app[i]
		k := det[i]
		f[2*i] = (j + k) / math.Sqrt2
		f[2*i+1] = (j - k) / math.Sqrt2
	}

	return f, nil
}
