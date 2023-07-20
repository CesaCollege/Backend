# Session 01 - Go Introduction

## History

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

## Pros & Cons

### Pros 👍

* Memory Safety
* Garbage Collection
* Structural Typing
* Concurrency
* Auto Generation
* Clean Formatting
* High Performance
* Low Resource Usage
* Open Source
* Error Handling

### Cons 👎

* OOP
* Lack of Libraries
* Small Community
* Lack of Generic Functions

## Companies

* Google
* Uber
* Docker
* Cloudflare
* Twitch
* DigitalOcean
* PayPal
* Dropbox
* Divar
* CafeBazaar
* Snapp
* Digikala
* Podro

---

## Syntax

Example Code:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```
[Check this out](E1.go)

### Variables

In Go, variables are explicitly declared and their types are specified at the time of declaration. This static typing helps catch errors early and promotes code reliability. You declare a variable in Go using the var keyword, followed by the variable name and its type.
```go
var message string
var message string = "Hello, world!"
message := "Hello, world!"
```
[Check this out](E2.go)

In Go, constant variables are values that are assigned once and cannot be changed during the execution of a program. They are defined using the const keyword and provide a way to declare and use fixed values that remain constant throughout the program's lifetime.
```go
const (
variable = value
)
const Pi = 3.14
```
In Go, the strconv package provides functions for converting strings to various data types and vice versa. The name "strconv" is short for "string conversion," and it offers a range of functionalities for handling string conversions in a convenient and efficient manner.
The strconv package in Go includes functions for converting strings to numeric types, such as integers and floats, as well as functions for converting numeric types back to strings. It also provides utilities for working with booleans and quoting or unquoting strings.
```go
num, _ := strconv.Atoi("123")
str := strconv.Itoa(456)
f, _ := strconv.ParseFloat("3.14", 64)
```
---
## Conditions

In Go, conditional statements provide a way to execute different blocks of code based on specific conditions. These statements allow your program to make decisions and alter its behavior dynamically. The most commonly used conditional statements in Go are if, else if, and else.
```go
package main

import "fmt"

func main() {
	age := 25

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}
}
```
[Check this out](E3.go)

---
### Task
Write a Go program that receives input from the user and prints the input if it is greater than the last number provided. Assume the first number is 0.

[Answer](T1.go)

---

### Functions
In Go, functions are essential building blocks that encapsulate reusable blocks of code. They allow you to define logic, perform computations, and execute specific tasks.
[Check this out](E4.go)

---

### Error Handling
In Go, error handling is a crucial aspect of writing reliable and robust code. Proper error handling helps identify and handle exceptional conditions or unexpected behavior, allowing your program to gracefully recover from errors and provide meaningful feedback to users. If you're new to Go, understanding the basics of error handling is essential.
[Check this out](E5.go)

---

### Array
In Go, an array is a fixed-size, ordered collection of elements of the same type. Arrays provide a way to store and access multiple values under a single variable name. 
```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Array is fixed-size buy slice is dynamic.

[Check this out](E6.go)

---
### Loop
In Go, loops are used to repeatedly execute a block of code until a certain condition is met. They provide a way to iterate over collections, perform iterative tasks, and control the flow of your program.

`For` loop:
```go
for i := 0; i < 5; i++ {
    // Code to be executed
}
```

`While` loop:
```go
i := 0
for i < 5 {
    // Code to be executed
    i++
}
```

`Do-While` loop:
```go
i := 0
for {
    // Code to be executed
    i++
    if i >= 5 {
        break
    }
}
```

`range` in loop:
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    // Code to be executed for each element
}
```
[Check this out](E7.go)

---
## Task

Get input from the user called "n". Loop "n" times and get an integer from the user during each iteration. At the end of the loop, print all the numbers entered by the user if they are even.

[Check this out](T2.go)

---

### Map
In Go, a map is a built-in data structure that provides an unordered collection of key-value pairs. Maps are also known as associative arrays, hash tables, or dictionaries in other programming languages. They are incredibly useful for storing and retrieving data based on unique keys. 
```go
var myMap map[string]int
myMap := map[string]int{"apple": 1, "banana": 2, "orange": 3}
```

[Check this out](E8.go)

### Struct
In Go, a struct is a composite data type that allows you to group together values of different types into a single entity. Structs provide a way to create complex data structures by defining custom types that contain multiple fields. 

#### Struct vs Class

* Definition
* Inheritance
* Methods
* Encapsulation and Visibility
* Polymorphism
* Constructors


Example:
```go
type Account struct {
	Name string
	Family string
	Balance int
}

```

[Check this out](E9.go)

[Check this out](E10.go)
