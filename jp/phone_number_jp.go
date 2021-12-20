package jp

import (
	"fmt"
	"github.com/szks-repo/gphone"
	"regexp"
	"strings"
	"unicode/utf8"
)

type JapanPhoneNumber struct {
	phoneNumber *phgo.PhoneNumber
}

var (
	mobilePhonePattern      = regexp.MustCompile(`^0[1-9]0`)
	fixedPhonePattern       = regexp.MustCompile(`^0[1-9]{3}`)
	regexpHighLevelService  = regexp.MustCompile(`0[1-9]{2}0`)
	regexpInterceptPhone    = regexp.MustCompile(``)
	importantPhonePattern   = regexp.MustCompile(`^1`)
	fixedPhoneInCityPattern = regexp.MustCompile(`^[2-9]]`)
)

func NewJapanPhoneNumber(ph *phgo.PhoneNumber) *JapanPhoneNumber {
	return &JapanPhoneNumber{
		phoneNumber: ph,
	}
}

type PhoneType struct {
	Pattern         string
	Name            string
	sepIndexPattern []int
	IsFree          bool
}

func (jp *JapanPhoneNumber) GetPhoneType() *PhoneType {
	if mobilePhonePattern.MatchString(jp.phoneNumber.Value()) {
		pht := &PhoneType{
			Pattern:         mobilePhonePattern.String(),
			Name:            "Mobile",
			sepIndexPattern: []int{2, 6, 10},
		}
		if utf8.RuneCountInString(jp.phoneNumber.Value()) > 12 {
			//TODO Consider
		}
		if strings.HasPrefix(jp.phoneNumber.Value(), "0800") {
			pht.IsFree = true
		}

		return pht
	}

	if fixedPhonePattern.MatchString(jp.phoneNumber.Value()) {
		return &PhoneType{
			Pattern:         fixedPhonePattern.String(),
			Name:            "Fixed",
			sepIndexPattern: []int{2, 5, 9},
		}
	}
	if strings.HasPrefix(jp.phoneNumber.Value(), "0550") {
		return &PhoneType{
			Pattern:         fixedPhonePattern.String(),
			Name:            "Fixed",
			sepIndexPattern: []int{3, 5, 9},
		}
	}

	if regexpHighLevelService.MatchString(jp.phoneNumber.Value()) {
		pht := &PhoneType{
			Pattern:         "",
			Name:            "",
			sepIndexPattern: []int{3, 5, 9},
		}
		if strings.HasPrefix(jp.phoneNumber.Value(), "0120") {
			pht.IsFree = true
		}
	}

	return &PhoneType{
		Pattern:         "",
		Name:            "",
		sepIndexPattern: nil,
		IsFree:          false,
	}
}

func (jp *JapanPhoneNumber) GetPrefecture() {

}

type Separator string

const (
	SepHyphen   Separator = "-"
	SepDot      Separator = "."
	SepParentis Separator = "()"
)

func (jp *JapanPhoneNumber) Separate(sep Separator, sepIndexPattern ...[]int) string {
	var separated []string
	base := strings.Split(jp.phoneNumber.Value(), "")

	if len(base) == 3 {
		return jp.phoneNumber.Value()
	}

	phType := jp.GetPhoneType()
	pattern := phType.sepIndexPattern
	if len(sepIndexPattern) > 0 && sepIndexPattern[0] != nil {
		pattern = sepIndexPattern[0]
	}

	var part string
	for i := range base {
		part += base[i]
		for _, idx := range pattern {
			if i == idx {
				separated = append(separated, part)
				part = ""
				break
			}
		}
	}

	if sep == SepHyphen || sep == SepDot {
		return strings.Join(separated, string(sep))
	}

	if sep == SepParentis && len(separated) == 3 {
		return fmt.Sprintf("%s(%s)%s", separated[0], separated[1], separated[2])
	}

	return "UNKNOWN"
}

