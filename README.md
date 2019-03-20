# AnyThingMock
go impl mock any variable, i use this on unit test, and love it so much. pls may sure mock type of variable is same as origin one. otherwise it will panic.

https://godoc.org/github.com/Petelin/AnyThingMock

## before this package
```golang
old := symbol
symbol = "Destroy World!"
defer func() {
	symbol = old
}()

// new way to write
rec := new(Mock).On(&symbol, "Destroy World!")
defer rec()
```

## useage

```golang
package mock

import "fmt"
import "testing"

var foo = func() {
	fmt.Println("i am func foo")
}

var bar = func() {
	fmt.Println("i am func bar")
}

var symbol string = "Hello World!"

func printAll() {
	fmt.Println(symbol)
	foo()
	bar()
	fmt.Println("---------")
}

func TestMockOne(t *testing.T)  {
	fmt.Println("raw value: ", symbol)

	rec := new(Mock).On(&symbol, "Destroy World!")
	fmt.Println("moc value: ", symbol)

	rec()
	fmt.Println("rec value: ", symbol)
}

func TestMock(t *testing.T) {
	fmt.Println("raw value:")
	printAll()

	mock := new(Mock)
	mock.On(&symbol, "Destory World!")
	mock.On(&foo, func() { fmt.Println("i am dummy foo") })
	mock.On(&bar, func() { fmt.Println("i am dummy bar") })

	fmt.Println("mock value:")
	printAll()

	mock.Recover()
	fmt.Println("recover value:")
	printAll()
}
```

