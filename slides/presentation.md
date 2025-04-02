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

class: center, middle

# Structs in Go
class: center, middle


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
    response := NewUrlResponse("https://example.com", 200, "Success", "")
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
    ResponseBody string // Public field
    error        string // Private field
}

func NewUrlResponse(url string, statusCode uint, responseBody string, err string) *UrlResponse {
    return &UrlResponse{Url: url, statusCode: statusCode, ResponseBody: responseBody, error: err}
}

func main() {
    response := NewUrlResponse("https://example.com", 200, "Success", "")
    fmt.Println(response.Url)          // Works
    fmt.Println(response.ResponseBody) // Works
    // fmt.Println(response.statusCode) // Does not work (private field)
    // fmt.Println(response.error)      // Does not work (private field)
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

response := UrlResponse{Url: "https://api.example.com", StatusCode: 200, ResponseBody: "Success", Error: ""}
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
// Useless
func (r UrlResponse) UpdateStatusCopy(statusCode uint) {
    r.StatusCode = statusCode
}

func main() {
    response := UrlResponse{Url: "https://api.example.com", StatusCode: 200, ResponseBody: "Success", Error: ""}
    
    response.UpdateStatusCopy(404)
    fmt.Println(response.StatusCode) // Still 200
    
    response.UpdateStatus(404)
    fmt.Println(response.StatusCode) // Now 404
}
```

---


### More on null-safety

You can't assign `null` to a variable unless you explicitly declare it as nullable.
If you do so, you have to be careful in how you handle it.

```dart
class Person {
  Person(this.age, this.name, this.height, this.surname);

  int age;
  String name;
  String? surname;
  double height;

  String getName() => this.name;
  String? getSurname() => this.surname;
}

void main() {
  Person recruta = Person(18, "Recruta", 1.8, null);

  if (recruta.getSurname() != null) {
    print(recruta.getSurname()!.length); // ignores null-safety, prints nothing
  }

  print(recruta.getSurname()?.length); // prints null
}

```

---

### Asynchronous programming

Some things in life take time. How awkward would it be if UI interaction in _uni_ was suspended while the app was waiting for a response from _Sigarra_, the fastest information system in the world?

```dart
import 'package:http/http.dart' as http;

Future<http.Response> fetchLectures() {
    final uri = Uri.parse('sigarra.up.pt/feup/pt/mob_hor_geral.estudante');
    return http.get(uri);
}

// If a function is marked as async, it can use the await keyword
void main() async {
    // Wait for the execution of the fetchLectures function before continuing
    final response = await fetchLectures();
    final json = jsonDecode(response.body);
    print(json);
}
```

---

class: center, middle, inverse

## So now I know the basics (kinda)

#### How do I build a Flutter app?

---

### Flutter: Getting started

For developing Flutter apps, you need:

- The Flutter SDK (which includes the Dart SDK)
- Platform-specific SDKs and/or IDEs
  - Android: Android Studio + Android SDK
  - iOS: Xcode + iOS SDK
  - Linux: C/C++ toolchain
  - Web: Chrome
  - ...
- A lot of patience for the first build

> In Flutter's defense, the first build is the only one that takes a long time, since you can use hot-reload and hot-restart to quickly iterate on your code later on. Also, transpilation to native code is not an easy task and also takes time in other frameworks, such as React Native.

---

### Flutter essentials: Widgets

- In Flutter, everything is a widget expressed in code
- Widgets are arranged in a tree structure and are rendered on a canvas provided by the platform
- Widgets (and its children) are rebuilt everytime their state changes, e.g. when a button is pressed
- The Flutter framework provides a set of built-in widgets, but you can also create your own

<img src="./assets/widget_tree.gif" style="width: 80%;">

---

### Flutter essentials: Stateless Widgets

- As the name say, widgets without a state
- These are immutable (cannot be changed)
- Suitable for static UI elements

```dart
class PersonInfo extends StatelessWidget {
    PersonInfo(this.age, this.name, this.height);

    int age;
    String name;
    double height;

    @override
    Widget build(BuildContext context) {
        return Column(
            children: [
              Text(name),
              Text(age),
              Text(height)
            ]
        );
    }
}
```

---

### Flutter essentials: Stateful Widgets

- As the name say, widgets with a state
- These are mutable
- Suitable for UI elements that change overtime

```dart
class PostLike extends StatefulWidget {
    @override
    PostLikeState createState() => PostLikeState();
}

class PostLikeState extends State<PostLike> {
    int _likes = 0;

    @override
    Widget build(BuildContext context) {
        return Row(
            children: [
                Text("Likes: $_likes"),
                LikeButton(
                    onPressed: () => setState(() => _likes = likes + 1),
                ),
            ],
        );
    }
}
```

---

### Flutter essentials: Layout

Understanding how things are position in the screen and how they will be displayed is a key step in order to master Flutter.

- `Row` and `Column` are two built-in widgets that arrange other widgets horizontally or vertically.
- `Container` is another built-in widget that acts like a box where you can fit other widgets.

<img src="./assets/schedule_card_1.png" style="width: 45%;">

---

### Flutter essentials: Layout

Understanding how things are position in the screen and how they will be displayed is a key step in order to master Flutter.

- `Row` and `Column` are two built-in widgets that arrange other widgets horizontally or vertically.
- `Container` is another built-in widget that acts like a box where you can fit other widgets.

<img src="./assets/schedule_card_2.png" style="width: 45%;">

---

### Flutter essentials: Layout

Understanding how things are position in the screen and how they will be displayed is a key step in order to master Flutter.

- `Row` and `Column` are two built-in widgets that arrange other widgets horizontally or vertically.
- `Container` is another built-in widget that acts like a box where you can fit other widgets.

<img src="./assets/schedule_card_3.png" style="width: 45%;">

---

### Flutter essentials: Layout

Understanding how things are position in the screen and how they will be displayed is a key step in order to master Flutter.

- `Row` and `Column` are two built-in widgets that arrange other widgets horizontally or vertically.
- `Container` is another built-in widget that acts like a box where you can fit other widgets.

<img src="./assets/schedule_card_4.png" style="width: 45%;">

---

class: center, middle, inverse

## Tired of my bla bla?

#### Let's build a widget

## <img src="./assets/example.png" style="width: 40%;">

---

class: inverse

### Part 1

- Clone the workshop's repo

```sh
git clone git@github.com:DGoiana/flutter-ws.git

```

- Checkout to part1 branch

```sh
git checkout part1

```

- Implement Meal class (meal.dart)
- Implement Meal card (meal_card.dart)

---

### Meal Class

```dart
class Meal{
  Meal({
    required this.name,
    required this.category,
    required this.region,
    this.imageURL
  });

  // name
  // category
  // region
  // optional URL
}

```

---

### Meal Card

```dart

import 'package:flutter/material.dart';

class MealCard extends StatelessWidget {
  const MealCard({super.key, required this.meal});

  final Meal meal;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(10),
      padding: const EdgeInsets.all(10),
      decoration: BoxDecoration(
        border: Border.all(color: Colors.black),
      ),
      child: _;
  }
}

