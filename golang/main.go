package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab5"
)

func prnt(ylist, xlist []float64) {
	for i, v := range ylist {
		fmt.Printf("x = %f  y = %f\n", xlist[i], v)
	}
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//лабораторная работа №2.
	fmt.Println("Горюнов Александр Алексеевич")

	//лабароторная работа №3 + 4 - 5 вариант.
	//задача 1.
	var a = -2.5
	var b = 3.4

	prnt(lab4.TaskA(a, b, 3.5, 6.5, 0.6))

	//задача 2.
	var xList = []float64{2.89, 3.54, 5.21, 6.28, 3.48}

	prnt(lab4.TaskB(xList, a, b), xList)

	//лабараторная работа №5 - произвольный вариант.
	dog, err := lab5.NewDog(20, "Паркурчик", "None")
	checkForError(err)
	anotherDog, err := lab5.NewDog(10, "Бандит", "None")
	checkForError(err)

	err = dog.SetAge(3)
	checkForError(err)

	anotherDog.SetName("Конь")
	dog.SetName("Прикол")

	fmt.Printf("dog's age is %d\n", dog.GetAge())
	fmt.Printf("dog's name is %s\n", dog.GetName())
	dog.PetThedog()

	fmt.Printf("another dog's age is %d\n", anotherDog.GetAge())
	fmt.Printf("it's breed is %s\n", anotherDog.GetBreed())
	anotherDog.PetThedog()
}
