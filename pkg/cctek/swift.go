package cctek

import (
	"strings"

	"gopkg.in/yaml.v2"
)

// SwiftData struct describes the swift data of a record.
type SwiftData struct {
	Institution string `yaml:"institution"`
	City        string `yaml:"city"`
	Branch      string `yaml:"branch"`
	Code        string `yaml:"-"`
}

func (s SwiftData) String() string {
	return "INSTITUTION: " + s.Institution + "\tCITY: " + s.City + "\tBRANCH: " + s.Branch + "\tCODE: " + s.Code + "\n"
}

// SwiftRecords struct
type SwiftRecords map[string]SwiftData

// Get returns the SwiftData associated with the SwiftCode
func (r SwiftRecords) Get(key string) SwiftData {
	if sdata, exists := r[key]; exists {
		return sdata
	} else {
		return SwiftData{}
	}
}

// Parse parses yaml data as bytes
func (r *SwiftRecords) Parse(data []byte) error {
	err := yaml.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	r.Enhance()
	return nil
}

// Enhance enhances the SwiftData with the code.
func (r *SwiftRecords) Enhance() {
	for k, v := range *r {
		v.Code = k
		(*r)[k] = v
	}
}

// String returns a string representation of the swift records
func (r *SwiftRecords) String() string {
	var s = ""
	for _, v := range *r {
		s += v.String() //+ "\n"
	}
	return s
}

// returns the country code from a swiftcode
func countryCodeFromSwiftCode(swiftcode string) string {
	return strings.ToLower(swiftcode[4:6])
}
