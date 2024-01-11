package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

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
	fmt.Println("----------------------------------------------------------------------------")

	sites := []string{"https://pokedex-three-smoky.vercel.app/", "https://space-app-dun.vercel.app/", "https://ola-mundo-react-swart.vercel.app/", "https://aleatorio404.vercel.app/"}

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("----------------------------------------------------------------------------")
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		return
	}

	fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
}
