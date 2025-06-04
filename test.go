package main

import (
	"fmt"
	"runtime"
	"sync"
	"../king"
)

func main() {
	runtime.GOMAXPROCS(1)
	var N int
	fmt.Scanf("%d\n", &N)
	println(N)
	firstInput := make(chan []byte)
	prevOutput := firstInput
	writeChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(N)
	runner := func(candidate king.ICandidate) {
		defer wg.Done()
		leader := candidate.SelectLeader()
		writeChannel <- leader
	}
	co:=0
	for it := 0; it < N-1; it++ {
		nextOutput := make(chan []byte)
		var input <-chan []byte
		var output chan<- []byte
		input = prevOutput
		output = nextOutput
		var value int
		fmt.Scanf("%d", &value)
		if(value>co){
			co = value
		}
		go runner(king.NewCandidate(value, input, output))
		prevOutput = nextOutput
	}
	var value int
	fmt.Scanf("%d", &value)
	if(value>co){
		co = value
	}
	input := prevOutput
	output := firstInput
	go runner(king.NewCandidate(value, input, output))
	for it := 0; it < N; it++ {
		leader := <-writeChannel
		if leader != co {
			println("err", leader, N,co)
		}
	}
	wg.Wait()
}
