# Learning go

by Jon Bodner

## Chapter 1: Setting Up Your Go Environment

### What is Go

- A **compiled**, **statically typed** language created at Google.
- Designed for:
  - Simplicity (minimal syntax)
  - C-like performance
  - High productivity (integrated tooling)
  - Easy concurrency via **goroutine** and **channels**
- Code is organized into **packages**, compiled into static binaries.

### Installing Go

- Download from: <https://go.dev/dl/>
- `$GOPATH` still exists but is **conceptually deprecated** since Go Modules.
- During installation, Go creates:
  - `~/go/bin` → installed binaries
  - `~/go/pkg/mod` → module cache

- Verify installation:

```bash
mqgo version
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

**_A nil map can be read from but cannot be written to._**

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
- **Struct fields starting with Uppercase letters are exported (public).**
  lowercase = package-private.

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
**_Always handle errors explicitly_**. This is core to Go’s philosophy.

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

**_Go is a call-by-value language_**, the values passed to functions are copies. For non-pointer types like primitives, structs, and arrays, this means that the called function cannot modify the original. **Since the called function has a copy of the original data, the original data's immutability is guaranteed**.

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
  **_Performance rarely justifies pointers_** → **prefer value semantics**.

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

### Reducing the Garbage Collector workload

- Most values live on the **stack** or in **CPU registers** (since Go 1.17).
- The heap is GC-managed → allocations here cost more.
- Fewer heap allocations = less GC pressure = better performance.
- Concept: **Mechanical sympathy**
  Understanding how hardware affects performance.
- `GOMEMLIMIT`
  Allows controlling max memory usage (bytes), influencing GC activity.

---

## Chapter 7: Types, Methods, and Interfaces

### TLDR types, methods, interfaces

- Methods add behavior to types (value/pointer receivers).
- Interfaces are implicit and describe behavior only.
- Embedding ≠ inheritance (use for composition).
- Accept interfaces, return concrete types.
- Type assertions + type switches = use sparingly.
- Interfaces enable clean dependency injection.

### Types in Go

- A type can be declared at any scope.
- **Concrete type** = data + behavior (methods).
- **Abstract type** = behavior only → **interfaces**.

### Methods

- Methods are functions with a **receiver**, declared at package level.
- Cannot be overloaded.
- Receivers:
  - **Value receiver** → method gets a copy.
  - **Pointer receiver** → method can mutate original value.

Go automatically adjusts the call:

```go
c.Inc()     // value → pointer method → (&c).Inc()
c.String()  // pointer → value method → (*c).String()
```

#### Go Idiom

**Never write getters/setters unless they satisfy an interface.**

### Functions vs Methods

- Use a **method** when logic depends on struct state.
- Use a **function** when logic only depends on input parameters.
- Methods = good for: configuration, shared state, domain behavior.
- Functions = pure logic.

### Iota as Enumerations

Go has no true enums, but uses `iota` for incrementing constants.

```go
type MailCategory int

const (
    Uncategorized MailCategory = iota
    Personal
    Spam
    Social
    Ads
)
```

### Embedding (Composition)

- Embedding a type inside a struct promotes its methods.
- Provides code reuse without inheritance.

```go
type A struct { X int }
type B struct { A }     // B “has” A, and gets A’s methods
```

⚠️ Embedding ≠ inheritance.

### Interfaces

- The **only abstract type** in Go.
- **Implicit implementation** → no “implements” keyword.

```go
type Stringer interface {
    String() string
}
```

If a type has a method `String() string`, it implements `Stringer`.

#### Interface Design

- Names often end with `-er`: Reader, Writer, Formatter…
- Interfaces express **behavior**, not data.

### Interfaces = Type-Safe Duck Typing

> **_“If it has the right methods, it fits the interface.”_**

This enables:

- **decorator pattern**
- wrappers
- mocks & testing
- dependency injection

### Embedding Interfaces

Interfaces can embed other interfaces:

```go
type Reader interface  { Read(p []byte) (n int, err error) }
type Closer interface  { Close() error }

