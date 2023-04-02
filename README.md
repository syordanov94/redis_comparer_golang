# Golang-Redis Library Comparer

This is a project, written in Go, that performs some basic perfomance comparisons between 2 of the most used Redis modules used in Golang: [redigo](https://github.com/gomodule/redigo) and [goredis](https://github.com/redis/go-redis)

The benchmarks performed compare the follwing operations:

- SET operation for a string using each connector
- GET operation for a string using each connector
- Both SET and GET operations for a string using each connector

**Note**
Bear in mind that this ONLY tests the Redis functions above and with a limited secuential test. The results between each connector may differ depending on the scenario one uses them in.

## Prerequisites

- Golang installed
- _Recomended but not mandatory_ VS Code or a similiar IDE 

## How to install and Run the project

- First you will have to clone the project from this github repository

```bash
git clone https://github.com/syordanov94/redis_comparer_golang.git
```

- Once cloned you will need to create a specific file calle _config.json_ in the root folder that will contain the Redis DB access configuration that you will need in order to access the DB. This config file will have the following structure:

```json
{
    "redis_host":"",
    "redis_port":00000, 
    "redis_password":""
}
```

- Download and update all module dependencies

```bash
go mod tidy
```

- Once upgraded, you can run the _main.go_ file that performs a simple SET and GET, to verify that the project is running correctly

```bash
go run main.go
```


## How to test the project

Now that the project is configured and ready, you can run the benchmarks defined inside. You can use a command like the following:

```bash
go test -bench=. -benchtime=20x -benchmem
```

Once ran, this will produce an output like the following:

![myimage-alt-tag](resources/benchmark_output.png)