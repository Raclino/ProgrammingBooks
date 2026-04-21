# Concurrency in Go

by Katherine Cox-Buday

## Chapter 1: An Introduction to Concurrency

### What is concurrency?

Concurrency is about **structuring a program so that multiple tasks can make progress independently**.

It does **not** necessarily mean that those tasks are running at the exact same time.

A concurrent program:

- breaks work into independent units
- allows those units to be coordinated safely
- improves responsiveness, composability, and resource utilization

### Concurrency vs Parallelism

These two terms are related, but they are **not the same**.

#### Concurrency

Concurrency is about the **design** of a program.

It means:

- multiple tasks are in progress
- tasks may be paused, resumed, or interleaved
- the program is organized to handle more than one thing at once

#### Parallelism

Parallelism is about the **execution** of a program.

It means:

- multiple tasks are literally running at the same time
- usually on multiple CPU cores

A program can be:

- **concurrent but not parallel**
- **parallel but not well-structured**
- **both concurrent and parallel**

A common way to think about it:

> Concurrency is about dealing with many things at once.  
> Parallelism is about doing many things at once.

### Why concurrency is difficult

Concurrency introduces complexity because multiple parts of a program may:

- access the same memory
- depend on timing
- interact in non-deterministic ways

This makes programs harder to:

- reason about
- test
- debug

A concurrent bug may appear:

- only sometimes
- only under load
- only on certain machines
- only when timing happens in an unlucky way

### The main risks of concurrent programming

#### Race conditions

A race condition happens when the behavior of a program depends on the timing or order of operations between concurrent tasks.

The result can become unpredictable.

#### Data races

A data race is a specific kind of race condition:

- two or more goroutines access the same memory
- at least one access is a write
- there is no proper synchronization

Data races are dangerous because they can produce corrupted or inconsistent state.

#### Deadlocks

A deadlock happens when multiple parts of a program are waiting on each other forever.

No progress can be made.

Typical example:

- goroutine A waits for goroutine B
- goroutine B waits for goroutine A

#### Livelocks

A livelock happens when goroutines are still active, but they keep reacting to each other in a way that prevents real progress.

The program is not blocked, but it is still stuck.

#### Starvation

Starvation happens when a goroutine never gets the resources it needs to make progress.

For example:

- one goroutine constantly gets access to a lock
- another goroutine keeps waiting indefinitely

### Determinism vs non-determinism

Sequential code is usually easier to reason about because execution order is obvious.

Concurrent code is different:

- scheduling is not fully under your control
- execution order may vary from run to run
- the same code may produce different timing behavior

This is one of the core reasons concurrency is hard.

### Is concurrency always worth it?

No.

Concurrency is useful when it helps with:

- responsiveness
- throughput
- coordination of independent tasks
- handling I/O efficiently
- modeling real-world systems more naturally

But it also adds:

- mental overhead
- debugging difficulty
- synchronization complexity
- more failure modes

You should not introduce concurrency unless it solves a real problem.

### Concurrency is not a silver bullet

Adding concurrency does not automatically make a program:

- faster
- simpler
- more scalable

In some cases, concurrent code can be:

- slower because of coordination overhead
- harder to maintain
- more bug-prone than a simple sequential solution

Good concurrent design is about **correctness first**, then performance.

### The role of abstraction

One of the main ideas in concurrent programming is to use abstractions that make coordination safer and easier to reason about.

In Go, the main tools are:

- **goroutines** for lightweight concurrent execution
- **channels** for communication and synchronization
- **select** for coordinating multiple channel operations
- synchronization primitives from `sync` when needed

The book strongly pushes toward a design where concurrency is **structured and communicated clearly**, rather than scattered through the codebase.

### Communication vs shared memory

A central idea in Go is:

> **_Do not communicate by sharing memory; instead, share memory by communicating._**

This means:

- instead of many goroutines mutating the same data directly
- prefer passing data between goroutines through channels
- this reduces some classes of bugs and often makes ownership clearer

This does **not** mean shared memory is forbidden.
It means channel-based communication is often the cleaner default.

### Safety and simplicity

A good concurrent program should aim for:

- clear ownership of data
- minimal shared mutable state
- explicit synchronization
- predictable lifecycle of goroutines
- easy cancellation and cleanup

The simpler the concurrency model, the easier it is to trust.

### Concurrency in Go

