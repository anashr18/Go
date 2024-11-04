# Go
Go related experiments
go run
go build
### Package
Only package with name main and func main is executable.
All the other packages are libraries. They are not executable.
These lib can be built and imported in other Go programs.
```go build package```
Files in the same package do not have to be imported into each other.
go main.go state.go
### Array
- array>> fixed size
- slice>> dynamic size
### Key Differences between Slices and Arrays:
- Slices:
  - Passed by reference.
  - Modifying a slice affects the underlying array.
- Arrays:
  - Passed by value.
  - Modifications do not affect the original array unless a pointer is used.
- Slices are mutable, but append() may allocate a new array if the capacity is exhausted.This behavior can feel like immutability because the original slice remains unchanged if the capacity changes.
- The key takeaway is to always reassign the result of append() to avoid holding onto the old slice.
- This design ensures performance efficiency (since Go minimizes allocations) while also maintaining flexibility to grow slices when needed.
### Go is always Pass by value 
- Go always passes variables by value, but what gets copied depends on the type:
  - Primitive types: The value itself is copied.
  - Pointers: The pointer is copied, but it still points to the original value.
  - Slices and Maps: A copy of the slice/map reference is made, but it still points to the same underlying data.
  - Structs: A copy of the struct is passed. To modify the original, pass a pointer to the struct.
- This behavior ensures that Go is simple and predictable, but it can appear as if Go is pass-by-reference in some cases (like with slices or maps). However, it’s important to remember that everything is still pass-by-value—you’re just passing a copy of a reference in those cases.
## Comparison Table of Pointer Syntax

| **Language** | **Pointer Access (Explicit)** | **Pointer Access (Syntactic Sugar)** |
| ------------ | ----------------------------- | ------------------------------------ |
| Go           | `(*p).field`                  | `p.field` (automatic dereferencing)  |
| C++          | `(*p).field`                  | `p->field` (arrow operator)          |


## Summary of Readers and Writers (Go, Python, C++)

| **Use Case**                | **Go Reader**         | **Go Writer**  | **Python Reader**         | **Python Writer**          | **C++ Reader**                   | **C++ Writer**                   |
| --------------------------- | --------------------- | -------------- | ------------------------- | -------------------------- | -------------------------------- | -------------------------------- |
| Reading from a string       | `strings.NewReader()` | N/A            | `io.StringIO()`           | `io.StringIO().write()`    | `std::istringstream`             | N/A                              |
| Reading from a file         | `os.Open()`           | `os.Create()`  | `open(file, 'r')`         | `open(file, 'w')`          | `std::ifstream`                  | `std::ofstream`                  |
| Reading from HTTP response  | `http.Get().Body`     | `http.Post()`  | `requests.get(url).text`  | `requests.post(url, data)` | **cURL library** or sockets      | **cURL library** or sockets      |
| Writing to standard output  | N/A                   | `os.Stdout`    | N/A                       | `sys.stdout.write()`       | N/A                              | `std::cout`                      |
| Writing to a file           | N/A                   | `os.Create()`  | N/A                       | `open(file, 'w').write()`  | N/A                              | `std::ofstream`                  |
| Writing to a buffer         | N/A                   | `bytes.Buffer` | `io.BytesIO()`            | `io.BytesIO().write()`     | `std::ostringstream`             | `std::ostringstream`             |
| Reading from a buffer       | `bytes.NewReader()`   | N/A            | `io.BytesIO().getvalue()` | N/A                        | `std::istringstream`             | N/A                              |
| Reading from terminal input | `os.Stdin`            | N/A            | `input()`                 | N/A                        | `std::cin`                       | N/A                              |
| Network connection (TCP)    | `net.Conn`            | `net.Conn`     | `socket.socket().recv()`  | `socket.socket().send()`   | **`boost::asio`** or **sockets** | **`boost::asio`** or **sockets** |

## Why Channels Can Be Slower Than Mutexes
Blocking Nature:

Each goroutine must wait for a send or receive to complete before proceeding. If multiple goroutines try to update a shared value via a channel, only one can access the value at a time, causing others to wait.
This serialization ensures safety, but it can become a bottleneck when you have high concurrency (many goroutines competing for access).
Overhead of Communication:

Passing data through channels involves communication overhead because the data must be passed back and forth between goroutines.
This can be slower than using a mutex, where the data remains in-place and only access is coordinated.
Goroutine Scheduling:

When a goroutine sends or receives from a channel, it may block and be put to sleep by the Go scheduler, only to be rescheduled later. This scheduling overhead can add latency compared to the simple locking mechanism of a mutex.

## Use fresh for hotloading
``go install github.com/pilu/fresh@latest``
