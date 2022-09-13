package main

import "fmt"

func fun1() {
	fmt.Println("Hello World")
}

func fun2(name string) {
	fmt.Println("Hello ", name)
	fmt.Printf("Hello %s", name)
}

func fun3(name string, age int) {
	fmt.Println("Hello : ", name)
	fmt.Printf("Hello: %s, Age: %d", name, age)
}

func fun4(name string, age, eid int) {
	fmt.Println("Hello : ", name)
	fmt.Printf("Hello: %s, Age: %d, EID: %d", name, age, eid)
}

func fun5(name string, age, eid int) int {
	// fmt.Println("Hello : ", name)
	// fmt.Printf("Hello: %s, Age: %d, EID: %d", name, age, eid)
	return age
}

func fun6(name string, age, eid int) (string, int) {
	// fmt.Println("Hello : ", name)
	// fmt.Printf("Hello: %s, Age: %d, EID: %d", name, age, eid)
	return name, age
}

func fun7(name string, age, eid int) (string, int, int) {
	// fmt.Println("Hello : ", name)
	// fmt.Printf("Hello: %s, Age: %d, EID: %d", name, age, eid)
	return name, age, eid
}

func fun8(name string, age, eid int) (name1 string, age1, eid1 int) {

	name1 = name
	age1 = age
	eid1 = eid

	return
}

func fun88(name string, age, eid int) (string, int, int) {

	name1 := name
	age1 := age
	eid1 := eid
	// fmt.Println("Hello : ", name)
	// fmt.Printf("Hello: %s, Age: %d, EID: %d", name, age, eid)
	return name1, age1, eid1
}

func main() {
	defer fun1()
	defer fun2("Alice")
	// fun3("Alice", 30)
	// fun4("Alice", 30, 231234234)

	// age := fun5("Alice", 30, 231234234)
	// fmt.Println(age)

	//fun6("Alice", 30, 231234234)
	//fmt.Println(age)

	// name, age, eid := fun7("Alice", 30, 231234234)
	// fmt.Println(name, age, eid)

	name, age, eid := fun8("Alice", 30, 231234234)
	fmt.Println(name, age, eid)
}
