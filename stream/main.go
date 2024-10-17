package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	text := "Hola, ¿en qué puedo ayudarte hoy?\n\n- Si tienes alguna pregunta sobre programación, estaré encantado de responder.\n- También puedo ayudarte a resolver problemas relacionados con datos, bases de datos, o a escribir código en diferentes lenguajes.\n\nDéjame saber en qué estás trabajando y qué necesitas."
	StreamingSim(text)
}
func StreamingSim(text string) {
	chunks := strings.Split(text, "\n")
	for _, chunk := range chunks {
		for i, letter := range chunk {
			if i == len(chunk)-1 {
				fmt.Printf("%c\n", letter)
			} else {
				fmt.Printf("%c", letter)
			}
			time.Sleep(10000000 * 2)
		}
	}
}
