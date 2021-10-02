# Chapter 3. Go Control Flow
----

Go borrows several of its control flow syntax from C-family of languages.

it supports all of the expected control structures:

- ``if ... else``
- ``switch``
- ``for``
- ``goto``


## The ``if`` statement

The statement conditionally executes a code block when the Boolean expression that follows the ``if`` keyword evaluates to ``true``.

- The parentheses around the test expression are not necessary.
While the following ``if`` statement will compile, it is ```not idiomatic```.

```go
    if (num0 > 100 || num0 < 900) { 
        fmt.Println("Currency: ", num0) 
        printCurr(num0) 
      } 
```

- The curly braces for the code block are always required.

```go
     if num0 > 100 || num0 < 900 { printCurr(num0)} 
```
- prefered idiomatic layout is to use multiple lines encourage good style and clarity
```go
    if num0 > 100 || num0 < 900 { 
         printCurr(num0)
    } 
```

## The ``else`` block

- The if statement may include and optional `` else`` block, wich is executed when the expression in the ``if`` block evaluates to ``false``
- The code in the ``else`` block must be wrapped in curly braces using multiple lines.

```go
if num0 > 100 || num0 < 900 { 
    fmt.Println("Currency: ", num0) 
    printCurr(num0) 
} else { 
    fmt.Println("Currency unknown") 
}  
```

- The ``else`` keyword may be immediately followed by another ``if `` statement forming an ``if ... else ... if`` chain.

```go
if CAD.Number == number { 
    fmt.Printf("Found: %+v\n", CAD) 
} else if FJD.Number == number { 
    fmt.Printf("Found: %+v\n", FJD) 
} 
```

- The ``if ... else ... if`` statemement chain can grow as long as needed and may be terminated by an optional ``else`` statement to expresss all other untested options 

```go
func printCurr(number int) { 
if CAD.Number == number { 
        fmt.Printf("Found: %+v\n", CAD) 
    } else if FJD.Number == number { 
        fmt.Printf("Found: %+v\n", FJD) 
    } else if JMD.Number == number { 
        fmt.Printf("Found: %+v\n", JMD) 
    } else if USD.Number == number { 
        fmt.Printf("Found: %+v\n", USD) 
    } else { 
        fmt.Println("No currency found with number", number) 
    }
}

```

## The ``if`` Statement initialization

- At runtime, the initialization is executed before the test expression is evaluated, as illustrated in this code snippet

```go
if num1 := 388; num1 > 100 || num1 < 900 { 
  fmt.Println("Currency:", num1) 
  printCurr(num1) 
}  
```

- The initialization follows normal variable declaration and initialization rules.

- The scope of the initialized variables is bound to the if statement block, beyond which they become unreachable
