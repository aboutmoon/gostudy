package main

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type Book struct {
}

func (b *Book) Read() {
	println("read book")
}

func (b *Book) Write() {
	println("write book")
}
func main() {
	b := new(Book)

	var r Reader

	r = b
	r.Read()

	var w Writer

	w = r.(Writer)

	w.Write()
}
