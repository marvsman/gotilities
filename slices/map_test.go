package slices

import (
	"reflect"
	"testing"
)

func Test_Map(t *testing.T) {
	type dataType struct {
		Name      string
		Age       int
		MaxEnergy int
	}

	type args[T any, V any] struct {
		fn func(T) V
	}
	type testCase[T any, V any] struct {
		name       string
		args       args[T, V]
		want       []V
		shouldFail bool
	}

	testData := []dataType{
		{Name: "Alice", Age: 30, MaxEnergy: 100},
		{Name: "Bob", Age: 25, MaxEnergy: 90},
		{Name: "Charlie", Age: 35, MaxEnergy: 110},
	}

	test1 := testCase[dataType, int]{
		name: "Test age map",
		args: args[dataType, int]{
			fn: func(dataType dataType) int {
				return dataType.Age
			},
		},
		want:       []int{30, 25, 35},
		shouldFail: false,
	}
	test2 := testCase[dataType, int]{
		name: "Test energy map",
		args: args[dataType, int]{
			fn: func(dataType dataType) int {
				return dataType.MaxEnergy
			},
		},
		want:       []int{100, 90, 110},
		shouldFail: false,
	}
	test3 := testCase[dataType, string]{
		name: "Test names",
		args: args[dataType, string]{
			fn: func(dataType dataType) string {
				return dataType.Name
			},
		},
		want:       []string{"Alice", "Bob", "Charlie"},
		shouldFail: false,
	}
	test4 := testCase[dataType, int]{
		name: "Test names",
		args: args[dataType, int]{
			fn: func(dataType dataType) int {
				return dataType.MaxEnergy
			},
		},
		want:       []int{110, 90, 90},
		shouldFail: true,
	}

	t.Run(test1.name, func(t *testing.T) {
		got := Map(testData, test1.args.fn)
		if !reflect.DeepEqual(got, test1.want) && !test1.shouldFail {
			t.Errorf("Map() = %v, want %v", got, test1.want)
		}
	})
	t.Run(test2.name, func(t *testing.T) {
		got := Map(testData, test2.args.fn)
		if !reflect.DeepEqual(got, test2.want) && !test2.shouldFail {
			t.Errorf("Map() = %v, want %v", got, test2.want)
		}
	})
	t.Run(test3.name, func(t *testing.T) {
		got := Map(testData, test3.args.fn)
		if !reflect.DeepEqual(got, test3.want) && !test3.shouldFail {
			t.Errorf("Map() = %v, want %v", got, test3.want)
		}
	})
	t.Run(test3.name, func(t *testing.T) {
		got := Map(testData, test4.args.fn)
		if !reflect.DeepEqual(got, test4.want) && !test4.shouldFail {
			t.Errorf("Map() = %v, want %v", got, test4.want)
		}
	})
}

func Test_Sum(t *testing.T) {
	testCases := []struct {
		name  string
		input any
		fn    any
		want  any
	}{
		{
			name:  "Sum of ints",
			input: []int{1, 2, 3, 4, 5},
			fn:    func(x int) int { return x },
			want:  15,
		},
		{
			name:  "Sum of floats",
			input: []float64{1.5, 2.5, 3.0},
			fn:    func(x float64) float64 { return x },
			want:  7.0,
		},
		{
			name:  "Sum of empty ints",
			input: []int{},
			fn:    func(x int) int { return x },
			want:  0,
		},
	}
	for _, tc := range testCases {
		switch input := tc.input.(type) {
		case []int:
			got := Sum(input, tc.fn.(func(int) int))
			if got != tc.want {
				t.Errorf("%s: Sum() = %v, want %v", tc.name, got, tc.want)
			}
		case []float64:
			got := Sum(input, tc.fn.(func(float64) float64))
			if got != tc.want {
				t.Errorf("%s: Sum() = %v, want %v", tc.name, got, tc.want)
			}
		}
	}
}

