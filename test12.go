package main

import "time"
import "runtime"

func main() {
   runtime.GOMAXPROCS(2)
   //abcGen()
   go abcGen()
   println("this comes first")
   time.Sleep(100 * time.Millisecond)
}

func abcGen() {
   for l := byte('a'); l <= byte('z'); l++ {
      go println(string(l))
   }
}
