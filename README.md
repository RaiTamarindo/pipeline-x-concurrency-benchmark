# Pipeline x Concurrency Benchmark

This project has benchmarks to measure performance metrics to answer the following question:

> Is go routines better used as stages of one pipeline or as each with it's own flow?

### How this benchmark was made

Problems *stream* like and that can be splitted on stages are good candidates to solve with machine parallelism.

The mocked problem used in this benchmark is a series of jobs, where each job consists of an array of integers to be incremented and an IO call. To do this I defined a `worker` function that do this job:

```go
func worker(job []int, dev, mean float64) []int {
	timeout := time.Duration(math.Round(rand.NormFloat64()*dev + mean))
	for i := range job {
		job[i]++
	}
	time.Sleep(timeout * time.Millisecond)

	return job
}
```

The worker simulates IO calls by sleeping for certain period. That period is generated from normal distribution given a mean and its deviation. In this benchmark, I've tested each approach (pipeline and concurrent) with `20 ms`, `70 ms` and `120 ms` of average "IO" time with 10% of deviation.

The entire flow consists on submit the job to four stages of work, each implemented by `worker` function

- Pipeline approach

  It was created four goroutines, one for each stage of work. The fist one simulates pooling of jobs by creating empty arrays and the following stages uses a channel to get job from previous stage and another to produce for the next.

- Concurrent approach

  It was created four goroutines, but here each one runs all stages in a quarter of all jobs.

### How to start

`go test -bench=. -benchmem`

### Results on my machine

| Strategy | Work time (ms) | Execution Time (ms) | Memory (B) |
|---|--:|--:|--:|
| Pipeline | 20 | 2259802.72  | 829504 |
| Pipeline | 70 |  7888779.38  | 821472 |
| Pipeline | 120 |  13518403.76 | 823712 |
| Concurrent | 20 |  2280110.63 | 820992 |
| Concurrent | 70 | 7910958.20  | 822144 |
| Concurrent | 120 | 13590690.06  | 820128 |