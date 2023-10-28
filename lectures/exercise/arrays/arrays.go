//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type ShoppingList struct {
  name  string
  price int
}

func main() {
  var sum int
  shoplist := [...]ShoppingList{
    { name : "apples", price : 10 },
    { name : "lemons", price : 20 },
    { name : "grapes", price : 30 },
    { name : "vines", price : 40 },
  }

  for i := 0; i < len(shoplist); i++ {
    fmt.Println(shoplist[i].price)
    sum += shoplist[i].price
  }
  fmt.Println(sum)
}
