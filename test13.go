package main

import "time"
import "runtime"

func main() {
   runtime.GOMAXPROCS(2)
   ch := make(chan string)
   //abcGen()
   go abcGen(ch)
   go printer(ch)
   println("this comes first")
   time.Sleep(100 * time.Millisecond)
}

func abcGen(ch chan string) {
   for l := byte('a'); l <= byte('z'); l++ {
      ch <- string(l)
   }
}
func printer(ch chan string) {
   for {
      println(<-ch)
   }
}
