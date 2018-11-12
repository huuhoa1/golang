package main

func main() {
    emp := enhancedMessagePrinter{}
    emp.message="foo"
    emp.printMessage()
}

type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() { //declaring a struct method
    println(mp.message)
}

type enhancedMessagePrinter struct {
    messagePrinter //anonymous field
}

