package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Expressoes() {
	var AExpressaoRegular string
	var AAlterarTexto string

	AReader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite um texto qualquer:")
	ATextoDigitado, err := AReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nDigite uma Expressão Regular para pesquisa (sem espaços):")
	if _, err := fmt.Scanln(&AExpressaoRegular); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nDigite uma palavra para alterar o texto na busca (sem espaços):")
	if _, err := fmt.Scanln(&AAlterarTexto); err != nil {
		fmt.Println("Error:", err)
	}

	ARegEx, _ := regexp.Compile(AExpressaoRegular)
	AMatchString := ARegEx.MatchString(ATextoDigitado)
	AFindString := ARegEx.FindString(ATextoDigitado)
	AFindStringIndex := ARegEx.FindStringIndex(ATextoDigitado)
	AFindAllString := ARegEx.FindAllString(ATextoDigitado, 100)
	AReplaceAllString := ARegEx.ReplaceAllString(ATextoDigitado, AAlterarTexto)

	fmt.Println("\nResultados:")
	fmt.Println("Texto digitado:", ATextoDigitado)
	fmt.Println("RegEx digitado:", AExpressaoRegular)

	fmt.Println("\nEncontrou expressão:", func() string {
		if AMatchString {
			return "SIM"
		}
		return "NÃO"
	}())

	fmt.Println("\nTexto encontrado:", AFindString)
	fmt.Println("\nPrimeira posição encontrada:", AFindStringIndex)
	fmt.Println("\nQuantidade encontrada na busca:", len(AFindAllString))
	fmt.Println("\nAlteração no texto:", AReplaceAllString)

	fmt.Println("Pressione 'Enter' para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	for {
		Expressoes()
	}
}
