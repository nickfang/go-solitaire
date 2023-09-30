package main
import "fmt"
func main() {

  array := [2]int{1, 2}
  newArray := make([]int, len(array)+2)

  // Copy the elements from the old array to the new array.
  copy(newArray[:len(array)], array[:])

  // Append the new elements to the new array.
  newArray[len(array)] = 3
  newArray[len(array)+1] = 4

  fmt.Println(newArray)
}
