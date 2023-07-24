package lesson7

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_sortString(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Test Case 1",
			args: []string{"banana", "apple", "orange"},
			want: []string{"apple", "banana", "orange"},
		},
		{
			name: "Test Case 2",
			args: []string{"яблоко", "груша", "апельсин"},
			want: []string{"апельсин", "груша", "яблоко"},
		},
		{
			name: "Test Case 3",
			args: []string{"", "hello", "world", ""},
			want: []string{"", "", "hello", "world"},
		},
		{
			name: "Test Case 4",
			args: []string{""},
			want: []string{""},
		},
		{
			name: "Test Case 5",
			args: []string{"banana", "яблоко", "orange"},
			want: []string{"banana", "orange", "яблоко"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortINT(t *testing.T) {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

	numbers = sortINT(numbers)

	for i := range numbers {
		if numbers[i] != expected[i] {
			t.Errorf("Expected %d at index %d, but got %d", expected[i], i, numbers[i])
		}
	}
}

func BenchmarkSortInts(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortINT(data)
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	data := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Float64()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortfloat64(data)
	}
}

/*
BenchmarkSortInts
BenchmarkSortInts-8       	   10000	    279326 ns/op
BenchmarkSortFloat64s
BenchmarkSortFloat64s-8   	   10000	    304615 ns/op
279326 ns/op и 304615 ns/op - это среднее время выполнения одной операции (итерации) функции в наносекундах.
Таким образом, в среднем, sort.Ints() выполняется за примерно 279326 наносекунду, а sort.Float64s() - за примерно 304615 наносекунды.

Эти результаты позволяют сравнивать производительность двух функций и определить, какая из них выполняется быстрее.
В данном случае, функция sort.Ints() быстрее функции sort.Float64s().
*/
