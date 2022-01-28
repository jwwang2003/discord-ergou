package helpers

const (
	PRIMARY   int = 9489145
	SECONDARY int = 13538264
	ERROR     int = 16007990
	WARNING   int = 16754470
	INFO      int = 2733814
	SUCCESS   int = 6732650
)

func ResolveColor(color [3]int) int {
	return (color[0] << 16) + (color[1] << 8) + color[2]
}