Go was designed with concurrency as a first-class concern.

This is one of the language's major strengths:

- goroutines are lightweight
- channels are built into the language
- the standard library includes useful synchronization tools
- the runtime scheduler helps manage goroutines efficiently

This makes Go especially well-suited for:

- servers
- pipelines
- network services
- background workers
- streaming systems
- concurrent CLI tools

### Key takeaways

- Concurrency is about **structure**, not necessarily simultaneous execution.
- Parallelism is about **simultaneous execution**.
- Concurrent programs are harder to reason about because of **non-determinism**.
- The main hazards are:
  - race conditions
  - data races
  - deadlocks
  - livelocks
  - starvation
- Concurrency should be introduced to solve a real problem, not by default.
- In Go, channels and goroutines provide a strong model for structuring concurrent programs.
- Correctness and clarity matter more than premature optimization.

---

## Chapter 2: Modeling Your Code - Communicating Sequential Processes

### Main idea

This chapter explains the concurrency model Go is built around.

The main point is that **concurrency is about structuring code**, while **parallelism is about execution**.

- **Concurrency**: multiple tasks can progress independently
- **Parallelism**: multiple tasks run at the same time on multiple CPU cores

A program can be concurrent without being parallel.

### CSP

Go is heavily influenced by **CSP** (_Communicating Sequential Processes_).

The idea is to model a program as independent units of work that **communicate explicitly**, instead of having many concurrent units directly share and mutate the same memory.

This leads to the core Go principle:

> **_Do not communicate by sharing memory; instead, share memory by communicating._**

The goal is to make concurrent code easier to reason about and safer.

### Goroutines

A **goroutine** is a lightweight concurrent function execution managed by the Go runtime.

```go
go myFunction()
```

Goroutines are cheap to create compared to traditional threads, which makes concurrency much more practical in Go.

But starting goroutines is easy: the hard part is designing how they communicate, synchronize, and stop correctly.

### Channels

A _channel_ is used for communication and synchronization between goroutines.

```go
messages := make(chan string)

go func() {
messages <- "hello"
}()

fmt.Println(<-messages)
```

Channels are important because they:

- pass data between goroutines
- synchronize execution
- reduce the need for shared mutable state

### Shared memory vs communication

Go does not forbid shared memory and locks, but it encourages a model where concurrent parts of the program exchange data explicitly through channels.

This often makes:

- ownership clearer
- synchronization more explicit
- code easier to reason about

### Runtime and scheduler

Goroutines are not raw OS threads.

The Go runtime schedules goroutines onto underlying threads, which is why goroutines are lightweight and scalable.

This means:

- you structure work with goroutines
- the runtime handles scheduling
- actual parallelism depends on the machine and available CPU cores

### Key takeaways

- **Concurrency** = structure
- **Parallelism** = execution
- Go’s concurrency model is strongly influenced by CSP
- In Go, concurrency is often modeled through communication, not shared memory
- **Goroutines** are lightweight concurrent tasks
- **Channels** are used to communicate and synchronize between goroutines
- The main challenge in concurrency is not launching work, but coordinating it correctly

> In Go, I should first think about how concurrent tasks communicate and synchronize, not just how to run code at the same time.

## Chapter 3: Go’s Concurrency Building Blocks

### Main idea

This chapter introduces the main concurrency primitives provided by Go.

The core idea is that Go gives two broad ways to coordinate concurrent work:

- **memory access synchronization** with the `sync` package
- **communication and synchronization** with **channels** and `select`

In general, Go encourages:

- simple goroutines
- explicit communication
- limited shared mutable state

### Goroutines

A **goroutine** is a lightweight concurrent unit of execution.

```go
go myFunction()
```

### Important ideas

- every Go program starts with at least one goroutine: `main`
- goroutines are lightweight compared to OS threads
- launching a goroutine is easy, but coordinating it correctly is the hard part
- goroutines run in the same address space, so shared memory must still be synchronized

A common way to wait for goroutines to finish is `sync.WaitGroup`.

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
defer wg.Done()
fmt.Println("hello")
}()

