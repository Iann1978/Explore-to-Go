package main

import (
	"fmt"
	"reflect"
)

type Bird interface {
	Fly()
	Eat()
}

type Parrot struct {
}

func (p *Parrot) Fly() {
	fmt.Println("Parrot  is flying")
}

func (p *Parrot) Eat() {
	fmt.Println("Parrot is eating")
}

type WhosBird struct {
	Bird
	Name string
}

func main() {

	var bird Bird
	bird = &Parrot{}

	tp := reflect.TypeOf(WhosBird{})
	fmt.Println(tp.Kind())
	newBird := reflect.New(tp)
	fmt.Println(newBird.Type().Kind())
	newBird.Elem().FieldByName("Bird").Set(reflect.ValueOf(bird))
	newBird.Elem().FieldByName("Name").Set(reflect.ValueOf("daogi"))
	newBird.Elem().MethodByName("Fly")

	// fmt.Println(tp)

	// bird.Fly()
}
