package main

import "fmt"
import "time"

func main(){
  fmt.Println("Hello, World!")
  fmt.Println("go" + "lang")
  fmt.Println("20+1 = ", 20 + 1)
  fmt.Println("10.5+1.3 = ", 10.5 + 1.3)
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true) 

  fmt.Println("----- Variables -----")

  var a = "initial"
  fmt.Println(a)

  var b,c int = 1,8
  fmt.Println(b,c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  f:= "apple"
  fmt.Println(f)

  var apl string = "apple"
  fmt.Println(apl)

  fmt.Println("----- For -----")

  i := 1
  for i <=3 {
    fmt.Println(i)
    i = i + 1
  }

  fmt.Println("----- For 2 -----")
  for j := 0; j < 3; j++ {
    fmt.Println(j)
  }

  fmt.Println("----- If/Else -----")

  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
      fmt.Println("7 is odd")
    }

  if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

  if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

  if num := 99; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
     

  fmt.Println("----- Switch -----")

    k := 2
    fmt.Print("Write ", i, " as ")
    switch k {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

