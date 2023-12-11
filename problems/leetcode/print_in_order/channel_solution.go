package print_in_order

type FooChannel struct {
	ch2 chan bool
	ch3 chan bool
}

func NewFooChannel() *FooChannel {
	f := &FooChannel{
		ch2: make(chan bool),
		ch3: make(chan bool),
	}
	return f
}

func (f *FooChannel) first(printFirst func()) {
	printFirst()
	f.ch2 <- true
}

func (f *FooChannel) second(printSecond func()) {
	<-f.ch2
	printSecond()
	f.ch3 <- true
}

func (f *FooChannel) third(printThird func()) {
	<-f.ch3
	printThird()
}
