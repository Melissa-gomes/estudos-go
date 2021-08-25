package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

//	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
// "%d" -----> diz que esta esperando receber um inteiro que vai modificar a variavel command
// Scan ----> Usado para captura de input
/* if command == 1 {
fmt.Println("Iniciando Monitoramento...")
} else if command == 2 {
	fmt.Println("Aqui estão os logs")
	} else if command == 0 {
		fmt.Println("Fechando o programa. Bye")
		} else {
			fmt.Println("Não reconheço este comando")
			}	*/

func main() {

	displayIntro()

	for {
		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Aqui estão os logs")
		case 0:
			fmt.Println("Fechando o programa. Bye")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando")
			os.Exit(-1)
		}
	}

}

func displayIntro() {
	var nome string = "Melissa"
	sobrenome := "Gomes"
	var version float32 = 1.2

	fmt.Println("Olá, meu nome é", nome, sobrenome)
	fmt.Println("Este programa está na versão", version)
}

func displayMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido pelo usuario foi", command)
	fmt.Println("")
	return command
}

func initMonitoring() {
	fmt.Println("Iniciando Monitoramento...")
	sites := readData()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			testingSites(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testingSites(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status code:", res.StatusCode)
	}
}

func readData() []string {

	var sites []string
	data, err := os.Open("sites.txt")

	// data, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(data)

	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	data.Close()

	return sites
}
