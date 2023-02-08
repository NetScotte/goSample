package main

import "fmt"

type Food interface {
	GetName() string
}

type Fruit struct {
	Name string
}

func (f *Fruit) GetName() string {
	return f.Name
}

type Peat struct {
	Name string
}

func (p *Peat) GetName() string {
	return p.Name
}

func Eat(f Food) {
	fmt.Println(f.GetName())
}

func main() {
	f := &Fruit{Name: "苹果"}
	p := &Peat{Name: "猪肉"}
	Eat(f)
	Eat(p)
}
