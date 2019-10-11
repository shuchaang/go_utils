package go_pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}

func TestNewPool(t *testing.T) {

	gp, e := NewPool(&Options{
		MinWorker:    1,
		MaxWorker:    10,
		JobQueueSize: 10,
		IdleTimeOut:  1 * time.Second,
	})

	if e != nil {
		t.Log(e)
		return
	}
	for index := 0; index < 1000; index++ {
		wg.Add(1)
		idx := index
		gp.SupplyAsync(func() {
			fmt.Println(idx, time.Now())
			time.Sleep(1 * time.Second)
			wg.Done()
		})
	}
	wg.Wait()
}
