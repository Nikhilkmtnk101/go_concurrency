package print_in_order

import "sync"

type FooMutex struct {
	mutex2 sync.Mutex
	mutex3 sync.Mutex
}

func NewFooMutex() *FooMutex {
	f := &FooMutex{}
	f.mutex2.Lock()
	f.mutex3.Lock()
	return f
}

func (f *FooMutex) first(printFirst func()) {
	printFirst()
	f.mutex2.Unlock()
}

func (f *FooMutex) second(printSecond func()) {
	f.mutex2.Lock()
	printSecond()
	f.mutex2.Unlock()
	f.mutex3.Unlock()
}

func (f *FooMutex) third(printThird func()) {
	f.mutex3.Lock()
	printThird()
	f.mutex3.Unlock()
}
