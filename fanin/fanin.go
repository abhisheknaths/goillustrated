package fanin

func Merge(streams ...<-chan string) <-chan string {
	out := make(chan string)
	done := make(chan struct{})
	for _, s := range streams {
		go func(x <-chan string) {
			defer func() {
				done <- struct{}{}
			}()
			for data := range x {
				out <- data
			}
		}(s)
	}
	go func() {
		for i := 0; i < len(streams); i++ {
			<-done
		}
		close(out)
	}()
	return out
}
