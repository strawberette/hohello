package main

import (
	"errors"
	"fmt"
	// "math"
)

// func main() {
// var firstName = "Noemi"
// // var cuteHeart = "<3"
// // isNoemi := firstName == "Noemi"
// // if isNoemi == true {
// // fmt.Println(firstName, cuteHeart)}
// // if firstName == "Noemi" {
// // 	fmt.Println(firstName, cuteHeart)}
// // number := 40
// // fmt.Println("Hello "+firstName+" "+cuteHeart+" ", number)

// for i := 0; i<10; i++ {
// 	fmt.Println(firstName)
// }

// a := 1
// b := 7
// c := 9

// sum := a + b + c
// fmt.Println(sum)

// item1 := 1.95
// item2 := 4.95
// item3 := 7.99

// totalPrice := item1 + item2 + item3
// vatRate := 0.2
// vat := totalPrice * vatRate
// vatPrice := totalPrice + vat

// fmt.Println("Price before vat: ", math.Round(totalPrice*100)/100)
// fmt.Println("Price after vat: ", math.Round(vatPrice*100)/100)
// age := 19
// if age >= 18 {
// 	fmt. Println("You can buy some glue")
// } else {
// 	fmt.Println("Sorry, we cannot sell glue to you")
// }
// dansMood := "happy"

// isDanHappy := dansMood == "happy"
// fmt.Println(isDanHappy)

// isDanSad := dansMood == "sad"
// fmt.Println(isDanSad)
// totalSheep := 5
// sheepNumber := 1

// for sheepNumber <= totalSheep {
// 	fmt.Println(sheepNumber, "sheep")
// 	sheepNumber++
// }
// fmt.Println("All ships sheared")

// totalSheep := 5

// for sheepNumber := 1; sheepNumber <= totalSheep; sheepNumber++ {
// 	fmt.Println(sheepNumber, "sheep")

// }
// fmt.Println("All ships sheared")

// sum := 0
// for i := 0; i<10; i++ {
// 	sum += i
// }
// fmt.Println(sum)

// 	tot := 1
// 	for i := 1; i<=10; i++ {
// 		tot *= i
// 		fmt.Println(tot)
// 	}
// }
// func sum(number1 int, number2 int) int {
// 	return number1 + number2
// }

// func greet(name string) error {
// 	if name == " " {
// 		return errors.New("empty name")
// 	}
// 	fmt.Println("Hello and welcome", name)
// 	return nil
// }

func add(num1 int, num2 int) (int, error) {
	if (num1 < 0) || (num2 < 0) {
		return -1, errors.New("Negative number")
	}
	result := num1 + num2
	return result, nil
}

func main() {
	// number1 := 7
	// number2 := 15
	// tot := sum(number1, number2)
	// fmt.Println(tot)

	// user := " "
	// err := greet(user)

	// if err != nil {
	// 	fmt.Println("Try again: supply user name")
	// }
	number1 := 7
	number2 := 15
	tot, err := add(number1, number2)

	if err == nil {
		fmt.Println(tot)
	} else {
		fmt.Println("The number you chose is a negative number")
	}
}
