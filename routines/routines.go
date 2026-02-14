package routines

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func Routines() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\ntotal execution time: %v", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res++
	}
	wg.Done()
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nthe current results are: %v", results)
	m.RUnlock()
}
