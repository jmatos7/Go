# 📘 Learning Go (Golang)

## 🚀 Objective

This repository documents my journey learning **Go (Golang)** — from fundamental concepts to advanced topics like concurrency and backend development.

The goal is to deeply understand Go’s design philosophy, performance model, and concurrency system while building practical projects.

---

## 📦 Installing Go

1. Download Go from the official website:  
   👉 https://go.dev/dl/

2. Run the installer and follow the default setup instructions.

3. Verify the installation:

```bash
go version
```

If correctly installed, you should see something like:

```bash
go version go1.22.0 windows/amd64
```

---

## 🧠 What is Go?

Go is a:

- Compiled language  
- Statically typed language  
- Simple and minimalistic  
- Designed with built-in concurrency  
- Created at Google in 2007  

Go is commonly used for:

- Backend development  
- REST APIs  
- Cloud-native applications  
- Distributed systems  
- DevOps tooling  

---

# 📚 Learning Progress

---

## 1️⃣ Basic Program Structure

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### What I learned:

- `package main` defines an executable program
- `import` allows access to standard libraries
- `func main()` is the entry point
- `fmt.Println()` prints to the console

---

## 2️⃣ Variables and Types

```go
var name string = "John"
age := 21
```

### What I learned:

- Go is statically typed
- Type inference with `:=`
- Difference between explicit declaration and short declaration
- Basic types: `int`, `float64`, `string`, `bool`

---

## 3️⃣ Functions

```go
func greet(name string) string {
    return "Hello " + name
}
```

### What I learned:

- Parameters must be typed
- Functions can return values
- Go supports multiple return values

Example:

```go
func divide(a, b float64) (float64, error)
```

---

## 4️⃣ Slices

```go
names := []string{"Alice", "Bob", "Carlos"}
```

### What I learned:

- Slices are dynamic arrays
- Using `append()`
- Looping with `for` and `range`

Example:

```go
for index, value := range names {
    fmt.Println(index, value)
}
```

---

## 5️⃣ Structs

```go
type Task struct {
    ID        int
    Name      string
    Completed bool
}
```

### What I learned:

- Creating custom data types
- Organizing structured data
- Preparing for JSON and API work

---

## 6️⃣ Concurrency – Goroutines & Channels

### Goroutines

```go
go myFunction()
```

- Runs a function concurrently

---

### Channels

```go
c := make(chan string)

c <- "message"   // send
msg := <-c       // receive
```

### What I learned:

- Safe communication between goroutines
- Automatic synchronization
- Go’s CSP model (Communicating Sequential Processes)
- No manual thread handling required

---

# 🎮 Project 1 – Number Guessing Game

## Goal

- Generate a random number
- Create multiple players
- Use goroutines
- Use channels to communicate results
- Synchronize execution

## Concepts Practiced

- `math/rand`
- `time.Sleep`
- Concurrent execution
- Channel-based communication
- Waiting for multiple goroutines

---

# 🛠 Running the Project

Run directly:

```bash
go run main.go
```

Build executable:

```bash
go build main.go
```

Run executable:

```
./main      (Linux/macOS)
main.exe    (Windows)
```

---

# 📘 Mini Go Web Server

## 🚀 Objective

This project is a mini HTTP server built in Go using only the standard library (`net/http`).  

It demonstrates:

- Handling HTTP requests
- Dynamic routes with query parameters
- Returning JSON responses
- HTTP status codes
- Simple server structure ready for expansion

---

## 🏗 How to Run

```bash
go run main.go
```

Open your browser or use a tool like curl:

- Home route:  
    - http://localhost:8080/

- Hello route with query parameter:  
    - http://localhost:8080/hello?name=Joao

---

## 📚 Concepts Learned

### 1️⃣ Handlers

Handlers are functions that respond to HTTP requests:

```go
func(w http.ResponseWriter, r *http.Request)
```
- w → used to write the response
- r → contains request data (URL, headers, body, etc.)

### 2️⃣ Query Parameters

Example: /hello?name=Joao

```go
name := r.URL.Query().Get("name")
```
- Reads parameters from the URL
- Defaults to "Guest" if none is provided

### 3️⃣ HTTP Methods & Status Codes

```go
if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
```

- Validates the request method
- Sends correct HTTP status codes (200, 405, etc.)

### 4️⃣ Returning JSON

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(response)
```

- w.Header().Set("Content-Type", "application/json")
- w.WriteHeader(http.StatusOK)
- json.NewEncoder(w).Encode(response)
