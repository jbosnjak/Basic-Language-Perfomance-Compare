package main

import "fmt"
import "time"

func main() {

    timestart := time.Now()

    for c := 0; c < 1000000; c++ {
        bubble()
    }

    timeend := time.Now();

    fmt.Print(timeend.Sub(timestart));
}

func bubble() {
    array := [...]int16{3,4,1,3,5,1,92,2,4124,424,52,12}

    for i := 0; i < len(array); i++ {
        for y := 0; y < len(array) - 1; y++ {
            if array[y+1] < array[y] {
                t := array[y]
                array[y] = array[y+1]
                array[y+1] = t
            }
        }
    }
}