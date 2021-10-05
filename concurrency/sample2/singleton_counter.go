package singleton_counter

// Channels
var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var doneCh chan bool = make(chan bool)

func init() {
	var count int

	// Do when program runs
	go func(addCh <-chan bool, getCountCh <-chan chan int, doneCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-doneCh:
				break
			}
		}
	}(addCh, getCountCh, doneCh)
}

// Singleton
type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	doneCh <- true
	defer func() {
		close(addCh)
		close(getCountCh)
		close(doneCh)
	}()
}
