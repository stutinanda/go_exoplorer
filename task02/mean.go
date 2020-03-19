// Calc mean of two numbers.
package main

import (
  "fmt" //formatting package
)


func main () {

  // FIRST WAY
  var x int  // by default this value is set to 0
  var y int
  // %v prints values
  fmt.Printf("this is x=%v \n",x)
  fmt.Printf("this is y=%v", y)
  x = 1
  y = 2
  // %T prints type of x
  fmt.Printf("x= %v, type of %T\n", x, x)
  fmt.Printf("y= %v, type of %T\n", y, y)
  var mean1 int
  mean1 = (x + y )/ 2
  fmt.Printf("result for ints: %v type of %T\n", mean1, mean1)



  // SECOND WAY
  var a float64 = 1
  var b float64 = 2.0
  var mean2 float64
  mean2 = (a + b) / 2.0
  fmt.Printf("result for floats: %v type of %T\n", mean2, mean2)

  // THIRD WAY
  p := 1.0
  q := 2.0
  mean3 := (p + q)/2.0
  fmt.Printf("result for p:= : %v type of %T\n", mean3, mean3)

  // FOURTH WAY

  r , s := 1.0, 2.0
  mean4 := (r + s)/2.0
  fmt.Printf("result for r, s initiated together := : %v type of %T\n", mean4, mean4)



}
