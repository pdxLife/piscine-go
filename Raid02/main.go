package main

import (
	"os"
	"github.com/01-edu/z01"
)

const n int = 9

func main() { //основа программы
	info := "Error" //инфо по ошибке
	mssg := "" //месседж пустой
	args := os.Args
	//arguments := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	arguments := []string(args[1:]) //срез 1 аргумента со строки
	lenA := 0 //находим длину, со скольки строк состоит массив
	for index := range arguments {
		lenA = index + 1
	}
	field := make([][]rune, n) //создаем массив, в котором содержится массив рун
	for i := range field {
		field[i] = make([]rune, n) //проходя по каждой строке, создаем массив рун
	}

	if lenA == n { //проверка длины аргументов
		for index := range field {
			field[index] = []rune(arguments[index]) //если длина правильная, создаем поле аргументов (матрицу)
		}
		if FindSolution(field) == false { //если решения нет, то выводи ошибку 
			mssg = info 
		}
		
	} else {
		mssg = info //если переданные неравны 9, то ошибка
	}

	if mssg == "" { //если месседж пустой, то распечатай поле 
		PrintField(field)
	} else { //если не пустой, то распечатай 
		for _, l := range []rune(mssg) {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}


func FindSolution(field [][]rune) bool{ //функция решения
	x := -1 //столбец, это координаты позиций в которой самое мин. количество возможных вариантов
	y := -1 //строка
	min := 10 //приравнивание 10 потому что, минимум вариантов девять
	for row := 0; row < n; row++ { //цикл строк
		for column := 0; column < n; column++ { //цикл столбцев
			if field[row][column] == '.' { //если поле не заполнены
				counter := 0 //то каунтер обнуляем
				for v := '1'; v < rune('1'+n); v++ { //то в цикле это поле будет заполнять пустые позиции от 1 до 9
					field[row][column] = rune(v) //приравниваем по циклу
					if CheckPos(field, row, column) {  //проверяем по функции, если он проходит 
						counter++
						}
					field[row][column] = '.' //после прохождения проверки, обнуляем ячейку
				}
				if min > counter { //если варианты, меньше мин. значения
					min = counter //то минимум становится пред. мин значением
					x = row
					y = column
				}
			}
		}
	}
	if x == -1 {  //в случае все ячейки не пустые
		return true //то возраващаем, ячейки заполнены полностью
	}
	for v := '1'; v < rune('1'+n); v++ { 
			field[x][y] = rune(v)
			if CheckPos(field, x, y) { //если он проходит проверку
				if FindSolution(field) { //если функция выполняется (когда все ячейки заполнены)
					return true //возращаемся
				}
			}
		}
	field[x][y] = '.' //обнуляем ячейки, если судоку не решается
	return false //решения нет
}

func PrintField(field [][]rune) { //функция распечатки поля
	for i := range field {
		for _, l := range field[i] {
			z01.PrintRune(l) //значение ячейки
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n') 
	}
}


func CheckPos(field [][]rune, row, column int) bool { //проверка позиций
	value := field[row][column] 
	count := 0
	//Check in column
	for i := 0; i < n; i++ {
		if field[i][column] == value {
			count++
		}
	}
	//Check in row
	for j := 0; j < n; j++ {
		if field[row][j] == value {
			count++
		}
	}
	//square 1,2,3,4,5,6,7,8,9
	for i := (row / 3) * 3; i < (row/3+1)*3; i++ { 
		for j := (column / 3) * 3; j < (column/3+1)*3; j++ { 
			if field[i][j] == value {
				count++
			}
		}
	}
	if count == 3 {
		return true
	}
	return false
}