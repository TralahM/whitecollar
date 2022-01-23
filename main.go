package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tralahm/whitecollar/pkg/cctek"
)

func main() {
	argsLen := len(os.Args[1:])
	card := "4407830106254205"
	var cardrecords cctek.CardRecords
	if argsLen > 0 {
		dirname := string(os.Args[1][0])
		filename := string(os.Args[1][0:4])
		target := "data/" + dirname + "/" + filename + ".yml"
		data, err := ioutil.ReadFile(target)
		if err != nil {
			log.Fatalln(err)
		}
		if err := cardrecords.Parse(data); err != nil {
			log.Fatalln(err)
		}
		// fmt.Printf("%+v\n", cardrecords)
		fmt.Printf("%s\n", cardrecords.Get(string(os.Args[1][0:6])).String())
		fmt.Println("Length of records: ", len(cardrecords))
		fmt.Println("Number of CMD line args: ", argsLen)
	} else {
		data, err := ioutil.ReadFile("data/4/4300.yml")
		if err != nil {
			log.Fatalln(err)
		}
		if err := cardrecords.Parse(data); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%+v\n", cardrecords)
		fmt.Println("Length of records: ", len(cardrecords))
		fmt.Println("Number of CMD line args: ", argsLen)
	}

	fmt.Println("Check card: ", card, " isValid: ", cctek.LuhnCheck(card))
}
