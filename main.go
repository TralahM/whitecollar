package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tralahm/whitecollar/pkg/cctek"
)

// Parser interface
type Parser interface {
	Parse([]byte) error
}

// readAndParseData reads a file called target and parses the records into the
// parser.
func readAndParseData(target string, parser Parser) {
	data, err := ioutil.ReadFile(target)
	if err != nil {
		log.Fatalln(err)
	}
	if err := parser.Parse(data); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	argsLen := len(os.Args[1:])
	card := "4407830106254205"
	var cardrecords cctek.CardRecords
	var swiftrecords cctek.SwiftRecords

	if argsLen > 0 {
		dirname := string(os.Args[1][0])
		filename := string(os.Args[1][0:4])
		target := "data/" + dirname + "/" + filename + ".yml"
		readAndParseData(target, &cardrecords)
		// fmt.Printf("%+v\n", cardrecords)
		details := cardrecords.Get(string(os.Args[1][0:6]))
		fmt.Printf("%s\n", details.String())

		swiftpath := "data/swift/" + details.SwiftFilename()
		readAndParseData(swiftpath, &swiftrecords)
		fmt.Printf("SwiftRecords:\n%s\n", swiftrecords.String())
		fmt.Println("Length of records: ", len(cardrecords))
		fmt.Println("Number of CMD line args: ", argsLen)
	} else {
		readAndParseData("data/4/4407.yml", &cardrecords)
		// fmt.Printf("%+v\n", cardrecords)
		details := cardrecords.Get("440783")
		fmt.Printf("%s\n", details.String())

		swiftpath := "data/swift/" + details.SwiftFilename()
		readAndParseData(swiftpath, &swiftrecords)
		fmt.Printf("SwiftRecords:\n %s\n", swiftrecords.String())
		fmt.Println("Length of records: ", len(cardrecords))
		fmt.Println("Number of CMD line args: ", argsLen)
	}

	fmt.Println("Check card: ", card, " isValid: ", cctek.LuhnCheck(card))
}
