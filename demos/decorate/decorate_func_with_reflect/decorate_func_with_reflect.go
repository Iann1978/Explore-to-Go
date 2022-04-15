package main

import (
	"fmt"
	"reflect"
)

type FuncTomEat func()
type FuncWhoEat func(string)
type FuncWhoDoWhat func(string, string)

func TomEat() {
	fmt.Print("Tom is eating.")
}

func WhoEat(who string) {
	fmt.Print(who, " is eating.")
}

func WhoDoWhat(who string, what string) {
	fmt.Print(who, " is ", what, ".")
}

func ISeeTomEat(f FuncTomEat) FuncTomEat {
	return func() {
		fmt.Print("I see ")
		f()
	}
}

func ISeeWhoDoWhat(f FuncWhoDoWhat) FuncWhoDoWhat {
	return func(who string, what string) {
		fmt.Print("I see ")
		f(who, what)
	}
}

func ISee(out interface{}, in interface{}) {
	outptr := reflect.ValueOf(out).Elem()
	// tomeat := reflect.ValueOf(TomEat)

	infunc := reflect.ValueOf(in)

	f := reflect.MakeFunc(infunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Print("I see ")
			out = infunc.Call(in)

			return
		})

	outptr.Set(f)
}

func main() {

	iseetomeat := ISeeTomEat(TomEat)
	iseetomeat()
	fmt.Println()

	var iseetomeat2 FuncTomEat
	ISee(&iseetomeat2, TomEat)
	iseetomeat2()
	fmt.Println()

	whodowhat := WhoDoWhat
	whodowhat("Join", "working")
	fmt.Println()

	var iseewhodowhat FuncWhoDoWhat
	ISee(&iseewhodowhat, WhoDoWhat)
	iseewhodowhat("Join", "working")
	fmt.Println()

	// iseewhodowhat := ISeeWhoDoWhat(WhoDoWhat)
	// iseewhodowhat("Tom", "eating")

}
