package render

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

type BenchConfig struct {
	Width, Height   int
	TestSequential  bool
	TestWorkerPool  bool
	TestRowParallel bool
	WorkerCounts    []int
}

var (
	SmallImageConfig = BenchConfig{
		Width: 400, Height: 400,
		TestSequential: true, TestWorkerPool: true, TestRowParallel: true,
		WorkerCounts: []int{1, 2, 4, 8},
	}

	LargeImageConfig = BenchConfig{
		Width: 2048, Height: 2048,
		TestSequential: false, TestWorkerPool: true, TestRowParallel: true,
		WorkerCounts: []int{4, 8, 16},
	}

	HugeImageConfig = BenchConfig{
		Width: 8192, Height: 8192,
		TestSequential: false, TestWorkerPool: true, TestRowParallel: true,
		WorkerCounts: []int{8, 16},
	}
)

func BenchmarkSequential(b *testing.B) {
	config := SmallImageConfig
	for n := 0; n < b.N; n++ {
		GenerateImage(config.Width, config.Height)
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	config := SmallImageConfig
	for n := 0; n < b.N; n++ {
		GenerateImageWithWorkerPool(config.Width, config.Height)
	}
}

func BenchmarkRowParallel(b *testing.B) {
	config := SmallImageConfig
	for n := 0; n < b.N; n++ {
		GenerateImageInParallelByRow(config.Width, config.Height)
	}
}

func TestPerformanceComparison(t *testing.T) {
	config := SmallImageConfig

	fmt.Printf("\n🖥️  System: %d CPU cores\n", runtime.NumCPU())
	fmt.Printf("📏 Testing %dx%d images\n", config.Width, config.Height)
	fmt.Printf("%-20s %12s %12s\n", "Method", "Time", "Speedup")
	fmt.Printf("%-20s %12s %12s\n", "------", "----", "-------")

	var sequentialTime time.Duration

	if config.TestSequential {
		start := time.Now()
		GenerateImage(config.Width, config.Height)
		sequentialTime = time.Since(start)
		fmt.Printf("%-20s %12s %12s\n", "Sequential",
			sequentialTime.Round(time.Millisecond), "1.00x")
	}

	if config.TestWorkerPool {
		start := time.Now()
		GenerateImageWithWorkerPool(config.Width, config.Height)
		workerPoolTime := time.Since(start)

		speedup := "N/A"
		if config.TestSequential {
			speedup = fmt.Sprintf("%.2fx", float64(sequentialTime)/float64(workerPoolTime))
		}
		fmt.Printf("%-20s %12s %12s\n", "Worker Pool",
			workerPoolTime.Round(time.Millisecond), speedup)
	}

	if config.TestRowParallel {
		start := time.Now()
		GenerateImageInParallelByRow(config.Width, config.Height)
		rowParallelTime := time.Since(start)

		speedup := "N/A"
		if config.TestSequential {
			speedup = fmt.Sprintf("%.2fx", float64(sequentialTime)/float64(rowParallelTime))
		}
		fmt.Printf("%-20s %12s %12s\n", "Row Parallel",
			rowParallelTime.Round(time.Millisecond), speedup)
	}
}

func TestWorkerScaling(t *testing.T) {
	config := SmallImageConfig

	fmt.Printf("\n⚙️  Worker Scaling (%dx%d)\n", config.Width, config.Height)
	fmt.Printf("%-10s %12s %12s\n", "Workers", "Time", "Efficiency")
	fmt.Printf("%-10s %12s %12s\n", "-------", "----", "----------")

	var baselineTime time.Duration

	for i, workers := range config.WorkerCounts {
		start := time.Now()
		GenerateImageWithWorkerPoolCustom(config.Width, config.Height, workers)
		elapsed := time.Since(start)

		if i == 0 {
			baselineTime = elapsed
		}

		speedup := float64(baselineTime) / float64(elapsed)
		efficiency := speedup / float64(workers) * 100

		fmt.Printf("%-10d %12s %11.1f%%\n", workers,
			elapsed.Round(time.Millisecond), efficiency)
	}
}

func TestQuick(t *testing.T) {
	config := HugeImageConfig

	start := time.Now()
	GenerateImageWithWorkerPool(config.Width, config.Height)
	elapsed := time.Since(start)

	pixels := float64(config.Width * config.Height)
	pixelsPerSec := pixels / elapsed.Seconds()

	fmt.Printf("Generated %dx%d in %s (%.0f pixels/sec)\n",
		config.Width, config.Height, elapsed.Round(time.Millisecond), pixelsPerSec)
}
