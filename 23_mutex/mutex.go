package main

import "sync"

type Post struct {
	View int
	mu   sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.View++
}

func main() {

	var wg sync.WaitGroup

	myPost := Post{View: 0}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
}
