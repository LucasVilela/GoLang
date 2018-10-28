##Naming

- If t he name b eg ins with an upper-case letter, it is exported
- Package names always lowercase
- Camel case

## Declarations

- `var, const, type, and func.`

## Assignments

The value held by a variable is update d by an assignment statement, which in its simplest form has a variable on the left of the = sign and an expression on the right.

```go
     x = 1 //named variable
     *p = true // indirect variable
     person.name = "bob" // struct field
     count[x] = count[x] * scale // array or slice or map element
```

Instead of having to repeat and re-evaluate you can use _assignment_ operator

```go
count[x] = count[x] * scale

count[x] *= scale
```

Increment or Decrement

```go
v := 1
v++
v--
```

## Tupe Assignment

Allow several variables to be assigned at once. All the values are evaluated before any of variables are updated, making useful when some of the variables appear on both sides of the assignment, for example, swapping the values of two variables

```go
x,y = y, x
```

```go
func greatestCommonDivisor(x, y int) int {
  for y != 0 {
    x, y = y, x%y
}
	return x
}
```

Also make assignments more compact

```go
a, b, c, d = 10, 30, 20, 19
```

_Even thought they are easier they are more_ **complex**

> Certain expressions have multiple results, and produce several values, make sure to have as many variables as the function result
>
> ```go
> 	f, err = os.Open("foo.txt")  // function call returns two values
> ```
>
> ```go
> v, ok = m[key] // map lookup
> v, ok = x.(T) // type assertion
> v, ok = <-ch // channel receive
> ```
>
> We can assing unwanted values with the blank identifier
>
> ```go
> _, err = io.Copy(dst, src) // discard byte count
>     _, ok = x.(T)              // check type but discard result
> ```
>
> .

### Assignability

Assignment statements are an explicit for m of assignment, but t here are many places in a prog ram w here an assignment occurs _implicitly_

explicitly

```go
nature :=[]string{"ocean","forest","desert"}
```

implicitly

```go
nature[0] = "ocean"
nature[1] = "forest"
nature[2] = "desert"
```

## Variables

General form
`var name type = expression`

- `var names []string` ?

* Type declarations
  The t yp e of a var i able or expression defines t he charac ter ist ics of t he values it may t ake on, such as their size.

Used to create a new _named type_ example:

`type Lucas string`

```go
  type Celsius float64
     type Fahrenheit float64
     const (
         AbsoluteZeroC Celsius = -273.15
         FreezingC     Celsius = 0
         BoilingC      Celsius = 100
     )
     func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
     func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

This case the units are both numbers but they can't be compared because they are different types. useful to avoid errors like combining numbers that are in different scales

```go
    fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
    boilingF := CToF(BoilingC)
    fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
    fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
```

Comparisson operators can be uses but two values from different name types can't be compared

```go
var c Celsius
var f Fahrenheit
fmt.Println(c == 0)
fmt.Println(f >= 0)
fmt.Println(c == f)
fmt.Println(c == Celsius(f)) // "true"!
```

## Packages and Files

Packages in Go serve the same purposes as libraries or modules in ot her languages, supporting modularity, encapsulation, separate compilation, and reuse.

Controlled by directory, `$GOPATH/src/gopl.io/ch1/helloworld.` for example for use one function.

_Exported identifiers start with an upper-case letter._
