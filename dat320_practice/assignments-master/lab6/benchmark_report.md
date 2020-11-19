# Benchmark Report

## CPU and Memory Benchmarks for all Three Stacks

```console
$ go test -v -run none -bench Benchmark -memprofilerate=1 -benchmem
TODO(student)
```

1. How much faster than the slowest is the fastest stack?
    - [ ] a) 2x-3x
    - [ ] b) 3x-4x
    - [ ] c) 6x-7x
    - [ ] d) 10x-11x

2. Which stack requires the most allocated memory?
    - [ ] a) CspStack
    - [ ] b) SliceStack
    - [ ] c) SafeStack
    - [ ] d) UnsafeStack

3. Which stack requires the least amount of allocated memory?
    - [ ] a) CspStack
    - [ ] b) SliceStack
    - [ ] c) SafeStack
    - [ ] d) UnsafeStack

## CPU Profile of BenchmarkCspStack

```console
$ go test -v -run none -bench BenchmarkCspStack -cpuprofile=csp-stack.prof
TODO(student)
$ go tool pprof csp-stack.prof
TODO(student)
```

4. Which function accounts for the most CPU usage?
    - [ ] a) `runtime.pthread_mutex_lock`
    - [ ] b) `dat320/lab6/stack.(*CspStack).run`
    - [ ] c) `runtime.pthread_cond_wait`
    - [ ] d) `runtime.pthread_cond_signal`

5. From this top 10 listing, what can you say about the underlying implementation of a CSP-based stack?
    - [ ] a) It is implemented using condition variables and locks
    - [ ] b) It is implemented using monitors
    - [ ] c) It is implemented using locks
    - [ ] d) It is implemented using semaphores

## Memory Profile of BenchmarkSafeStack

```console
$ go test -v -run none -bench BenchmarkSafeStack -memprofile=safe-stack.prof -tags solution
TODO(student)
$ go tool pprof safe-stack.prof
TODO(student)
```

6. Which function accounts for all memory allocations in the `SafeStack` implementation?
    - [ ] a) `Size`
    - [ ] b) `NewSafeStack`
    - [ ] c) `Push`
    - [ ] d) `Pop`

7. Which line in `SafeStack` does the actual memory allocation?
    - [ ] a) `type SafeStack struct {`
    - [ ] b) `ss.top = &Element{value, ss.top}`
    - [ ] c) `value, ss.top = ss.top.value, ss.top.next`
    - [ ] d) `top  *Element`

## Visualizing the SafeStack Call Graph

Examine the call graph visualization and answer the questions below.

8. What is the root node of the call graph that ends with `runtime.pthread_cond_wait`?
    - [ ] a) `runtime.park_m`
    - [ ] b) `runtime.schedule`
    - [ ] c) `runtime.mstart`
    - [ ] d) `runtime.mcall`

9. Which of the following consume the most CPU time?
    - [ ] a) `runtime.park_m`
    - [ ] b) `runtime.usleep`
    - [ ] c) `runtime.mstart`
    - [ ] d) `runtime.lock`
