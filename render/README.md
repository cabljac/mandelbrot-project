
```sh
go test -run=TestPerformanceComparison  # Compare methods
go test -run=TestWorkerScaling          # Test worker counts
go test -bench=BenchmarkWorkerPool       # Benchmark worker pool
go test -run=TestQuick                   # Quick performance check
```