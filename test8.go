package main
import "fmt"
func main() {
   //foo := myStruct{"bar"} //like a constructor on stack
   //foo := new(myStruct) //like a constructor on heap
   foo := &myStruct{"bar"} //like a constructor on heap
   //foo.myField = "bar"
   fmt.Println(foo)
}

type myStruct struct {
   myField string //field
}
