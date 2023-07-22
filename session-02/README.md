# CESA Backend - Session 2

In this tutorial, we will cover various essential aspects of the Go programming language (Golang). Golang is known for its simplicity, efficiency, and powerful features, making it a popular choice for building scalable and high-performance applications. We'll explore the following topics:

- [CESA Backend - Session 2](#cesa-backend---session-2)
  - [1. Structs](#1-structs)
  - [2. Methods](#2-methods)
  - [2.5. Embedding](#25-embedding)
    - [Embedding Syntax](#embedding-syntax)
    - [Accessing Embedded Fields and Methods](#accessing-embedded-fields-and-methods)
    - [Method Overriding](#method-overriding)
    - [Anonymous Struct Embedding](#anonymous-struct-embedding)
  - [3. Exported vs Unexported](#3-exported-vs-unexported)
  - [4. Package Management](#4-package-management)
  - [5. Modules](#5-modules)
  - [6. Interfaces](#6-interfaces)

## 1. Structs

Structs in Golang are composite data types that allow you to group together different types of variables under a single name. They provide a convenient way to represent and manage related data. To define a struct, use the `type` keyword followed by the struct's name and its fields.

```go
// Define a struct representing a person
type Person struct {
    Name string
    Age  int
}
```

You can then create instances of the struct as follows:

```go
func main() {
    person := Person{
        Name: "John",
        Age:  30,
    }

    fmt.Println(person.Name, person.Age)
}
```

Every field in a struct will be initialized with a zero value if not specified:

```go
func main() {
    person := Person{}
    fmt.Println(person.Name, person.Age) // "", 0
}
```

## 2. Methods

Methods in Golang are functions associated with a specific struct type. They allow you to add behavior or functionality to a struct. To define a method, you need to specify the receiver before the method name, which indicates on which struct type the method operates.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Usage of the method:

```go
func main() {
    rect := Rectangle{Width: 5, Height: 10}
    area := rect.Area()
    fmt.Println("Area:", area)
}
```

Methods, if used with a pointer receiver can modify their receiver as well, note that if the method accepts a value receiver, nothing will change:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) DoubleHeightFaulty() {
    r.Height = r.Height * 2
}

func (r *Rectangle) DoubleHeight() {
    r.Height = r.Height * 2
    // equals: (*r).Height = (*r).Height * 2, it's a syntactic sugar
} 
```

NOTE: treat the receiver just like a normal function parameter that just happens to come before the function name.

## 2.5. Embedding

Struct embedding in Golang allows a struct to include fields and methods of another struct, effectively creating a relationship known as composition. By embedding one struct into another, you can reuse the embedded struct's properties and methods without explicitly invoking them. This feature provides a way to achieve code reuse and create more maintainable and extensible code.

### Embedding Syntax

To embed a struct into another, simply include the name of the struct as a field within the target struct, without specifying a field name. Let's see an example:

```go
// Define a basic struct
type Person struct {
    Name string
    Age  int
}

// Define a struct that embeds the Person struct
type Employee struct {
    Person // Embedded struct
    EmployeeID int
}
```

In this example, the `Employee` struct embeds the `Person` struct, making the `Employee` struct inherit all the fields and methods of the `Person` struct.

### Accessing Embedded Fields and Methods

When a struct is embedded into another, its fields and methods are automatically promoted to the embedding struct. This means you can access them directly from the embedding struct, without needing to specify the embedded struct's name.

```go
func main() {
    emp := Employee{
        Person: Person{Name: "John", Age: 30},
        EmployeeID: 12345,
    }

    // Accessing the fields from the embedded struct
    fmt.Println("Name:", emp.Name)
    fmt.Println("Age:", emp.Age)

    // Accessing the method from the embedded struct
    emp.Introduce()
}

// Method defined for the embedded Person struct
func (p Person) Introduce() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
```

### Method Overriding

When a method is defined both in the embedded struct and the embedding struct, the method in the embedding struct takes precedence. This means that the method in the embedded struct can be overridden by redefining it in the embedding struct.

```go
type Dog struct {
    Name string
}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}

type Labrador struct {
    Dog // Embedding the Dog struct
}

func (l Labrador) Speak() {
    fmt.Println("Bark! Bark!")
}
```

In this example, we have an embedded `Dog` struct with a `Speak` method. We then have a `Labrador` struct that embeds the `Dog` struct. The `Labrador` struct has its own `Speak` method, which overrides the method of the embedded `Dog` struct.

### Anonymous Struct Embedding

You can achieve anonymous struct embedding by not providing a name for the embedded struct within the embedding struct. This allows you to directly access the fields and methods of the embedded struct without referencing its name.

```go
type Car struct {
    Brand string
}

// Anonymous struct embedding
type Vehicle struct {
    Car
    Type string
}
```

In this example, the `Vehicle` struct embeds the `Car` struct anonymously. Now, you can access the fields of the `Car` struct directly from the `Vehicle` struct, like this:

```go
func main() {
    vehicle := Vehicle{
        Car:  Car{Brand: "Toyota"},
        Type: "Sedan",
    }

    // Accessing the embedded struct fields directly
    fmt.Println("Brand:", vehicle.Brand)
    fmt.Println("Type:", vehicle.Type)
}
```

Anonymous struct embedding can lead to more concise code when you want to reuse the properties and methods of an existing struct.


## 3. Exported vs Unexported

In Golang, identifiers (variables, consts, functions, and structs) can be exported or unexported. The naming convention dictates that identifiers starting with an uppercase letter are exported and accessible from other packages, while those starting with a lowercase letter are unexported and only accessible within the same package.

```go
// Exported identifier (accessible from other packages)
const ExportedConst = 42

// Unexported identifier (only accessible within this package)
const unexportedConst = 100
```

NOTE: In go, privacy can only be defined around packages, this means that nothing inside a package is private to other things inside the same package, and an unexported element inside a package will not at all be visible nor accessible to other packages.

Unused unexported methods, functions, etc. will probably be underlined by your golang language server and/or linter for a good reason: if they are not used in their own package, they won't be used in any other package at all.

## 4. Package Management

Golang has built-in support for package management. It uses the `go` command to fetch, build, and install packages and their dependencies. Packages are used to organize and share code effectively.

To create a new Go module for your project, you can use the following command:

```bash
go mod init module_name
```

This will initialize a new module in your project and create a `go.mod` file that keeps track of the module's dependencies.

Checkout [here](./parent/child1/child1.go) and [here](./parent/child2/child2.go) to see how these packages will be managed and used in go.  

## 5. Modules

Go modules provide a way to manage dependencies for your projects effectively. Modules are collections of related packages that are versioned together. The `go.mod` file, located in the root of your project, specifies the module's name, dependencies, and their versions.

To add a new dependency to your project, use the following command:

```bash
go get package_name
```

For example:

```bash
go get github.com/gin-gonic/gin
```

This command will fetch the specified package and add it to your `go.mod` file.

## 6. Interfaces

Interfaces in Golang define a set of method signatures. Any type that implements all the methods of an interface automatically satisfies that interface. Interfaces allow you to write more flexible and modular code.

```go
// Defining an interface
type Shape interface {
    Area() float64
}

// Implementing the interface for the Rectangle struct
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Using the interface:

```go
func CalculateArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    rect := Rectangle{Width: 5, Height: 10}
    CalculateArea(rect)
}
```

Interfaces allow you to work with different types that have similar behavior, making your code more versatile.

```go
type Saver interface {
    Save([]byte) 
}

type postgresDB struct {
    storage []byte
}

func NewPostgresDB() Saver {
    return &postgresDB{}
}

func (p *postgresDB) Save(d []byte) {
    p.storage = append(p.storage, d...)
    fmt.Println("saved in postgres db!")
}

func (p *postgresDB) SomeMethod() {...}

type mongoDB struct {
    storage []byte
}

func NewMongoDB() Saver {
    return &mongoDB{}
}

func (m *mongoDB) Save(d []byte) {
    m.storage = append(m.storage, d...)
    fmt.Println("saved in mongo db!")
}

type redisDB struct {
    storage []byte
}

func NewRedisDB() Saver {
    return &redisDB{}
}

func (r *redisDB) Save(d []byte) {
    r.storage = append(r.storage, d...)
    fmt.Println("saved in redis db!")
}

func (r *redisDB) AnotherMethod() {...}

func main() {
    pg := NewPostgresDB()
    mongo := NewMongoDB()
    redis := NewRedisDB()
    storages := []Saver{pg, mongo, redis}

    // let's say we have some data
    data := []byte("some data")

    for _, s := range storages {
        s.Save(data)
    } 
}
```

Checkout [here](./interface/main.go) for a more complete example

In order to satisfy and implement an interface, you only need to implement the mandatory functions (e.g. the functions that are contained in the interface definition, `Save` in the example above) and you are free to have as much additional methods as you want.

Interfaces also enables us to `mock` different dependencies in order to facilitate unit tests. More on this later!

---

This tutorial has covered some essential aspects of the Go programming language, including structs, methods, exported vs. unexported identifiers, package management, modules, and interfaces. These concepts are fundamental to building efficient and maintainable Go applications. Happy coding with Golang!