wg.Wait()
```

**Closure gotcha**
Closures capture variables by reference, not by value.

This is a classic bug:

```go
for _, salutation := range []string{"hello", "greetings", "good day"} {
    go func() {
      fmt.Println(salutation)
    }()
}
```

The safer version is to pass the value explicitly:

```go
for _, salutation := range []string{"hello", "greetings", "good day"} {
    go func(s string) {
      fmt.Println(s)
    }(salutation)
}
```

### The sync package

The `sync` package is used for **low-level memory synchronization**.

It is useful when multiple goroutines must coordinate access to shared memory.

#### `WaitGroup`

Used to wait for a group of goroutines to complete.

Main methods:

- `Add(n)`
- `Done()`
- `Wait()`

#### `Mutex`

Used to protect a critical section so only one goroutine can access it at a time.

```go
var mu sync.Mutex
var count int

mu.Lock()
count++
mu.Unlock()
```

Use it when multiple goroutines read/write the same shared state.

#### `RWMutex`

A read-write mutex:

- multiple readers allowed at the same time
- only one writer at a time
- no readers while writing

Useful when:

- reads are frequent
- writes are rare

#### `Cond`

A condition variable used to signal goroutines that some condition has changed.
Less commonly used directly in everyday Go code than channels, but useful for advanced coordination around shared state.

Main methods:

- `Wait()`
- `Signal()`
- `Broadcast()`

#### `Once`

Ensures a piece of code runs exactly one time.

```go
var once sync.Once

once.Do(func() {
fmt.Println("initialized once")
})
```

Useful for:

- lazy initialization
- singleton-like setup
- one-time startup logic

`Pool`

`sync.Pool` is a concurrent-safe object pool.

Useful for reusing temporary objects and reducing allocations/GC pressure.

Typical use case:

- short-lived reusable buffers
- high-throughput paths

But it should not be the default solution for everything.

### Channels

A channel is a typed conduit used to send and receive values between goroutines.

```go
messageStream := make(chan string)

go func() {
messageStream <- "hello"
}()

