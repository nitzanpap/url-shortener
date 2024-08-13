package pkg

import "math"

const (
	NUM_OF_CHARS_IN_URL_ID = 4
)

var NumOfPossibleUrls int

func init() {
	NumOfPossibleUrls = int(math.Pow(62, NUM_OF_CHARS_IN_URL_ID))
}
