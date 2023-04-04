# Golang-Redis Library Comparer

This is a project, written in Go, that performs some basic perfomance comparisons between 2 of the most used Redis modules used in Golang: [redigo](https://github.com/gomodule/redigo) and [goredis](https://github.com/redis/go-redis)

The benchmarks performed compare the follwing operations:

- SET operation for a string using each connector
- GET operation for a string using each connector
- Both SET and GET operations for a string using each connector

**Note**
Bear in mind that this ONLY tests the Redis functions above and with a limited secuential test. The results between each connector may differ depending on the scenario one uses them in.

## Prerequisites

- Golang 1.19+ installed
- Docker installed
- Docker Compose installed
- _Recomended but not mandatory_ VS Code or a similiar IDE 

## How to install and Run the project

- First you will have to clone the project from this github repository

```bash
git clone https://github.com/syordanov94/redis_comparer_golang.git
```

- Download dependencies

```bash
go mod vendor
```

## How to test the project
Use Makefile:
```bash
make env-up
make test
make bench
make env-down
```
Or commands:
```bash
docker-compose up -d
docker exec rediscomparer-go-app go test ./... -v -count=1
docker exec rediscomparer-go-app go test ./... -v -bench='(Redigo|Goredis)' -benchmem -benchtime=1000000x -count=10 | tee bench-redis-1000000x-sequence.txt
benchstat bench-redis-1000000x-sequence.txt
docker-compose down --remove-orphans -v
```
Result:
```text
=== RUN   TestRedigo
--- PASS: TestRedigo (0.00s)
=== RUN   TestGoredis
--- PASS: TestGoredis (0.00s)
goos: linux
goarch: amd64
pkg: rediscomparer
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkRedigoGet
BenchmarkRedigoGet-12           	   10000	     26765 ns/op	      88 B/op	       5 allocs/op
BenchmarkRedigoSet
BenchmarkRedigoSet-12           	   10000	     28315 ns/op	      64 B/op	       3 allocs/op
BenchmarkRedigoGetAndSet
BenchmarkRedigoGetAndSet-12     	   10000	     57800 ns/op	     152 B/op	       8 allocs/op
BenchmarkGoredisGet
BenchmarkGoredisGet-12          	   10000	     30770 ns/op	     200 B/op	       8 allocs/op
BenchmarkGoredisSet
BenchmarkGoredisSet-12          	   10000	     33330 ns/op	     265 B/op	      10 allocs/op
BenchmarkGoredisGetAndSet
BenchmarkGoredisGetAndSet-12    	   10000	     62447 ns/op	     465 B/op	      18 allocs/op
PASS
ok  	rediscomparer	2.407s
```
```text
name                 time/op
RedigoGet-12         30.8µs ±18%
RedigoSet-12         30.1µs ±14%
RedigoGetAndSet-12   59.3µs ±13%
GoredisGet-12        32.6µs ± 8%
GoredisSet-12        33.6µs ± 7%
GoredisGetAndSet-12  70.7µs ±11%

name                 alloc/op
RedigoGet-12          88.0B ± 0%
RedigoSet-12          64.0B ± 0%
RedigoGetAndSet-12     152B ± 0%
GoredisGet-12          200B ± 0%
GoredisSet-12          265B ± 0%
GoredisGetAndSet-12    465B ± 0%

name                 allocs/op
RedigoGet-12           5.00 ± 0%
RedigoSet-12           3.00 ± 0%
RedigoGetAndSet-12     8.00 ± 0%
GoredisGet-12          8.00 ± 0%
GoredisSet-12          10.0 ± 0%
GoredisGetAndSet-12    18.0 ± 0%
```
