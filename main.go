// Гипотеза Коллатца
// Найдите число шагов, за которые получится единица, используя следующий процесс:
// берём любое натуральное число n больше единицы.
// Если оно чётное, то делим его на 2,
// а если нечётное, то умножаем на 3 и прибавляем 1

// Программа выводит количество шагов по гипотезе Коллаца для всех натуральных чисел до N включительно
// N вводится пользователем
package main

import (
	"errors"
	"fmt"
)

// Создаем map где будут хранится все найденные значения
var numbers = make(map[int]int)

// Проверка того что N > 1
func nNisCorrect(N *int) (bool, error) {
	if *N <= 1 {
		return false, errors.New("N must be type int && N > 1")
	}
	return true, nil
}

// Сколько шагов требуется для того чтобы получить из N единицу по гипотезе Коллатца
func FindSteps(N int) (cnt int) {
	val, ok := numbers[N]
	if ok { // если N уже считалось, берём количество шагов из map
		return val
	} else { // если такое N ещё не встречалось то считаем следующий шаг рекурсии
		if N%2 == 0 {
			numbers[N] = FindSteps(N/2) + 1
			return numbers[N]
		} else {
			numbers[N] = FindSteps(N*3+1) + 1
			return numbers[N]
		}
	}
}

func main() {

	// Пользователь вводит натуральное число
	var N int
	var err error

	// Проверка того что введённо правильное N
	for {
		fmt.Print("Input positive N: ")
		_, err = fmt.Scanln(&N)
		fmt.Println()

		// Проверка правильности считанного значения
		if err != nil {
			fmt.Println("This value is incorrect, please try another")
			continue
		} else {
			// Проверка того что введённое число больше 1
			_, err = nNisCorrect(&N)
			if err == nil { // если после 2 проверок ошибок не обнаружено, то выходим из цикла ввода значения
				break
			} else {
				fmt.Println("This value is incorrect, please try another")
			}
		}
	}

	numbers[2] = 1 // При N = 2  нам необходим 1 шаг чтобы  дойти до 1
	fmt.Println("  N  | Value")
	for i := 2; i <= N; i++ {
		j := FindSteps(i)
		fmt.Printf("%5d | %d", i, j)
		fmt.Println()
	}
}
