# Go Language Cheat Sheet

## Table of Contents
1. [String Manipulation](#string-manipulation)
2. [Heap (Min and Max)](#heap-min-and-max)
3. [Math Operations](#math-operations)
4. [Sorting](#sorting)
5. [Binary Search](#binary-search)
6. [BFS Traversal](#bfs-traversal)
7. [Dynamic Programming](#dynamic-programming)
8. [For Loops](#for-loops)
9. [Operations Missing from Go's Standard Library](#operations-missing-from-gos-standard-library)

## String Manipulation

```go
import (
    "strings"
    "fmt"
    "strconv"
)

// Concatenation
s1 := "Hello"
s2 := "World"
result := s1 + " " + s2  // "Hello World"

// Using fmt.Sprintf
result = fmt.Sprintf("%s %s", s1, s2)  // "Hello World"

// Substrings
s := "Hello, World!"
substr := s[7:12]  // "World"

// String functions
upper := strings.ToUpper(s)  // "HELLO, WORLD!"
lower := strings.ToLower(s)  // "hello, world!"
trimmed := strings.TrimSpace(" Hello ")  // "Hello"
replaced := strings.Replace(s, "World", "Go", 1)  // "Hello, Go!"
split := strings.Split(s, ", ")  // ["Hello", "World!"]

// Joining strings
parts := []string{"Hello", "World", "of", "Go"}
joined := strings.Join(parts, " ")  // "Hello World of Go"

// Conversion between types and strings
intVar := 42
floatVar := 3.14
boolVar := true

// Convert to string
intString := strconv.Itoa(intVar)  // "42"
floatString := fmt.Sprintf("%.2f", floatVar)  // "3.14"
boolString := strconv.FormatBool(boolVar)  // "true"

// Convert from string
parsedInt, _ := strconv.Atoi("42")  // 42
parsedFloat, _ := strconv.ParseFloat("3.14", 64)  // 3.14
parsedBool, _ := strconv.ParseBool("true")  // true
```

## Heap (Min and Max)

Go's `container/heap` package provides an interface for heap operations. Here's how to implement min and max heaps for strings and integers:

```go
import (
    "container/heap"
    "fmt"
)

// Min Heap for Integers
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Max Heap for Integers
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Min Heap for Strings
type StringMinHeap []string

func (h StringMinHeap) Len() int           { return len(h) }
func (h StringMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StringMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *StringMinHeap) Push(x interface{}) { *h = append(*h, x.(string)) }
func (h *StringMinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Usage example
func main() {
    // Integer Min Heap
    intMinHeap := &IntMinHeap{3, 1, 4, 1, 5, 9}
    heap.Init(intMinHeap)
    heap.Push(intMinHeap, 2)
    fmt.Printf("Minimum: %d\n", (*intMinHeap)[0])
    
    // String Min Heap
    stringMinHeap := &StringMinHeap{"banana", "apple", "cherry"}
    heap.Init(stringMinHeap)
    heap.Push(stringMinHeap, "date")
    fmt.Printf("Minimum: %s\n", (*stringMinHeap)[0])
}
```

## Math Operations

Go's `math` package provides various mathematical functions:

```go
import (
    "math"
    "fmt"
)

// Power of a number
pow := math.Pow(2, 3)  // 8 (2^3)

// Square root
sqrt := math.Sqrt(16)  // 4

// Absolute value
abs := math.Abs(-10)  // 10

// Rounding
round := math.Round(3.7)  // 4
floor := math.Floor(3.7)  // 3
ceil := math.Ceil(3.2)    // 4

// Trigonometric functions
sin := math.Sin(math.Pi / 2)  // 1
cos := math.Cos(0)            // 1
tan := math.Tan(math.Pi / 4)  // 1

// Logarithms
log := math.Log(math.E)    // 1 (natural logarithm)
log10 := math.Log10(100)   // 2

// Constants
pi := math.Pi    // 3.141592653589793
e := math.E      // 2.718281828459045

// Max and Min
max := math.Max(10, 5)  // 10
min := math.Min(10, 5)  // 5
```

## Sorting

Go's `sort` package provides functions for sorting slices and user-defined collections.

```go
import (
    "sort"
    "fmt"
)

// Sorting integers
ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
sort.Ints(ints)
fmt.Println(ints)  // [1 1 2 3 4 5 6 9]

// Sorting strings
strings := []string{"banana", "apple", "cherry", "date"}
sort.Strings(strings)
fmt.Println(strings)  // [apple banana cherry date]

// Reverse sorting
sort.Sort(sort.Reverse(sort.IntSlice(ints)))
fmt.Println(ints)  // [9 6 5 4 3 2 1 1]

// Custom sorting
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 20},
}

// Sort by age
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
fmt.Println(people)  // [{Charlie 20} {Alice 25} {Bob 30}]

// Sort by name
sort.Slice(people, func(i, j int) bool {
    return people[i].Name < people[j].Name
})
fmt.Println(people)  // [{Alice 25} {Bob 30} {Charlie 20}]

// Searching (binary search)
ints = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
index := sort.SearchInts(ints, 5)
fmt.Println(index)  // 4 (index of 5 in the sorted slice)

// Note: The slice must be sorted before using Search functions
```

Note on the search algorithm: Go's `sort` package uses binary search for its `Search` functions (like `SearchInts`, `SearchStrings`, etc.). Binary search has a time complexity of O(log n), making it efficient for searching in sorted slices.

The sorting algorithm used internally by Go's `sort` package is a hybrid algorithm:
1. It uses quicksort for larger slices (more than 12 elements).
2. For smaller slices or partitions, it switches to insertion sort.
3. It also uses heapsort as a fallback to guarantee O(n log n) worst-case complexity.

This hybrid approach combines the average-case efficiency of quicksort with the good performance of insertion sort on small slices and the worst-case guarantee of heapsort.

## Binary Search

Binary search is an efficient algorithm for searching a sorted array by repeatedly dividing the search interval in half.

```go
import (
    "fmt"
)

// Iterative Binary Search
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1 // Not found
}

// Recursive Binary Search
func binarySearchRecursive(arr []int, target, left, right int) int {
    if left > right {
        return -1 // Not found
    }

    mid := left + (right-left)/2
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        return binarySearchRecursive(arr, target, mid+1, right)
    } else {
        return binarySearchRecursive(arr, target, left, mid-1)
    }
}

// Usage
func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    target := 7

    // Iterative
    result := binarySearch(arr, target)
    fmt.Printf("Iterative: Element found at index: %d\n", result)

    // Recursive
    result = binarySearchRecursive(arr, target, 0, len(arr)-1)
    fmt.Printf("Recursive: Element found at index: %d\n", result)
}
```

## BFS Traversal

Breadth-First Search (BFS) is an algorithm for traversing or searching tree or graph data structures.

```go
import (
    "fmt"
)

// Graph represented as an adjacency list
type Graph struct {
    vertices int
    adjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
    return &Graph{
        vertices: vertices,
        adjList:  make(map[int][]int),
    }
}

func (g *Graph) AddEdge(v, w int) {
    g.adjList[v] = append(g.adjList[v], w)
    g.adjList[w] = append(g.adjList[w], v)
}

func (g *Graph) BFS(start int) {
    visited := make(map[int]bool)
    queue := []int{start}

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]

        if !visited[vertex] {
            visited[vertex] = true
            fmt.Printf("%d ", vertex)

            for _, neighbor := range g.adjList[vertex] {
                if !visited[neighbor] {
                    queue = append(queue, neighbor)
                }
            }
        }
    }
}

// Usage
func main() {
    g := NewGraph(4)
    g.AddEdge(0, 1)
    g.AddEdge(0, 2)
    g.AddEdge(1, 2)
    g.AddEdge(2, 3)

    fmt.Println("BFS traversal starting from vertex 0:")
    g.BFS(0)
}
```

## Dynamic Programming

Dynamic Programming (DP) is a method for solving complex problems by breaking them down into simpler subproblems. Here's an example of solving the Fibonacci sequence using DP.

```go
import (
    "fmt"
)

// Recursive Fibonacci (inefficient)
func fibRecursive(n int) int {
    if n <= 1 {
        return n
    }
    return fibRecursive(n-1) + fibRecursive(n-2)
}

// Dynamic Programming: Top-down approach (Memoization)
func fibMemo(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    if val, exists := memo[n]; exists {
        return val
    }
    memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
    return memo[n]
}

// Dynamic Programming: Bottom-up approach (Tabulation)
func fibDP(n int) int {
    if n <= 1 {
        return n
    }
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

// Usage
func main() {
    n := 10

    fmt.Printf("Recursive Fibonacci(%d) = %d\n", n, fibRecursive(n))

    memo := make(map[int]int)
    fmt.Printf("Memoized Fibonacci(%d) = %d\n", n, fibMemo(n, memo))

    fmt.Printf("DP Fibonacci(%d) = %d\n", n, fibDP(n))
}
```

This example demonstrates three approaches to solving the Fibonacci sequence:
1. Recursive (inefficient for large n)
2. Top-down DP with memoization
3. Bottom-up DP with tabulation

Dynamic Programming is crucial in solving many optimization problems efficiently, especially in competitive programming scenarios.

## For Loops

Go's `for` loop is versatile and can be used in various ways. Here are common patterns and use cases:

```go
import (
    "fmt"
)

func main() {
    // Basic for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // While-like loop
    j := 0
    for j < 5 {
        fmt.Println(j)
        j++
    }

    // Infinite loop
    k := 0
    for {
        fmt.Println(k)
        k++
        if k >= 5 {
            break
        }
    }

    // Iterating over a slice or array
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Iterating over a map
    scores := map[string]int{"Alice": 95, "Bob": 80, "Charlie": 90}
    for key, value := range scores {
        fmt.Printf("%s scored %d\n", key, value)
    }

    // Iterating in reverse
    for i := 10; i > 0; i-- {
        fmt.Println(i)
    }

    // Loop with multiple variables
    for i, j := 0, 10; i < j; i, j = i+1, j-1 {
        fmt.Printf("i = %d, j = %d\n", i, j)
    }

    // Loop with continue
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println(i)
    }
}
```

Key points about for loops in Go:

1. Go has only one looping construct: the `for` loop.
2. The basic `for` loop has three components: init statement, condition expression, and post statement.
3. The init and post statements are optional.
4. You can use `break` to exit a loop early and `continue` to skip to the next iteration.
5. The `range` form of the for loop iterates over slices, arrays, maps, strings, and channels.
6. You can create infinite loops using `for {}` or `for true {}`.
7. Go doesn't have a do-while loop, but you can simulate it using a for loop with a break statement.
8. Labels can be used with `break` and `continue` to control outer loops from within nested loops.

Remember that in Go, there are no parentheses surrounding the three components of the for statement, but the braces `{ }` are required.

## Operations Missing from Go's Standard Library

Go's standard library is designed to be minimal and focused. As a result, some operations that are commonly found in other languages' standard libraries are not included in Go. Here are some examples and how to work around them:

1. Set Data Structure
   - Go doesn't have a built-in Set type.
   - Workaround: Use a map with bool values.

   ```go
   set := make(map[string]bool)
   set["item"] = true
   if set["item"] {
       fmt.Println("Item exists")
   }
   ```

2. Optional/Nullable Types
   - Go doesn't have built-in optional types like Some languages (e.g., Rust's Option or Scala's Option).
   - Workaround: Use pointers or a custom struct.

   ```go
   type Optional struct {
       value int
       isSet bool
   }
   ```

3. Ternary Operator
   - Go doesn't have a ternary operator (?:).
   - Workaround: Use an if-else statement.

   ```go
   result := ""
   if condition {
       result = "true"
   } else {
       result = "false"
   }
   ```

4. Generic Max/Min Functions
   - Go doesn't have generic max/min functions in its standard library.
   - Workaround: Write your own function or use math.Max/Min for float64.

   ```go
   func Max(x, y int) int {
       if x > y {
           return x
       }
       return y
   }
   ```

5. String Joining with Separator
   - While Go has strings.Join, it doesn't have a direct equivalent to Python's " ".join(list).
   - Workaround: Use strings.Join with a slice.

   ```go
   words := []string{"Hello", "World"}
   sentence := strings.Join(words, " ")
   ```

6. String Multiplication
   - Go doesn't have a string multiplication operator like Python's "*".
   - Workaround: Use strings.Repeat.

   ```go
   repeatedString := strings.Repeat("Go", 3)  // "GoGoGo"
   ```

7. Built-in List Comprehensions
   - Go doesn't have list comprehensions like Python.
   - Workaround: Use a for loop or create a helper function.

   ```go
   squares := make([]int, 10)
   for i := range squares {
       squares[i] = i * i
   }
   ```

8. Built-in Functional Programming Operations
   - Go doesn't have built-in map, filter, or reduce functions.
   - Workaround: Write your own functions or use a third-party library.

   ```go
   // Simple map function
   func Map(vs []int, f func(int) int) []int {
       vsm := make([]int, len(vs))
       for i, v := range vs {
           vsm[i] = f(v)
       }
       return vsm
   }
   ```

9. Complex Number Type
   - While Go does have a complex number type, it lacks some operations found in languages like Python.
   - Workaround: Use the math/cmplx package for additional operations.

   ```go
   c := complex(3, 4)
   abs := cmplx.Abs(c)
   ```

10. Date Parsing and Formatting
    - Go's time parsing and formatting is different from many other languages.
    - Workaround: Use the specific layout string required by Go.

    ```go
    const layout = "2006-01-02"
    date, _ := time.Parse(layout, "2023-05-15")
    ```