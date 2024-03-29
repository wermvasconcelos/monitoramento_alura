package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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

	//sites := []string{"https://pokedex-three-smoky.vercel.app/", "https://space-app-dun.vercel.app/", "https://ola-mundo-react-swart.vercel.app/", "https://aleatorio404.vercel.app/"}
	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("----------------------------------------------------------------------------")
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao monitorar o site ", site, "\n Erro:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
		return
	}

	fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	registraLog(site, false)
}

func lerSitesDoArquivo() []string {
	sites := []string{}

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		if err == io.EOF {
			break
		}

		sites = append(sites, strings.TrimSpace(linha))
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log.")
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log")
		return
	}

	fmt.Println("Imprimindo logs...")

	fmt.Println(string(arquivo))
}
