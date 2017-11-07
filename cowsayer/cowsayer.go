package cowsayer

import "strings"

func Simplesay(args []string) string {
	if len(args) <= 1 {
		args = append(args, "hey")
	}
	args[0] = "oh"
	return strings.Join(args, " ")
}
