package main

import "strings"

func normalize(t string) string {
	t = strings.ReplaceAll(t, " ", "+")
	t = strings.ReplaceAll(t, ":", "%3A")
	return t
}
