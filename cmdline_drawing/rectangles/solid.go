package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	for row := 0; row < height; row++ {

		// Zeischne eine Zeile auf die Konsole
		for col := 0; col < width; col++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
