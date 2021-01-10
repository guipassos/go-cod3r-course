package main

import "math"

// Iniciando com letra maiúscula é PUBLICO (visivel dentro e fora do pacote)!
// Iniciando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)!

// Por exemplo...
// Ponto - gerará algo público
// ponto ou (iniciando com underline) _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) { // função privada disponível apenas dentro do pacote
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 { // função pública disponível fora do pacote
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