type ReadCloser interface {
    Reader
    Closer
}
```

### Accept Interfaces, Return Structs (Important Rule)

This chapter’s key idiom:

- **Accept interfaces** → caller can inject different behaviors
- **Return concrete types** → you keep control over implementation

This makes code easier to test, extend, mock.

### Interfaces and `nil`

An interface is internally: **(type, value)**.

- Interface is **nil only if both are nil**.
- A typed nil value still makes the interface **non-nil**, causing surprises:

```go
var r io.Reader = (*os.File)(nil)
fmt.Println(r == nil) // false → type is set
```

### Interfaces Are Comparable

- Interfaces can be compared, but comparing unsupported values may `panic`.
- Avoid comparing if type might contain non-comparable values.

### Empty Interface (`any`)

- `interface{}` or `any` (alias) accepts any type.
- Avoid overuse; loses type safety.

### Type Assertions & Type Switches

#### Type Assertion

```go
i := any(10)
v := i.(int) // panic if not int
```

Use comma-ok form to avoid panic:

```go
v, ok := i.(int)
```

#### Type Switch

```go
switch v := i.(type) {
case int:
case string:
case MyType:
default:
}
```

Use sparingly → often a sign of poor abstraction.

### Function Types → Bridge to Interfaces

You can use function types to adapt functions to interfaces (adapter pattern):

```go
type LoggerAdapter func(string)

func (lg LoggerAdapter) Log(msg string) { lg(msg) }
```

### Dependency Injection with Interfaces (Key Go Pattern)

Interfaces let you inject behavior instead of hard-coding dependencies.

#### Example (simplified)

```go
type DataStore interface {
    UserNameForID(id string) (string, bool)
}

type Logger interface {
    Log(msg string)
}

type SimpleLogic struct {
    ds DataStore
    l  Logger
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
    return SimpleLogic{l: l, ds: ds}
}
```

Benefits:

- decoupling
- testability
- mocks or adapters
- clear architecture

#### Tools

- **Wire** (Google) generates dependency injection wiring automatically.

---

## Chapter 8: Generics

_**Why Generics Exist**_

- Generics were added to:
  - reduce **code duplication**

  - increase **type safety**

- They allow writing **generic algorithms** that work on multiple types.

- Generics in Go are **intentionally limited** (by design).

- Go generics are about abstraction of algorithms, not abstraction of behavior.

### TLDR Generics

- Generics abstract **algorithms**, not behavior.
- Use them for reusable data structures and slice algorithms.
- Type constraints control allowed operations.
- Interfaces and generics are complementary.
- Go generics are powerful but deliberately limited.

### Basic Generic Functions

A generic function uses **type parameters**:

```go
func Identity[T any](v T) T {
    return v
}
```

- T is a **type parameter**

- any means “any type”

- Type arguments are usually inferred by the compiler

Usage:

```go
x := Identity(10)      // T = int
y := Identity("hello") // T = string
```

### Generics for Abstract Algorithms

Generics shine for algorithms like:

- map
- filter
- reduce
- search
- stack / queue
- min / max

Example: generic Map over slices

```go
func Map[T any, R any](in []T, fn func(T) R) []R {
    out := make([]R, len(in))
    for i, v := range in {
        out[i] = fn(v)
    }
    return out

}
```

### Generic Data Structures

Generics are useful for reusable data structures:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(v T) {
    s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    v := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return v, true
}
```

### Type Constraints

A type constraint restricts what types a generic parameter can be.

```go
func Equal[T comparable](a, b T) bool {
    return a == b
}
```

Common constraints:

`any` → any type
`comparable` → supports `==` and `!=`

### Interfaces as Constraints

Any interface can be used as a constraint:

```go
type Stringer interface {
    String() string
}

func Print[T Stringer](v T) {
    fmt.Println(v.String())
}
```

Constraints describe **what operations are allowed**, not behavior.

### Type Terms and `~` (Underlying Types)

The `~` operator allows matching underlying types:

```go
type Integer interface {
    ~int | ~int32 | ~int64
}
```

This allows custom types:

```go
type UserID int

func Add[T Integer](a, b T) T {
    return a + b
}
```

Without `~`, `UserID` would not match.

### Operators and Constraints

Operators are **not generic by default**.

You must explicitly allow them via constraints.

```go
type Number interface {
    ~int | ~float64
}
```

Only operations allowed by the constraint can be used.

### Type Inference

- Go usually infers type arguments automatically.
- Explicit type arguments are rarely needed.

