package main

import (
	"fmt"
	m "math" //pode-se usar um alias como por exemplo: m "math" e assim chamar apenas m.Pow ao invés de math.Pow
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido automaticamente pelo compilador

	// forma reduzida de criar uma var (também é a mais usada)
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
