package main

import (
	"fmt"
	"net/http"
)

type I interface {
	m() string
}
type F func() string

func (f F) m() string {
	return f()
}
func print(i I) {
	fmt.Printf("%s\n", i.m())
}

type IX interface {
	x() string
}
type X func() string

func (f X) x() string {
	return f()
}

type IY interface {
	y() string
}

type Y func() string

func (f Y) y() string {
	return f()
}

type II interface {
	IX
	IY
}

func printXY(i II) {
	fmt.Printf("x:%s y:%s\n", i.x(), i.y())
}

type subject struct {
	id   int
	name string
}

func main() {

	print(F(func() string { return "x" }))
	printXY(&struct {
		IX
		IY
	}{
		X(func() string { return "hello" }),
		Y(func() string { return "world" }),
	})

	subj := subject{name: "world"}
	fmt.Printf("hello %s\n", subj.name)
	subj.id++
	fmt.Printf("id %d", subj.id)


	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}