package main

import "fmt"

func main() {
	//var nome string = "Gustavo"
	nome := "Gustavo"
	versao := 1.0
	fmt.Println("Bem vindo", nome)
	fmt.Println("Programa na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não reconheço este comando")
	}

}
