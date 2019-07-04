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
	time.Sleep(timeout * time.Microsecond)

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
			stage1 <- worker(job, dev, mean)
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

func concurrentWay(count, size int, dev, mean float64) {
	stage1 := make(chan []int)
	stage2 := make(chan []int)
	stage3 := make(chan []int)
	stage4 := make(chan []int)

	go func() {
		for i := 0; i < count; i += 4 {
			j0 := make([]int, size)
			j1 := worker(j0, dev, mean)
			j2 := worker(j1, dev, mean)
			j3 := worker(j2, dev, mean)
			worker(j3, dev, mean)
		}
		close(stage1)
	}()

	go func() {
		for i := 1; i < count; i += 4 {
			j0 := make([]int, size)
			j1 := worker(j0, dev, mean)
			j2 := worker(j1, dev, mean)
			j3 := worker(j2, dev, mean)
			worker(j3, dev, mean)
		}
		close(stage2)
	}()

	go func() {
		for i := 2; i < count; i += 4 {
			j0 := make([]int, size)
			j1 := worker(j0, dev, mean)
			j2 := worker(j1, dev, mean)
			j3 := worker(j2, dev, mean)
			worker(j3, dev, mean)
		}
		close(stage3)
	}()

	go func() {
		for i := 3; i < count; i += 4 {
			j0 := make([]int, size)
			j1 := worker(j0, dev, mean)
			j2 := worker(j1, dev, mean)
			j3 := worker(j2, dev, mean)
			worker(j3, dev, mean)
		}
		close(stage4)
	}()

	<-stage1
	<-stage2
	<-stage3
	<-stage4
}

func main() {}
