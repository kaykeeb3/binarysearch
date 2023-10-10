package main

import "fmt"

// Função que implementa a busca binária
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1 // Retorna -1 se o elemento não for encontrado
}

func main() {
    // Exemplo de uso
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    target := 7

    index := binarySearch(arr, target)

    if index != -1 {
        fmt.Printf("O elemento %d foi encontrado no índice %d\n", target, index)
    } else {
        fmt.Printf("O elemento %d não foi encontrado no array\n", target)
    }
}
