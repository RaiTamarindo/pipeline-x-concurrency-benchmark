package main

import (
	"math"
	"math/rand"
	"time"
)

func worker(job []int, dev, mean float64) []int {
	timeout := time.Duration(math.Round(rand.NormFloat64()*dev + mean))
	for i := range job {
		job[i]++
	}
	time.Sleep(timeout * time.Millisecond)

	return job
}

func pipelineWay(count, size int, dev, mean float64) {
	stage1 := make(chan []int)
	stage2 := make(chan []int)
	stage3 := make(chan []int)
	stage4 := make(chan []int)

	go func() {
		for i := 0; i < count; i++ {
			job := make([]int, size)
			stage1 <- worker(job[:], dev, mean)
		}
		close(stage1)
	}()

	go func() {
		for job := range stage1 {
			stage2 <- worker(job, dev, mean)
		}
		close(stage2)
	}()

	go func() {
		for job := range stage2 {
			stage3 <- worker(job, dev, mean)
		}
		close(stage3)
	}()

	go func() {
		for job := range stage3 {
			worker(job, dev, mean)
		}
		close(stage4)
	}()

	<-stage4
}

func main() {

	pipelineWay(10, 1000, 10, 100)

}
