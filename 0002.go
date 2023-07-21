package main

import "fmt"

func main() {
    x := 1
    y := 2
    z := 0
    tot := 2

    for x + y < 4000000 {
        z = x + y
        x = y
        y = z
        if z % 2 == 0 {
            tot += z
        }
    }

    fmt.Println(tot)
}
