package gphone

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	

	ph, err := New("0120-111-333")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	fmt.Println(ph)
}
