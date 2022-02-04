# Interfaces en GO

Una de las principales implementaciones de la composición es el uso de interfaces. Una interfaz define un comportamiento de un tipo. Una de las interfaces que más se usan en la biblioteca estándar de Go es fmt.Stringer:
```
type Stringer interface {
    String() string
}
```
La primera línea de código define un type llamado Stringer. Luego indica que es una interfaz. Al igual cuando se define una struct, Go utiliza llaves ({}) para rodear la definición de la interfaz. En comparación con la definición de estructuras, solo defi_nimos el compo_rtamiento de las interfaces; es decir, “qué puede hacer este tipo”.

En el caso de la interfaz Stringer, el único comportamiento es el método String(). El método no toma argumentos y muestra una cadena.

A continuación, veremos código que tiene el comportamiento fmt.Stringer:
```
package main

import "fmt"

type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	fmt.Println(a.String())
}
```
Lo primero que hacemos es crear un nuevo tipo llamado Article. Este tipo tiene un campo Title y un campo Author, y ambos son de la cadena de tipo de datos:

```
type Article struct {
	Title string
	Author string
}
```
A continuación, definimos un method llamado String en el tipo Article. El método String mostrará una cadena que representa el tipo Article:

```
func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}
```
A continuación, en nuestra función main, creamos una instancia del tipo Article y la asignamos a la variable llamada a. Proporcionamos los valores de "Understanding Interfaces in Go" para el campo Title y "Sammy Shark"para el campo Author:

```
a := Article{
	Title: "Understanding Interfaces in Go",
	Author: "Sammy Shark",
}
```
A continuación, imprimimos el resultado del método String invocando fmt.PrintIn y pasando el resultado de la invocación del método a.String():

```
fmt.Println(a.String())
```
Después de ejecutar el programa, verá el siguiente resultado:
```
The "Understanding Interfaces in Go" article was written by Sammy Shark.
```
Hasta ahora no usamos una interfaz, pero creamos un tipo que tuvo un comportamiento. Ese comportamiento coincidió con la interfaz fmt.Stringer. A continuación, veremos la forma de usar ese comportamiento para hacer que nuestro código sea más reutilizable.

## Definir una interfaz
Ahora que nuestro tipo está definido con el comportamiento deseado, podemos ver la forma de usar ese comportamiento.

Antes de hacer eso, sin embargo, veremos lo que deberíamos hacer si deseáramos invocar el método String desde el tipo Article en una función:
```
package main

import "fmt"

type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)
}

func Print(a Article) {
	fmt.Println(a.String())
}
```
En este código, añadimos una nueva función llamada Print que toma un Article como argumento. Observe que lo único que la función Print hace es invocar el método String. Debido a esto, podríamos definir una interfaz que se pasaría a la función:

```
package main

import "fmt"

type Article struct {
	Title string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
```
Aquí creamos una interfaz llamada Stringer:

```
type Stringer interface {
	String() string
}
```
La interfaz Stringer solo tiene un método, llamado String(), que muestra una string. Un método es una función especial que tiene ámbito en un tipo específico en Go. A diferencia de una función, un método solo puede invocarse desde la instancia del tipo sobre el que se definió.

A continuación actualizamos la firma del método Print para tomar un Stringer y no un tipo concreto de Article. Debido a que el compilador reconoce que una interfaz Stringer define el método String, solo aceptará los tipos que también tienen el método String.

Ahora podemos usar el método Print con cualquier cosa que se adecue a la interfaz Stringer. Crearemos otro tipo para demostrar esto:

```
package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
```
Ahora, añadimos un segundo tipo llamado Book. También tiene el método String definido. Esto significa que además se adecua a la interfaz Stringer. Debido a esto, podemos enviarlo también a nuestra función Print:
```
The "Understanding Interfaces in Go" article was written by Sammy Shark.
The "All About Go" book was written by Jenny Dolphin. It has 25 pages.
```
Hasta ahora, demostramos la forma de usar una interfaz única. Sin embargo, para una interfaz puede haber más de un comportamiento definido. A continuación, veremos la forma en que podemos hacer que nuestras interfaces sean más versátiles declarando más métodos.

## Varios comportamientos en una interfaz
Uno de los objetivos principales de escribir código en Go es escribir tipos pequeños y concisos, componerlos de modo que conformen tipos más grandes y complejos. Sucede lo mismo cuando se componen interfaces. Para ver la forma de crear una interfaz, primero comenzaremos definiendo solo una interfaz. Definiremos dos formas, Circle y Square, y ambas definirán un método llamado Area. Este método mostrará el área geométrica de sus respectivas formas:

```
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

type Sizer interface {
	Area() float64
}

func main() {
	c := Circle{Radius: 10}
	s := Square{Height: 10, Width: 5}

	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n", l)
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}
```
Debido a que cada tipo declara el método Area, podemos crear una interfaz que defina ese comportamiento. Crearemos la siguiente interfaz Sizer:

```
type Sizer interface {
	Area() float64
}
```
A continuación definiremos una función llamada Less que toma dos Sizer y muestra el más pequeño:

```
func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}
```
Observe que no solo aceptamos ambos argumentos como el tipo Sizer, sino también mostramos el resultado como Sizer. Esto significa que ya no mostramos un Square ni un Circle, sino la interfaz Sizer.

Por último, imprimimos el que tenía el área más pequeña:
```
{Width:5 Height:10} is the smallest
```
A continuación, añadiremos otro comportamiento a cada tipo. Esta vez, añadiremos el método String() que muestra una cadena. Esto satisfará la interfaz fmt.Stringer:

```
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func main() {
	c := Circle{Radius: 10}
	PrintArea(c)

	s := Square{Height: 10, Width: 5}
	PrintArea(s)

	l := Less(c, s)
	fmt.Printf("%v is the smallest\n", l)

}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
```
Debido a que los tipos Circle y Square implementan los métodos Area y String, podemos crear otra interfaz para describir ese conjunto más amplio de comportamientos. Para hacer esto, crearemos una interfaz llamada Shaper. Compondremos lo siguiente con las interfaces Sizer y fmt.Stringer:

```
type Shaper interface {
	Sizer
	fmt.Stringer
}
```
**Nota:** Se considera que corresponde intentar dar nombre a su interfaz con finalización en er; fmt.Stringer e io.Writer son algunos ejemplos. Por eso, dimos a nuestra interfaz el nombre Shaper y no Shape.

Ahora podemos crear una función llamada PrintArea que toma Shaper como argumento. Esto significa que podemos invocar ambos métodos en el valor pasado para los métodos Area y String:

```
func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
```
Si ejecutamos el programa, veremos el siguiente resultado:
```
area of Circle {Radius: 10.00} is 314.16
area of Square {Width: 5.00, Height: 10.00} is 50.00
Square {Width: 5.00, Height: 10.00} is the smallest
```
Acabamos de ver la forma de podemos crear interfaces más pequeñas y hacerlas más grandes según sea necesario. Aunque podríamos haber comenzado con la interfaz más grande y haberla pasado a todas nuestras funciones, se considera mejor enviar solo la interfaz más pequeña a una función que sea necesaria. Esto normalmente da como resultado un código más claro, ya que cualquier cosa que acepte una interfaz específica más pequeña solo tiene la intención de funcionar con ese comportamiento definido.

Por ejemplo, si pasamos Shaper a la función Less podemos suponer que invocará los métodos Area y String. Sin embargo, ya que solo queremos invocar el método Area, hace que la función Less sea clara porque sabemos que solo podemos invocar el método Area de cualquier argumento que se le pase.
