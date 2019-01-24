package Mock

var foo = func(){
	fmt.Println("i am func foo")
}

var bar = func(){
	fmt.Println("i am func bar")
}

var symbol string = "Hello World!"

func TestMock(t *testing.T) {
	fmt.Println("raw value:")
	printAll()

	mock := new(Mock)
	mock.On(&symbol, "Destory World!")
	mock.On(&foo, func(){fmt.Println("i am dummy foo")})
	mock.On(&bar, func(){fmt.Println("i am dummy bar")})


	fmt.Println("mock value:")
	printAll()

	mock.Recover()
	fmt.Println("recover value:")
	printAll()
}

func printAll()  {
	fmt.Println(symbol)
	foo()
	bar()
}