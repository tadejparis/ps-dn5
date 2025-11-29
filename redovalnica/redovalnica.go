package redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena > 10 && ocena < 0 {
		fmt.Println("Ocena mora biti v intervalu [0,10].")
		return
	}
	e, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Student s to vpisno stevilko ne obstaja.")
		return
	}

	e.ocene = append(studenti[vpisnaStevilka].ocene, ocena)
	studenti[vpisnaStevilka] = e
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	_, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	var total = 0.0
	for i := range studenti[vpisnaStevilka].ocene {
		total += float64(studenti[vpisnaStevilka].ocene[i])
	}
	return total / float64(len(studenti[vpisnaStevilka].ocene))
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, s.ime, s.priimek, s.ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, s := range studenti {
		var avg = povprecje(studenti, vpisna)
		var uspeh string
		if avg >= 9 {
			uspeh = "Odličen študent!"
		} else if avg < 6 {
			uspeh = "Neuspešen študent"
		} else {
			uspeh = "Povprečen študent"
		}

		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", s.ime, s.priimek, avg, uspeh)
	}
}
