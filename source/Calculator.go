package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
    fmt.Printf("%v", promt)
    input,_ := reader.ReadString('\n')
    value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    if err != nil {
        message,_ := fmt.Scanf("%v must number only", promt)
        panic(message)
    }
    
    return value
}

func getOperator() string {
    fmt.Print("operator is ( + - * /): ")
    op,_ := reader.ReadString('\n')
    return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
    return value1 + value2
}

func remove(value1, value2 float64) float64 {
    return value1 - value2
}

func multipile(value1, value2 float64) float64 {
    return value1 * value2
}

func divide(value1, value2 float64) float64 {
    return value1 / value2
}

func main() {
    value1 := getInput("value1 = ")
    value2 := getInput("value2 = ")
    
    var Result float64
    
    switch operator := getOperator(); operator {
        case "+":
            Result = add(value1, value2)
        case "-":
            Result = remove(value1, value2)
        case "*":
            Result = multipile(value1, value2)
        case "/":
            Result = divide(value1, value2)
        default:
	        panic("wrong operator")
    }
    
    fmt.Printf("Result is %v", Result)
}