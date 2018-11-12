package main


func main() {
   foo := 2
   if foo == 1 {
      println("bar")
   } else {
      println("buz")
   }

   foo = 1
   switch {
   case foo == 1:
      println("one")
   case foo > 2:
      println("two")
   }

   for i:=0; i < 5; i++ {
      println(i)
   }

   m := make(map[string]string)
   m["first"] = "foo"
   m["second"] = "foo2"
   m["third"] = "foo3"

   for k,v := range m {
      println(k, v)
   }
}

