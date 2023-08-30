package cnh_requisitos

import "testing"

func TestIntroduction(t *testing.T) {
	name = "igor"
	age = 22
	command = 1
	Introduction()
}

func TestMenorIdade(t *testing.T) {
	age = 17
	Introduction()
}

func TestCase2(t *testing.T) {
	command = 2
	category()
}

func TestCase3(t *testing.T) {
	command = 3
	category()
}

func TestCase4(t *testing.T) {
	age = 21
	command = 4
	category()
}

func TestCase5(t *testing.T) {
	age = 21
	command = 5
	condition = "s"
	infrigiment = "n"
	category()
}

func TestCase6(t *testing.T) {
	age = 21
	command = 6
	condition = "s"
	infrigiment = "n"
	category()
}

func TestCase5CantGet(t *testing.T) {
	age = 20
	command = 5
	category()
}

func TestCase6CantGet(t *testing.T) {
	age = 20
	command = 6
	category()
}

func TestCantGetCondition(t *testing.T) {
	age = 21
	command = 6
	condition = "n"
	category()
}

func TestCantGetInfrigiment5(t *testing.T) {
	age = 21
	command = 5
	condition = "s"
	infrigiment = "s"
	category()
}

func TestCantGetInfrigiment6(t *testing.T) {
	age = 21
	command = 6
	condition = "s"
	infrigiment = "s"
	category()
}
