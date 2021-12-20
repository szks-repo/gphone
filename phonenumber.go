package gphone

import (
	"errors"
	"fmt"
	"golang.org/x/text/width"
	"regexp"
	"strings"
	"unicode/utf8"
)

type PhoneNumber struct {
	inputNumber string
	c14nNumber  string
	countryCode string
}

var c14nPattern = regexp.MustCompile(`^[0-9]+$`)

func c14n(ph string) string {
	ph = width.Narrow.String(ph)
	ph = strings.TrimPrefix(ph, " ")
	//TODO Consider i18n
	if strings.HasPrefix(ph, "+XX") {
		//
	}
	ph = strings.Replace(ph, "-", "", -1)
	ph = strings.Replace(ph, ".", "", -1)
	ph = strings.Replace(ph, " ", "", -1)
	ph = strings.Replace(ph, "(", "", -1)
	ph = strings.Replace(ph, ")", "", -1)

	return ph
}

func New(number string) (*PhoneNumber, error) {
	if number == "" {
		return nil, errors.New("invalid number")
	}

	c14ed := c14n(number)
	if !c14nPattern.MatchString(c14ed) {
		return nil, errors.New("invalid character detected")
	}

	//119
	//110
	//117
	if utf8.RuneCountInString(c14ed) == 3 {
		return &PhoneNumber{
			inputNumber: number,
			c14nNumber: c14ed,
		}, nil
	}

	//

	return &PhoneNumber{
		inputNumber: number,
		c14nNumber: c14ed,
	}, nil
}

func (ph *PhoneNumber) Value() string {
	return ph.c14nNumber
}

func (ph *PhoneNumber) GetI18nNumber(countryCode string) string {
	return fmt.Sprintf("+%s-%s", countryCode, ph.c14nNumber)
}

const (
	CountryCodeJP = "81"
	CountryCodeUS = ""
	CountryCodeXX = ""
	CountryCodeYY = ""
	CountryCodeZZ = ""
)

type Country struct {
	Name string
	Code string
}

func (ph *PhoneNumber) GetCountry() *Country {
	return &Country{}
}