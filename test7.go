package main

func main() {
   addFunc := func(terms ...int) (int, int) { // passing a slice of int 
   result := 0
   for _, term := range terms {
      result += term
   }
   return len(terms),result
}
   numTerms, sum := addFunc(1,3,5,9)
   println("Added:", numTerms, "terms for a total of", sum)
}
