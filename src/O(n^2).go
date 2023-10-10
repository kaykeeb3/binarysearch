package main

import "fmt"

// Função que implementa a ordenação por inserção
func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1

        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }

        arr[j+1] = key
    }
}

func main() {
    // Exemplo de uso
    arr := []int{12, 11, 13, 5, 6}
    
    fmt.Println("Array antes da ordenação:", arr)
    
    insertionSort(arr)
    
    fmt.Println("Array após a ordenação:", arr)
}
