# Learning go

by Jon Bodner

## Chapter 1: Setting Up Your Go Environment

### What is Go

- A **compiled**, **statically typed** language created at Google.
- Designed for:
  - Simplicity (minimal syntax)
  - C-like performance
  - High productivity (integrated tooling)
  - Easy concurrency via **goroutines** and **channels**
- Code is organized into **packages**, compiled into static binaries.

### Installing Go

- Download from: https://go.dev/dl/
- `$GOPATH` still exists but is **conceptually deprecated** since Go Modules.
- During installation, Go creates:

  - `~/go/bin` → installed binaries
  - `~/go/pkg/mod` → module cache

- Verify installation:

```bash
go version
go env
```

#### Go Toolchain (the go command)

Main commands:

`go run file.go` → compiles + runs.
`go build` → compiles a binary.
`go install` → installs a binary into ~/go/bin.
`go mod init` → initializes a module.
`go mod tidy` → cleans unused dependencies.
`go get` → adds or updates a dependency.
`go fmt` ./... → formats code (Go standard).
`go vet` ./... → detects common or dangerous patterns.
`go test ./...` → runs tests.

The toolchain includes:

- Compiler (cmd/compile)

- Linker (cmd/link)

- Formatter (gofmt)

- Simple linter (go vet)

- Test runner (go test)

- Module management

### First Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

### Workspace Layout

#### With Go Modules (recommended)

You can work in any folder, no need for `$GOPATH`.

Steps to create a new module/project:

1. `mkdir myproject`
2. `cd myproject`
3. `go mod init github.com/user/myproject`

The `go.mod` file:

- defines the module
- lists the dependencies + their versions
- dependencies are downloaded into `~/go/pkg/mod` (global cache)

#### Old workflow: `$GOPATH`

- Historically, all code had to be inside $GOPATH/src.
- Still supported, but should be avoided when using Modules.

---

## Chapter 2: Predeclared Types and Declarations

### Predeclared Types Overview

Go provides a set of built-in (predeclared) types available in all packages:

- **Boolean**: `bool` (`true` / `false`)
- **Numeric**:
  - Integers: `int`, `int8`, `int16`, `int32`, `int64`
  - Unsigned: `uint`, `uint8` (alias `byte`), `uint16`, `uint32`, `uint64`
  - Floats: `float32`, `float64`
  - Complex: `complex64`, `complex128`
- **String**: immutable sequences of bytes
- **Rune**: `rune` (alias for `int32`), used for Unicode code points

⚠️ **Note:**  
`int` and `uint` are **architecture-dependent**:

- 32-bit on 32-bit systems
- 64-bit on 64-bit systems

All of these types have **zero values**, e.g. `0`, `""`, `false`, `nil` (only for reference types).

### Variable Declarations

#### Using `var`

```go
var x int
var name string
var ok bool
```

- Variables declared with var get their zero value automatically.

- You can also assign explicitly:

```go
var count int = 10
```

#### Short Declaration (:=)

```go
x := 42
msg := "hello"
ok := true
```

Rules:

- Only allowed inside functions.

- Automatically infers the type.

- Cannot redeclare an existing variable in the same scope unless at least one new variable appears.

#### Multiple Declarations

```go
var a, b, c int
x, y := 1, 2
```

Or with different types:

```go
var (
    name string
    age  int
    ok   bool
)
```

#### Constants (`const`)

Must be compile-time values.

Cannot hold runtime results.

```go
const Pi = 3.14
const Greeting = "Hello"
```

- Cannot be declared with :=

- Only numbers, strings, booleans, and runes allowed.

#### Constant Groups

```go
const (
    A = 1
    B = 2
    C = 3
)
```

#### Untyped Constants

Go allows untyped numeric constants:

```go
const n = 42
```

- `n` can be used as an `int`, `float64`, `etc`. depending on context.

This gives constants more flexibility than variables.

Example:

```go
var x int = n
var y float64 = n
```

#### The iota Identifier

Used for auto-incrementing constants:

```go
const (
    Red = iota  // 0
    Green       // 1
    Blue        // 2
)
```

Resets for each `const` block.

#### Type Conversions

Go does not allow implicit type conversions.

```go
var x int = 10
var y float64 = float64(x)
```

**_Conversions are explicit and required._**

#### Blank Identifier (`_`)

Used to ignore a value:

```go
value, _ := someFunction()
```

Useful for:

- ignoring unused returns

- discarding values in multiple assignment

- imports solely for their side effects

#### Predeclared Identifiers

Go includes built-in identifiers accessible everywhere:

- `len`, `cap`, `new`, `make`

- `append`, `copy`

- `delete`, `complex`, `real`, `imag`

- `panic`, `recover`

Example:

