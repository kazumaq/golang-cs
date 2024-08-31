
# Python Cheatsheet

## Table of Contents
1. [String Manipulation](#string-manipulation)
2. [Heap (Min and Max)](#heap-min-and-max)
3. [Math Operations](#math-operations)
4. [Sorting](#sorting)
5. [Binary Search](#binary-search)
6. [BFS Traversal](#bfs-traversal)
7. [Dynamic Programming](#dynamic-programming)
8. [For Loops](#for-loops)
9. [Python-specific Operations](#python-specific-operations)
10. [Printing and Formatting Data](#printing-and-formatting-data)
11. [Classes](#classes)
12. [Functions](#functions)
13. [Abstract Base Classes and Protocols](#abstract-base-classes-and-protocols)
14. [Python References and Resources](#python-references-and-resources)

## String Manipulation

```python
# Concatenation
s1 = "Hello"
s2 = "World"
result = s1 + " " + s2  # "Hello World"

# Using f-strings (Python 3.6+)
result = f"{s1} {s2}"  # "Hello World"

# Substrings
s = "Hello, World!"
substr = s[7:12]  # "World"

# String methods
upper = s.upper()  # "HELLO, WORLD!"
lower = s.lower()  # "hello, world!"
trimmed = " Hello ".strip()  # "Hello"
replaced = s.replace("World", "Python", 1)  # "Hello, Python!"
split = s.split(", ")  # ["Hello", "World!"]

# Joining strings
parts = ["Hello", "World", "of", "Python"]
joined = " ".join(parts)  # "Hello World of Python"

# Conversion between types and strings
int_var = 42
float_var = 3.14
bool_var = True

# Convert to string
int_string = str(int_var)  # "42"
float_string = f"{float_var:.2f}"  # "3.14"
bool_string = str(bool_var)  # "True"

# Convert from string
parsed_int = int("42")  # 42
parsed_float = float("3.14")  # 3.14
parsed_bool = bool("True")  # True

# String formatting
name = "Alice"
age = 30
formatted = "Name: {}, Age: {}".format(name, age)  # "Name: Alice, Age: 30"

# F-strings (Python 3.6+)
formatted = f"Name: {name}, Age: {age}"  # "Name: Alice, Age: 30"

# String multiplication
repeated = "Python" * 3  # "PythonPythonPython"

# Checking substrings
contains = "World" in s  # True

# String methods for checking content
starts_with = s.startswith("Hello")  # True
ends_with = s.endswith("!")  # True
is_alpha = "Hello".isalpha()  # True
is_digit = "123".isdigit()  # True

# Finding substrings
index = s.find("World")  # 7
rindex = s.rfind("o")  # 8 (last occurrence)

# Counting occurrences
count = s.count("l")  # 3

# Alignment and padding
left_aligned = "Python".ljust(10)  # "Python    "
right_aligned = "Python".rjust(10)  # "    Python"
centered = "Python".center(10)  # "  Python  "

# Removing characters
stripped = "   Python   ".strip()  # "Python"
left_stripped = "   Python   ".lstrip()  # "Python   "
right_stripped = "   Python   ".rstrip()  # "   Python"

# Case conversion
title_case = "hello world".title()  # "Hello World"
swapped_case = "Hello World".swapcase()  # "hELLO wORLD"

# Encoding and decoding
encoded = "Pythön".encode("utf-8")  # b'Pyth\xc3\xb6n'
decoded = encoded.decode("utf-8")  # "Pythön"
```

## Heap (Min and Max)

Python's `heapq` module provides an implementation of the heap queue algorithm, which is a min heap. To create a max heap, you can invert the values.

```python
import heapq

# Min Heap
min_heap = []
heapq.heappush(min_heap, 4)
heapq.heappush(min_heap, 1)
heapq.heappush(min_heap, 7)
heapq.heappush(min_heap, 3)

print(min_heap)  # [1, 3, 7, 4]

smallest = heapq.heappop(min_heap)  # 1
print(smallest)  # 1
print(min_heap)  # [3, 4, 7]

# Peek at the smallest item without removing it
print(min_heap[0])  # 3

# Convert a list to a heap in-place
numbers = [4, 1, 7, 3]
heapq.heapify(numbers)
print(numbers)  # [1, 3, 7, 4]

# Max Heap (by inverting values)
max_heap = []
heapq.heappush(max_heap, -4)
heapq.heappush(max_heap, -1)
heapq.heappush(max_heap, -7)
heapq.heappush(max_heap, -3)

print([-x for x in max_heap])  # [7, 4, 1, 3]

largest = -heapq.heappop(max_heap)  # 7
print(largest)  # 7
print([-x for x in max_heap])  # [4, 3, 1]

# Peek at the largest item without removing it
print(-max_heap[0])  # 4

# Convert a list to a max heap in-place
numbers = [4, 1, 7, 3]
numbers = [-x for x in numbers]
heapq.heapify(numbers)
print([-x for x in numbers])  # [7, 4, 1, 3]

# Heap operations on custom objects
class Person:
    def __init__(self, name, priority):
        self.name = name
        self.priority = priority
    
    def __lt__(self, other):
        return self.priority < other.priority

people_heap = []
heapq.heappush(people_heap, Person("Alice", 3))
heapq.heappush(people_heap, Person("Bob", 1))
heapq.heappush(people_heap, Person("Charlie", 2))

highest_priority = heapq.heappop(people_heap)
print(highest_priority.name)  # "Bob"

# Get the n largest or smallest elements
numbers = [4, 1, 7, 3, 8, 2, 9, 5, 6]
largest_3 = heapq.nlargest(3, numbers)  # [9, 8, 7]
smallest_3 = heapq.nsmallest(3, numbers)  # [1, 2, 3]
```

Key points about heaps in Python:
1. The `heapq` module implements a min heap.
2. To create a max heap, invert the values when pushing and popping.
3. Heapify operation converts a list into a heap in-place.
4. Custom objects can be used in heaps by implementing the `__lt__` method.
5. `heapq.nlargest()` and `heapq.nsmallest()` are efficient for getting the n largest or smallest elements.

Remember that while these operations are common in Python, always refer to the official Python documentation for the most up-to-date and comprehensive information on string manipulation and heap operations.

## Math Operations

Python provides a rich set of mathematical operations through its built-in functions and the `math` module.

```python
import math

# Basic arithmetic
add = 5 + 3  # 8
subtract = 5 - 3  # 2
multiply = 5 * 3  # 15
divide = 5 / 3  # 1.6666666666666667
floor_divide = 5 // 3  # 1
modulo = 5 % 3  # 2
power = 5 ** 3  # 125

# Math functions
sqrt = math.sqrt(16)  # 4.0
pow = math.pow(2, 3)  # 8.0
abs_value = abs(-10)  # 10

# Rounding
round_value = round(3.7)  # 4
floor = math.floor(3.7)  # 3
ceil = math.ceil(3.2)  # 4

# Trigonometric functions
sin = math.sin(math.pi / 2)  # 1.0
cos = math.cos(0)  # 1.0
tan = math.tan(math.pi / 4)  # 0.9999999999999999

# Logarithms
log = math.log(math.e)  # 1.0 (natural logarithm)
log10 = math.log10(100)  # 2.0

# Constants
pi = math.pi  # 3.141592653589793
e = math.e  # 2.718281828459045

# Complex numbers
c = 3 + 4j
magnitude = abs(c)  # 5.0
real_part = c.real  # 3.0
imag_part = c.imag  # 4.0

# Infinity and NaN
inf = math.inf
nan = math.nan

# Check for infinity and NaN
is_inf = math.isinf(inf)  # True
is_nan = math.isnan(nan)  # True

# Degrees and radians conversion
degrees = math.degrees(math.pi)  # 180.0
radians = math.radians(180)  # 3.141592653589793

# Hyperbolic functions
sinh = math.sinh(1)  # 1.1752011936438014
cosh = math.cosh(1)  # 1.5430806348152437
tanh = math.tanh(1)  # 0.7615941559557649

# Factorial
factorial = math.factorial(5)  # 120

# Greatest Common Divisor (GCD)
gcd = math.gcd(48, 18)  # 6

# Least Common Multiple (LCM)
lcm = math.lcm(48, 18)  # 144 (Python 3.9+)

# Exponentiation and logarithms
exp = math.exp(1)  # 2.718281828459045
log2 = math.log2(8)  # 3.0

# Statistical functions
numbers = [1, 2, 3, 4, 5]
mean = sum(numbers) / len(numbers)  # 3.0
geometric_mean = math.prod(numbers) ** (1 / len(numbers))  # 2.605171084697352
```

## Sorting

Python provides built-in sorting capabilities for lists and other iterable objects.

```python
# Sorting a list in-place
numbers = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]
numbers.sort()
print(numbers)  # [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]

# Sorting in descending order
numbers.sort(reverse=True)
print(numbers)  # [9, 6, 5, 5, 5, 4, 3, 3, 2, 1, 1]

# Creating a new sorted list
original = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]
sorted_numbers = sorted(original)
print(sorted_numbers)  # [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]
print(original)  # [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5] (unchanged)

# Sorting strings
fruits = ["banana", "apple", "cherry", "date"]
sorted_fruits = sorted(fruits)
print(sorted_fruits)  # ['apple', 'banana', 'cherry', 'date']

# Case-insensitive sorting
mixed_case = ["Apple", "banana", "Cherry", "date"]
sorted_case_insensitive = sorted(mixed_case, key=str.lower)
print(sorted_case_insensitive)  # ['Apple', 'banana', 'Cherry', 'date']

# Sorting with a custom key function
def sort_by_last_letter(s):
    return s[-1]

sorted_by_last = sorted(fruits, key=sort_by_last_letter)
print(sorted_by_last)  # ['banana', 'apple', 'date', 'cherry']

# Sorting dictionaries
scores = {"Alice": 85, "Bob": 92, "Charlie": 78, "David": 95}
sorted_by_score = sorted(scores.items(), key=lambda x: x[1], reverse=True)
print(sorted_by_score)  # [('David', 95), ('Bob', 92), ('Alice', 85), ('Charlie', 78)]

# Sorting objects
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def __repr__(self):
        return f"Person('{self.name}', {self.age})"

people = [
    Person("Alice", 25),
    Person("Bob", 30),
    Person("Charlie", 20),
    Person("David", 35)
]

# Sort by age
sorted_by_age = sorted(people, key=lambda p: p.age)
print(sorted_by_age)
# [Person('Charlie', 20), Person('Alice', 25), Person('Bob', 30), Person('David', 35)]

# Sort by name
sorted_by_name = sorted(people, key=lambda p: p.name)
print(sorted_by_name)
# [Person('Alice', 25), Person('Bob', 30), Person('Charlie', 20), Person('David', 35)]

# Using the operator module for multiple sorting criteria
from operator import attrgetter

sorted_multiple = sorted(people, key=attrgetter('age', 'name'))
print(sorted_multiple)
# [Person('Charlie', 20), Person('Alice', 25), Person('Bob', 30), Person('David', 35)]

# Partial sorting with heapq
import heapq

numbers = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]
three_smallest = heapq.nsmallest(3, numbers)
print(three_smallest)  # [1, 1, 2]

three_largest = heapq.nlargest(3, numbers)
print(three_largest)  # [9, 6, 5]
```

Key points about sorting in Python:
1. The `sort()` method sorts lists in-place, while `sorted()` returns a new sorted list.
2. Use the `reverse=True` parameter for descending order.
3. The `key` parameter allows custom sorting criteria.
4. For dictionaries, sort by keys or values using `lambda` functions or `operator.itemgetter()`.
5. When sorting objects, use `lambda` functions or `operator.attrgetter()` to specify attributes for sorting.
6. The `heapq` module provides efficient partial sorting with `nsmallest()` and `nlargest()`.

Remember to refer to the official Python documentation for the most up-to-date and comprehensive information on math operations and sorting techniques.

## Binary Search

Binary search is an efficient algorithm for searching a sorted array by repeatedly dividing the search interval in half. Python's `bisect` module provides binary search functionality, but we'll also implement it manually for educational purposes.

```python
import bisect

def binary_search(arr, target):
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = (left + right) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return -1  # Target not found

# Manual implementation
arr = [1, 3, 5, 7, 9, 11, 13, 15]
target = 7
result = binary_search(arr, target)
print(f"Target {target} found at index: {result}")  # Target 7 found at index: 3

# Using bisect module
index = bisect.bisect_left(arr, target)
if index != len(arr) and arr[index] == target:
    print(f"Target {target} found at index: {index}")
else:
    print(f"Target {target} not found")

# Finding insertion point for a new element
new_element = 8
insert_point = bisect.bisect_left(arr, new_element)
print(f"Insert point for {new_element}: {insert_point}")  # Insert point for 8: 4

# Inserting a new element while maintaining sort order
bisect.insort_left(arr, new_element)
print(f"Array after insertion: {arr}")  # Array after insertion: [1, 3, 5, 7, 8, 9, 11, 13, 15]

# Recursive binary search
def binary_search_recursive(arr, target, left, right):
    if left > right:
        return -1
    
    mid = (left + right) // 2
    if arr[mid] == target:
        return mid
    elif arr[mid] < target:
        return binary_search_recursive(arr, target, mid + 1, right)
    else:
        return binary_search_recursive(arr, target, left, mid - 1)

result = binary_search_recursive(arr, target, 0, len(arr) - 1)
print(f"Target {target} found at index: {result}")  # Target 7 found at index: 3

# Binary search with custom comparison function
from dataclasses import dataclass

@dataclass
class Person:
    name: str
    age: int

def compare_person_age(person, age):
    return person.age - age

people = [
    Person("Alice", 25),
    Person("Bob", 30),
    Person("Charlie", 35),
    Person("David", 40)
]

target_age = 35
index = bisect.bisect_left(people, target_age, key=lambda p: p.age)
if index != len(people) and people[index].age == target_age:
    print(f"Person with age {target_age} found: {people[index].name}")
else:
    print(f"No person with age {target_age} found")
```

## BFS Traversal

Breadth-First Search (BFS) is an algorithm for traversing or searching tree or graph data structures. It explores all the vertices at the present depth before moving on to the vertices at the next depth level.

```python
from collections import deque

# Graph represented as an adjacency list
class Graph:
    def __init__(self):
        self.graph = {}
    
    def add_edge(self, u, v):
        if u not in self.graph:
            self.graph[u] = []
        self.graph[u].append(v)
    
    def bfs(self, start):
        visited = set()
        queue = deque([start])
        visited.add(start)
        
        while queue:
            vertex = queue.popleft()
            print(vertex, end=" ")
            
            for neighbor in self.graph.get(vertex, []):
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append(neighbor)
        print()

# Example usage for graph
g = Graph()
g.add_edge(0, 1)
g.add_edge(0, 2)
g.add_edge(1, 2)
g.add_edge(2, 0)
g.add_edge(2, 3)
g.add_edge(3, 3)

print("BFS traversal starting from vertex 2:")
g.bfs(2)  # Output: 2 0 3 1

# BFS for binary tree
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def bfs_tree(root):
    if not root:
        return
    
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        for _ in range(level_size):
            node = queue.popleft()
            print(node.val, end=" ")
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        print()  # New line for each level

# Example usage for binary tree
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.left.left = TreeNode(4)
root.left.right = TreeNode(5)

print("BFS traversal of binary tree:")
bfs_tree(root)
# Output:
# 1
# 2 3
# 4 5

# BFS for maze solving
def solve_maze(maze, start, end):
    rows, cols = len(maze), len(maze[0])
    queue = deque([(start, [])])
    visited = set([start])
    
    while queue:
        (x, y), path = queue.popleft()
        
        if (x, y) == end:
            return path + [(x, y)]
        
        for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:  # Right, Down, Left, Up
            nx, ny = x + dx, y + dy
            if 0 <= nx < rows and 0 <= ny < cols and maze[nx][ny] == 0 and (nx, ny) not in visited:
                queue.append(((nx, ny), path + [(x, y)]))
                visited.add((nx, ny))
    
    return None  # No path found

# Example maze (0 is path, 1 is wall)
maze = [
    [0, 0, 0, 0, 0],
    [1, 1, 0, 1, 0],
    [0, 0, 0, 0, 0],
    [0, 1, 1, 1, 0],
    [0, 0, 0, 1, 0]
]

start = (0, 0)
end = (4, 4)

path = solve_maze(maze, start, end)
if path:
    print("Path found:", path)
else:
    print("No path found")
```

Key points about Binary Search and BFS in Python:

1. Binary Search:
   - Works on sorted arrays
   - Has O(log n) time complexity
   - Can be implemented iteratively or recursively
   - Python's `bisect` module provides built-in binary search functionality

2. BFS Traversal:
   - Uses a queue data structure (often implemented with `collections.deque` for efficiency)
   - Explores all neighbors at the present depth before moving to the next level
   - Useful for finding the shortest path in unweighted graphs
   - Can be adapted for various problems like maze solving or level-order traversal of trees

Remember to refer to the official Python documentation and algorithm resources for more detailed information on these topics.

## Dynamic Programming

Dynamic Programming (DP) is a method for solving complex problems by breaking them down into simpler subproblems. It is applicable to problems exhibiting the properties of overlapping subproblems and optimal substructure.

```python
# Fibonacci sequence using DP
def fibonacci_dp(n):
    if n <= 1:
        return n
    
    dp = [0] * (n + 1)
    dp[1] = 1
    
    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]
    
    return dp[n]

print(fibonacci_dp(10))  # Output: 55

# Memoization approach (Top-down)
def fibonacci_memo(n, memo={}):
    if n in memo:
        return memo[n]
    if n <= 1:
        return n
    
    memo[n] = fibonacci_memo(n-1, memo) + fibonacci_memo(n-2, memo)
    return memo[n]

print(fibonacci_memo(10))  # Output: 55

# Longest Common Subsequence
def lcs(X, Y):
    m, n = len(X), len(Y)
    L = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if X[i-1] == Y[j-1]:
                L[i][j] = L[i-1][j-1] + 1
            else:
                L[i][j] = max(L[i-1][j], L[i][j-1])
    
    return L[m][n]

print(lcs("ABCDGH", "AEDFHR"))  # Output: 3

# 0/1 Knapsack Problem
def knapsack(W, wt, val, n):
    K = [[0 for _ in range(W + 1)] for _ in range(n + 1)]
    
    for i in range(n + 1):
        for w in range(W + 1):
            if i == 0 or w == 0:
                K[i][w] = 0
            elif wt[i-1] <= w:
                K[i][w] = max(val[i-1] + K[i-1][w-wt[i-1]], K[i-1][w])
            else:
                K[i][w] = K[i-1][w]
    
    return K[n][W]

val = [60, 100, 120]
wt = [10, 20, 30]
W = 50
n = len(val)
print(knapsack(W, wt, val, n))  # Output: 220

# Coin Change Problem
def coin_change(coins, amount):
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0
    
    for coin in coins:
        for x in range(coin, amount + 1):
            dp[x] = min(dp[x], dp[x - coin] + 1)
    
    return dp[amount] if dp[amount] != float('inf') else -1

print(coin_change([1, 2, 5], 11))  # Output: 3

# Longest Increasing Subsequence
def lis(arr):
    n = len(arr)
    dp = [1] * n
    
    for i in range(1, n):
        for j in range(i):
            if arr[i] > arr[j] and dp[i] < dp[j] + 1:
                dp[i] = dp[j] + 1
    
    return max(dp)

print(lis([10, 22, 9, 33, 21, 50, 41, 60, 80]))  # Output: 6
```

## For Loops

Python's `for` loop is versatile and can be used in various ways. Here are common patterns and use cases:

```python
# Basic for loop with range
for i in range(5):
    print(i)  # Prints 0, 1, 2, 3, 4

# For loop with start, stop, and step
for i in range(2, 10, 2):
    print(i)  # Prints 2, 4, 6, 8

# Iterating over a list
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(fruit)

# Enumerating a list
for index, fruit in enumerate(fruits):
    print(f"Index {index}: {fruit}")

# Iterating over a dictionary
person = {"name": "Alice", "age": 30, "city": "New York"}
for key in person:
    print(f"{key}: {person[key]}")

# Iterating over dictionary items
for key, value in person.items():
    print(f"{key}: {value}")

# List comprehension
squares = [x**2 for x in range(10)]
print(squares)  # [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]

# Dictionary comprehension
square_dict = {x: x**2 for x in range(5)}
print(square_dict)  # {0: 0, 1: 1, 2: 4, 3: 9, 4: 16}

# Nested for loops
for i in range(3):
    for j in range(3):
        print(f"({i}, {j})", end=" ")
    print()

# For loop with else clause
for i in range(5):
    print(i)
else:
    print("Loop completed normally")

# Breaking out of a loop
for i in range(10):
    if i == 5:
        break
    print(i)

# Continuing to the next iteration
for i in range(10):
    if i % 2 == 0:
        continue
    print(i)

# Iterating over multiple sequences with zip
names = ["Alice", "Bob", "Charlie"]
ages = [25, 30, 35]
for name, age in zip(names, ages):
    print(f"{name} is {age} years old")

# Reversing a sequence
for item in reversed(fruits):
    print(item)

# Using a for loop with a function
def process_item(item):
    return item.upper()

results = [process_item(fruit) for fruit in fruits]
print(results)

# Flattening a nested list
nested_list = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
flattened = [item for sublist in nested_list for item in sublist]
print(flattened)  # [1, 2, 3, 4, 5, 6, 7, 8, 9]

# Iterating over a string
for char in "Python":
    print(char)

# Using itertools for more complex iterations
import itertools

# Cartesian product
for item in itertools.product([1, 2], [3, 4]):
    print(item)

# Combinations
for combo in itertools.combinations([1, 2, 3, 4], 2):
    print(combo)

# Permutations
for perm in itertools.permutations([1, 2, 3], 2):
    print(perm)
```

Key points about Dynamic Programming and For Loops in Python:

1. Dynamic Programming:
   - Solves complex problems by breaking them down into simpler subproblems
   - Uses memoization (top-down) or tabulation (bottom-up) to store intermediate results
   - Applicable to problems with overlapping subproblems and optimal substructure
   - Common DP problems include Fibonacci, LCS, Knapsack, Coin Change, and LIS

2. For Loops:
   - Versatile construct for iterating over sequences (lists, tuples, strings, etc.)
   - Can be used with `range()` for numeric iterations
   - Supports enumeration, dictionary iteration, and multiple sequence iteration with `zip()`
   - List and dictionary comprehensions provide concise ways to create new sequences
   - Can be combined with `break`, `continue`, and `else` clauses for more control
   - The `itertools` module provides additional iteration tools for complex scenarios

Remember to refer to the official Python documentation for more detailed information on these topics and best practices for using them effectively.

## Python-specific Operations

Python has a rich standard library and many built-in features. Here are some operations and features that are specific to Python or implemented differently compared to other languages:

```python
# List comprehensions
squares = [x**2 for x in range(10)]
print(squares)  # [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]

# Set comprehensions
even_squares = {x**2 for x in range(10) if x % 2 == 0}
print(even_squares)  # {0, 4, 16, 36, 64}

# Dictionary comprehensions
square_dict = {x: x**2 for x in range(5)}
print(square_dict)  # {0: 0, 1: 1, 2: 4, 3: 9, 4: 16}

# Generator expressions
gen = (x**2 for x in range(10))
print(next(gen))  # 0
print(next(gen))  # 1

# Unpacking
a, *b, c = [1, 2, 3, 4, 5]
print(a, b, c)  # 1 [2, 3, 4] 5

# Multiple assignment
x, y = 10, 20
x, y = y, x  # Swapping variables
print(x, y)  # 20 10

# Slicing with step
numbers = list(range(10))
print(numbers[::2])  # [0, 2, 4, 6, 8]
print(numbers[::-1])  # [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]

# Ternary operator
x = 5
result = "positive" if x > 0 else "non-positive"
print(result)  # positive

# Chained comparisons
x = 5
print(1 < x < 10)  # True

# String multiplication
print("Python" * 3)  # PythonPythonPython

# Set operations
set1 = {1, 2, 3, 4, 5}
set2 = {4, 5, 6, 7, 8}
print(set1 | set2)  # Union: {1, 2, 3, 4, 5, 6, 7, 8}
print(set1 & set2)  # Intersection: {4, 5}
print(set1 - set2)  # Difference: {1, 2, 3}

# Enumerate
for i, value in enumerate(['a', 'b', 'c']):
    print(i, value)

# Zip
for name, age in zip(['Alice', 'Bob', 'Charlie'], [25, 30, 35]):
    print(f"{name} is {age} years old")

# Any and all
print(any([False, False, True]))  # True
print(all([True, True, False]))  # False

# Lambda functions
square = lambda x: x**2
print(square(5))  # 25

# Map, filter, reduce
from functools import reduce
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, numbers))
evens = list(filter(lambda x: x % 2 == 0, numbers))
sum_all = reduce(lambda x, y: x + y, numbers)
print(squared, evens, sum_all)  # [1, 4, 9, 16, 25] [2, 4] 15

# Context managers
with open('example.txt', 'w') as f:
    f.write('Hello, World!')

# Decorators
def uppercase_decorator(func):
    def wrapper():
        result = func()
        return result.upper()
    return wrapper

@uppercase_decorator
def greet():
    return "hello, world!"

print(greet())  # HELLO, WORLD!

# Type hints (Python 3.5+)
def add(x: int, y: int) -> int:
    return x + y

# F-strings (Python 3.6+)
name = "Alice"
age = 30
print(f"{name} is {age} years old")

# Walrus operator (Python 3.8+)
if (n := len(numbers)) > 5:
    print(f"List is long ({n} elements)")

# Dataclasses (Python 3.7+)
from dataclasses import dataclass

@dataclass
class Point:
    x: float
    y: float

p = Point(1.0, 2.0)
print(p)  # Point(x=1.0, y=2.0)
```

## Printing and Formatting Data

Python provides various ways to print and format data, from simple print statements to more complex string formatting techniques.

```python
# Basic print
print("Hello, World!")

# Print multiple items
x, y = 10, 20
print("x =", x, "and y =", y)

# Print without newline
print("Hello", end=" ")
print("World")  # Prints "Hello World"

# Formatted string literals (f-strings, Python 3.6+)
name = "Alice"
age = 30
print(f"{name} is {age} years old")

# Older .format() method
print("{} is {} years old".format(name, age))

# Even older % formatting
print("%s is %d years old" % (name, age))

# Specifying field width and alignment
for i in range(1, 11):
    print(f"{i:2d} {i*i:3d} {i*i*i:4d}")

# Formatting floats
pi = 3.14159
print(f"Pi is approximately {pi:.2f}")

# Formatting percentages
percentage = 0.75
print(f"{percentage:.1%}")  # Prints: 75.0%

# Formatting with comma as thousand separator
large_number = 1234567890
print(f"{large_number:,}")  # Prints: 1,234,567,890

# Left, right, and center alignment
print(f"{'left':<10}{'center':^10}{'right':>10}")

# Padding with zeros
print(f"{42:05d}")  # Prints: 00042

# Using format() with dictionaries
person = {'name': 'Alice', 'age': 30}
print("%(name)s is %(age)d years old" % person)
print("{name} is {age} years old".format(**person))

# Repr vs Str
class Point:
    def __init__(self, x, y):
        self.x, self.y = x, y
    def __str__(self):
        return f"({self.x}, {self.y})"
    def __repr__(self):
        return f"Point({self.x}, {self.y})"

p = Point(1, 2)
print(str(p))  # Prints: (1, 2)
print(repr(p))  # Prints: Point(1, 2)

# Printing to a file
with open('output.txt', 'w') as f:
    print("This goes to the file", file=f)

# Flushing the output
import time
for i in range(5):
    print(i, end=' ', flush=True)
    time.sleep(1)

# Using pprint for pretty-printing complex structures
from pprint import pprint
complex_dict = {'a': [1, 2, 3, 4], 'b': {'x': 10, 'y': 20}}
pprint(complex_dict)

# Tabulate for creating tables
from tabulate import tabulate
data = [["Alice", 25], ["Bob", 30], ["Charlie", 35]]
headers = ["Name", "Age"]
print(tabulate(data, headers=headers, tablefmt="grid"))

# Colorama for colored console output
from colorama import Fore, Back, Style
print(Fore.RED + "This is red text")
print(Back.GREEN + "This has a green background")
print(Style.RESET_ALL)  # Reset all colors and styles

# Using string.Template for simple substitutions
from string import Template
t = Template('$name is $age years old')
print(t.substitute(name=name, age=age))

# Formatted string literals with debugging (Python 3.8+)
x = 10
y = 20
print(f"{x=}, {y=}")  # Prints: x=10, y=20
```

Key points about Python-specific Operations and Printing/Formatting in Python:

1. Python-specific Operations:
   - List, set, and dictionary comprehensions offer concise ways to create and transform data
   - Generator expressions provide memory-efficient iteration
   - Multiple assignment and unpacking simplify variable manipulation
   - Slicing with steps allows for powerful sequence manipulation
   - Built-in functions like `zip`, `enumerate`, `any`, and `all` simplify common operations
   - Lambda functions, map, filter, and reduce offer functional programming capabilities
   - Context managers (with statement) simplify resource management
   - Decorators allow modifying or enhancing functions and classes
   - Type hints improve code readability and enable static type checking
   - Newer Python versions introduce features like f-strings and the walrus operator

2. Printing and Formatting:
   - The `print()` function is versatile, allowing multiple arguments and custom separators/end characters
   - F-strings (formatted string literals) offer a concise and readable way to embed expressions in strings
   - The `.format()` method and %-formatting provide alternative ways to format strings
   - Various format specifiers allow precise control over the appearance of numbers and strings
   - The `pprint` module helps with pretty-printing complex data structures
   - Third-party libraries like `tabulate` and `colorama` can enhance console output
   - File objects can be used as the `file` argument in `print()` to redirect output

Remember to refer to the official Python documentation for more detailed information on these topics and best practices for using them effectively.

## Classes

In Python, classes are used to create user-defined data structures. Classes define a set of attributes and methods that the objects of the class will have.

```python
# Basic class definition
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def greet(self):
        return f"Hello, my name is {self.name} and I'm {self.age} years old."

# Creating an instance of the class
alice = Person("Alice", 30)
print(alice.greet())  # Output: Hello, my name is Alice and I'm 30 years old.

# Class with class variable and static method
class Math:
    pi = 3.14159  # Class variable
    
    @staticmethod
    def square(x):
        return x * x

print(Math.pi)  # Output: 3.14159
print(Math.square(5))  # Output: 25

# Inheritance
class Employee(Person):
    def __init__(self, name, age, employee_id):
        super().__init__(name, age)
        self.employee_id = employee_id
    
    def work(self):
        return f"{self.name} is working. Employee ID: {self.employee_id}"

bob = Employee("Bob", 35, "E12345")
print(bob.greet())  # Output: Hello, my name is Bob and I'm 35 years old.
print(bob.work())  # Output: Bob is working. Employee ID: E12345

# Multiple inheritance
class A:
    def method(self):
        print("Method from A")

class B:
    def method(self):
        print("Method from B")

class C(A, B):
    pass

c = C()
c.method()  # Output: Method from A (Method Resolution Order)

# Property decorators
class Circle:
    def __init__(self, radius):
        self._radius = radius
    
    @property
    def radius(self):
        return self._radius
    
    @radius.setter
    def radius(self, value):
        if value < 0:
            raise ValueError("Radius cannot be negative")
        self._radius = value
    
    @property
    def area(self):
        return 3.14159 * self._radius ** 2

circle = Circle(5)
print(circle.radius)  # Output: 5
circle.radius = 10
print(circle.area)  # Output: 314.159

# Magic methods (dunder methods)
class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def __str__(self):
        return f"({self.x}, {self.y})"
    
    def __add__(self, other):
        return Point(self.x + other.x, self.y + other.y)
    
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

p1 = Point(1, 2)
p2 = Point(3, 4)
print(p1)  # Output: (1, 2)
print(p1 + p2)  # Output: (4, 6)
print(p1 == Point(1, 2))  # Output: True

# Abstract Base Class
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

class Square(Shape):
    def __init__(self, side):
        self.side = side
    
    def area(self):
        return self.side ** 2

# Data Classes (Python 3.7+)
from dataclasses import dataclass

@dataclass
class Product:
    name: str
    price: float
    quantity: int = 0

    def total_cost(self) -> float:
        return self.price * self.quantity

product = Product("Apple", 0.5, 10)
print(f"Total cost: ${product.total_cost()}")  # Output: Total cost: $5.0

# Class and static methods
class MyClass:
    @classmethod
    def class_method(cls):
        print("This is a class method")
    
    @staticmethod
    def static_method():
        print("This is a static method")

MyClass.class_method()  # Output: This is a class method
MyClass.static_method()  # Output: This is a static method
```

## Functions

Functions in Python are first-class objects, which means they can be assigned to variables, passed as arguments, and returned from other functions.

```python
# Basic function definition
def greet(name):
    return f"Hello, {name}!"

print(greet("Alice"))  # Output: Hello, Alice!

# Function with default parameter
def power(base, exponent=2):
    return base ** exponent

print(power(3))  # Output: 9
print(power(3, 3))  # Output: 27

# Function with variable-length arguments
def sum_all(*args):
    return sum(args)

print(sum_all(1, 2, 3, 4, 5))  # Output: 15

# Function with keyword arguments
def print_info(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")

print_info(name="Alice", age=30, city="New York")

# Lambda functions (anonymous functions)
square = lambda x: x**2
print(square(5))  # Output: 25

# Higher-order functions
def apply_twice(func, arg):
    return func(func(arg))

print(apply_twice(square, 3))  # Output: 81

# Closure
def make_multiplier(n):
    def multiplier(x):
        return x * n
    return multiplier

times_three = make_multiplier(3)
print(times_three(5))  # Output: 15

# Decorator
def uppercase_decorator(func):
    def wrapper():
        result = func()
        return result.upper()
    return wrapper

@uppercase_decorator
def greet():
    return "hello, world!"

print(greet())  # Output: HELLO, WORLD!

# Generator function
def countdown(n):
    while n > 0:
        yield n
        n -= 1

for num in countdown(5):
    print(num)  # Output: 5 4 3 2 1

# Function with type hints (Python 3.5+)
def add(x: int, y: int) -> int:
    return x + y

print(add(3, 5))  # Output: 8

# Recursive function
def factorial(n: int) -> int:
    if n == 0 or n == 1:
        return 1
    return n * factorial(n - 1)

print(factorial(5))  # Output: 120

# Partial functions
from functools import partial

def power(base, exponent):
    return base ** exponent

square = partial(power, exponent=2)
cube = partial(power, exponent=3)

print(square(4))  # Output: 16
print(cube(4))  # Output: 64

# Function with mutable default argument (be careful!)
def append_to(element, to=[]):
    to.append(element)
    return to

print(append_to(1))  # Output: [1]
print(append_to(2))  # Output: [1, 2] (not [2]!)

# Correct way to use mutable default argument
def append_to_correct(element, to=None):
    if to is None:
        to = []
    to.append(element)
    return to

print(append_to_correct(1))  # Output: [1]
print(append_to_correct(2))  # Output: [2]

# Unpacking arguments
def print_coordinates(x, y, z):
    print(f"X: {x}, Y: {y}, Z: {z}")

coords = [1, 2, 3]
print_coordinates(*coords)  # Output: X: 1, Y: 2, Z: 3

# Function attributes
def func():
    pass

func.attribute = 42
print(func.attribute)  # Output: 42

# Nested functions
def outer(x):
    def inner(y):
        return x + y
    return inner

add_five = outer(5)
print(add_five(3))  # Output: 8
```

Key points about Classes and Functions in Python:

1. Classes:
   - Used to create user-defined data structures
   - Support inheritance, including multiple inheritance
   - Can have class variables, instance variables, and methods
   - Special methods (magic methods) allow customization of behavior
   - Properties can be used for getter/setter functionality
   - Abstract base classes define interfaces
   - Data classes provide a concise way to create classes that mainly store data

2. Functions:
   - First-class objects in Python
   - Can have default parameters, variable-length arguments, and keyword arguments
   - Lambda functions provide a way to create small anonymous functions
   - Closures allow functions to capture and carry their environment
   - Decorators modify or enhance functions
   - Generator functions use `yield` to create iterators
   - Type hints improve readability and enable static type checking
   - Recursive functions call themselves
   - Partial functions create new functions with some arguments preset

Remember to refer to the official Python documentation for more detailed information on these topics and best practices for using them effectively.

## Abstract Base Classes and Protocols

Python uses Abstract Base Classes (ABCs) and Protocols to define interfaces. ABCs are part of the `abc` module, while Protocols are part of the `typing` module (Python 3.8+).

```python
# Abstract Base Classes (ABCs)
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

    @abstractmethod
    def perimeter(self):
        pass

class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height

    def perimeter(self):
        return 2 * (self.width + self.height)

# rect = Shape()  # This would raise TypeError
rect = Rectangle(5, 3)
print(f"Area: {rect.area()}, Perimeter: {rect.perimeter()}")

# Abstract properties
class Animal(ABC):
    @property
    @abstractmethod
    def sound(self):
        pass

class Dog(Animal):
    @property
    def sound(self):
        return "Woof!"

dog = Dog()
print(dog.sound)  # Output: Woof!

# Abstract class methods and static methods
class AbstractClass(ABC):
    @classmethod
    @abstractmethod
    def class_method(cls):
        pass

    @staticmethod
    @abstractmethod
    def static_method():
        pass

class ConcreteClass(AbstractClass):
    @classmethod
    def class_method(cls):
        return "Implemented class method"

    @staticmethod
    def static_method():
        return "Implemented static method"

# Protocols (Python 3.8+)
from typing import Protocol, runtime_checkable

@runtime_checkable
class Drawable(Protocol):
    def draw(self) -> None:
        ...

class Circle:
    def draw(self) -> None:
        print("Drawing a circle")

class Square:
    def draw(self) -> None:
        print("Drawing a square")

def draw_shape(shape: Drawable) -> None:
    shape.draw()

circle = Circle()
square = Square()

draw_shape(circle)  # Output: Drawing a circle
draw_shape(square)  # Output: Drawing a square

print(isinstance(circle, Drawable))  # Output: True
print(isinstance(square, Drawable))  # Output: True

# Structural subtyping with Protocols
from typing import Iterable, TypeVar

T = TypeVar('T')

def first(collection: Iterable[T]) -> T:
    return next(iter(collection))

print(first([1, 2, 3]))  # Output: 1
print(first("hello"))    # Output: h
print(first({1: 'a', 2: 'b'}))  # Output: 1

# Combining ABCs and Protocols
from typing import runtime_checkable, Protocol
from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def make_sound(self):
        pass

@runtime_checkable
class Swimable(Protocol):
    def swim(self) -> None:
        ...

class Fish(Animal, Swimable):
    def make_sound(self):
        return "Blub"

    def swim(self):
        print("Swimming")

fish = Fish()
print(fish.make_sound())  # Output: Blub
fish.swim()  # Output: Swimming
print(isinstance(fish, Swimable))  # Output: True
```

## Python References and Resources

This section provides links to official Python documentation and other valuable resources for learning and reference.

### Official Python Documentation

1. [Python Official Website](https://www.python.org/)
   - The main hub for all things Python, including downloads, documentation, and community resources.

2. [Python Language Reference](https://docs.python.org/3/reference/index.html)
   - The authoritative reference for Python's syntax and "core semantics".

3. [Python Standard Library](https://docs.python.org/3/library/index.html)
   - Documentation for all modules and functions in Python's extensive standard library.

4. [Python HOWTOs](https://docs.python.org/3/howto/index.html)
   - In-depth documents on specific topics, like "Sorting HOW TO" and "Logging HOWTO".

5. [Python Enhancement Proposals (PEPs)](https://www.python.org/dev/peps/)
   - Index of Python Enhancement Proposals, including PEP 8 (Style Guide for Python Code).

### Learning Resources

6. [The Python Tutorial](https://docs.python.org/3/tutorial/index.html)
   - An official tutorial for beginners to get started with Python.

7. [Real Python](https://realpython.com/)
   - A website with a wide range of Python tutorials, from basics to advanced topics.

8. [Python for Everybody](https://www.py4e.com/)
   - A free course by Dr. Charles Severance, covering Python basics and more.

### Community and Additional Resources

9. [Python Package Index (PyPI)](https://pypi.org/)
   - The official repository for third-party Python packages.

10. [Awesome Python](https://awesome-python.com/)
    - A curated list of awesome Python frameworks, libraries, software and resources.

11. [Python Discord](https://discord.gg/python)
    - A large community Discord server for Python discussions and help.

12. [Talk Python To Me Podcast](https://talkpython.fm/)
    - A podcast on Python and related technologies.

### Books

13. "Python Crash Course" by Eric Matthes
    - A fast-paced, thorough introduction to Python for beginners.

14. "Fluent Python" by Luciano Ramalho
    - A comprehensive guide to writing effective Python code.

### Tools and Editors

15. [PyCharm](https://www.jetbrains.com/pycharm/)
    - A full-featured IDE for Python development by JetBrains.

16. [Visual Studio Code](https://code.visualstudio.com/docs/languages/python)
    - A popular, free editor with excellent Python support.

17. [Jupyter Notebooks](https://jupyter.org/)
    - An open-source web application for creating and sharing documents with live code, equations, visualizations, and narrative text.

### Online Python Environments

18. [Python Tutor](http://pythontutor.com/)
    - Visualize Python code execution step by step.

19. [Repl.it](https://repl.it/languages/python3)
    - An online Python environment for coding and collaborating.

### Style Guides and Best Practices

20. [PEP 8 -- Style Guide for Python Code](https://www.python.org/dev/peps/pep-0008/)
    - The official style guide for Python code.

21. [Google Python Style Guide](https://google.github.io/styleguide/pyguide.html)
    - Google's Python style guide, which is widely respected in the community.

Remember to check these resources regularly, as the Python language and its ecosystem are continually evolving. These references will help you stay up-to-date with the latest developments and best practices in Python programming.

