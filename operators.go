package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a)

  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)

  fmt.Print(10*5)

  var x = 10
  x +=5
  fmt.Println(x)

  var x = 5
  var y = 3
  fmt.Println(x>y) // returns 1 (true) because 5 is greater than 3

//   x < 5 &&  x < 10
//   x < 5 || x < 4
//   !(x < 5 && x < 10)

  var x = 9
  var y = 8

  fmt.Printf("x = %b\n",x)
  fmt.Printf("y = %b\n",y)
    
  fmt.Printf("x & y is %b\n",x & y)
}