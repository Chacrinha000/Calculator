package main

import "fmt"

func main() {
	fmt.Println("Digite sua operação")
	var primeironumero int
	var segundonumero int
	var operador string
	fmt.Scanf("%d %s %d", &primeironumero, &operador, &segundonumero)
	if operador == "+" {
		var soma int
		soma = primeironumero + segundonumero
		fmt.Printf("O resultado e: %d", soma)
		return
	}
	if operador == "-" {
		var subtracao int
		subtracao = primeironumero - segundonumero
		fmt.Printf("O resultado e: %d", subtracao)
		return
	}
	if operador == "*" {
		var multiplicacao int
		multiplicacao = primeironumero * segundonumero
		fmt.Printf("O resultado e: %d", multiplicacao)
		return
	}
	if operador == "/" {
		var divisao int
		divisao = primeironumero / segundonumero
		fmt.Printf("O resultado e %d", divisao)
		return
	}

}
