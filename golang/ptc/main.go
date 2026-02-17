package main

import "fmt"

func main() {

	runVariables()
	runCondition()
	runLoop()
	runFunctionExample()
	runStructExample()
	runPointerExample()
	runArrayExample()
	runSliceExample()

}

// ######## VARIABLES & DATA TYPES ########

func runVariables() {

	fmt.Println("====== VARIABLES & DATA TYPES ======")

	var name string = "Shakil"
	age := 25
	height := 5.8
	isStudent := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	fmt.Println()
}

// ######## CONDITION (IF ELSE) ########

func runCondition() {

	fmt.Println("====== CONDITION EXAMPLE ======")

	number := 15

	if number%2 == 0 {
		fmt.Println(number, "is Even")
	} else {
		fmt.Println(number, "is Odd")
	}

	if number > 10 {
		fmt.Println("Number is greater than 10")
	}

	fmt.Println()
}

// ######## FOR LOOP ########

func runLoop() {

	fmt.Println("====== FOR LOOP EXAMPLE ======")

	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	fmt.Println()
}

// ######## FUNCTION EXAMPLE ########

func add(a int, b int) int {
	return a + b
}

func runFunctionExample() {

	fmt.Println("====== FUNCTION EXAMPLE ======")

	result := add(10, 20)
	fmt.Println("Sum:", result)

	fmt.Println()
}

// ######## STRUCT EXAMPLE ########

type Person struct {
	Name string
	Age  int
}

func runStructExample() {

	fmt.Println("====== STRUCT EXAMPLE ======")

	p1 := Person{
		Name: "Rahim",
		Age:  30,
	}

	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)

	fmt.Println()
}

// ######## POINTER EXAMPLE ########

func runPointerExample() {

	fmt.Println("====== POINTER EXAMPLE ======")

	value := 100
	ptr := &value

	fmt.Println("Original Value:", value)
	fmt.Println("Memory Address:", ptr)
	fmt.Println("Dereferenced Value:", *ptr)

	*ptr = 200

	fmt.Println("Updated Value:", value)
	fmt.Println()
}

// ######## ARRAY EXAMPLE ########

func runArrayExample() {

	fmt.Println("====== ARRAY EXAMPLE ======")

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}

	fmt.Println()
}

// ######## SLICE EXAMPLE ########

func runSliceExample() {

	fmt.Println("====== SLICE EXAMPLE ======")

	// # Create slice
	numbers := []int{10, 20, 30, 40, 50}

	// # Append value
	numbers = append(numbers, 60)

	fmt.Println("Slice:", numbers)

	// # Slice from array
	arr := [6]int{1, 2, 3, 4, 5, 6}
	part := arr[1:4]

	fmt.Println("Array:", arr)
	fmt.Println("Sliced Part:", part)

	// # Slice of struct

	type Product struct {
		Name  string
		Price int
	}

	products := []Product{
		{Name: "Laptop", Price: 50000},
		{Name: "Phone", Price: 20000},
		{Name: "Tablet", Price: 30000},
	}

	fmt.Println("\nProduct List:")
	for _, p := range products {
		fmt.Println("Name:", p.Name, "| Price:", p.Price)
	}

	fmt.Println()
}
