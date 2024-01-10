package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido! Tente Novamente.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	var nome string = "Wermenson"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Bem vindo ao nosso sistema de monitoramento!")
	fmt.Println("Este programa está na versão", versao)

}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)

	return comandoLido
}

func IniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site com URL inexistente
	site := "https://httpbin.org/status/404" // ou 200
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		return
	}

	fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)

}
