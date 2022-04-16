package strcolor

import "fmt"

const (
	RESET         = "\u001b[0m"
	BLACK         = "\u001b[30m"
	RED           = "\u001b[31m"
	GREEN         = "\u001b[32m"
	YELLOW        = "\u001b[33m"
	BLUE          = "\u001b[34m"
	MAGENTA       = "\u001b[35m"
	CYAN          = "\u001b[36m"
	WHITE         = "\u001b[37m"
	BLACKBRIGHT   = "\u001b[90m"
	REDBRIGHT     = "\u001b[91m"
	GREENBRIGHT   = "\u001b[92m"
	YELLOWBRIGHT  = "\u001b[93m"
	BLUEBRIGHT    = "\u001b[94m"
	MAGENTABRIGHT = "\u001b[95m"
	CYANBRIGHT    = "\u001b[96m"
	WHITEBRIGHT   = "\u001b[97m"
)

func Black(str string) string {
	return fmt.Sprintf("%s%s%s", BLACK, str, RESET)
}

func Red(str string) string {
	return fmt.Sprintf("%s%s%s", RED, str, RESET)
}

func Green(str string) string {
	return fmt.Sprintf("%s%s%s", GREEN, str, RESET)
}

func Yellow(str string) string {
	return fmt.Sprintf("%s%s%s", YELLOW, str, RESET)
}

func Blue(str string) string {
	return fmt.Sprintf("%s%s%s", BLUE, str, RESET)
}

func Magenta(str string) string {
	return fmt.Sprintf("%s%s%s", MAGENTA, str, RESET)
}

func Cyan(str string) string {
	return fmt.Sprintf("%s%s%s", CYAN, str, RESET)
}

func White(str string) string {
	return fmt.Sprintf("%s%s%s", WHITE, str, RESET)
}

func BlackBright(str string) string {
	return fmt.Sprintf("%s%s%s", BLACKBRIGHT, str, RESET)
}

func RedBright(str string) string {
	return fmt.Sprintf("%s%s%s", REDBRIGHT, str, RESET)
}

func GreenBright(str string) string {
	return fmt.Sprintf("%s%s%s", GREENBRIGHT, str, RESET)
}

func YellowBright(str string) string {
	return fmt.Sprintf("%s%s%s", YELLOWBRIGHT, str, RESET)
}

func BlueBright(str string) string {
	return fmt.Sprintf("%s%s%s", BLUEBRIGHT, str, RESET)
}

func MagentaBright(str string) string {
	return fmt.Sprintf("%s%s%s", MAGENTABRIGHT, str, RESET)
}

func CyanBright(str string) string {
	return fmt.Sprintf("%s%s%s", CYANBRIGHT, str, RESET)
}

func WhiteBright(str string) string {
	return fmt.Sprintf("%s%s%s", WHITEBRIGHT, str, RESET)
}
