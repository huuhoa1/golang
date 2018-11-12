package main

func main() {
   sayHello("Hello", "Go", "from", "PST")
}
func sayHello(messages ...string) { // passing a slice of strings
   for _, message := range messages {
      println(message)
   }
}
