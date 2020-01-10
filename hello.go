package main

import (
	"fmt"
	"math"
)

//Exercise Of Book - Intro Go Language - Caleb Doxsey - Novatec 2016

//Reader Jorge Dias - jorgediascpd@gmail.com - Ok

//package main

//import ("os", "fmt")

//Like a comment
//func main() {

// 	fmt.Println("Chapter1 - Start")
// 	fmt.Println("Hi Jorge!")

// 	fmt.Println("Chapter2 - Types")
// 	fmt.Println("32.132 X 42.452 = ", 32.132*42.452)

// 	var resultado = ((true && false) || (false && true) || !(false && false))
// 	fmt.Println("O resultado da expressão é ", resultado)

// 	var name = "Jorge Dias"
// 	fmt.Println(len(name))
// 	fmt.Println(name[0], " ", name[1], " ", name[2], " ", name[3], " ", name[4])

// 	fmt.Println("Chapter3 - Variable")
// 	// Other Method
// 	// var x string
// 	// x = "Hi "
// 	var x string = "Hi "
// 	fmt.Println("Your message is", x, name)

// 	var y string
// 	y = "First"
// 	fmt.Println(y)
// 	y = y + " Second"
// 	fmt.Println(y)
// 	y += " third"
// 	fmt.Println(y)

// 	var ya string
// 	ya = "third"

// 	fmt.Println("Variable 'y' is same variable 'ya'? ", y == ya)
// 	fmt.Println("")
// 	fmt.Println("Variable 'y' is same variable 'y'?", y == y)

// 	xy := "Use := para declarar variavel concisa"
// 	ex := 500
// 	ax := "Ex:   xy := "
// 	ms := "Dessa forma ele infere com base no conteudo"

// 	fmt.Println(xy, " ", " ", ax, ex, " ", ms)
// }

// //Test local variables
// const a string = "Variable local"

// func f() {
// 	//const a string = "Local variable"
// 	fmt.Println("Definido a variavel local.")
// }
// func g() {
// 	fmt.Println(a)
// }

// var (
// 	a = 5
// 	b = 10
// 	c = 15
// )

//Programa que lê um input e retorna o dobro

// fmt.Println("Enter the number: ")
// var input float64
// fmt.Scanf("%f", &input)
// output := input * 2
// fmt.Println("O dobro de ", input, " é ", output)

//Programa que convert fahrenheit em Celsius (C=(F-32)*5/9)
// var fahrenheit float64
// fmt.Println("Digite a temperatura em fahrenheit: ")
// fmt.Scanf("%f", &fahrenheit)
// output := (fahrenheit - 32) * 5 / 9
// fmt.Println("O resultado do Fahrenheit ", fahrenheit, " é ", output)

//Programa que converte pés para metros 1pé = 0,3048m
// var pes float64
// fmt.Println("Digite a altura em pés: ")
// fmt.Scanf("%f", &pes)
// fmt.Println("O resultado de ", pes, "pes é ", pes*0.3048, " em metros.")
// 	fmt.Println(`1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10`)

//For usando
//fmt.Println("Chapter4 - For If Switch")
// i := 1
// for i <= 10 {
// 	fmt.Println(i)
// 	i = i + 1
// }

// for i := 1; i <= 10; i++ {
// 	// fmt.Println("Contador em: ", i)
// 	if i%4 == 0 {
// 		fmt.Println(i, "é divisivel por 4!")
// 	} else if i%3 == 0 {
// 		fmt.Println(i, "é divisivel por 3!")
// 	} else if i%2 == 0 {
// 		fmt.Println(i, "is even!")
// 	} else {
// 		fmt.Println(i, "is odd")
// 	}
// }