fmt.Println(<-messageStream)
```

Channels are important because they provide both:

- **communication**
- **synchronization**

#### Unbuffered channels

An unbuffered channel requires sender and receiver to rendezvous.

This means:

- sending blocks until another goroutine receives
- receiving blocks until another goroutine sends

This makes unbuffered channels useful for synchronization.

#### Buffered channels

A buffered channel has a capacity:

```go
messageStream := make(chan string, 4)
```

This means sends can proceed without an immediate receiver, up to the channel capacity.
Buffered channels are useful when:

- producers and consumers do not run at the exact same speed
- you want limited decoupling between stages

But they **do not** remove the need for proper design.

#### Channel ownership

A very important Go idea is channel ownership.
A good rule is:

- the goroutine that creates a channel should usually own it
- the owner writes to it
- the owner closes it
- other goroutines typically only read from it

This helps avoid bugs such as:

- writing to a closed channel
- closing a channel multiple times
- confusion about lifecycle

#### Directional channels

Go allows restricting channel direction in function signatures:

```go
var readStream <-chan int
var writeStream chan<- int
```

This improves API clarity and type safety:

- receive-only channel: `<-chan T`
- send-only channel: `chan<- T`

#### Closing channels

A channel can be closed with:

```go
close(ch)
```

Important rules:

- receivers can still read remaining buffered values
- receiving from a closed channel returns the zero value after the channel is drained
- only the sender / owner should usually close the channel
- never close a channel from the receiver side unless ownership is explicit

A common form:

```go
v, ok := <-ch
```

- `ok == true`: value received normally
- `ok == false`: channel is closed and drained

`for range` **on channels**
A common pattern for consuming values until a channel is closed:

```go
for v := range ch {
fmt.Println(v)
}
```

This keeps receiving until the channel is closed.

#### The Select statement

`select` lets a goroutine wait on multiple channel operations at once.

```go
select {
case msg := <-c1:
fmt.Println(msg)
case c2 <- "hello":
fmt.Println("sent")
default:
fmt.Println("no channel ready")
}
```

Important ideas:

- if multiple cases are ready, one is chosen pseudo-randomly
- if no case is ready and there is no `default`, `select` blocks
- `default` makes the `select` non-blocking

Useful for:

- waiting on multiple channels
- timeouts
- cancellation
- multiplexing concurrent events

Special case:

```go
select {}
```

This blocks forever.

#### Channel operation behavior by state

| Operation | Channel state      | Result                            |
| --------- | ------------------ | --------------------------------- |
| Read      | `nil`              | Blocks forever                    |
| Write     | `nil`              | Blocks forever                    |
| Close     | `nil`              | Panic                             |
| Read      | Open and empty     | Blocks until a value is available |
| Write     | Open and full      | Blocks until space is available   |
| Read      | Closed and drained | Returns zero value immediately    |
| Write     | Closed             | Panic                             |
| Close     | Closed             | Panic                             |

**Important**:

- Reading from a closed channel is valid.
- Writing to a closed channel causes a panic.
- Closing a `nil` channel or an already closed channel also causes a panic.
- Operations on a `nil` channel block forever.

### GOMAXPROCS

`GOMAXPROCS` controls how many OS threads can execute Go code simultaneously.

Main idea:

- it affects parallelism
- not the logical structure of concurrency

So:

- goroutines are about concurrency
- `GOMAXPROCS` influences how much true parallel execution is possible

Most of the time, you do not need to tweak it manually unless you have a specific performance reason.

### Key takeaways

- **Goroutines** are lightweight concurrent tasks
- `sync` primitives are used for **shared memory synchronization**
- **Channels** are used for **communication + synchronization**
- Prefer clear ownership and simple coordination
- **Unbuffered channels** synchronize sender and receiver directly
- **Buffered channels** allow limited decoupling
- **Directional channels** improve API clarity
- `select` is used to coordinate multiple channel operations
- `GOMAXPROCS` affects parallel execution, not the concurrency model itself

## Chapter 4: Concurrency Patterns in Go

### Main ideas

This part of the chapter introduces recurring concurrency patterns that help structure Go programs safely.

The goal is not just to launch goroutines, but to make them:

- easier to reason about
- easier to stop cleanly
- easier to compose
- less likely to leak or block forever

The patterns covered here are especially important because they appear again and again in real Go code:

- limiting access to shared data (**confinement**)
- structuring event loops (**for-select loop**)
- ensuring goroutines can terminate (**preventing goroutine leaks**)
- combining cancellation signals (**or-channel**)
- propagating failures cleanly (**error handling**)

### Confinement

#### Definition

**Confinement** means making sure that a piece of data is only accessible from one concurrent context, or from a very limited and controlled set of goroutines.

The idea is simple:

**if data is not shared freely, it does not need heavy synchronization.**

Instead of protecting shared mutable state everywhere with mutexes, you design the program so that ownership is clear.

This usually makes code:

- safer
- easier to reason about
- less prone to race conditions

#### Why confinement matters

A lot of concurrency bugs come from multiple goroutines being able to:

- read the same state
- mutate the same state
- depend on timing

Confinement reduces this risk by making ownership explicit.

Rather than saying:

- “everyone can access this, but be careful”

you say:

- “only this goroutine owns this data”

That is often much simpler.

#### Types of confinement

##### 1. Ad hoc confinement

This is confinement by convention.

The code relies on discipline and documentation:

- one goroutine is _supposed_ to own the data
- other goroutines are _supposed_ not to touch it directly

This can work, but it is weaker because the compiler does not help much.

Example idea:

- a slice is created in one goroutine
- only that goroutine is expected to mutate it
- other goroutines only receive copies or derived values

##### 2. Lexical confinement

This is stronger and preferred.

The scope of the variable itself limits who can access it.

If a channel or variable is created inside a function and only a read-only view is exposed, then the API itself enforces the confinement.

This is one of the cleanest patterns in Go.

Example:

```go
chanOwner := func() <-chan int {
  results := make(chan int, 5)
  
  go func() {
    defer close(results)
    for i := 0; i < 5; i++ {
      results <- i
    }
  }()

  return results
}

consumer := func(results <-chan int) {
  for result := range results {
    fmt.Printf("Received: %d\n", result)
  }
}

