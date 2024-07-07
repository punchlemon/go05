package piscine

import "ft"

func PrintWordsTables(a []string) {
	for _, s := range a {
		for _, r := range s {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}