```go
arr := []int{1, 2, 3}
l := len(arr)
```

## Chapter 3: Composite Types

### Overview

Composite types in Go are types made of other types.  
The main composite types are:

- **Arrays**
- **Slices**
- **Maps**
- **Structs**

They are core building blocks for data modeling and collection handling in Go.

### Arrays

- Fixed-length sequences of elements of the same type.
- Length is **part of the type** → `[3]int` and `[4]int` are different types.
- Rarely used directly in Go because slices are more flexible.

Example:

```go
var a [3]int        // [0, 0, 0]
b := [3]int{1, 2, 3}
```

### Slices

- **The most commonly used collection type in Go.**
- Built on top of arrays but with dynamic size.
- A slice is a lightweight descriptor containing:
  - pointer to underlying array
  - length
  - capacity

Example:

```go
s := []int{1, 2, 3}
```

Key behaviors:

- Slices are **_reference-like_** → assigning or passing a slice does not copy the data.

- append may reallocate the underlying array if capacity is exceeded.

- Slicing shares the same underlying array:

```go
sub := s[1:3]  // shares memory with s
```

Append growth:

- Go typically doubles the capacity when growing a slice.
- This amortizes append cost and keeps it efficient.

If reallocation happens, old slices still point to the old array.

### Maps

- Unordered key-value store.

- Keys must be comparable (`==` allowed).

- Created with `make` or literal syntax:

```go
m := map[string]int{"a": 1, "b": 2}
```

Key points:

- Maps are references to internal hash tables → copying a map value does not copy its contents.

- Reading a missing key returns the zero value of the value type.

- Use the “_comma ok_” idiom:

```go
v, ok := m["key"]
```

==A nil map can be read from but cannot be written to.==

### Structs

- Groups named fields under a single type.

```go
type Person struct {
    Name string
    Age int
}
```

⚠️ **Important notes:**

- Structs are **value types** → assigning or passing copies all fields.

- Preferred over maps for APIs: safer, typed, self-documenting.

- Zero value of a struct = zero value of its fields.
- **Struct fields starting with uppercase letters are exported (public).**
  Lowercase = package-private.

#### Composite Literals

Used to initialize composite types concisely:

```go
p := Person{Name: "Alice", Age: 30}
s := []int{1, 2, 3}
m := map[string]bool{"active": true}
```

Useful for tests and readable initialization.

### The `make` Function

Used to initialize slices, maps, and channels:

```go
s := make([]int, 0, 10)
m := make(map[string]int)
```

Difference with new:

`new(T)` allocates memory and returns `*T`.

`make(T)` initializes runtime structures for slices, maps, channels.

### Zero Values of Composite Types

- Arrays → all elements zeroed.

- Slices → `nil`, len=0, cap=0.

- Maps → `nil` (cannot write).

- Structs → each field gets its zero value.

---

## Chapter 4: Control Structures

Go's control structures are intentionally minimal.  
They favor clarity, explicit behavior, and avoid hidden pitfalls found in other languages.  
The main constructs are:

- `if`
- `for`
- `switch`
- `defer`

### If Statements

- Parentheses are **not** required.
- Braces `{}` are **mandatory**.
- Can include an optional **short statement** before the condition.

Example:

```go
if x := compute(); x > 10 {
    fmt.Println("big")
} else {
    fmt.Println("small")
}
```

Key points:

- The variable `x` exists only inside the `if/else` block.

- Avoid deeply nested `if` statements; return early instead (common Go style).

### For Loops

Go has **one** looping construct: `for`.
It covers classical loops, while-loops, and infinite loops.

#### Classic loop

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

#### While-style loop

```go
for x < 100 {
    x += 10
}
```

#### Infinite loop

```go
for {
    fmt.Println("looping")
}
```

#### Range loop

Used to iterate over slices, arrays, strings, maps, and channels.

```go
for i, v := range nums {
    fmt.Println(i, v)
}
```

Key notes:

- For `strings`, range returns runes (Unicode-aware).

- For maps, iteration order is not guaranteed.

### Switch Statements

More powerful than in many languages:

- No implicit fallthrough.

- Each `case` compares values or boolean expressions.

- Can include a short statement like `if`.

#### Value switch

```go
switch day {
case "Mon", "Tue":
    fmt.Println("weekday")
case "Sat", "Sun":
    fmt.Println("weekend")
default:
    fmt.Println("unknown")
}
```

#### Expressionless switch (like chained if)

```go
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}
```

#### Fallthrough

Explicit when needed:

```go
switch n {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
}
```

### Defer

- Schedules a function to run after the current function returns.

- Most common use: closing resources (files, connections, etc.).

- Deferred calls execute LIFO (stack order).

Example:

