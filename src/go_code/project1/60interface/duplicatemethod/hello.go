package main

func main() {

}

type BInterface interface {
	Test01()
	Test02()
}

type CInterface interface {
	Test01()
	Test03()
}

type DInterface interface {
	BInterface
	CInterface
	//test03()
}