//Switch
// for i := 0; i <= 8; i++ {
// 	switch i {
// 	case 0:
// 		fmt.Println("Zero")
// 	case 1:
// 		fmt.Println("One")
// 	case 2:
// 		fmt.Println("Two")
// 	case 3:
// 		fmt.Println("Tree")
// 	case 4:
// 		fmt.Println("Four")
// 	case 5:
// 		fmt.Println("Five")
// 	case 6:
// 		fmt.Println("Six")
// 	case 7:
// 		fmt.Println("Seven")
// 	default:
// 		fmt.Println("Unknown Number")
// 	}
// }
// 	f := 0
// 	b := 0
// 	fb := 0
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 {
// 			if i%5 == 0 {
// 				fmt.Println(i, "FizzBuzz")
// 				fb++
// 			}
// 			fmt.Println(i, "Fizz")
// 			f++
// 		} else if i%5 == 0 {
// 			fmt.Println(i, "  Buzz")
// 			b++
// 		} else {
// 			fmt.Println(i, "----")
// 		}
// 	}
// 	fmt.Println("Fizz=", f, " Buzz=", b, "FizzBuzz=", fb, "Total:", fb+f+b)

// Chapter 5 - Arrays slice e map
//fmt.Println("Chapter5 - Arrays slice e map")
// var x [5]int
// x[4] = 100
// fmt.Println(x)

// var x [5]float64
// x[0] = 27
// x[1] = 32
// x[2] = 43
// x[3] = 53
// x[4] = 20
// var total float64 = 0
// for i := 0; i < 5; i++ {
// 	fmt.Println("Total :", total, " + ", x[i])
// 	total += x[i]
// }
// fmt.Println("A média de tudo é ", total/float64(len(x))) //Aqui convertemos int em float64, pq len só usa int.

// x := [5]float64{20, 30, 40, 50, 60}
// var total float64 = 0

// for _, value := range x {
// 	total += value
// }
// fmt.Println(total / float64(len(x)))

// var total float64 = 0
// x := [5]float64{
// 	21,
// 	32,
// 	73,
// 	// 	44, // Sim, podemos comentar nessa estrutura um item do array
// 	85,
// }
// for _, value := range x { //Aqui não usamos contador, usamos a propria variavel com _, value e range
// 	total += value
// }
// fmt.Println("A média é", total/float64(len(x)))

// Usar fatia, não colocamos a informação do array, isso poderá ser alterado.
// var x []float64

// x := make([]float64, 5)
// x := make([]float64, 5, 10)

// [low : high]    [0:5]  [:5] [3:] [de:ate]
// arr := [5]float64{1, 2, 3, 4, 5}
// x := arr[1:3]
// fmt.Println(x)
// y := arr[0:]
// fmt.Println(y)
// z := arr[:2]
// fmt.Println(z)

//Escreva um programa que encontre o menor numero dentro de uma array
// x := [9]float64{
// 	52,
// 	60,
// 	80,
// 	90,
// 	60,
// 	10,
// 	25,
// 	29,
// 	27,
// }
// var menor float64 = 1000
// for _, value := range x {
// 	fmt.Println("Analisando ", value)
// 	if value < menor {
// 		menor = value
// 		fmt.Println(menor, "é menor!")
// 	}
// }
// fmt.Println("O vencedor de menor número foi ", menor)

//Append - Juntando arrays
// slice1 := []int{4, 5, 6}
// slice2 := append(slice1, 7, 8, 9)
// fmt.Println(slice1, "----", slice2)

// copy - usado para copiar dst -> src
// slice1 := []int{1, 2, 3}
// slice2 := make([]int, 2)
// copy(slice2, slice1) //(dst, src)
// fmt.Println(slice1, "----", slice2)

// Trabalhando com mapas, na verdade mapas são arrays não sequenciais que podem ser com strings,

// x := make(map[string]int) // x é um para de string para ints - Use o make para inicializar o mapa.
// x["key"] = 10
// fmt.Println(x["key"])
// delete(x, "key") //Use o delete para apagar
// fmt.Println(x)

//Exemplo de mapa
// elements := make(map[string]string)
// elements["H"] = "Hydrogen"
// elements["He"] = "Heliun"
// elements["Li"] = "Lithiun"
// elements["Be"] = "Beryllium"
// elements["B"] = "Boron"
// elements["C"] = "Carbon"
// elements["N"] = "Nitrogen"
// elements["O"] = "Oxygen"
// elements["F"] = "Fluorine"
// elements["Ne"] = "Neon"
// fmt.Println(elements["Li"])

// fmt.Println(elements["Un"]) //Aqui não existe e não retorna nada.

// name, ok := elements["Un"] // Para retornar quando não existe nada usamos essa função.
// fmt.Println(name, ok)

// // Usamos essa função com frequência para retornar valores existentes.
// if name, ok := elements["N"]; ok {
// 	fmt.Println(name, ok)
// }

//Outra forma de usar mapa
// elements := map[string]map[string]string{
// 	"H": map[string]string{
// 		"name":  "Hydrogen",
// 		"state": "gas",
// 	},
// 	"He": map[string]string{
// 		"name":  "Helium",
// 		"state": "gas",
// 	},
// 	"Li": map[string]string{
// 		"name":  "Lithium",
// 		"state": "solid",
// 	},
// }
// if el, ok := elements["Li"]; ok {
// 	fmt.Println(el["name"], el["state"])
// }
// fmt.Println(elements)

//}

/////////////////////////////////////////////////////////////////////////////////////////////
//   Até aqui é preciso habilitar a função main no inicio para funcionar as funções acima.///
/////////////////////////////////////////////////////////////////////////////////////////////

// Chapter 6 - Functions

//sample function
// func average(xs []float64) float64 {
// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	return total / float64(len(xs))
// }
// func main() {
// 	fmt.Println("Chapter6 - Functions")
// 	arrayMedia := []float64{98, 93, 77, 82, 83}
// 	fmt.Println(average(arrayMedia))
// }

//Variática - recebe vários números usando ...

// func add(args ...int) int { //Aqui está as ... antes de int
// 	total := 0
// 	for _, v := range args {
// 		total += v
// 	}
// 	return total
// }

// func main() {
// 	fmt.Println(add(1, 2, 3, 50))
// }

// Closure - Função dentro de outra função

// func main() {
// 	add := func(x, y int) int { // Aqui foi criada uma função dentro da função principal armazenada dentro da variável add
// 		return x + y
// 	}
// 	fmt.Println(add(1, 2)) // Aqui foi chamada a função enclausurada com o nome da variável.
// }

//Outra Clausure

// func main() {
// 	x := 0
// 	increment := func() int { //Função camorrenta escondida aqui dentro da variável.
// 		x++
// 		return x
// 	}
// 	fmt.Println(increment()) // cada chamada mesmo sem parâmetro retorna o valor de x+1
// 	fmt.Println(increment())
// }

//Recursão - Função que chama ela mesma...

// func factorial(x uint) uint {
// 	if x == 0 {
// 		return 1
// 	}
// 	return x * factorial(x-1)
// }
// func main() {
// 	fmt.Println(factorial(5))
// }

//Defer, panic ou recovery

//Defer coloca para ser executado por ultimo antes de sair da função.
//Usamos muito o defer para colocar a função de fechar logo após a de abrir.
// // f, _ := os.Open(filename)
// // defer f.Close()

// func first() {
// 	fmt.Println("First")
// }
// func second() {
// 	fmt.Println("Second")
// }
// func main() {
// 	defer first() // Aqui você fala para ele executar por ultimo usando o defer.
// 	second()
// }

// Saída
// Second
// First
// Process exiting with code: 0

//Panic e recover
// Para usar o panic e recover é preciso acionar antes o defer, para ele chamar depois que ocorrer o panic

// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}()
// 	panic("PANIC")
// }

//Ponteiros - use para apontar onde está armazenada a informação.
// func zero(xPtr *int) {
// 	*xPtr = 0
// }
// func main() {
// 	x := 5
// 	zero(&x)
// 	fmt.Println(x) //x é 0
// }

//Execcícios, calcule um programa que receba um inteiro e calcule sua metade e retorne verdadeiro para par e falso para impar.

// func main() {
// 	fmt.Println(parOuImpar(12))
// }

// func parOuImpar(n int) (int, bool, int) { // Aqui colocamos 2 retornos um int e um bool
// 	metade := n / 2
// 	x := false // Se não declarar aqui, de denro do if para fora não vai...
// 	y := 50
// 	if metade%2 == 0 {
// 		x = true
// 		y = 1
// 	} else {
// 		x = false
// 		y = 0
// 	}
// 	return y, x, n
// }

// Receba uma lista de numeros e retorne o maior usando variatico ...
// Aqui usei o for já com o limite do array.

// func variatico(n ...int) int {
// 	maior := 0
// 	for _, value := range n {
// 		if value > maior {
// 			maior = value
// 		}
// 	}
// 	return maior
// }
// func main() {
// 	fmt.Println(variatico(1, 2, 3, 1, 10, 500, 23, 200))
// }

// Gerador de números impares.
// func makeOddGenerator() func() uint {
// 	i := uint(1)
// 	return func() (ret uint) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }
// func main() {
// 	nextEven := makeOddGenerator()
// 	fmt.Println(nextEven())
// 	fmt.Println(nextEven())
// 	fmt.Println(nextEven())
// 	fmt.Println(nextEven())
// }

// //Sequencia de Fibonacci
// func fib(m, n int) int { // Aqui recebe a anterior e a anteanterior
// 	seq := n + m
// 	return seq
// }
// func main() { //Aqui tem um for para enviar a anterior e a anteanterior
// 	atual := 1
// 	anterior := 0
// 	anteAnterior := 0
// 	for i := 0; i < 20; i++ {
// 		fmt.Print(atual, "-")
// 		anteAnterior = anterior
// 		anterior = atual
// 		atual = fib(anterior, anteAnterior)
// 	}
// }

/////////////////////////////////////////////////////////////////////
//// Chapter 7 - Estruturas e Interfaces                         ////
/////////////////////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func distance(x1, y1, x2, y2 float64) float64 {
// 	a := x2 - x1
// 	b := y2 - y1
// 	return math.Sqrt(a*a + b*b)
// }
// func rectangleArea(x1, y1, x2, y2 float64) float64 {
// 	l := distance(x1, y1, x2, y2)
// 	w := distance(x1, y1, x2, y2)
// 	return l * w
// }
// func circleArea(x, y, r float64) float64 {
// 	return math.Pi * r * r
// }
// func main() {
// 	var rx1, ry1 float64 = 0, 0
// 	var rx2, ry2 float64 = 10, 10
// 	var cx, cy, cr float64 = 0, 0, 5

// 	fmt.Println("Chapter7 - Struct and interface") // How to create structure.
// 	fmt.Println("Rectangle Área:", rectangleArea(rx1, ry1, rx2, ry2))
// 	fmt.Println("Circle Área:", circleArea(cx, cy, cr))
// }

// O programa acima está calculando a área de um retangulo e circulo, dessa forma é mais complexo de ver e dar manutenção,
// Ai que entra as estruturas e interfaces, para facilitar essa abordagem.

// // Representação
// type Circle struct {
// 	x, y, r float64
// }
// // Inicialização
// //var c Circle
// //ou
// //c := new(Circle)

// // A forma mais usada não é usando new, mas assim...
// //c := Circle{x:0, y:0, r: 5}
// // Se souber a ordem dos campos não é preciso especificar os nomes...
// c := Circle{0,0,5}
// // Se precisar de um ponteiro use &
// //c := &Circle{0,0,5}

// // Para acessarmos os campos da estrutura usamos o .
// c.x = 10
// c.y = 5

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
