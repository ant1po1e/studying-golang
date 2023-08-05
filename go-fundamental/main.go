package main

import (
	"fmt"
	// "errors"
)
	


func main() {
	//////////////////////////////////////// variable
	// var name string
	// name = "Antipole"
	// age := 8
	// fmt.Println(name, age)


	//////////////////////////////////////// if-else
	// score := 90
	// var grade string

	// if (score <= 70) {
	// 	grade = "C"
	// } else if (score <= 80) {
	// 	grade = "B"
	// } else if (score <= 90) {
	// 	grade = "A"
	// }

	// fmt.Println(grade)


	//////////////////////////////////////// switch conditional
	// number := 2

	// switch number {
	// case 1: fmt.Println("One")
	// case 2: fmt.Println("Two")
	// case 3: fmt.Println("Three")
	// }


	//////////////////////////////////////// for
	// for i := 1; i <= 20; i++ {
	// 	fmt.Println("waduh :", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("waduh", i)
	// 	i++
	// }

	//////////

	// title := "waduh icikiwir kumalala"
	// for i, letter := range title {
	// 	if (i % 2 == 0) {
	// 		fmt.Println("i :", i, "letter", string(letter))
	// 	}
	// }
	
	// for i, letter := range title {
	// 	letterStr := string(letter)
	// 	if (letterStr == "a" || letterStr == "i" || letterStr == "u" || letterStr == "e" || letterStr == "o" ) {
	// 		fmt.Println("i :", i, "letter", string(letter))
	// 	}

	//////////
		
	// 	switch letterStr {
	// 	case "a", "i", "u", "e", "o":
	// 		fmt.Println("i :", i, "letter", string(letter))

	// 	}
	// }


	//////////////////////////////////////// array
	// var languages [5] string
	// languages[0] = "JS"
	// languages[1] = "HTML"
	// languages[2] = "CSS"
	// languages[3] = "TS"
	// languages[4] = "GO"

	//////////

	// languages := [5] string{"JS", "HTML", "CSS", "TS", "GO"}
	// languages := [5] string{
	// 	"JS", 
	// 	"HTML", 
	// 	"CSS", 
	// 	"TS", 
	// 	"GO",
	// }
	// fmt.Println(languages)

	// for i, lang := range languages {
	// 	fmt.Println("index :", i, "language: ", lang)
	// }

	
	//////////////////////////////////////// slice
	// var gamingConsoles [] string
	
	// gamingConsoles = append(gamingConsoles, "PlayStation")
	// gamingConsoles = append(gamingConsoles, "Nitendo")
	// gamingConsoles = append(gamingConsoles, "XBOX")

	// gamingConsoles := [] string {
	// 	"PlayStation",
	// 	"Nitendo",
	// 	"XBOX",
	// }

	// for _, console := range gamingConsoles {
	// 	fmt.Println("Consoles: ", console)
	// }

	// fmt.Println(gamingConsoles)


	//////////////////////////////////////// map
	// var myMap map[string] int
	// myMap = map[string] int{}

	// myMap["Sucipto"] = 1
	// myMap["Asep"] = 4
	// myMap["Sigit"] = 3
	// myMap["Fulan"] = 2

	// fmt.Println(myMap)

	//////////
	// myMap := map[string] string{
	// 	"Chino": "SD",
	// 	"Azusa": "SMP",
	// 	"Rushia": "SMK",
	// }

	// for name, grade := range myMap {
	// 	fmt.Println("Name: ", name, "Grade: ", grade)
	// }
	
	// fmt.Println("====================")
	
	// delete(myMap, "Rushia")
	
	// for name, grade := range myMap {
	// 	fmt.Println("Name: ", name, "Grade: ", grade)
	// }
	
	// fmt.Println("====================")

	// value, isAvaible := myMap["Chino"]
	// fmt.Println(value)
	// fmt.Println(isAvaible)


	//////////////////////////////////////// slice of map
	// students := [] map[string] string {
	// 	{"name": "Sucipto", "score": "A"},
	// 	{"name": "Asep", "score": "B"},
	// 	{"name": "Budi", "score": "C"},
	// }

	// fmt.Println(students)
	
	// for _, student := range students {
	// 	fmt.Println(student["name"], " : ", student["score"])
	// }


	////////////////////////////////////////
	// scores := [10] int {80, 70, 75, 90, 95, 78, 88, 96, 100, 72}

	// var total int

	// for _, score := range scores {
	// 	total = total + score
	// }

	// length := len(scores)
	// avg := float64(total) / float64(length)

	// fmt.Println("Average: ", avg)

	////////////
	// var goodScores [] int

	// for _, score := range scores {
	// 	if (score >= 90) {
	// 		goodScores = append(goodScores, score)
	// 	}
	// }

	// fmt.Println(goodScores)


	//////////////////////////////////////// basic function
    // Print("This is a word")

	////////// return
	// result := Sum(69, 10)
	// fmt.Println(result)

	////////// return 2
	// luas, keliling := Calculate(11, 10)
	// fmt.Println(luas)
	// fmt.Println(keliling)

	//////////
	// scores := [] int {10, 5, 8, 9, 7}
	// total := Sum(scores)
	// fmt.Println(total)

	////////// array in function
	// result, err := Calculate(21, 16, "a")
	// if (err != nil) {
	// 	fmt.Println("Something went wrong")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result)


	//////////////////////////////////////// structure
	// user := User{
	// 	ID: 1, 
	// 	Name: "Stanford", 
	// 	Username: "Antipole", 
	// 	Email: "apolantipole@gmail.com", 
	// 	IsActive: true,
	// }

	/////
	// user := User{1, "Tystiana", "Azusa", "azusa@gmail.com", true}
	
	/////		
	// user.ID = 1
	// user.Name = "Stanford"
	// user.Username = "Antipole"
	// user.Email = "apolantipole@gmail.com"
	// user.IsActive = true
	
	////////// struct as parameter
	// user := User{1, "Tystiana", "Azusa", "azusa@gmail.com", true}
	// user2 := User{2, "Stanford", "Antipole", "antipole@gmail.com", true}
	
	// displayUser := DisplayUser(user)
	// displayUser2 := DisplayUser(user2)
	
	// fmt.Println(displayUser)
	// fmt.Println(displayUser2)
	
	////////// embedded struct
	// user := User{1, "Tystiana", "Azusa", "azusa@gmail.com", true}
	// user2 := User{2, "Stanford", "Antipole", "antipole@gmail.com", true}

	// member := []User{user, user2}
	// group := Group{"Jawir", user, member, true}
	
	
	// DisplayGroup(group)
	
	
	//////////////////////////////////////// method
	// user := User{1, "Tystiana", "Azusa", "azusa@gmail.com", true}
	// user2 := User{1, "Stanford", "Antipole", "antipole@gmail.com", true}
	
	// member := []User{user, user2}
	// group := Group{"Jawir", user, member, true}
	// // fmt.Println(user.Display())
	// // fmt.Println(user2.Display())

	/////
	// group.DisplayGroup()


	//////////////////////////////////////// pointer
	// var numberA int = 5
	// var numberB *int = &numberA
	
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	
	// numberA = 69
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	
	//////////
	// number := 5
	// fmt.Println("Default value: ", number)
	
	// Change(&number, 69)
	// fmt.Println("New value: ", number)

	//////////////////// pointer struct as parameter
	// student := Student{1, "Antipole"}
	
	// fmt.Println(student.Name)
	
	// Graduate(&student)
	
	// fmt.Println(student.Name)
	
	
	//////////////////// method with pointer receiver
	// student := Student{1, "Antipole"}
	
	// fmt.Println(student.Name)
	
	// student.Graduate()
	
	// fmt.Println(student.Name)
	
	////////// 
	// gamer := Gamer{Name: "Blue Archive"}
	
	// gamer.AddGame("Genshin Impact")
	// gamer.AddGame("Limbus Company")
	
	// for _, game := range gamer.Games {
		// 	fmt.Println(game)
		// }
	
		
	//////////////////////////////////////// interface
	rectangle := Rectangle{10, 7}
	wide := HowWide(rectangle)
	round := HowRound(rectangle)
	fmt.Println("Rectangle's Wide : ", wide)
	fmt.Println("Rectangle's Round : ", round)
	
	square := Square{6}
	wide = HowWide(square)
	round = HowRound(square)
	fmt.Println("Square's Wide : ", wide)
	fmt.Println("Square's Round : ", round)
	
	circle := Circle{14}
	wide = HowWide(circle)
	round = HowRound(circle)
	fmt.Println("Circle's Wide : ", wide)
	fmt.Println("Circle's Round : ", round)


}


