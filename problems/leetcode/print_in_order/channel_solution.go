package print_in_order

import "sync"

type FooChannel struct {
	mutex2 sync.Mutex
	mutex3 sync.Mutex
}

func NewFooChannel() *FooChannel {
	f := &FooChannel{
		mutex2: sync.Mutex{},
		mutex3: sync.Mutex{},
	}
	f.mutex2.Lock()
	f.mutex3.Lock()
	return f
}

func (f *FooChannel) first(printFirst func()) {
	printFirst()
}

func (f *FooChannel) second(printSecond func()) {
	printSecond()
}

func (f *FooChannel) third(printThird func()) {
	printThird()
}
