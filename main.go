package main

import (
	"fmt"
)

func separar(arr []int, menor, maior int) int {
	//Seleciona o ultimo indice do array
	ultimoValor := arr[maior]
	for j := menor; j < maior; j++ {
		//Verifica se o valor dentro do array é menor que o ultimo indice, caso sim ele inverte as posições do valor encontrado com o menor valor
		if arr[j] < ultimoValor {
			arr[j], arr[menor] = arr[menor], arr[j]
			menor++
		}
	}
	//Muda a posição dos menores valores para antes dos maiores valores
	arr[menor], arr[maior] = arr[maior], arr[menor]
	fmt.Println(menor)
}

func ordenar(arr []int, menor, maior int) {
	//Checa até que o menor valor seja maior
	if menor > maior {
		return
	}
	//Separa o array
	pivo := separar(arr, menor, maior)
	ordenar(arr, menor, pivo-1)
	ordenar(arr, pivo+1, maior)
}

// Declara o array e realiza a chamada da função para ordenar
func main() {
	arr := []int{22, 7, 9, 12, 8, 36, 45, 92, 54, 82, 63, 15, 4, 3}

	ordenar(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
