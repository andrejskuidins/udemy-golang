//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
  count := 50
  for i := 0; i < count; i++ {
    divisibilityBy3 := i%3 == 0
    divisibilityBy5 := i%5 == 0
    if divisibilityBy3 && divisibilityBy5 {
      fmt.Println("FizzBuzz")
    } else if divisibilityBy3 {
      fmt.Println("Fizz")
    } else if divisibilityBy5 {
      fmt.Println("Buzz")
    }
  }
}
