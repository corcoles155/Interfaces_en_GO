package main

import "fmt"

type servivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

/* Humano */
type mujer struct {
	edad int
	altura float32
	peso float32
	respisando bool
	pensando bool
	comiendo bool
	esMujer bool
	estaVivo bool
}

//hombre hereda de la estructura mujer
type hombre struct {
	mujer
}

/* Funciones de la estructura mujer,
para implementar de la interface humano basta con definir todas sus funciones */

func (m *mujer) respirar()  {
	m.respisando = true
}

func (m *mujer) comer()  {
	m.comiendo = true
}

func (m *mujer) pensar()  {
	m.pensando = true
}

func (m *mujer) sexo() string  {
	if m.esMujer {
		return  "Mujer"
	}
	return "Hombre"
}

func (m *mujer) estaVivo() bool  {
	return m.estaVivo
}

func humanosRespirando(hu humano)  {
	hu.respirar()
	fmt.Printf("Soy un/a %s y estoy respirando \n", hu.sexo())
}

/* Animal */
type perro struct {
	respisando bool
	comiendo bool
	carnivoro bool
	estaVivo bool
}

func (p *perro) respirar()  {
	p.respisando = true
}

func (p *perro) comer()  {
	p.comiendo = true
}

func (p *perro) esCarnivoro bool {
	return p.carnivoro
}

func (p *perro) estaVivo() bool  {
	return p.estaVivo
}

func animalesRespirando(an animal)  {
	an.respirar()
	fmt.Printf("Soy un animal y estoy respirando \n")
}

func animalesCarnivoros(an animal) int {
	if an.esCarnivoro() {
		return 1
	}
	return 0
}

func estoyVivo(v servivo) bool  {
	return v.estaVivo()
}

func main()  {
	/* Humano */
	pedro := new(hombre)
	pedro.estaVivo = true
	humanosRespirando(pedro)

	maria := new(mujer)
	maria.esMujer = true
	maria.estaVivo = true
	humanosRespirando(maria)

	/* Animal */
	totalCarnivoros := 0
	dogo := new(perro)
	dogo.carnivoro = true
	dogo.estaVivo = true
	animalesRespirando(dogo)
	totalCarnivoros=+animalesCarnivoros(dogo)

	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)

	/* Servivo */
	fmt.Printf("Estoy vivo = %t \n", estoyVivo(dogo))
}	
