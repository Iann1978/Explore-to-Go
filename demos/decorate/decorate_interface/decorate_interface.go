package main

import "fmt"

type Bird interface {
	Fly()
	Eat()
}

type Parrot struct{}

func (b *Parrot) Fly() {
	fmt.Println("Parrot is Fling.")
}

func (b *Parrot) Eat() {
	fmt.Println("Parrot is Eatting")
}

type WhosBird struct {
	Bird
	Who string
}

func (b *WhosBird) Fly() {
	fmt.Print(b.Who, "'s ")
	b.Bird.Fly()
}

func MakeWhosBird(who string, b Bird) Bird {
	return &WhosBird{
		Bird: b,
		Who:  who,
	}
}

func main() {
	tomsBird := MakeWhosBird("tom", &Parrot{})
	tomsBird.Fly()
	tomsBird.Eat()
}
