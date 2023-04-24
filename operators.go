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
}