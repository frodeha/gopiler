package main

type MyStruct struct {
}

type MyInterface interface {
}

func main() {
	// Hello this is a comment that should be ignored
	123
	"this is a string"

	test := 123
	var test2 int = 456
	_ = test2

	test(test2)

}

func test(i int) error {
	return nil
}
