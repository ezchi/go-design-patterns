package main

import "fmt"

type Athlete struct{}

func (a Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming")
}

type Animal struct{}

func (a Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	swimmer := CompositeSwimmerA{MySwim: Swim}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	shark := Shark{Swim: Swim}
	shark.Eat()
	shark.Swim()

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmerB.Train()
	swimmerB.Swim()
}
