package main

import "fmt"

func main() {

    channel1 := make(chan int)
    channel2 := make(chan int)
    channel3 := make(chan int)

    //First Way
    var channels_first []chan int
    channels_first = append(channels_first, channel1)
    channels_first = append(channels_first, channel2)
    channels_first = append(channels_first, channel3)

    fmt.Println("\nOutput for First slice of channels")
    for _, c := range channels_first {
        fmt.Println(c)
    }

    //Second Way
    channels_second := make([]chan int, 3)
    channels_second[0] = channel1
    channels_second[1] = channel2
    channels_second[2] = channel3

    fmt.Println("\nOutput for Second slice of channels")
    for _, c := range channels_second {
        fmt.Println(c)
    }
}
