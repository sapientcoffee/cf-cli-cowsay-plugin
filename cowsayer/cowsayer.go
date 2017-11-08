package cowsayer

import (
	"strings"

	neocowsay "github.com/Code-Hex/Neo-cowsay"
)

func Simplesay(args []string) string {
	if len(args) <= 1 {
		args = append(args, "hey")
	}
	args[0] = "oh"
	return Cow(strings.Join(args, " "))
}

func Cow(text string) string {
	say, err := neocowsay.Say(&neocowsay.Cow{
		Phrase:      text,
		Type:        "default",
		BallonWidth: 40,
	})
	if err != nil {
		panic(err)
	}
	return say
}
