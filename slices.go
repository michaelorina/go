package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)

  var myarray = [length]datatype{values} // An array
  myslice := myarray[start:end] // A slice made from the array

  myslice3 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice3)
  fmt.Printf("length = %d\n", len(myslice3))
  fmt.Printf("capacity = %d\n", cap(myslice3))

  // with omitted capacity
  myslice4 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice4)
  fmt.Printf("length = %d\n", len(myslice4))
  fmt.Printf("capacity = %d\n", cap(myslice4))

  prices := []int{10,20,30}
  prices[2] = 50
  fmt.Println(prices[0])
  fmt.Println(prices[2])
}