```go
Sum([]int{1, 2, 3})      // inferred
Sum[int]([]int{1, 2, 3}) // explicit (rare)
```

### Constant Limits in Generics

Generic functions cannot assume arbitrary constant values.

```go
// INVALID
func PlusOneThousand[T Integer](in T) T {
    return in + 1000
}
// VALID
func PlusOneHundred[T Integer](in T) T {
    return in + 100
}
```

Reason: constant must fit **all possible types** in the constraint.

### Generics + Interfaces (Important Distinction)

- **Interfaces** abstract behavior.
- **Generics** abstract data types and algorithms.

They solve **different problems** and are often used together.

### Comparable Types (More Notes)

- `comparable` includes:
  - basic types
  - structs whose fields are comparable
- Slices, maps, and functions **are not comparable**.

### What Go Generics Do NOT Have (By Design)

Go intentionally excludes:

- No Operator overloading
- No Type parameters on methods
- No Specialization (no overloaded generic + specific versions)
- No Currying
- No Compile-time metaprogramming

Generics in Go are **simple, explicit, and predictable**.

### Idiomatic Use of Generics

- Prefer `any` over `interface{}`
- Use generics only when duplication is real
- Avoid generic APIs that reduce readability
- Many Go APIs do not need generics

Rule of thumb:

**_If the generic version is harder to read than the duplicated code, don’t use generics._**

### When NOT to Use Generics

- Simple business logic
- One-off functions
- APIs where types are already clear
- When interfaces already solve the problem better

---

## Chapter 9: Errors

### Errors in Go: Core Philosophy

- Errors are **values**, not exceptions.
- Functions signal failure by **returning an error**.
- By convention:
  - `error` is the **last return value**
  - `nil` means success

```go
func ReadFile(path string) ([]byte, error) {
    if path == "" {
        return nil, errors.New("empty path")
    }
    return os.ReadFile(path)
}
```

### Creating Errors

Two common ways to create errors from strings:

- `errors.New("message")`
- `fmt.Errorf("formatted %s", value)`

Use `fmt.Errorf` when you need **contextual information**.

### Errors Are Values

- Errors can be:
  - stored in variables
  - passed to functions
  - compared
  - wrapped

- This encourages **explicit error handling**.

```go
if err != nil {
    return err
}
```

> **Idiomatic Go favors clarity over cleverness.**

### Sentinel Errors

Sentinel errors represent **specific, expected failure states**.

- Declared at **package level**
- Treated as **read-only**
- Named with `ErrXxx`

```go
var ErrNotFound = errors.New("not found")
```

Use them when the caller must **react differently** to a specific error.

⚠️ Avoid overusing sentinel errors — they couple callers to your package.

### Wrapping Errors (Very Important)

Wrapping adds **context** while preserving the original error.

```go
return fmt.Errorf("failed to load user: %w", err)
```

- `%w` wraps the error
- Creates an **error chain** (tree)

You can unwrap:

```go
errors.Unwrap(err)
```

Returns `nil` if there is no wrapped error.

### Checking Errors: `errors.Is` and `errors.As`

Use `errors.Is` to check if an error **matches a sentinel**, even when wrapped:

```go
if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

Use `errors.As` to extract a **specific error type**:

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    // access pathErr fields
}
```

Never compare errors with `==` unless you fully control them.

### Wrapping Multiple Errors

To combine several errors into one:

```go
err := errors.Join(err1, err2, err3)
```

- Useful for batch operations
- `errors.Is` and `errors.As` still work

### Wrapping Errors with `defer`

You can add context at the end of a function:

```go
func process() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("process failed: %w", err)
        }
    }()
    return doWork()
}
```

- Reduces repetitive wrapping
- Use carefully (can hide control flow)

### Panic and Recover

- `panic` stops normal execution.
- Deferred functions still run.
- Program exits after unwinding the stack.

```go
panic("something went very wrong")
```

`recover` regains control **only inside a deferred function**:

```go
defer func() {
    if r := recover(); r != nil {
        log.Println("recovered:", r)
    }
}()
```

⚠️ **Do not use panic for normal error handling.**

### Why Panic/Recover Is Discouraged

- Hides what can fail
- Makes control flow unclear
- Breaks explicit error contracts

> Idiomatic Go prefers **explicit errors** over generic recovery.

Special case:

- Panic is acceptable for **programmer bugs** (impossible states).

