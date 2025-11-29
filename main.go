package main

import (
	"github.com/tadejparis/ps-dn5/redovalnica"
)

func main() {

	var studenti = make(map[string]redovalnica.Student)

	studenti["63220001"] = redovalnica.Student{Ime: "Eva", Priimek: "Novak", Ocene: []int{8, 9, 7}}
	studenti["63220002"] = redovalnica.Student{Ime: "Janez", Priimek: "Kovač", Ocene: []int{10, 8, 9}}
	studenti["63220003"] = redovalnica.Student{Ime: "Bob", Priimek: "Primer", Ocene: []int{7, 6, 8}}

	redovalnica.IzpisVsehOcen(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti)

	redovalnica.DodajOceno(studenti, "63220001", 5)

	redovalnica.IzpisVsehOcen(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti)

}
