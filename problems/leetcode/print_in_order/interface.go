package print_in_order

type Foo interface {
	first(printFirst func())
	second(printSecond func())
	third(printThird func())
}