results := chanOwner()
consumer(results)
```

#### Why this example matters

- the goroutine that creates the channel owns it
- only the owner writes to it
- only the owner closes it
- consumers only receive from it

This prevents many common channel mistakes:

- writing to a closed channel
- closing a channel from the wrong side
- mixing responsibilities

#### Key idea of confinement

A good rule in Go is:

- keep mutable state local when possible
- pass data through channels when possible
- expose the smallest interface possible

Confinement is really about **ownership**.

If ownership is obvious, concurrency becomes much easier to understand.

### The for-select loop

#### Definition

A **for-select loop** is a common Go pattern used when a goroutine needs to:

- repeatedly wait for events
- react to multiple channels
- handle cancellation
- continue running until explicitly told to stop

General form:

```go
for {
  select {
  case <-done:
    return
  case msg := <-messages:
    fmt.Println(msg)
  }
}
```

This is one of the most common building blocks in Go concurrency.

#### Why it is useful

Without `select`, a goroutine can only block on one channel operation at a time.

With a `for-select` loop, a goroutine can continuously:

- wait for work
- listen for cancellation
- respond to multiple inputs
- avoid getting stuck on only one communication path

This makes it a natural structure for:

- workers
- background services
- stream processors
- long-running goroutines

#### Typical use case

A goroutine often needs two things:

- a way to receive useful input
- a way to know when to stop

That is exactly where the `for-select` loop shines.

Example:

```go
done := make(chan struct{})
messages := make(chan string)

go func() {
  for {
    select {
    case <-done:
      fmt.Println("worker stopping")
      return
    case msg := <-messages:
      fmt.Println("received:", msg)
    }
  }
}()
```

This goroutine:

- keeps running
- handles incoming messages
- exits cleanly when `done` is closed

#### Important subtlety

A `for-select` loop is often paired with a `done` channel because otherwise the goroutine may have no safe termination path.

That is a major theme of this chapter:

**every goroutine should have a clear exit strategy.**

#### Variants

A `for-select` loop can also:

- multiplex several inputs
- use a `default` case for non-blocking behavior
- handle timers or timeouts
- coordinate heartbeat and work channels

Example with two inputs:

```go
for {
  select {
  case <-done:
    return
  case msg := <-messages:
    fmt.Println("message:", msg)
  case err := <-errs:
    fmt.Println("error:", err)
  }
}
```

### Preventing goroutine leaks

#### What is a goroutine leak?

A goroutine leak happens when a goroutine:

- is still alive
- is blocked forever or effectively forever
- can no longer make useful progress
- is never cleaned up

This is similar in spirit to a memory leak, except the leaked resource is a goroutine.

Leaked goroutines are dangerous because they consume:

- memory
- scheduling overhead
- file descriptors or network resources in some cases
- mental clarity in the codebase

#### Common cause

The most common cause is this:

- a goroutine is waiting on a channel operation
- but nothing will ever unblock it

For example:

```go
func leak() {
  ch := make(chan int)

  go func() {
    fmt.Println(<-ch) // blocks forever
  }()
}
```

Here, the goroutine waits forever because nobody sends on `ch`.

#### Why leaks matter

Because goroutines are cheap, it is tempting to think leaks are not serious.

But in long-running systems, leaked goroutines can accumulate and cause:

- increased memory usage
- performance degradation
- hidden bugs that are hard to trace

Cheap does not mean free.

#### Core principle

**Never start a goroutine without knowing how it will stop.**

Before launching a goroutine, ask:

- what work is it waiting for?
- under what condition does it terminate?
- who signals shutdown?
- what happens if downstream code stops consuming?

If you cannot answer these clearly, the goroutine design is incomplete.

#### Using a done channel to prevent leaks

A common approach is to pass a `done` channel to any goroutine that may otherwise block indefinitely.

Example:

```go
doWork := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
  terminated := make(chan struct{})

  go func() {
    defer fmt.Println("doWork exited")
    defer close(terminated)

    for {
      select {
      case s := <-strings:
        fmt.Println(s)
      case <-done:
        return
      }
    }
  }()

  return terminated
}
```

This pattern matters because:

- the goroutine does useful work
- but it also listens for cancellation
- callers can close `done` to ensure cleanup

#### Important mindset

When designing pipelines or workers, always think in terms of cancellation propagation.

If upstream stops early, downstream or sibling goroutines must not be left hanging forever.

This is why cancellation channels are everywhere in Go concurrency patterns.

### The or-channel

#### Problem it solves

Sometimes a goroutine depends on **multiple cancellation signals**.

For example, maybe work should stop when:

- parent context is cancelled
- timeout fires
- some external worker fails
- another done channel closes

Instead of manually selecting over many done channels everywhere, the idea is to combine them into a single channel:

- if **any** input channel closes
- the combined channel closes too

This is the purpose of the **or-channel** pattern.

#### Main ideas

The or-channel pattern takes multiple channels and returns one channel that closes as soon as any of them closes.

This allows code to depend on a single cancellation signal.

Conceptually:

```go
or(done1, done2, done3)
```

means:

- close as soon as `done1` closes
- or `done2`
- or `done3`

#### Why it is useful

It improves composition.

Instead of wiring cancellation logic repeatedly in every part of the code, you can create one combined signal and pass it around.

This is especially useful in:

- recursive concurrency patterns
- pipelines
- systems with multiple possible shutdown sources

#### Example

A classic recursive version:

```go
var or func(channels ...<-chan interface{}) <-chan interface{}

