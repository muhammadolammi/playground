package chwg

import (
	"fmt"
	"sync"
	"time"
)

func doWork(t time.Duration, id int, resultch chan (string)) {
	fmt.Println("Work start.....")
	time.Sleep(t)
	resultch <- fmt.Sprintf("work %d done.\n", id)
}

func runWorks(works []Work, resultch chan (string)) {

	for _, work := range works {
		go doWork(work.t, work.id, resultch)
	}
}
func printResults(resultch chan (string), wg *sync.WaitGroup, worksCount int) {
	wg.Add(worksCount)
	go func() {
		for res := range resultch {
			fmt.Println(res)
			wg.Done()
		}
	}()
	wg.Wait()
	close(resultch)
}

func Chwg() {
	start := time.Now()
	resultch := make(chan (string))
	works := []Work{
		{
			t:  time.Second * 2,
			id: 1,
		},
		{
			t:  time.Second * 4,
			id: 2,
		},
	}
	runWorks(works, resultch)
	wg := sync.WaitGroup{}
	printResults(resultch, &wg, len(works))
	fmt.Printf("works took %v seconds.\n", time.Since(start))
}
