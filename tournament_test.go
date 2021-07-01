package Tournament

import (
	"bytes"
	"strings"
	"testing"
)

var happyTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "good",
		input: `
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`[1:], // [1:] = strip initial readability newline
	},{
		description: "ignore comments and newlines",
		input: `

Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;win
# Catastrophic Loss of the Californians
Courageous Californians;Blithering Badgers;loss

Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
Devastating Donkeys;Courageous Californians;draw


`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskians            |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`[1:],
	},
}

func TestTallyHappy(t *testing.T) {
	for _, tt := range happyTestCases {
		reader := strings.NewReader(tt.input)
		var buffer bytes.Buffer
		err := Tally(reader, &buffer)
		actual := buffer.String()
		// We don't expect errors for any of the test cases
		if err != nil {
			t.Fatalf("Tally for input named %q returned error %q. Error not expected.",
				tt.description, err)
		}
		if actual != tt.expected {
			t.Fatalf("Tally for input named %q was expected to return...\n%s\n...but returned...\n%s",
				tt.description, tt.expected, actual)
		}
	}
}