```go
func readFile() {
    f, _ := os.Open("test.txt")
    defer f.Close()
    // work with f
}
```

Order example:

```go
defer fmt.Println("first")
defer fmt.Println("second")
```

Output:

```bash
second
first
```

Key behaviors:

- Arguments to the deferred function are evaluated immediately, not at execution time.

- Useful for ensuring cleanup even when a function returns early.

### Panic and Recover (briefly)

Although covered later in the book, they are part of control flow:

- `panic` stops normal execution.

- `recover` regains control inside a deferred function.

Example:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}()
panic("boom")
```

---

## Chapter 5: Functions

Functions in Go are first-class citizens:  
they can be assigned to variables, passed as arguments, and returned from other functions.

Key concepts:

- simple function declarations
- multiple return values
- named return values
- variadic parameters
- functions as values
- closures

### Declaring Functions

Basic form:

```go
func add(a int, b int) int {
    return a + b
}
```

If multiple parameters share the same type:

```go
func add(a, b int) int
```

### Multiple Return Values

Very common in Go, especially for error handling.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
```

⚠️ Important:
==Always handle errors explicitly==. This is core to Go’s philosophy.

### Named Return Values (optional)

You can name return values and return without arguments:

```go
func sum(a, b int) (result int) {
    result = a + b
    return
}
```

Use sparingly; can reduce clarity if overused.

### Variadic Functions

Accept a variable number of arguments:

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

Call:

```go
sum(1, 2, 3, 4)
```

Important:
`nums` is a slice inside the function.

### Functions as Values

You can assign functions to variables:

```go
f := func(x int) int {
    return x * 2
}
fmt.Println(f(10))
```

Useful for callbacks, passing behavior, functional patterns.

### Closures

Functions that capture variables from their surrounding scope.

```go
func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

c := counter()
c() // 1
c() // 2
```

Closures retain access to captured variables even after the outer function returns.

### Defer with Functions

Often used inside functions for cleanup:

```go
func example() {
    defer fmt.Println("done")
    fmt.Println("working")
}
```

Output:

```bash
working
done
```

### Errors as Values

Go does not use exceptions.
Errors are returned as values (usually the last return value).

```go
data, err := readFile()
if err != nil {
    return err
}
```

This is foundational to idiomatic Go.

---

## Chapter 6: Pointers

A pointer is a variable that contains the address where another variable is stored.
In Go, the zero value for a pointer is _nil_.

Pointer arithmetic is not allowed in Go like you could do in C / C++.

The **&** is the _address_ operator. It precedes a value type and returns the address where the value is stored:

```go
x := "hello"
pointerToX := &x
```

The **\*** is the _indirection operator_. It precedes a variable of pointer type and returns the pointed-to value. This is called **dereferencing**:

```go
x := 10
pointerToX := &x
fmt.Println(pointerToX) //prints a memory address
fmt.Println(*pointerToX) // prints 10
z := 5 + *pointerToX
fmt.Println(z) //prints 15
```

### A pointer indicate a mutable parameters

==_Go is a call-by-value language_==, the values passed to functions are copies. For nonpointer types like primitives, structs, and arrays, this means that the called function cannot modify the original. **Since the called function has a copy of the original data, the original data's immutability is guaranteed**.

If a pointer is passed to a function, the function gets a copy of the pointer.
This still points to the original data, which means that the original data can be modified by the called function.

### Pointers are a Last Resort

The only time you should use a pointer param to modify a variable is when the function expects an interface --> JSON

```go
f := struct {
    Name string `json:"name"`
    Age int `json:"age"`
}{}
err := json.Unmarshal([]bytes(`{"name": "Bob", "age": 30}`), &f)
```

Unmarshal() take 2 arguments, a slice of bytes and an any. The value passed in for the any parameter must be a pointer or an error is return.

### Pointer passing performance
- Pointers have **constant size**, regardless of data size.
- Passing large structs by value is expensive → pointers can reduce copying.
- For **small/medium structs**, passing by value is usually faster and safer.
- Rule of thumb:
  - **< ~10 MB** → returning/passing by value is fine.
  - **Huge structs** → pointer passing helps.
- In practice:
==Performance rarely justifies pointers== → **prefer value semantics**.

### The difference between Maps and Slices

- **Maps are reference types**
  - Internally: implemented as a pointer to a runtime hash table.
  - Passing a map → copies the pointer, **not** the contents.
  - Mutations inside a function affect the original map.

- **Slices are also reference-like, but:**
  - Slice header (ptr + len + cap) is copied.
  - Underlying array may be shared or reallocated on append.

- **API design caution**
  - Maps expose no structure → unclear contract.
  - Avoid using maps directly in public APIs.
  - Prefer **structs** for clarity, type safety, validation, and future changes.

---
