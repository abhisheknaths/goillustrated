package fanout

import (
	"context"
	"fmt"
	"sync"

	"github.com/abhisheknaths/goillustrated/utils"
	"github.com/bxcodec/faker/v4"
)

// Split spawns workers to process in channel
// Keeps max number of workers to n at any time
func Split(in <-chan string, n int) (<-chan struct{}, *utils.AtomicCounter) {
	token := make(chan struct{}, n)
	done := make(chan struct{})
	var ok bool = true
	wg := &sync.WaitGroup{}
	ac := new(utils.AtomicCounter)
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for ok {
			select {
			case <-ctx.Done():
				ok = false
			case token <- struct{}{}:
				// faker.Name is not thread safe, moved out from inside of worker
				randomWorkerName := faker.Name()
				wg.Add(1)
				go worker(randomWorkerName, in, token, cancel, wg, ac)
			}
		}
		// wait for any remaining go routines to complete
		wg.Wait()
		close(done)
	}()
	return done, ac
}

// worker does not wait on in channel until it is closed
// but exists after reading one time
func worker(workerName string, in <-chan string, token chan struct{}, cancel context.CancelFunc, wg *sync.WaitGroup, ac *utils.AtomicCounter) {
	defer func() {
		wg.Done()
		<-token
	}()
	val, ok := <-in
	if !ok {
		// unlike closing an already closed channel, calling cancel() multiple times does not panic
		cancel()
	} else {
		ac.Add()
		fmt.Println(fmt.Sprintf(`%s processed by random worker: %s`, val, workerName))
	}
}
