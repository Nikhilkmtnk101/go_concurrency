package print_in_order

import "sync"

type FooCond struct {
	mutex2 sync.Mutex
	mutex3 sync.Mutex
}

func NewFooCond() *FooCond {
	f := &FooCond{
		mutex2: sync.Mutex{},
		mutex3: sync.Mutex{},
	}
	f.mutex2.Lock()
	f.mutex3.Lock()
	return f
}

func (f *FooCond) first(printFirst func()) {
	printFirst()
}

func (f *FooCond) second(printSecond func()) {
	printSecond()
}

func (f *FooCond) third(printThird func()) {
	printThird()
}
