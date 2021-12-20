package jp

import (
	"fmt"
	"github.com/szks-repo/gphone"
	"testing"
)

func TestNewJp(t *testing.T) {

	ph, err := gphone.New("090 (1234) 1234") 
	if err != nil {
		t.Error(err)
	}

	jph := NewJapanPhoneNumber(ph)

	fmt.Println(jph.Separate(SepHyphen))
	fmt.Println(jph.Separate(SepDot))
	fmt.Println(jph.Separate(SepParentis))
}