package main

import "fmt"

func add(a int, b int) (int, int){
    
	//fmt.Println(a * b)
    return (a + b), (a * b)
}


func main() {
	fmt.Println(add(10, 20))

}

