package main

import "fmt"

type Say func()

func SayHello() {
	fmt.Print("Hello ")
}

func SayHi() {
	fmt.Print("Hi ")
}

func SayTo(say Say, to string) Say {
	return func() {
		say()
		fmt.Print(to)
	}
}

func WhoSay(who string, say Say) Say {
	return func() {
		fmt.Print(who, " say ")
		say()

	}
}

func main() {
	sayHelloToTom := SayTo(SayHello, "tom")
	joinSayHi := WhoSay("join", SayHi)

	sayHelloToTom()
	fmt.Println()

	joinSayHi()
	fmt.Println()

	joinSayHeeloToTom := WhoSay("join", SayTo(SayHello, "tom"))
	joinSayHeeloToTom()
	fmt.Println()

	joinSayHeeloToTom1 := SayTo(WhoSay("join", SayHello), "tom")
	joinSayHeeloToTom1()
}
