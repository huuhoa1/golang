package main
import "fmt"
func main() {
   foo := newMyStruct()
   foo.myMap["bar"] = "baz"
   fmt.Println(foo)
}

type myStruct struct {
   myMap map[string]string //field
}

func newMyStruct() *myStruct {
   result := myStruct{} // go keeps track objects created inside the function!
   result.myMap = map[string]string{}
   return &result
}