### `panic(nil)` Edge Case

- `panic(nil)` still panics.
- `recover()` returns `nil`.
- This makes debugging harder → **avoid it**.

### Stack Traces and Errors

The standard `error` type **does not include stack traces**.

Options:

- Let panics print stack traces (last resort)
- Use third-party libraries (e.g. CockroachDB errors)
- Or log stack traces at the boundary (HTTP, CLI, worker)

In Go, stack traces are usually handled at **process boundaries**, not everywhere.

### Error Handling Guidelines (Important)

- Handle errors **once**, at the right level.
- Add context **when crossing boundaries** (IO, DB, HTTP).
- Don’t log and return the same error.
- Prefer wrapping over creating new errors.
- Avoid sentinel errors unless they are part of the API contract.

### Chapter Summary

- Errors are values, returned explicitly.
- Wrap errors to add context.
- Use `errors.Is` / `errors.As`, not `==`.
- Panic is for unrecoverable bugs only.
- Keep error handling explicit and readable.

---

## Chapter 10: Modules, Packages, and Imports

### Packages

- A **package** is the smallest unit of code organization in Go.
- All `.go` files in the same directory belong to the same package.
- Package names:
  - lowercase
  - short
  - descriptive
- Package name acts as a **namespace** when imported.

```go
package user
```

- Visibility rules:
  - **Exported** identifiers → start with uppercase
  - **Unexported** → lowercase (package-private)

Guideline:

> A package should represent **one clear responsibility**.

### Imports

- Imports make identifiers from another package available.
- Standard form:

```go
import "fmt"
```

- Multiple imports use a block:

```go
import (
    "fmt"
    "log"
)
```

- Identifiers are accessed via the package name:

```go
fmt.Println("hello")
```

### Import Aliases

- Used to:
  - avoid name collisions
  - improve readability

```go
import f "fmt"
```

Common in large projects or when importing multiple packages with the same name.

### Blank Imports (`_`)

- Import a package **only for its side effects**.
- Executes the package’s `init()` function.

```go
import _ "github.com/lib/pq"
```

Typical use cases:

- database drivers
- plugin registration
- framework integrations

### init Functions

- `init()` runs automatically when the package is imported.
- No parameters, no return values.
- Can appear multiple times per package.

Typical uses:

- registration
- configuration
- validation

Guideline:

> Avoid business logic in `init()`; keep behavior explicit.

### Modules

- A **module** is a collection of related packages.
- Defined by a `go.mod` file at the module root.
- Modules replace the old `$GOPATH` workflow.

```bash
go mod init github.com/user/project
```

A project usually corresponds to **one module**.

### go.mod

The `go.mod` file defines:

- module path
- Go version
- direct dependencies

```go
module github.com/user/project

go 1.22

require github.com/go-chi/chi v5.0.10
```

- Indirect dependencies are added automatically by Go tooling.

### Module Paths

- Module paths usually match the repository URL.
- Used by Go tooling to locate and download code.
- Imports use:
  - module path
  - plus package subpath

```go
import "github.com/user/project/internal/service"
```

### go.sum

- Contains cryptographic checksums of all dependencies.
- Ensures **reproducible builds**.
- Always committed to version control.
- Automatically maintained by Go tools.

### Minimal Version Selection (MVS)

Go uses **Minimal Version Selection** to resolve dependencies.

Key ideas:

- Go chooses the **minimum version** required to satisfy all dependencies.
- No complex dependency resolution trees.
- No version conflicts at build time.
- Builds are deterministic and fast.

This is a major reason Go dependency management is simple and reliable.

### Updating Dependencies

Common commands:

```bash
go get pkg@latest
go get pkg@v1.2.3
go mod tidy
```

- `go mod tidy`:
  - removes unused dependencies
  - adds missing ones
  - keeps `go.mod` and `go.sum` clean

### Compatible vs Incompatible Versions

- Go follows **semantic versioning**.
- Minor / patch updates should be backward-compatible.
- Breaking changes require a **major version bump**.

For major versions ≥ v2:

- version must appear in the module path

```go
import "github.com/user/lib/v2"
```

This prevents accidental breaking upgrades.

### Versioning and Publishing Modules

