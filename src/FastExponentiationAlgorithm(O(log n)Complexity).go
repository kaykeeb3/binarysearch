package main

import "fmt"

// Função que implementa a exponenciação rápida
func power(base, exponent int) int {
    if exponent == 0 {
        return 1
    }
    if exponent%2 == 0 {
        halfPower := power(base, exponent/2)
        return halfPower * halfPower
    }
    return base * power(base, exponent-1)
}

func main() {
    // Exemplo de uso
    base := 2
    exponent := 5

    result := power(base, exponent)
    fmt.Printf("%d elevado a %d é igual a %d\n", base, exponent, result)
}
