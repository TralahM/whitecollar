package cctek

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Country structure
type Country struct {
	AlphaV2 string `yaml:"alpha_2"`
	AlphaV3 string `yaml:"alpha_3"`
	Name    string `yaml:"name"`
}

// String returns the string representation of a Country.
func (c Country) String() string {
	return "Country: " + c.Name + "\tAlpha2: " + c.AlphaV2 + "\t" + "Alpha3: " + c.AlphaV3 + "\n"
}

// Card structure
type Card struct {
	Bin      string  `yaml:"bin"`
	Brand    string  `yaml:"brand"`
	Type     string  `yaml:"type"`
	Issuer   string  `yaml:"issuer"`
	Category string  `yaml:"category"`
	Country  Country `yaml:"country"`
}

// CardRecords is a Map of BIN to Card details
type CardRecords map[string]Card

// Parse parses yaml bytes into CardRecords
func (c *CardRecords) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// Get returns a Card represented by the key BIN
func (c CardRecords) Get(key string) Card {
	if card, exists := c[key]; exists {
		return card
	} else {
		return Card{}
	}
}

// String returns the string representation of a Card.
func (card Card) String() string {
	ret := "BIN: " + card.Bin + "\nBRAND: " + card.Brand + "\nCategory: " + card.Category + "\nIssuer: " + card.Issuer + "\nType: " + card.Type + "\n" + card.Country.String() + "\n"
	return ret
}

// sum adds digits in array
func sum(digits []int64) int64 {
	var s int64 = 0
	for i := 0; i < len(digits); i++ {
		s += digits[i]
	}
	return s
}

// LuhnCheck checks a credit card Number for validity with Luhn's Algorithm
func LuhnCheck(cardNumber string) bool {
	// odd digits
	fmt.Println("Card Number: ", cardNumber)
	var odds []int64 = make([]int64, 0)
	var dbl_odds []int64 = make([]int64, 0)
	var dbl_odds_less_nine []int64 = make([]int64, 0)
	var evens []int64 = make([]int64, 0)
	for i := 0; i < len(cardNumber)-1; i++ {
		v, _ := strconv.ParseInt(string(cardNumber[i]), 10, 64)
		odds = append(odds, v)
		i++
	}

	for i := 1; i < len(cardNumber)-1; i++ {
		v, _ := strconv.ParseInt(string(cardNumber[i]), 10, 64)
		evens = append(evens, v)
		i++
	}

	fmt.Println("odds: ", odds, " length: ", len(odds))
	fmt.Println("evens: ", evens)

	for i := 0; i < len(odds); i++ {
		dbl_odds = append(dbl_odds, odds[i]*2)
	}

	fmt.Println("dbl_odds: ", dbl_odds)

	for i := 0; i < len(dbl_odds); i++ {
		if dbl_odds[i] < 10 {
			dbl_odds_less_nine = append(dbl_odds_less_nine, dbl_odds[i])
		} else {
			dbl_odds_less_nine = append(dbl_odds_less_nine, dbl_odds[i]-9)
		}
	}

	fmt.Println("dbl_odds_less_nine: ", dbl_odds_less_nine)

	// Checksum
	checksum := 10 - ((sum(evens) + sum(dbl_odds_less_nine)) % 10)
	fmt.Println("sum evens: ", sum(evens))
	fmt.Println("sum odds: ", sum(dbl_odds_less_nine))
	fmt.Println("Calculated checksum: ", checksum)
	if strconv.FormatInt(checksum, 10) == string(cardNumber[(len(cardNumber)-1)]) {
		return true
	}
	return false
}
