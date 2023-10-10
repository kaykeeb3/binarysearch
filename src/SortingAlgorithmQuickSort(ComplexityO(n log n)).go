package main

import "fmt"

// Função que implementa o QuickSort
func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0]
    var less, equal, greater []int

    for _, num := range arr {
        if num < pivot {
            less = append(less, num)
        } else if num == pivot {
            equal = append(equal, num)
        } else {
            greater = append(greater, num)
        }
    }

    return append(append(quickSort(less), equal...), quickSort(greater)...)
}

func main() {
    // Exemplo de uso
    arr := []int{12, 4, 5, 6, 7, 3, 1, 15}
    
    fmt.Println("Array antes da ordenação:", arr)
    
    sortedArr := quickSort(arr)
    
    fmt.Println("Array após a ordenação:", sortedArr)
}
