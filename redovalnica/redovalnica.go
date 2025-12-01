package redovalnica

import "fmt"

// Student je struct, sestavljen iz imena, priimka in seznama ocen.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno v seznam ocen študentu, podanem z vpisno številko
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

	e.Ocene = append(studenti[vpisnaStevilka].Ocene, ocena)
	studenti[vpisnaStevilka] = e
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	_, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}

	var total = 0.0
	for i := range studenti[vpisnaStevilka].Ocene {
		total += float64(studenti[vpisnaStevilka].Ocene[i])
	}
	return total / float64(len(studenti[vpisnaStevilka].Ocene))
}

// IzpisVsehOcen izpiše celotno redovalnico, torej seznam ocen vsakega študenta
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, s.Ime, s.Priimek, s.Ocene)
	}
}

// IzpisiKoncniUspeh uspeh vsakega študenta glede na povprečje in število ocen
func IzpisiKoncniUspeh(studenti map[string]Student, minStOcen int) {
	for vpisna, s := range studenti {
		var avg = povprecje(studenti, vpisna)
		var uspeh string
		if len(studenti[vpisna].Ocene) < minStOcen {
			uspeh = "Neuspešen študent!"
		} else if avg >= 9 {
			uspeh = "Odličen študent!"
		} else if avg < 6 {
			uspeh = "Neuspešen študent"
		} else {
			uspeh = "Povprečen študent"
		}

		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", s.Ime, s.Priimek, avg, uspeh)
	}
}
