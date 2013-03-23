functgo
=======

Go Implementation of higher-order functions {map, reduce, filter}

Intro
=====

The idea behind this project is to provide a way for Go programmers to use ubiquitous functional tools, such as Map, Reduce and Filter.
The project aim is pretty simple.

Examples
========

Map
---

    package main

    import ( "fmt"
         "reflect"
         "github.com/neurogeek/functgo"
    )

    func main() {

      // Anon function that calculates the double of
      // a number
      d := func(a int64) int64 { return (a * 2) }

      // Calling FMap from functgo. First argument is a Slice, second
      // the reflect.Value of the function we want to apply to the elements
      // in the slice
      r := functgo.FMap([]int64{1, 2, 3}, reflect.ValueOf(d)).([]int64)

      // In order to actually use range, we need to assert the type 
      // of the resulting array, hence the .([]int64) part of the line above.

      for _, v := range r {
        fmt.Println(v)
      }
    }


Reduce
------

    package main

    import ( "fmt"
         "reflect"
         "github.com/neurogeek/functgo"
    )

    func main() {

      // Anon. function that returns the sum of two numbers
      f := func(a int64, b int64) int64 { return a + b }

      // Call functgo.FReduce on a Slice as the first argument, and the 
      // reflect.Value of the reduce function.
      m := FReduce([]int64{1,2,3,4}, reflect.ValueOf(f))

      // Result is a single element
      fmt.Println("Reduction of ", []int64{1,2,3,4}, " is ", m)
    }

Cont.
...
