package main

import (
	"fmt"
	//"strconv"
)

// 11. Functions
func mult(a int, b int) (int, string) {
	res := a * b
	return res, "Successfully returned a result!"
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}


func main() {
	// 1. Hello World
	// fmt.Println("Hello World!")

	// 2. Variables
	/*var b bool
	fmt.Println(b)

	something := "1" + "2"
	fmt.Println(something)*/

	// 3. Constants
	/* const b int = 0
	b = 2
	fmt.Println(b)*/

	// 4. Loops
	/*i := 0
	fmt.Println("First Loop:")
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nSecond Loop:")
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nThird Loop:")
	limit := 10
	for i := 0; i < limit; i++ {
		if i > limit/2 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("\nFourth Loop:")
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\nFinished Looping...")*/

	// 5. If statements

	/*var a int = 5
	if num := a; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is in single digits")
	} else {
		fmt.Println("num is in double digits")
	}*/

	// 6. Switch statements

	/*whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool!")
		case int:
			fmt.Println("I'm an int!")
		case string:
			fmt.Println("I'm a string!")
		default:
			fmt.Println("Who am I?!", t)
		}
	}

	whatAmI(true)
	whatAmI("String")
	whatAmI(99)
	whatAmI(3.4)*/

	// 7. Initialize Array

	/*arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}*/

	// 8. Slices
	/*fmt.Println("Making slice s")
	s := make([]string, 100)
	for i := 0; i < len(s); i++ {
		s[i] = strconv.Itoa(i)
	}
	fmt.Println(s)

	fmt.Println("\ns[1:10]")
	fmt.Println(s[1:10])

	fmt.Println("\nt")
	t := append(s, "101")
	fmt.Println(t)

	fmt.Println("\nPrinting s again..")
	fmt.Println("", s)

	fmt.Println("\nAppending two slices together")
	a := []int{1, 2, 3, 4}
	fmt.Println("a = ", a)
	b := []int{5, 6, 7, 8}
	fmt.Println("b = ", b)
	fmt.Println("Appending b to a...")
	a = append(a, b...)
	fmt.Println("a = ", a)*/

	// 9. Maps
	/*m := make(map[int]string)
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"

	fmt.Println("This is a number to string map")
	fmt.Println(m)

	fmt.Println("The string for 1 is: ", m[1])
	fmt.Println("The string for 2 is: ", m[2])
	fmt.Println("The string for 3 is: ", m[3])

	fmt.Println("Deleting m[3]...")
	delete(m, 3)

	_, present := m[3]
	fmt.Println("Is the value for m[3] present?")
	if present {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No :(")
	}*/

	// 10. Range
	/*s := []int{1, 2, 3, 4}

	for i, v := range s {
		fmt.Printf("%d + %d = %d\n", i, v, i + v)
	}

	for i, c := range "Stephen Terepaii Tangi O'Sullivan" {
		fmt.Printf("%d, %s == %x\n", i, string(c), c)
	}*/

	// 11. Functions
	/*
	fmt.Println("Multiplying 1 and 2")
	fmt.Println(mult(1, 2))

	// Variadic
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("Adding up these numbers: %d\n", s)
	fmt.Println(sum(s...))*/

	// 12. Closures
	/*nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4 
	fmt.Println(nextInt()) // 5

	anotherNextInt := intSeq()
	fmt.Println(anotherNextInt()) // 1, new state separate from nextInt()*/

	// 13. Recursion
	fmt.Println(fact(1))
}