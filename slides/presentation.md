class: center, middle, inverse, small-images

# Go Simple! Go Fast!

### An introduction to Golang and concurrency

<div style="display: flex; justify-content: center; margin-top: 3em; align-items: center; gap: 1em;">
<img src="./assets/.png">
<div style="font-size: 2.5em; padding-inline: 0.5em;">❤️</div>
<img src="./assets/flutter-logo.png">
</div>

---

### Your Hosts

<div style="display: flex; justify-content: center; gap: 20px;">
    <div style="text-align: center;">
        <img src="./assets/andre.jpeg" style="height: 150px;">
        <p><strong>Andre Moreira</strong><br>Software Engineer @ VertsaPlay</p>
    </div>
    <div style="text-align: center;">
        <img src="./assets/eduardo.jpeg" style="height: 150px;">
        <p><strong>Eduardo Guedes</strong><br>Software Engineer @ VertsaPlay</p>
    </div>
</div>

<div style="text-align: center; margin-top: 80px;">
    <img src="./assets/vertsa.jpeg" style="width: 400px;">
</div>

---


class: center, middle, inverse

## Don't let this be a monologue

#### Ask questions whenever you want

---

### What is Flutter?

- A framework developed by Google for building near-native **cross-platform** apps
- Uses **Dart**, a statically-typed, compiled, null-safe programming-language built with UI development in mind
- Provides facilities such as **hot-reload** and **hot-restart** for fast development
- Easy to learn

<img src="./assets/flutter_arch.png" style="width: 100%;">

---

### Dart 101

- Dart is an **imperative programming language**
- It is **object-oriented** as well
- It supports **null-safety**

```dart
class Person {
  Person(this.age, this.name, this.height, this.surname);

  int age;
  String name;
  String? surname;
  double height;

  String getName() => this.name;

  String? getSurname() => this.surname
}

void main() {
  Person recruta = Person(18, "Recruta", 1.8);

  print(recruta.getName());
}
```

---

class: center, middle

# Basic Types in Go

---

## Integers (int, uint)

- `int` (signed integers, platform-dependent size)
- `int8`, `int16`, `int32`, `int64` (fixed-size signed integers)
- `uint` (unsigned integers, platform-dependent size)
- `uint8` (alias for `byte`), `uint16`, `uint32`, `uint64` (fixed-size unsigned integers)
- Used for counting, indexing, and mathematical operations

```go
var a int = 42
var b uint = 100
var c int64 = -5000
fmt.Println(a, b, c)
```

---

## Floating Point Numbers

- `float32` (single precision, ~7 decimal digits)
- `float64` (double precision, ~15 decimal digits)
- Used for precise mathematical calculations and measurements

```go
var pi float64 = 3.14159
var temp float32 = 36.6
fmt.Println(pi, temp)
```

---

## Strings

- Immutable sequence of bytes
- Supports UTF-8 encoding
- Can be concatenated using `+`
- Length can be determined using `len()`
- Characters can be accessed as bytes

```go
var name string = "GoLang"
fmt.Println(len(name))  // String length
fmt.Println(name[0])    // Access character (byte)
fmt.Println(name + " is awesome!")
```

---

## Arrays

- Fixed-size collection of elements of the same type
- Cannot be resized after declaration
- Not commonly used

```go
var arr [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Println(arr)
fmt.Println(len(arr))  // Get the length of the array
```

---

## Slices

- Dynamic array with flexible length
- Built-in functions: `append()`, `len()`, `cap()`, `copy()`
- More powerful than arrays as they can grow dynamically

```go
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)  // Adding elements
fmt.Println(nums)          // [1 2 3 4 5]
fmt.Println(len(nums))     // Length of slice
fmt.Println(cap(nums))     // Capacity of slice
```

---

## Maps

- Key-value pairs (similar to dictionaries/hash tables in other languages)
- Unordered collection of elements
- Dynamic size (can grow and shrink)
- Keys must be comparable types (string, int, etc.)

```go
myMap := map[string]int{"one": 1, "two": 2}
myMap["three"] = 3  // Add new element
value, exists := myMap["key"]  // Check if key exists
delete(myMap, "one")  // Remove element
fmt.Println(len(myMap))  // Get the length of the map
```

---

## Pointers