func Test_Avg(t *testing.T) {
	testCases := []struct {
		name  string
		input any
		fn    any
		want  any
	}{
		{
			name:  "Avg of ints",
			input: []int{2, 4, 6, 8},
			fn:    func(x int) int { return x },
			want:  5,
		},
		{
			name:  "Avg of floats",
			input: []float64{2.0, 4.0, 6.0},
			fn:    func(x float64) float64 { return x },
			want:  4.0,
		},
		{
			name:  "Avg of empty ints",
			input: []int{},
			fn:    func(x int) int { return x },
			want:  0,
		},
	}
	for _, tc := range testCases {
		switch input := tc.input.(type) {
		case []int:
			got := Avg(input, tc.fn.(func(int) int))
			if got != tc.want {
				t.Errorf("%s: Avg() = %v, want %v", tc.name, got, tc.want)
			}
		case []float64:
			got := Avg(input, tc.fn.(func(float64) float64))
			if got != tc.want {
				t.Errorf("%s: Avg() = %v, want %v", tc.name, got, tc.want)
			}
		}
	}
}

func Test_Min(t *testing.T) {
	testCases := []struct {
		name  string
		input any
		fn    any
		want  any
	}{
		{
			name:  "Min of ints",
			input: []int{5, 2, 8, 1, 9},
			fn:    func(x int) int { return x },
			want:  1,
		},
		{
			name:  "Min of floats",
			input: []float64{3.5, 2.1, 4.8},
			fn:    func(x float64) float64 { return x },
			want:  2.1,
		},
		{
			name:  "Min of empty ints",
			input: []int{},
			fn:    func(x int) int { return x },
			want:  0,
		},
	}
	for _, tc := range testCases {
		switch input := tc.input.(type) {
		case []int:
			got := Min(input, tc.fn.(func(int) int))
			if got != tc.want {
				t.Errorf("%s: Min() = %v, want %v", tc.name, got, tc.want)
			}
		case []float64:
			got := Min(input, tc.fn.(func(float64) float64))
			if got != tc.want {
				t.Errorf("%s: Min() = %v, want %v", tc.name, got, tc.want)
			}
		}
	}
}

func Test_Max(t *testing.T) {
	testCases := []struct {
		name  string
		input any
		fn    any
		want  any
	}{
		{
			name:  "Max of ints",
			input: []int{5, 2, 8, 1, 9},
			fn:    func(x int) int { return x },
			want:  9,
		},
		{
			name:  "Max of floats",
			input: []float64{3.5, 2.1, 4.8},
			fn:    func(x float64) float64 { return x },
			want:  4.8,
		},
		{
			name:  "Max of empty ints",
			input: []int{},
			fn:    func(x int) int { return x },
			want:  0,
		},
	}
	for _, tc := range testCases {
		switch input := tc.input.(type) {
		case []int:
			got := Max(input, tc.fn.(func(int) int))
			if got != tc.want {
				t.Errorf("%s: Max() = %v, want %v", tc.name, got, tc.want)
			}
		case []float64:
			got := Max(input, tc.fn.(func(float64) float64))
			if got != tc.want {
				t.Errorf("%s: Max() = %v, want %v", tc.name, got, tc.want)
			}
		}
	}
}

func Test_Reduce(t *testing.T) {
	testCases := []struct {
		name    string
		input   any
		fn      any
		initial any
		want    any
	}{
		{
			name:    "Reduce sum ints",
			input:   []int{1, 2, 3, 4},
			fn:      func(acc, x int) int { return acc + x },
			initial: 0,
			want:    10,
		},
		{
			name:    "Reduce product floats",
			input:   []float64{1.0, 2.0, 3.0},
			fn:      func(acc, x float64) float64 { return acc * x },
			initial: 1.0,
			want:    6.0,
		},
		{
			name:    "Reduce concat strings",
			input:   []string{"a", "b", "c"},
			fn:      func(acc, x string) string { return acc + x },
			initial: "",
			want:    "abc",
		},
		{
			name:    "Reduce empty ints",
			input:   []int{},
			fn:      func(acc, x int) int { return acc + x },
			initial: 42,
			want:    42,
		},
	}
	for _, tc := range testCases {
		switch input := tc.input.(type) {
		case []int:
			got := Reduce(input, tc.fn.(func(int, int) int), tc.initial.(int))
			if got != tc.want {
				t.Errorf("%s: Reduce() = %v, want %v", tc.name, got, tc.want)
			}
		case []float64:
			got := Reduce(input, tc.fn.(func(float64, float64) float64), tc.initial.(float64))
			if got != tc.want {
				t.Errorf("%s: Reduce() = %v, want %v", tc.name, got, tc.want)
			}
		case []string:
			got := Reduce(input, tc.fn.(func(string, string) string), tc.initial.(string))
			if got != tc.want {
				t.Errorf("%s: Reduce() = %v, want %v", tc.name, got, tc.want)
			}
		}
	}
}
