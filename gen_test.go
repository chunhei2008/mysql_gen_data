package main

import (
	"testing"
)

func Benchmark_genchar10(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_char(10)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genchar100(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_char(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genchar1000(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_char(1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genvarchar10(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_varchar(10, 20)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genvarchar100(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_varchar(100, 200)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genvarchar1000(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_varchar(1000, 2000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

/*
func Benchmark_autoincr(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_autoincr(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genvarint(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_varint(100, 200)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genint(b *testing.B) {
	gen := NewGen()
	genchar := gen.gen_int(100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-genchar
	}
	b.StopTimer()
}

func Benchmark_genseed(b *testing.B) {
	gen := NewGen()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-gen.seed
	}
	b.StopTimer()
}
*/
