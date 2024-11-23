package math

import "testing"

func TestAverage(t *testing.T) {
	average := Average([]float64{2, 1})
	if average != 1.5 {
		t.Error("Expected 1.5 got ", average)
	}
}

type testValues struct {
	values []float64
	average float64
}

var testpairs = []testValues {
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 5}, 3},
	{[]float64{1, 3}, 2},
}
func TestAverageWithValues(t *testing.T) {
	for _, pair := range testpairs {
		avg := Average(pair.values)
		if avg != pair.average {
			t.Error("For ", pair.values, ", Expected ", pair.average, ", Got ", avg)
		}
	}
}