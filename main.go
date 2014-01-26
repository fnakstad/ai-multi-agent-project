package main

import "fmt"

func main() {
    // TODO: Set these up as CLI arguments later
    iterations := 10000;

    // Read in reward matrix

    for i:=1; i <= iterations; i++ {

        // For each agent
        //  choose action according to Q-values
        //  perform action
        //  observe r (result?) s' (resulting state?)
        //  update the q-values for the state/action
    }

    // TODO: Print run report
    fmt.Println("End report:")
}

type Agent struct {
    Points int
    epsilon float32
    alpha float32
} 

func(a Agent) ChooseAction() string {
    return "hey"
}