- Modules are versioned using Git tags.
- Publishing a module = pushing a tagged version to a repository.
- No central registry like npm.
- Modules are fetched directly from VCS hosts (GitHub, GitLab, etc.).

### Retracting a Module Version

- Allows marking a version as **bad** without deleting it.
- Declared in `go.mod`:

```go
retract v1.2.3
```

Used when:

- a release contains critical bugs
- users should avoid a specific version

### Overriding Dependencies (`replace`)

- Temporarily override a dependency.

```go
replace github.com/user/lib => ../lib
```

Common use cases:

- local development
- testing forks
- debugging dependencies

Guideline:

> Avoid committing `replace` directives meant only for local use.

### Workspaces (`go.work`)

- Allows working on **multiple modules simultaneously**.
- Useful for:
  - monorepos
  - developing multiple related modules together

```bash
go work init
go work use ./moduleA ./moduleB
```

Workspaces avoid `replace` hacks during development.

### Using Private Repositories

- Go supports private modules.
- Requires proper authentication (SSH, tokens).
- You may need to configure:
  - `GOPRIVATE`
  - proxy settings

```bash
export GOPRIVATE=github.com/myorg/*
```

### Module Proxy Servers

- By default, Go uses a **module proxy** (`proxy.golang.org`).
- Proxies:
  - cache modules
  - improve reliability
  - improve security

You can:

- disable proxies
- use private proxies
- configure with environment variables

```bash
GOPROXY=direct
```

### Vendoring

- Vendoring copies dependencies into your repository.
- Enabled via:

```bash
go mod vendor
```

Use cases:

- reproducible builds in restricted environments
- compliance or auditing requirements

Most projects **do not need vendoring**.

### pkg.go.dev

- Official Go documentation site.
- Automatically generates documentation from source code.
- Public modules appear automatically.
- Good documentation = good GoDoc comments.

### internal Packages

- Any package under `internal/`:
  - can only be imported by code within the parent module

- Enforced by the compiler.

Used to:

- protect internal APIs
- prevent misuse

### cmd Directory

- Convention for application entry points.
- Each subdirectory builds a separate binary.

```txt
cmd/api/main.go
cmd/worker/main.go
```

### Idiomatic Project Layout

Typical layout:

```txt
cmd/
internal/
go.mod
go.sum
```

- `cmd/` → binaries
- `internal/` → application code
- modules manage dependencies and versions

### Chapter Summary

- Packages organize code; modules organize packages.
- `go.mod` defines module identity and dependencies.
- MVS guarantees simple and predictable dependency resolution.
- Semantic import versioning prevents breaking upgrades.
- Workspaces and replace help local development.
- Go modules favor stability, simplicity, and explicitness.

---

## Chapter 11: Tooling

Go places a strong emphasis on built-in tooling.
Unlike many ecosystems, most essential tools are part of the standard Go toolchain.

Tooling is one of Go’s biggest strengths.

### The `go` Command

The `go` CLI is the central entry point for:

- Building (`go build`)
- Running (`go run`)
- Testing (`go test`)
- Formatting (`go fmt`)
- Vetting (`go vet`)
- Dependency management (`go mod`)
- Generating code (`go generate`)
- Installing tools (`go install`)

The Go philosophy:
**Convention over configuration.**

### go fmt

- Automatically formats code according to Go standards.
- No configuration.
- Eliminates style debates.

```bash
go fmt ./...
```

> **Formatting is not optional in idiomatic Go.**

### go vet

- Static analyzer.
- Detects suspicious constructs:
  - unreachable code
  - incorrect format strings
  - shadowed variables
  - misused struct tags

```bash
go vet ./...
```

It does not replace a linter but catches common mistakes.

### go test

- Built-in testing framework.
- Tests live in `_test.go` files.
- Supports:
  - benchmarks
  - coverage
  - race detection

```bash
go test ./...
go test -race ./...
go test -cover ./...
```

The race detector is extremely valuable in concurrent code.

### Benchmarks

Defined using `BenchmarkXxx(b *testing.B)`.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(1, 2)
    }
}
```

Run with:

```bash
go test -bench=.
```

Used to measure performance regressions.

### Code Coverage

```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Helps visualize which parts of code are tested.

### The Race Detector

Detects data races in concurrent programs.

```bash
go test -race
```

Very important for production backend systems.

### go generate

Used to generate code automatically.

It does nothing by default unless you add special comments:

