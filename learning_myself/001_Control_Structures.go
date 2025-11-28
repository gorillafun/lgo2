package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 { // 2で割った余りが0なら (偶数)
            fmt.Println(i, "even")
        } else { // それ以外なら (奇数)
            fmt.Println(i, "odd")
        }
    }
}