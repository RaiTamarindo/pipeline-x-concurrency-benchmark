package main

import "testing"

func benchmark(b *testing.B, way int, millis int) {
	var process func(count, size int, dev, mean float64)
	mean := float64(millis)
	dev := float64(millis) / 10
	if way == 0 {
		process = pipelineWay
	} else {
		process = concurrentWay
	}
	for i := 0; i < b.N; i++ {
		process(100, 1000, dev, mean)
	}
}

func BenchmarkPipelineWay100(b *testing.B) { benchmark(b, 0, 100) }
func BenchmarkPipelineWay200(b *testing.B) { benchmark(b, 0, 200) }
func BenchmarkPipelineWay300(b *testing.B) { benchmark(b, 0, 300) }

func BenchmarkConcurrentWay100(b *testing.B) { benchmark(b, 0, 100) }
func BenchmarkConcurrentWay200(b *testing.B) { benchmark(b, 0, 200) }
func BenchmarkConcurrentWay300(b *testing.B) { benchmark(b, 0, 300) }
