package main

import "fmt"
import "os"
import "net/http"

func main() {
	//var nome string = "Gustavo"
	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não reconheço este comando")
		os.Exit(-1)
	}

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Não reconheço este comando")
	// }
}

func exibeIntroducao() {
	nome := "Gustavo"
	versao := 1.0
	fmt.Println("Bem vindo", nome)
	fmt.Println("Programa na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	response, _ := http.Get(site)
	fmt.Println(response.StatusCode)
}