```go
//go:generate stringer -type=State
```

Then run:

```bash
go generate ./...
```

Use cases:

- Generating mocks
- Generating string methods for enums
- Generating database code
- Embedding static assets

Important:
`go generate` is not dependency-aware.
It is a developer convenience tool.

### go install

Used to install binaries (including CLI tools):

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Installs into `$GOBIN` or `$GOPATH/bin`.

Common pattern:
Pin tool versions explicitly in documentation or Makefile.

### go list

Inspects modules and packages.

```bash
go list ./...
go list -m all
```

Useful for debugging dependency graphs.

### go doc and pkg.go.dev

- `go doc` shows documentation in terminal.
- Public packages are automatically indexed on [https://pkg.go.dev](https://pkg.go.dev).

Good documentation = exported symbols + proper comments.

### Makefiles

Go doesn’t require Makefiles, but they are commonly used to:

- Standardize commands
- Avoid long CLI commands
- Install tools locally
- Wrap Docker commands
- Manage migrations

Example:

```make
run:
	go run ./cmd/app

test:
	go test ./...

lint:
	golangci-lint run
```

Makefiles improve team consistency.

### Linters (External Tools)

Common tool:

- `golangci-lint` (meta-linter)

Combines many linters:

- ineffassign
- staticcheck
- govet
- unused
- errcheck

Not built into Go but widely adopted in industry.

### Staticcheck

More advanced static analysis than `go vet`.

Catches:

- unused code
- subtle bugs
- incorrect API usage

Often run through `golangci-lint`.

### Profiling (pprof)

Go includes built-in profiling support:

- CPU profiling
- Memory profiling
- Goroutine profiling

Used for performance tuning in production systems.

### Embedding Files (go:embed)

Allows embedding static files into the binary.

```go
//go:embed templates/*
var templates embed.FS
```

Useful for:

- HTML templates
- SQL migrations
- Config files

Very common in REST APIs.

### Tool Philosophy in Go

Key ideas:

- Tooling is part of the language.
- Formatting is standardized.
- Testing is first-class.
- Performance tooling is built-in.
- Simplicity over complex build systems.

Compared to ecosystems like JavaScript:
Go relies far less on external tooling.

### Practical Backend Takeaways

For a production REST API, you will typically use:

- `go test -race`
- `golangci-lint`
- `go generate` (mocks, stringer, SQL tools)
- `go build`
- `pprof` for profiling
- Makefile to unify everything

Tooling discipline is a major part of writing professional Go.

---

## Chapter 12: Concurrency

Concurrency in Go is built around:

- Goroutines
- Channels
- Select
- Synchronization primitives (Mutex, WaitGroup, etc.)
- Context

Go's philosophy:
"Do not communicate by sharing memory; share memory by communicating."

### Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

Start one with the `go` keyword:

```go
go doSomething()
```

Goroutines are:

- Very cheap (thousands can run)
- Scheduled by Go runtime (not OS threads directly)

Important:
The main function exiting stops all goroutines.

Example:

```go
go fmt.Println("Hello")
```

⚠ This may not print if main exits immediately.

### WaitGroup

Used to wait for multiple goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()
    fmt.Println("Working")
}()

wg.Wait()
```

Rules:

- Call Add BEFORE launching goroutine
- Always call Done (usually with defer)
- Wait blocks until counter reaches 0

### Data Races

A data race occurs when:

- Two goroutines access the same variable
- At least one write
- No synchronization

Example:

```go
var counter int

go func() {
    counter++
}()

