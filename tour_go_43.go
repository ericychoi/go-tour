package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    wc := make(map[string]int)
    for _,v := range(strings.Fields(s)) {
        _, ok := wc[v];
        if (!ok) {
            wc[v] = 1
        } else {
            wc[v]++
        }
    }
    return wc
}

func main() {
    wc.Test(WordCount)
}
