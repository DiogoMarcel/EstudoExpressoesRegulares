package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var AExpressaoRegular string
	var AAlterarTexto string

	AReader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite um texto qualquer:")
	ATextoDigitado, err := AReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("")

	fmt.Println("Digite uma Expressão Regular para pesquisa (sem espaços):")
	if _, err := fmt.Scanln(&AExpressaoRegular); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("")

	fmt.Println("Digite uma palavra para alterar o texto na busca (sem espaços):")
	if _, err := fmt.Scanln(&AAlterarTexto); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("")

	ARegEx, _ := regexp.Compile(AExpressaoRegular)
	AMatchString := ARegEx.MatchString(ATextoDigitado)
	AFindString := ARegEx.FindString(ATextoDigitado)
	AFindStringIndex := ARegEx.FindStringIndex(ATextoDigitado)
	AFindAllString := ARegEx.FindAllString(ATextoDigitado, 100)
	AReplaceAllString := ARegEx.ReplaceAllString(ATextoDigitado, AAlterarTexto)

	fmt.Println("Resultados:")
	fmt.Println("Texto digitado:", ATextoDigitado)
	fmt.Println("RegEx digitado:", AExpressaoRegular)
	fmt.Println("")
	fmt.Println("Encontrou expressão:", func() string {
		if AMatchString {
			return "SIM"
		}
		return "NÃO"
	}())
	fmt.Println("")
	fmt.Println("Texto encontrado:", AFindString)
	fmt.Println("")
	fmt.Println("Primeira posição encontrada:", AFindStringIndex)
	fmt.Println("")
	fmt.Println("Quantidade encontrada na busca:", len(AFindAllString))
	fmt.Println("")
	fmt.Println("Alteração no texto:", AReplaceAllString)

	fmt.Println("Pressione 'Enter' para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
