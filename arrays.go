package main
import ("fmt")


var array_name = [length]datatype{values} // here length is defined

var array_name = [...]datatype{values} // here length is inferred

array_name := [length]datatype{values} // here length is defined

array_name := [...]datatype{values} // here length is inferred

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)

  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)

  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])

  prices[2] = 50
  fmt.Println(prices)
}