or = func(channels ...<-chan interface{}) <-chan interface{} {
  switch len(channels) {
  case 0:
    return nil
  case 1:
    return channels[0]
  }

  orDone := make(chan interface{})

  go func() {
    defer close(orDone)

    switch len(channels) {
    case 2:
      select {
      case <-channels[0]:
      case <-channels[1]:
      }
    default:
      select {
      case <-channels[0]:
      case <-channels[1]:
      case <-channels[2]:
      case <-or(append(channels[3:], orDone)...):
      }
    }
  }()

  return orDone
}
```

Usage idea:

```go
sig := func(after time.Duration) <-chan interface{} {
  c := make(chan interface{})
  go func() {
    defer close(c)
    time.Sleep(after)
  }()
  return c
}

start := time.Now()
<-or(
  sig(2*time.Hour),
  sig(5*time.Minute),
  sig(1*time.Second),
  sig(1*time.Hour),
)

fmt.Printf("done after %v\n", time.Since(start))
```

The combined channel closes as soon as the earliest signal fires.

#### Key takeaway

The or-channel pattern is about **composability of cancellation**.

It helps transform many stop signals into one stop signal.

### Error handling

#### Problem

In concurrent programs, errors are harder to handle than in sequential code.

In sequential code:

- function returns value + error
- caller handles it immediately

In concurrent code:

- work may happen in another goroutine
- result and error may arrive later
- multiple goroutines may fail independently
- you may need to stop everything when one part fails

So concurrent error handling needs structure.

#### Main principle

Errors should be treated as part of the communication flow.

If a goroutine produces data, it often also needs a way to communicate:

- success
- failure
- termination

A common Go approach is to send a struct containing both result and error.

#### Example: result + error together

```go
type Result struct {
  Error    error
  Response *http.Response
}

checkStatus := func(done <-chan struct{}, urls ...string) <-chan Result {
  results := make(chan Result)

  go func() {
    defer close(results)

    for _, url := range urls {
      var result Result
      resp, err := http.Get(url)
      result = Result{
        Error:    err,
        Response: resp,
      }

      select {
      case <-done:
        return
      case results <- result:
      }
    }
  }()

  return results
}
```

Consumer side:

```go
done := make(chan struct{})
defer close(done)

urls := []string{
  "https://www.google.com",
  "https://badhost",
  "https://www.example.com",
}

for result := range checkStatus(done, urls...) {
  if result.Error != nil {
    fmt.Println("error:", result.Error)
    continue
  }
  fmt.Println("response:", result.Response.Status)
}
```

#### Why this pattern is good

It keeps the data flow explicit.

Instead of:

- hidden shared error variables
- unclear synchronization
- partially handled failures

you send a value that fully describes the outcome of one unit of work.

That aligns well with Go’s concurrency philosophy.

#### Important design question

When an error happens in concurrent code, you need to decide:

- should this goroutine stop only itself?
- should it report the error and continue?
- should it cancel sibling goroutines?
- should the whole pipeline shut down?

This is not syntax. It is design.

A good concurrent design makes this decision explicit.

#### Common rule of thumb

When possible:

- keep error reporting close to the work that produced it
- communicate errors through channels instead of shared state
- make cancellation possible if one failure should stop the whole operation

### Key takeaways

- **Confinement** reduces shared mutable state by making ownership explicit
- **Lexical confinement** is stronger than ad hoc confinement because the API itself restricts access
- A **for-select loop** is the standard way to build long-running goroutines that handle work and cancellation
- A goroutine should always have a clear termination path
- **Goroutine leaks** happen when goroutines are left blocked forever or with no useful exit path
- A **done channel** is a common way to propagate cancellation
- The **or-channel** pattern combines multiple cancellation signals into one
- In concurrent code, **errors should be part of the communication flow**
- A common pattern is to send a struct containing both result data and an `error`

The main thing I should remember is:

**good concurrent Go code is usually explicit about ownership, cancellation, and communication.**
