# GO

Statically typed programming language.
Great garbage collector
Go routines.

# Features

1. Simplicity:-  easy to learn and understand for both beginners and experienced.

2. Concurrency:- Built in support for concurrency.

Concurrency handles multiple tasks by interleaving their execution, while parallelism involves simultaneously executing independent tasks using multiple resources.

3. Garbage collection: Go has automatic memory management, which frees developers from having to worry about memory allocation and deallocation

4. Strong typing: Go is a statically typed language, which helps catch errors at compile time rather than at runtime.

# Complilation flow

1. The scanner, which converts the source code into a list of tokens, for use by the parser.

2. The parser, which converts the tokens into an Abstract Syntax Tree to be used by code generation.

3. The code generation, which converts the Abstract Syntax Tree to machine code.

# Variables 
  Variables in programming are used to store and manipulate data values. They provide a way to give names to memory locations where data can be stored.
  (it is storing )

# Identifiers

 In Go language, an identifier can be a variable name, function name, constant, statement labels, package name, or types.
 (it is representing the storing )
 
# Keywords
 Keywords or Reserved words are the words in a language that are used for some internal process or represent some predefined actions. 
 there are 25 keywords present in golang.


# DataTypes
  
  DataTypes will be used to define the what is the valid type that variable can store in it.

1. Basic type: Numbers, strings and booleans come under this category.
2. Aggregate type: Array and structs come under this category.
3. Reference type: Pointers, slices, maps, functions, and channels come under this category.
4. Interface type


# Constants

 once the value of constant is defined, it cannot be modified further

1. Untyped  [](./examples.go#L4)
2. Typed     [](./examples.go#L5)

# Operators

Operators allow us to perform different kinds of operations on operands. 

1. Arithmetic Operators (+,-,/,%,*)
2. Relational Operators (>=,<=,==,>,<)
3. Logical Operators (&&,||,!)
4. Bitwise Operators (&,|,<<,>>)
5. Assignment Operators (=,:=,+=,-=)
6. Misc Operators (*)

# Decision Making

1. if som{

}

2. if some{

}else{

}

3. if some{

} else if (other){

}else{

}

# Function

Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code which ultimately saves the excessive use of memory

1. Call by value: 
   
   any changes made inside functions are not reflected in actual parameters of the caller.

2. Call by reference:
     
     any changes made inside the function are actually reflected in actual parameters of the caller.

# Variadic parameters

Variadic functions can be used instead of passing a slice to a function

useful when the arguments to your function are most likely not going to come from a single data structure (i.e. slice or array etc.)


# Anonymous Function

An anonymous function is a function which doesn’t contain any name. It is useful when you want to create an inline function. ( function literal)

# Main Function

The main() function is a special type of function and it is the entry point of the executable programs.
It does not take any argument nor return anything.

# Init()

does not take any argument nor return anything. This function is present in every package and this function is called when the package is initialized


# Blank Identifier(underscore)

Golang has a special feature to define and use the unused variable using Blank Identifier. 

# Defer

 defer function or method call arguments evaluate instantly, but they don’t execute until the nearby functions returns. 

 multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) order 

 # Structures

  Any real-world entity which has some set of properties/fields can be represented as a struct.

  Go does not support inheritance, but you can use composition to achieve similar results.

# Zeroth Value
   default value assigned to variables of certain types when they are declared without an explicit initialization.


# Go Routine
  A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program

# Select Statement
 the select statement is just like switch statement, but in the select statement, case statement refers to communication, i.e. sent or receive operation on the channel.

 # Channel 

  channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. 

  Send operation:- Mychannel <- element

  Receive operation:- element := <-Mychannel




1. What sparked your interest in programming with Golang?
2. Concurrency vs. parallelism in Golang's Goroutines?
3. Golang's error handling benefits?
4. Impact of simplicity and readability in your code?
5. Example of using the defer statement?
6. Practical use of interfaces in your projects?
7. Explanation of Channels and their use?
8. Influence of Golang's garbage collector?
9. Practical example of using pointers?
10. How have modules improved dependency management?
11. Scenario where Golang's performance was advantageous?
12. Role of static typing in code quality?



1. Interest in Golang: I was drawn to Golang for its speed, simplicity, and built-in concurrency support.

2. Concurrency vs. Parallelism: Concurrency is about managing tasks that might overlap in time, while parallelism is about executing tasks simultaneously. Goroutines enable concurrency in Golang.

3. Error Handling: Golang's error handling uses explicit return values for errors, promoting clear and concise error management.

Simplicity and Readability: Golang's clean syntax and straightforward design make code easier to understand and maintain.

defer Example: Using defer, I can ensure resources are cleaned up, like closing files, when a function exits.

Interfaces in Projects: Interfaces allow me to create flexible code that works with different types that satisfy the same contract.

Channels: Channels enable safe communication between Goroutines, avoiding race conditions.

Garbage Collector Impact: The garbage collector automatically manages memory, reducing the need for manual memory management.

Pointer Example: Pointers are used to modify values directly in memory, like updating elements in a large array.

Module Dependency Improvement: Modules simplify package management by providing versioned dependencies and better control.

Performance Advantage: Golang's compiled nature and lightweight Goroutines make it great for building high-performance concurrent applications.

Static Typing and Code Quality: Static typing catches type-related errors early, improving code reliability and maintainability.





