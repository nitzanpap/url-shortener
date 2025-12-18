package pkg

import "math"

const (
	NumOfCharsInURLID = 4
)

var NumOfPossibleUrls int

func init() {
	NumOfPossibleUrls = int(math.Pow(62, NumOfCharsInURLID))
}
