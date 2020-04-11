package main

import (
	"fmt"
)

func main() {
	var idade float64
	var laco [1]int
	parar := ""
	for i := 0; i < len(laco); i++ {
		fmt.Print("Digite sua idade: ")
		fmt.Scan(&idade)
		fmt.Printf("Sua idade no planeta Terra em segundos é %f \n", idade*31557600.0)
	}
	for parar != "q" {	
		fmt.Print("Digite um número de 1 a 7 para ver sua idade ou 'q' para parar a aplicação: ")
		fmt.Scan(&parar)
		if parar == "1" {
			fmt.Printf("Sua idade em Mercúrio é %f \n", idade*0.2408467)
		}
		if parar == "2" {
			fmt.Printf("Sua idade em Vênus é %f \n", idade*0.61519726)
		}
		if parar == "3" {
			fmt.Printf("Sua idade em Marte é %f \n", idade*1.8808158)
		}
		if parar == "4" {
			fmt.Printf("Sua idade em Júpiter é %f \n", idade*11.862615)
		}
		if parar == "5" {
			fmt.Printf("Sua idade em Saturno é %f \n", idade*29.447498)
		}
		if parar == "6" {
			fmt.Printf("Sua idade em Urano é %f \n", idade*84.016846)
		}
		if parar == "7" {
			fmt.Printf("Sua idade em Netuno é %f \n", idade*164.79132)
		}

	}
	fmt.Print("Você parou a aplicação \n")
}
