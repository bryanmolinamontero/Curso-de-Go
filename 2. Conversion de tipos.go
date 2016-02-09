package main

import (
	"fmt" ; "strconv"
 )

func main(){



	// esta es una conversion de un entero a string
	edad := 22
	edad_str := strconv.Itoa(edad)

	fmt.Println("La edad es: " + edad_str )

	// pero una conversion de un string a un entero retorna dos valores
	//el primero el valor como tal, y el segundo si por si acaso hay algun error

	edad2 := "22"
	//edad_int,err := strconv.Atoi(edad2) -> esto es correcto pero como no estoy usando la variable err en ningun lado no permite compilar
	edad_int,_ := strconv.Atoi(edad2) // en vez de err pongo un _ que significa q esa variable no la quiero ni la voy a usar asi que la desecho

	fmt.Println(edad_int + 10)

}