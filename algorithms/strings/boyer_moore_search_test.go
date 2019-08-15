package strings

import (
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		Text          string
		Pattern       string
		ExpectedIndex int
	}{
		{Text: "ACBBCAB", Pattern: "AB", ExpectedIndex: 5},
		{Text: "ACBBCAB", Pattern: "AC", ExpectedIndex: 0},
		{Text: "ACBBCAB", Pattern: "CA", ExpectedIndex: 4},
	}
	for _, test := range tests {
		if gotIndex := Index(test.Text, test.Pattern); gotIndex != test.ExpectedIndex {
			t.Fatalf("given text:%s, pattern:%s, got index:%d, expected index:%d",test.Text,test.Pattern,gotIndex,test.ExpectedIndex)
		}
	}

}
