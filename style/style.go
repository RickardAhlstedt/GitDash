package style

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

var colors = map[string]tcell.Color{}

var currentTheme string

func SetTheme(name string, custom map[string]string) {
	colors = map[string]tcell.Color{}

	currentTheme = name

	if base, ok := themes[name]; ok {
		for k, v := range base {
			if c, ok := parseHex(v); ok {
				colors[k] = c
			}
		}
	}

	for key, hex := range custom {
		if c, ok := parseHex(hex); ok {
			colors[key] = c
		}
	}
}

func SetCustomColors(user map[string]string) {
	for key, hex := range user {
		if c, ok := parseHex(hex); ok {
			colors[key] = c
		}
	}
}

func Color(name string) tcell.Color {
	if c, ok := colors[name]; ok {
		return c
	}
	return tcell.ColorWhite
}

func parseHex(hex string) (tcell.Color, bool) {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	hexValue, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return tcell.ColorDefault, false
	}
	return tcell.NewHexColor(int32(hexValue)), true
}

func DumpTheme() {
	fmt.Printf("🎨 Theme: %s\n", currentTheme)
	if len(colors) == 0 {
		fmt.Println("⚠️  No colors loaded.")
		return
	}
	for name, color := range colors {
		fmt.Printf("  %-8s: #%06x\n", name, color.Hex())
	}
}
