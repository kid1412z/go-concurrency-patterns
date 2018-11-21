## Go Concurrency Patterns

Golang concurrency patterns from Rob Pike Google I/O 2012.

### Concurrency is NOT parallelism

Concurrency is not parallelism, although it enables parallelism.
See [golang.org/s/concurrency-is-not-parallelism](golang.org/s/concurrency-is-not-parallelism)

### Distinction

Go is the latest on the Newsqueak-Alef-Limbo branch, distinguished by first-class channels.

### Main Points

|Point|Files|
|:--- |:----|
|Goroutine basics| [goroutine1](https://github.com/kid1412z/go-concurrency-patterns/blob/master/goroutine1.go) [goroutine2](https://github.com/kid1412z/go-concurrency-patterns/blob/master/goroutine2.go) |
|Channel basics|[channel1.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/channel1.go)  [channel2.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/channel2.go)  |
|Generator and multiplexing| [](https://github.com/kid1412z/go-concurrency-patterns/blob/master/generator1.go) [generator2.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/generator2.go) [generator3.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/generator3.go) [generator4.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/generator4.go) |
|Channel chain| [daisychain.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/daisychain.go) |
|Select basics| [select1.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/select1.go) [select2](https://github.com/kid1412z/go-concurrency-patterns/blob/master/select2.go) [select3.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/select3.go) [select4.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/select4.go)|
|Concurrent Requests| [google1.go](https://github.com/kid1412z/go-concurrency-patterns/blob/master/google1.go) [google2](https://github.com/kid1412z/go-concurrency-patterns/blob/master/google2.go) [google3](https://github.com/kid1412z/go-concurrency-patterns/blob/master/google3.go) [google4](https://github.com/kid1412z/go-concurrency-patterns/blob/master/google4.go) |


## Resouces

* [https://www.youtube.com/watch?v=f6kdp27TYZs](https://www.youtube.com/watch?v=f6kdp27TYZs)
* [https://talks.golang.org/2012/concurrency.slide](https://www.youtube.com/watch?v=f6kdp27TYZs)
* [http://golang.org/s/concurrency-is-not-parallelism](http://golang.org/s/concurrency-is-not-parallelism)

