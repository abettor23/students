package main

import (
	"fmt"
)

func main() {
	fmt.Println("Помощник старосты.")

	// применил еще не изученный оператор return, для избежания ошибок. разобрался в его применении
	fmt.Println("Введите количество студентов:")
	var studentsAmount int
	fmt.Scan(&studentsAmount)
	if studentsAmount <= 0 {
		fmt.Println("Введено неверное количество студентов. Попробуйте снова.")
		return
	}

	fmt.Println("Введите количество групп:")
	var groupAmount int
	fmt.Scan(&groupAmount)
	if groupAmount <= 0 {
		fmt.Println("Введено неверное количество групп. Попробуйте снова.")
		return
	}

	fmt.Println("Введите порядковый номер студента:")
	var studentID int
	fmt.Scan(&studentID)
	if studentID <= 0 || studentID > studentsAmount {
		fmt.Println("Введен неверный порядковый номер студента. Попробуйте снова.")
		return
	}

	// определяем номер группы, в которую попадет студент
	groupID := studentID % groupAmount
	if groupID == 0 {
		groupID = groupAmount
	}

	fmt.Println("Студент с порядковым номером", studentID, "попадет в группу номер", groupID)
}