go func() {
    counter++
}()
```

Use:

```bash
go run -race main.go
```

Race detector only catches races that occur at runtime.

### Mutex

Protect shared memory.

```go
var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()
```

Best practice:

```go
mu.Lock()
defer mu.Unlock()
```

Use mutex when:

- Multiple goroutines modify shared state

`sync.RWMutex` if you need to read and write at the same time

### Channels

Channels allow safe communication between goroutines.

```go
ch := make(chan int)
```

#### Sending

```go
ch <- 42
```

#### Receiving

```go
v := <-ch
```

Unbuffered channels:

- Block until sender and receiver are ready

Buffered channels:

```go
ch := make(chan int, 3)
```

Allow up to 3 values before blocking.

### Channel Direction

Function parameters can restrict channel usage:

```go
func send(ch chan<- int)
func receive(ch <-chan int)
```

Improves type safety.

### Closing Channels

Only the sender should close a channel.

```go
close(ch)
```

Reading from closed channel:

- Returns zero value
- `ok` becomes false

```go
v, ok := <-ch
```

Never close a channel from receiver side.

Never close twice (panic).

### Range over Channels

```go
for v := range ch {
    fmt.Println(v)
}
```

Loop ends when channel is closed.

### Select

Allows waiting on multiple channel operations.

```go
select {
case v := <-ch1:
    fmt.Println(v)
case ch2 <- 42:
    fmt.Println("sent")
default:
    fmt.Println("no communication")
}
```

If multiple cases are ready:

- One is chosen randomly

Used for:

- Timeouts
- Cancellation
- Fan-in patterns

### Timeouts

```go
select {
case result := <-ch:
    fmt.Println(result)
case <-time.After(2 * time.Second):
    fmt.Println("timeout")
}
```

### Context

Used for cancellation and deadlines.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

Pass context to goroutines:

```go
select {
case <-ctx.Done():
    return
}
```

Critical for:

- HTTP servers
- Database calls
- External API calls

### Fan-out / Fan-in Pattern

Fan-out:

- Multiple workers reading from same channel

Fan-in:

- Merge results into single channel

Common worker pool pattern.

### Worker Pool Example

```go
jobs := make(chan int, 5)
results := make(chan int, 5)

for w := 0; w < 3; w++ {
    go worker(jobs, results)
}
```

Used to:

- Limit concurrency
- Control resource usage

### When to Use What

Use goroutines:

- For parallel tasks

Use channels:

- When you want communication-based concurrency

Use mutex:

- When protecting shared state

Prefer channels over shared memory when possible.

### Common Concurrency Bugs

- Forgetting `wg.Done()`
- Closing channel from wrong side
- Writing to closed channel (panic)
- Goroutine leaks (blocked forever)
- Ignoring context cancellation

### Concurrency vs Parallelism

Concurrency:

- Structure of program (multiple tasks at once)

Parallelism:

- Running simultaneously on multiple CPUs

Go supports both.

### Important Mental Model

Goroutines communicate via channels.

Avoid:

- Shared mutable state
- Overusing mutex
- Complex lock hierarchies

Prefer:

- Message passing
- Clear ownership of data

### Production Takeaways

- Always use context in HTTP handlers
- Use race detector in CI
- Control goroutine lifetimes
- Never leak goroutines
- Think about cancellation paths

Concurrency is powerful but dangerous.
Design carefully.

---

## Chapter 13:  Standard Library Packages

This chapter focuses on the most commonly used packages in everyday Go development.

The Go standard library philosophy:

- Small, composable building blocks
- Interfaces over inheritance
- Explicit error handling
- Minimal hidden behavior

Mastering these packages makes you productive without heavy frameworks.

### io

The `io` package defines the fundamental I/O abstractions.

Core interfaces:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

These interfaces are everywhere in the standard library.

Examples of types implementing them:

- `os.File`
- `http.Request.Body`
- `bytes.Buffer`
- `strings.Reader`

Very common helper:

```go
io.Copy(dst, src)
```

This works with anything that satisfies `Reader` and `Writer`.

Important mental model:

> If your function needs input data, accept `io.Reader`.
> If it produces output, accept `io.Writer`.

This makes your code flexible and testable.

Other useful helpers:

- `io.ReadAll`
- `io.MultiReader`
- `io.MultiWriter`
- `io.TeeReader`

### os

The `os` package provides operating system functionality.

Common usage:

#### Files

```go
f, err := os.Open("file.txt")
defer f.Close()
```

`*os.File` implements `io.Reader` and `io.Writer`.

#### Environment Variables

```go
port := os.Getenv("PORT")
```

Use environment variables for configuration in production.

#### Process Control

```go
os.Exit(1)
```

Important:
`os.Exit` does NOT run deferred functions.

### fmt

Used for formatted I/O.

Printing:

```go
fmt.Println("hello")
fmt.Printf("User: %s\n", name)
```

Creating formatted errors:

```go
fmt.Errorf("user %d not found", id)
```

Error wrapping:

```go
fmt.Errorf("failed to load user: %w", err)
```

Important formatting verbs:

- `%v` default representation
- `%+v` detailed (useful with structs)
- `%T` type
- `%w` wrap error

Production note:

`fmt` is convenient but not optimized for structured logging.

### log

The standard `log` package provides simple logging.

```go
log.Println("starting server")
```

You can configure prefix and flags:

```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
```

However:

⚠️ The standard `log` package is minimal and unstructured.

In modern production Go (1.21+), prefer:

### slog (structured logging)

```go
import "log/slog"