//////////////////////////////////////// basic function
// func Print(word string) {
// 	fmt.Println(word)
// }

////////// return
// func Sum(numOne int, numTwo int) int {
// 	return numOne + numTwo
// }

////////// return 2
// func Calculate(panjang int, lebar int) (int, int) {
// 	luas := panjang * lebar
// 	keliling := 2 * (panjang + lebar)
// 	return luas, keliling
// }

// func Calculate(panjang int, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)

// 	return
// }

//////////
// func Sum(numbers [] int) (total int) {
// 	for _, number := range numbers {
// 		total = total + number
// 	}

// 	return 
// }

////////// array in function 
// func Calculate(numOne, numTwo int, operation string) (result int, resultError error) {
// 	switch operation {
// 	case "+":
// 		result = numOne + numTwo
// 	case "-":
// 		result = numOne - numTwo
// 	case "*":
// 		result = numOne * numTwo
// 	case "/":
// 		result = numOne / numTwo
// 	default:
// 		resultError = errors.New("Unknown Operation")
// 	}

// 	return
// }


//////////////////////////////////////// structure
// type User struct {
// 	ID int
// 	Name string
// 	Username string
// 	Email string
// 	IsActive bool
// }

// type Group struct {
// 	Name string
// 	Admin User
// 	Member []User
// 	IsAvaible bool
// }

