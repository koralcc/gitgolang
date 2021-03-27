package main

import "fmt"

func main() {
	catCh := make(chan bool, 1)
	dogCh := make(chan bool, 1)
	fishCh := make(chan bool, 1)
	exitCh := make(chan bool, 1)
	fishCh <- true
	go CatSay(fishCh, catCh)
	go DogSay(catCh, dogCh)
	go FishSay(dogCh, fishCh, exitCh)
	<-exitCh
}

func CatSay(fishCh <-chan bool, catCh chan<- bool) {
	for i := 0; i < 100; i++ {
		if v := <-fishCh; v {
			fmt.Println("Cat")
			catCh <- true
		}
	}
}

func DogSay(catCh <-chan bool, dogCh chan<- bool) {
	for i := 0; i < 100; i++ {
		if v := <-catCh; v {
			fmt.Println("Dog")
			dogCh <- true
		}
	}
}

func FishSay(dogCh <-chan bool, fishCh chan<- bool, exitCh chan<- bool) {
	for i := 0; i < 100; i++ {
		if v := <-dogCh; v {
			fmt.Println("Fish")
			fishCh <- true
		}
	}
	exitCh <- true
}
