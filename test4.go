package main

func main() {
   message := "Hello"
   sayHello(message)
   sayHello1(&message)
   println(message)
}
func sayHello(message string) { //value passing
   println(message)
}
func sayHello1(message *string) { //reference passing
   println(*message)
   *message = "Hello Go"
}
