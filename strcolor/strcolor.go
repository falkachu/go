package strcolor

import "fmt"

func Red(str string) string {
	return fmt.Sprintf("\u001b[31m%s\u001b[0m", str)
}

func GreenBright(str string) string {
	return fmt.Sprintf("\u001b[92m%s\u001b[0m", str)
}

func YellowBright(str string) string {
	return fmt.Sprintf("\u001b[93m%s\u001b[0m", str)
}

func Blue(str string) string {
	return fmt.Sprintf("\u001b[34m%s\u001b[0m", str)
}
