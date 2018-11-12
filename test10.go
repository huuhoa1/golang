package main

func main() {
    mp := messagePrinter{"foo"}
    mp.printMessage()
}

type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() { //declaring a struct method
    println(mp.message)
}

