package cctek

import "gopkg.in/yaml.v2"

type Country struct {
	AlphaV2 string `yaml:"alpha_2"`
	AlphaV3 string `yaml:"alpha_3"`
	Name    string `yaml:"name"`
}

type Card struct {
	Bin      string  `yaml:"bin"`
	Brand    string  `yaml:"brand"`
	Type     string  `yaml:"type"`
	Issuer   string  `yaml:"issuer"`
	Category string  `yaml:"category"`
	Country  Country `yaml:"country"`
}

type CardRecords map[string]Card

func (c *CardRecords) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func (c CardRecords) Get(key string) Card {
	if card, exists := c[key]; exists {
		return card
	} else {
		return Card{}
	}
}

func (card Card) String() string {
	ret := "BIN: " + card.Bin + "\nBRAND: " + card.Brand + "\nCategory: " + card.Category + "\nIssuer: " + card.Issuer + "\nType: " + card.Type + "\nCountry: " + card.Country.Name + "\n"
	return ret
}
