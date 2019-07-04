package main

import "testing"

func benchmark(b *testing.B, way int, micros int) {
	var process func(count, size int, dev, mean float64)
	mean := float64(micros)
	dev := float64(micros) / 10
	if way == 0 {
		process = pipelineWay
	} else {
		process = concurrentWay
	}
	for i := 0; i < b.N; i++ {
		process(100, 1000, dev, mean)
	}
}

func BenchmarkPipelineWay20(b *testing.B)  { benchmark(b, 0, 20) }
func BenchmarkPipelineWay70(b *testing.B)  { benchmark(b, 0, 70) }
func BenchmarkPipelineWay120(b *testing.B) { benchmark(b, 0, 120) }

func BenchmarkConcurrentWay20(b *testing.B)  { benchmark(b, 1, 20) }
func BenchmarkConcurrentWay70(b *testing.B)  { benchmark(b, 1, 70) }
func BenchmarkConcurrentWay120(b *testing.B) { benchmark(b, 1, 120) }