- Store memory addresses of variables
- Allow passing references instead of values
- Useful for modifying variables across functions and passing heavy structs
- Zero value is nil

```go
var x int = 10
var p *int = &x  // p holds memory address of x
fmt.Println(*p)  // Dereferencing - prints 10

*p = 20  // Modify value through pointer
fmt.Println(x)  // Prints 20
```

---


class: center, middle

# Structs in Go

---

## Basics of Structs

- Custom data types with named fields
- Used to define objects with multiple properties

```go
type UrlResponse struct {
    Url          string
    StatusCode   uint
    ResponseBody string
    Error        string
}

response := UrlResponse{Url: "https://example.com", StatusCode: 200, ResponseBody: "Hello, World!", Error: ""}
fmt.Println(response.Url, response.StatusCode)
```

---

## Constructors in Go

- Go does **not** have traditional constructors like other languages.
- A common convention is to use a function named `NewObject` (e.g., `NewUrlResponse`) to initialize and return a new instance of a struct.
- This allows encapsulation and validation before object creation.

```go
type UrlResponse struct {
    Url          string
    StatusCode   uint
    ResponseBody string
    Error        string
}

// NewUrlResponse acts as a constructor function
func NewUrlResponse(url string, statusCode uint) *UrlResponse {
    return &UrlResponse{
      Url: url, 
      StatusCode: statusCode, 
    }
}

func main() {
    response := NewUrlResponse("https://example.com", 200)
    fmt.Println(response.Url)  // Works
    // Other operations with response
}
```

---

## Public and Private Fields

- Visibility of fields in Go is determined by their casing:
- **Public**: Fields starting with an uppercase letter are accessible outside the package.
- **Private**: Fields starting with a lowercase letter are only accessible within the same package.
- Private fields promote encapsulation and better struct design.

```go
type UrlResponse struct {
    Url          string // Public field
    statusCode   uint   // Private field
    responseBody string // Private field
    errorResponse        string // Private field
}

func main() {
    response := NewUrlResponse("https://example.com", 200, "Success", "")
    fmt.Println(response.Url)               // Works
    // fmt.Println(response.responseBody)   // Does not works (private field)
    // fmt.Println(response.statusCode)     // Does not work (private field)
    // fmt.Println(response.errorResponse)  // Does not work (private field)
}
```

---

## Struct Methods

- Fields can be accessed using dot notation
- Supports methods to define behavior

```go
type UrlResponse struct {
    Url          string
    StatusCode   uint
    ResponseBody string
    Error        string
}

func (r UrlResponse) Info() string {
    return fmt.Sprintf("URL: %s, Status: %d", r.Url, r.StatusCode)
}

// Syntactic Sugar for the following
func Info(r UrlResponse) string {
    return fmt.Sprintf("URL: %s, Status: %d", r.Url, r.StatusCode)
}

response := UrlResponse{Url: "https://api.example.com", StatusCode: 200}
fmt.Println(response.Info())
```

---

## Struct Methods: Pointer receivers

- Beware for pointer receiver!

```go
type UrlResponse struct {
    Url          string
    StatusCode   uint
    ResponseBody string
    Error        string
}

// This will update the instance
func (r *UrlResponse) UpdateStatus(statusCode uint) {
    r.StatusCode = statusCode
}

// This will update a copy the instance which is not returned
// Only usefull in cases where you want to guarantee no side effects happen
func (r UrlResponse) UpdateStatusCopy(statusCode uint) {
    r.StatusCode = statusCode
}

func main() {
    response := UrlResponse{StatusCode: 200}
    
    response.UpdateStatusCopy(404)
    fmt.Println(response.StatusCode) // Still 200
    
    response.UpdateStatus(404)
    fmt.Println(response.StatusCode) // Now 404
}
```

---

class: center, middle, inverse

## Usefull commands

---

## Running your program

- Beware for pointer receiver!

---

## Testing your program

- Beware for pointer receiver!

---

## Building your program

- Beware for pointer receiver!

---



class: center, middle, inverse

## So now I know the basics (kinda)

#### What do I do with this?

---

class: center, middle, inverse

## Mini Project 1

#### Simple URL checker


---

## Struct Methods: Pointer receivers

- Beware for pointer receiver!

---

class: center, middle, inverse

### Thank you for your attention :)
