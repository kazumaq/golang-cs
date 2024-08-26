# Go Language Cheat Sheet

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
