# AnyThingMock
go impl mock any variable

useage

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

func TestMockOne(t *testing.T)  {
	fmt.Println("raw value: ", symbol)

	rec := new(Mock).On(&symbol, "Destroy World!")
	fmt.Println("moc value: ", symbol)

	rec()
	fmt.Println("rec value: ", symbol)
}
```
