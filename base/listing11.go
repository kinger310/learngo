package main

import "fmt"

type user struct {
	name string
	age  int64
}

// notify implements a method with a value receiver
func (u user) notify() {
	fmt.Printf("User: %s Age: %d\n", u.name, u.age)
}

// changeAge implements a method with a pointer receiver.
func (u *user) changeAge(age int64) {
	u.age = age
}

func main() {
	bill := user{"bill", 23}
	bill.notify()

	lisa := &user{"lisa", 34}
	lisa.notify()

	bill.changeAge(32)
	bill.notify()

	lisa.changeAge(43)
	lisa.notify()
}
