package main

import "fmt"

// Função que calcula a soma dos elementos de um array
func sum(arr []int) int {
    result := 0
    for _, num := range arr {
        result += num
    }
    return result
}

func main() {
    // Exemplo de uso
    arr := []int{1, 2, 3, 4, 5}
    total := sum(arr)
    fmt.Printf("A soma dos elementos do array é: %d\n", total)
}