logger := slog.Default()
logger.Info("user created", "userID", id)
```

Structured logs:

- Machine-readable
- Better for observability systems
- Key-value based

Other popular logging libraries in industry:

- `zap` (Uber)
- `zerolog`
- `logrus` (older but still used)

Rule of thumb:

- `log` → simple apps
- `slog` / `zap` → production APIs

### strings

The `strings` package manipulates UTF-8 encoded strings.

Common functions:

- `strings.Contains`
- `strings.Split`
- `strings.Join`
- `strings.TrimSpace`
- `strings.ToLower`
- `strings.HasPrefix`
- `strings.ReplaceAll`

Performance tip:

For repeated concatenation, use `strings.Builder`:

```go
var sb strings.Builder
sb.WriteString("Hello ")
sb.WriteString("World")
result := sb.String()
```

Strings in Go are:

- Immutable
- UTF-8 encoded
- Byte slices underneath

### strconv

String ↔ numeric conversions.

Very common in HTTP parameter parsing.

```go
age, err := strconv.Atoi("42")
```

Other examples:

```go
strconv.Itoa(42)
strconv.ParseFloat("3.14", 64)
strconv.ParseBool("true")
```

Always handle errors.

### bufio

Buffered I/O.

Useful for reading files or streams efficiently.

```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}
```

Important limitation:

- `Scanner` has a max token size (default 64K)
- For large lines, use `bufio.Reader`

### encoding/json

Critical for building APIs.

#### Marshalling

```go
data, err := json.Marshal(user)
```

#### Unmarshalling

```go
err := json.Unmarshal(data, &user)
```

Struct tags:

```go
type User struct {
    Name string `json:"name"`
}
```

Important:

- Only exported fields (capitalized) are marshalled.
- Unknown fields are ignored by default.
- Use pointers to distinguish zero value vs missing field.

Advanced usage:

```go
decoder := json.NewDecoder(r)
decoder.DisallowUnknownFields()
```

For large payloads, prefer streaming with `Decoder`.

Production note:

- `encoding/json` is slower than some alternatives (e.g. `jsoniter`)
- But standard library is usually good enough.

### net/http

The core HTTP package.

#### Server

```go
http.HandleFunc("/users", handler)
http.ListenAndServe(":8080", nil)
```

Handler signature:

```go
func(w http.ResponseWriter, r *http.Request)
```

Important concepts:

- `http.Handler` interface
- Middleware pattern
- Request context via `r.Context()`

#### Client

```go
resp, err := http.Get(url)
```

⚠️ In production, do NOT use default client blindly.

Better:

```go
client := &http.Client{
    Timeout: 5 * time.Second,
}
```

Always close body:

```go
defer resp.Body.Close()
```

### time

Time handling and scheduling.

Current time:

```go
time.Now()
```

Parsing:

```go
time.Parse(time.RFC3339, input)
```

Sleeping:

```go
time.Sleep(time.Second)
```

Timers:

- `time.After`
- `time.NewTimer`
- `time.NewTicker`

Production note:

- Avoid `time.After` inside tight loops
- Use `Ticker` for repeated intervals

### path/filepath

OS-independent path handling.

```go
filepath.Join("dir", "file.txt")
```

Important:

- Use `filepath` (not `path`) for file system paths.

### sort

Sorting slices.

```go
sort.Ints(nums)
```

Custom sorting:

```go
sort.Slice(users, func(i, j int) bool {
    return users[i].Age < users[j].Age
})
```

### Practical Production Takeaways

In real-world Go backend development, you constantly use:

- io
- os
- fmt
- log / slog
- strings
- strconv
- encoding/json
- net/http
- time
- errors

The standard library is intentionally powerful enough for most APIs.

You do not need a heavy framework to build production-ready services.

Understanding these packages deeply is more important than knowing a web framework.

---