////////// struct as parameter
// func DisplayUser(user User) string {
// 	return fmt.Sprintf("Name: %s | Username: %s | Email: %s", user.Name, user.Username, user.Email)
// }

////////// embedded struct
// func DisplayGroup (group Group) {
//	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count: %d", len(group.Member))
// 	fmt.Println("")
// 	fmt.Println("Member Name: ")

// 	for _, user := range group.Member {
	// 		fmt.Println(user.Username)
	// 	}
	// }
	
//////////////////////////////////////// method
// func (user User) Display() string {
// 	return fmt.Sprintf("Name: %s | Username: %s | Email: %s", user.Name, user.Username, user.Email)
// }
	
/////
// func (group Group) DisplayGroup() {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count: %d", len(group.Member))
// 	fmt.Println("")
// 	fmt.Println("Member Name: ")
		
// 	for _, user := range group.Member {
// 		fmt.Println(user.Username)
// 	}
// }


//////////////////////////////////////// pointer
// func Change(old *int, new int) {
// 	*old = new
// }

//////////////////// pointer struct as parameter
// type Student struct {
	// 	ID int
	// 	Name string
	// }
	
	// func Graduate(student *Student) {
		// 	student.Name = student.Name + " S.T"
		// }
		
		
//////////////////// method with pointer receiver
// type Student struct {
// 	ID int
// 	Name string
// }

// func (student *Student) Graduate() {
// 	student.Name = student.Name + " S.T"
// }

//////////
// type Gamer struct {
// 	Name string
// 	Games [] string
// }

// func (gamer *Gamer) AddGame(game string) {
// 	gamer.Games = append(gamer.Games, game)
// }


//////////////////////////////////////// interface
type Shape interface {
	CalculateWide() int
	CalculateRound() int
}

type Rectangle struct {
	Wide int
	Length int
}

type Square struct {
	Side int
}

type Circle struct {
	Radius int
}

func (rectangle Rectangle) CalculateWide() int {
	return rectangle.Wide * rectangle.Length
} 

func (square Square) CalculateWide() int {
	return square.Side * square.Side
} 

func (circle Circle) CalculateWide() int {
	return 22 * (circle.Radius * circle.Radius) / 7
}

func (rectangle Rectangle) CalculateRound() int {
	return 2 * (rectangle.Wide + rectangle.Length)
} 

func (square Square) CalculateRound() int {
	return 2 * (square.Side + square.Side)
} 

func (circle Circle) CalculateRound() int {
	return 22 * (2 * circle.Radius) / 7
}

func HowWide (shape Shape) int {
	return shape.CalculateWide()
}

func HowRound (shape Shape) int {
	return shape.CalculateRound()
}


