package print_in_order

import "sync"

type FooCond struct {
	val  int
	cond *sync.Cond
}

func NewFooCond() *FooCond {
	f := &FooCond{
		val:  1,
		cond: sync.NewCond(&sync.Mutex{}),
	}
	f.cond.L.Lock()
	return f
}

func (f *FooCond) first(printFirst func()) {
	printFirst()
	f.val = 2
	f.cond.L.Unlock()
	f.cond.Signal()
}

func (f *FooCond) second(printSecond func()) {
	f.cond.L.Lock()
	for f.val != 2 {
		f.cond.Wait()
	}
	printSecond()
	f.val = 3
	f.cond.L.Unlock()
	f.cond.Signal()
}

func (f *FooCond) third(printThird func()) {
	f.cond.L.Lock()
	for f.val != 3 {
		f.cond.Wait()
	}
	printThird()
	f.cond.L.Unlock()
}
