package main

import (
	"fmt"
	"regexp"
	"sync"
)

var gamers = []string{"Tyrion Lannister", "Jon Snow", "Ned Stark", "Sandor Clegane"}

func main() {
	var wg sync.WaitGroup
	for _, gamer := range gamers {
		wg.Add(1)
		go func(gamer string) {
			defer wg.Done()
			re := regexp.MustCompile("Stark")
			out := re.FindString(gamer)
			if out == "" {
				fmt.Println(gamer, ": probably you are still alive")
			} else {
				fmt.Println(gamer, ": Ohh You are a Stark, you are dead")
			}
		}(gamer)
		for i := 0; i < 2; i++ {
			fmt.Println("doing something else twice for every gamer")
		}
	}
	wg.Wait()
}
