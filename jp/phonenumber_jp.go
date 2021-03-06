package jp

import (
	"errors"
	"fmt"
	"github.com/szks-repo/gphone"
	"regexp"
	"strings"
)

type JapanPhoneNumber struct {
	phoneNumber *gphone.PhoneNumber
	phoneType   *PhoneType
}

type PhoneType struct {
	Name            string
	sepIndexPattern []int
	IsFree          bool
}

const (
	phoneTypeMobile           = "Mobile"
	phoneTypeFixed            = "Fixed"
	phoneTypeImportant        = "Important"
	//特番?
	phoneTypeHighLevelService = "HighLevelService"
	phoneTypeRelay            = "Relay"
)

var (
	// 携帯電話
	// 11桁 090-xxxx-xxxx
	mobilePhonePattern      = regexp.MustCompile(`^0[1-9]0`)
	// 固定電話(市外局番)
	// 10桁
	fixedPhonePattern       = regexp.MustCompile(`^0[1-9]{3}`)
	// ナビダイヤル
	// フリーダイヤル
	highLevelServicePattern = regexp.MustCompile(`^0[1-9]{2}0`)
	// 中継電話会社経由
	relayPhonePattern       = regexp.MustCompile(`^00`)
	// 固定電話(市内) -> スコープ外とする
	//fixedPhoneInCityPattern = regexp.MustCompile(`^[2-9]`)
)

var real3DigitNumbers = []string{
	"104",
	"110",
	"115",
	"117",
	"118",
	"119",
	"171",
	"177",
}

func NewJapanPhoneNumber(ph *gphone.PhoneNumber) (*JapanPhoneNumber, error) {
	jph := &JapanPhoneNumber{
		phoneNumber: ph,
	}

	phtype, err := jph.GetPhoneType();
	if err != nil {
		return nil, err
	}
	jph.phoneType = phtype

	return jph, nil
}

func (jp *JapanPhoneNumber) GetPhoneType() (*PhoneType, error) {
	// 携帯電話
	if mobilePhonePattern.MatchString(jp.phoneNumber.Value()) {
		pht := &PhoneType{
			Name:            phoneTypeMobile,
			sepIndexPattern: []int{2, 6, 10},
		}

		if strings.HasPrefix(jp.phoneNumber.Value(), "0800") {
			pht.IsFree = true
		}

		return pht, nil
	}

	// 固定電話
	// 0550 御殿場市
	if fixedPhonePattern.MatchString(jp.phoneNumber.Value()) {
		return &PhoneType{
			Name:            phoneTypeFixed,
			sepIndexPattern: []int{2, 5, 9},
		}, nil
	}
	if strings.HasPrefix(jp.phoneNumber.Value(), "0550") {
		return &PhoneType{
			Name:            phoneTypeFixed,
			sepIndexPattern: []int{3, 5, 9},
		}, nil
	}

	// 0120 フリーダイヤル
	// 0570 ナビダイヤル
	//
	if highLevelServicePattern.MatchString(jp.phoneNumber.Value()) {
		phtype := &PhoneType{
			Name:            phoneTypeHighLevelService,
			sepIndexPattern: []int{3, 5, 9},
		}
		if strings.HasPrefix(jp.phoneNumber.Value(), "0120") {
			phtype.IsFree = true
		}

		return phtype, nil
	}

	// 104 電話番号の案内
	// 110 警察への通報
	// 115 電報受付
	// 117 時報
	// 118 海上保安機関への通報
	// 119 消防への通報
	// 171 災害用伝言ダイヤル
	// 177 天気予報
	if strings.HasPrefix(jp.phoneNumber.Value(), "1") && len(jp.phoneNumber.Value()) == 3 {
		for _, num := range real3DigitNumbers {
			if num == jp.phoneNumber.Value() {
				return &PhoneType{
					Name: phoneTypeImportant,
				}, nil
			}
		}

		return nil, errors.New("not present number")
	}

	// 中継電話
	if relayPhonePattern.MatchString(jp.phoneNumber.Value()) {
		return &PhoneType{
			Name: phoneTypeRelay,
		}, nil
	}

	return nil, errors.New("unknown or unsupported phone type")
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

	pattern := jp.phoneType.sepIndexPattern
	if len(sepIndexPattern) > 0 && sepIndexPattern[0] != nil {
		pattern = sepIndexPattern[0]
	}

	var part string
	for i := range base {
		part = part + base[i]
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

func (jp *JapanPhoneNumber) GetPrefecture() {

}