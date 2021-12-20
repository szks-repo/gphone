package gphone

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	
	tests := []struct{
		name string
		arg  string
	}{
		{
			name: "0120",
			arg: "0120-111-333",
		},
		{
			name: "080 parentis",
			arg:  "080(1234)1234",
		},
		{
			name: "080",
			arg: "080-1234-1234",
		},
		{
			name: "090",
			arg: "090-1234-1234",
		},
		{
			name: "070",
			arg: "070-1234-1234",
		},
		{
			name: "0800",
			arg: "0800-1234-1234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			ph, err := New(tt.arg)
			if err != nil {
				t.Errorf("err: %v", err)
			}
			fmt.Println("==========>", ph.Value())
		})
	}

}
