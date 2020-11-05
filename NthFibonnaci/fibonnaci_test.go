package main

import "testing"

func Test_fibo(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"initial-0", args{0}, 0},
		{"Input-10", args{10}, 55},
		{"Input-13", args{13}, 233},
		{"Input-14", args{14}, 377},
		{"Input-15", args{15}, 610},
		{"Input-44", args{44}, 701408733},
		{"Input-45", args{45}, 1134903170},
		{"Input-46", args{46}, 1836311903},
		{"Input-44", args{44}, 701408733},
		{"Input-(-8)", args{-8}, -1},
		{"Input-77", args{77}, 5527939700884757},
		{"Input-78", args{78}, 8944394323791464},
		{"Input-79", args{79}, 14472334024676221},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibo(tt.args.num); got != tt.want {
				t.Errorf("fibo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFibo(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibo(i)
	}
}

func Benchmark_fibo1(b *testing.B)  { benchmarkFibo(1, b) }
func Benchmark_fibo2(b *testing.B)  { benchmarkFibo(2, b) }
func Benchmark_fibo3(b *testing.B)  { benchmarkFibo(3, b) }
func Benchmark_fibo10(b *testing.B) { benchmarkFibo(10, b) }
func Benchmark_fibo45(b *testing.B) { benchmarkFibo(45, b) }
func Benchmark_fibo78(b *testing.B) { benchmarkFibo(78, b) }