```

---

### Flutter advanced: State management

In applications, it is common to have global state that is useful to sibling or distant widgets in the widget tree. For example, a user's authentication state is useful to many widgets in the app.

A solution known so far would be to pass the user authentication state as a parameter to every widget that needs it, which is error-prone.

To avoid that, fortunately, packages such as `provider` provide solutions to this problem.

<img src="./assets/provider.jpg" style="width: 70%;">

---

### Flutter advanced: Offline storage

Flutter provides a set of packages that allow you to store data locally on the device, such as `shared_preferences` (key-value storage) and `sqflite` (SQLite database).

```dart
final sharedPreferences = await SharedPreferences.getInstance();
final savedUsername = await sharedPreferences.getString("username");
```

```dart
final db = await openDatabase('my_db.db');
final savedUsername = await
    db
    .query("users", where: "id = ?", whereArgs: [1])
    .then((rows) => rows.first["username"]);
```

The choice of storage depends on the complexity of the data you want to store. For simple data, `shared_preferences` is enough. For more complex data, you should use `sqflite`.

---

### Flutter advanced: Architecture

Real-world applications are large and complex and thus require a well-defined architecture to be maintainable. Flutter is no exception.

A common architecture for reactive applications is **MVC** (Model-View-Controller):

- **Model**: the data layer, i.e. the classes that represent entities in the application (e.g. a `User` class). May also contain state management logic.
- **View**: the UI layer, i.e. the widgets that are rendered on the screen.
- **Controller**: the business logic layer, i.e. the classes that handle the application's logic.

<img src="./assets/mvc.png" style="width: 60%;">

---

### Flutter advanced: Best practices

- Use `const` whenever possible
- Keep the state local to the widget that needs it if possible, i.e. do not overuse global state management
- Extract widgets to separate files to keep your code clean and avoid heavy rebuilds
- Use a static code analyzer ("linter") to keep your code formatted, clean and consistent
- Do not put complex business logic in the UI layer
- Do not reinvent the wheel: use packages from the community whenever possible
- Test your code, especially the business logic layer

---

class: center, middle, inverse

# Questions?

---

class: center, middle, inverse

### So you are now a Flutter expert

#### Let's get our hands dirty. Can you build this (or better)?

<div style="display: flex; justify-content: center; align-items: center; gap: 2em">
<img src="./assets/practical/part2.gif" style="height: 50vh;">
</div>

---

class: inverse

### Part 2

- Checkout to part2 branch

```sh
git checkout part2
```

- Implement RecipesProvider (will memorize all reciped generated)
- Implement RandomMealPage (page to randomize meals)
- Implement FetchRandomMeal (function to call API)

---

### Fetch Random Meal

```dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'meal.dart';

Future<Meal> fetchRandomMeal() async {
  final response = await http
      .get(Uri.parse('https://www.themealdb.com/api/json/v1/1/random.php'));

  if (response.statusCode == 200) {
    final data = jsonDecode(response.body);
    final mealData = data['meals'][0];

    // retornar um objeto meal novo

  } else {
    throw Exception('Failed to load meal');
  }
}

```

---

### Recipes Provider

```dart
import 'package:flutter/material.dart';

class RecipesProvider extends ChangeNotifier {
  final List<String> _recipes = [];

  // cria uma função que retorna o nome das receitas

  // cria uma função que adiciona as receitas à lista
}

```

---

### Random Meal Page

```dart

class RandomMealPage extends StatefulWidget {
  const RandomMealPage({super.key});

  @override
  State<RandomMealPage> createState() => RandomMealPageState();
}

class RandomMealPageState extends State<RandomMealPage> {
  late Meal randomMeal;

  @override
  void initState() {
    super.initState();
    setRandomMeal();
  }

  Future<void> setRandomMeal() async {
    Meal meal = await fetchRandomMeal();

    // modificar estado
  }

```

---

### Random Meal Page II

```dart

@override
  Widget build(BuildContext context) {
    final recipeProvider = Provider.of<RecipesProvider>(context);

    return Column(
      children: [
        // adicionar botão
        MealCard(meal: randomMeal),
        Expanded(
          child: // adicionar lista
        ),
      ],
    );
  }
}

```

---

class: center, middle, inverse

### Thank you for your attention :)
