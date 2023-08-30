package cnh_requisitos

import "fmt"

var name string
var age int64
var condition string
var infrigiment string
var command int

func Introduction() {

	fmt.Println("Olá, seja bem vindo ao canal da CNH!!")
	fmt.Println("Digite seu nome: ")
	fmt.Scanln(&name)
	fmt.Println("Digite a sua idade: ")
	fmt.Scanln(&age)

	fmt.Println(name, age)

	if age >= 18 {
		fmt.Println("Ok ", name, " agora vamos para o próximo passo!")
		category()
	} else {
		fmt.Println("Você não tem idade suficiente para tirar a CNH!!!")
	}

}

func category() {
	fmt.Println("Digite o número da categoria desejada: ")
	fmt.Println("1: A - Moto")
	fmt.Println("2: B - Carro")
	fmt.Println("3: AB - Carro e Moto")
	fmt.Println("4: C -  Carretos de 3.500 a 6.000kg")
	fmt.Println("5: D - Ônibus")
	fmt.Println("6: E - Caminhão +6.000Kg")
	fmt.Scanln(&command)

	switch command {
	case 1:
		canGet()
	case 2:
		canGet()
	case 3:
		canGet()
	case 4:
		testPossibility()
	case 5:
		testPossibility()
	case 6:
		testPossibility2()
	default:
		fmt.Println("Comando inválido, digite um número de 1 à 6!")
		category()
	}
}

func canGet() {
	fmt.Println("Ok, realize o exame e agende suas aulas!")

}

func cantGet() {
	fmt.Println("Você não está apto para essa categoria!")
}

func qulificationDuration1() {
	fmt.Println("Você está habilitado na categoria B a mais de um ano? (s/n): ")
	fmt.Scanln(&condition)
}

func qulificationDuration2() {
	fmt.Println("Você está habilitado na categoria C ou D a mais de um ano? (s/n): ")
	fmt.Scanln(&condition)
}

func checkInfrigiment() {
	fmt.Println("Você cometeu alguma infração grave nos ultimos 12 meses? (s/n): ")
	fmt.Scanln(&infrigiment)
}

func testPossibility() {
	if age >= 21 {
		qulificationDuration1()
	}else {
		cantGet()
		return
	}
	if condition == "s" {
		checkInfrigiment()
	}else {
		cantGet()
		return
	}
	if infrigiment == "n" {
		canGet()
	}else {
		cantGet()
	}
}

func testPossibility2() {
	if age >= 21 {
		qulificationDuration2()
	}else {
		cantGet()
		return
	}
	if condition == "s" {
		checkInfrigiment()
	}else {
		cantGet()
		return
	}
	if infrigiment == "n" {
		canGet()
	}else {
		cantGet()
	}
}