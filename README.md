## The Problem

Read random integers, in the range of 1 to 10000, from a channel, and distributes those integers to multiple go-routines. Each go-routine should sleep for the millisecond duration matching the integer; so if the goroutine receives a value of 100; it will sleep for 100 milliseconds. The number of go routines should be a constant, such that it can be easily changed; and the number of integers that will be generated should be a constant value too. Wait for all of the go-routines to finish processing before existing.

## Design

The task suits _worker pool_ go concurrency pattern: from one side we have a random in channel from the other side we have number of goroutines(_workers_) that read int from the channel and do work(sleep) 

## Run
```
go run main.go -w 8 -r 100
```
Arguments:
 - -w number of goroutines(workers). 4 is default
 - -r upper limit of range [1, r]. 10000 is default.
