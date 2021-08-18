package http

import (
	"sync"
	"time"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/nivaldogmelo/web-api-tester/pkg/requests"
)

func createWorkerPool(noOfWorkers int, buffer chan root.Request) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, buffer)
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup, buffer chan root.Request) {
	for request := range buffer {
		requests.ExecuteRequest(request)
	}
	wg.Done()
}

func startJobs(done chan bool, buffer chan root.Request) {
	content, err := requests.GetAll()
	if err != nil {
		done <- false
	}
	for _, request := range content {
		buffer <- request
	}
	close(buffer)
	done <- true
}

func routine() {
	for {
		httpRequests := make(chan root.Request, 10)
		done := make(chan bool)
		go startJobs(done, httpRequests)
		noOfWorkers := 10
		createWorkerPool(noOfWorkers, httpRequests)
		<-done

		time.Sleep(300 * time.Second)
	}
}
