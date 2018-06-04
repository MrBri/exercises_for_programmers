package main

import (
	"reflect"
	"testing"

	"github.com/mrbri/tip"
)

type testpair struct {
	values []float64
	result []float64
}

var tests = []testpair{
	{[]float64{10, 15}, []float64{1.5, 11.5}},
	{[]float64{11.25, 15}, []float64{1.69, 12.94}},
}

func TestTip(t *testing.T) {
	for _, pair := range tests {
		v := tip.Tip(pair.values)
		if !reflect.DeepEqual(v, pair.result) {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

// package math

// import "testing"

// type testpair struct {
// 	values []float64
// 	result []float64
// }

// var tests = []testpair{
// 	{[]float64{10, 15}, []float64{1.5, 11.5}},
// 	{ []float64{1,1,1,1,1,1}, 1 },
// 	{ []float64{-1,1}, 0 },
// }

// func TestAverage(t *testing.T) {
// 	for _, pair := range tests {
// 		v := Average(pair.values)
// 		if v != pair.average {
// 			t.Error(
// 				"For", pair.values,
// 				"expected", pair.average,
// 				"got", v,
// 			)
// 		}
// 	}
// }
