package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("hogehoge", enparenthis))
	fmt.Println(expand("$hoge", enparenthis))
	fmt.Println(expand("hoge$fuga", enparenthis))
	fmt.Println(expand("piyo $hoge fuga", enparenthis))
	fmt.Println(expand("piyo $hoge fuga $foo bar", enparenthis))
	fmt.Println(expand("piyo $hoge fuga $foobar hogera", enparenthis))
}

func enparenthis(s string) string {
	return "(" + s + ")"
}

// expand search tokens followed by $ and replace those with given func
func expand(s string, f func(string) string) string {
	tokens, ok := parse(s); if !ok {
		return s
	}
	for _, token := range tokens {
		s = strings.Replace(s, "$"+token, f(token), -1)
	}
	return s
}

// parse parses string and get tokens followed by $
// if there is no tokens, bool returns false
func parse(arg string) ([]string, bool) {
	var tokens []string
	strs := strings.Split(arg, "$")
	if len(strs) <= 1 {
		return nil, false
	}

	// dump first element because a first element is invalid
	strs = strs[1:]

	for _, str := range strs {
		words := strings.Split(str, " ")
		if len(words[0]) > 0 {
			tokens = append(tokens, words[0])
		}
	}
	if len(tokens) < 0 {
		return nil, false
	}
	return tokens, true
}
