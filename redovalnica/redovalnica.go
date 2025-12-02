package redovalnica

import "fmt"


type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent s to vpisno številko ne obstaja")
		return
	}
	if ocena < 0 || ocena > 10 {
		fmt.Println("Neveljavna ocena. Ocena mora biti med 0 in 10.")
		return
	}
	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1.0
	}
	if len(student.ocene) < 6 {
		return 0.0
	}
	var povprecje float64 = 0
	var velikost int = len(student.ocene)
	for i := 0; i < velikost; i++ {
		povprecje += float64(student.ocene[i])
	}

	povprecje /= float64(velikost)

	return povprecje
}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v", vpisna, student.ime, student.priimek, student.ocene)
		fmt.Println()
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		povprecnaOcena := povprecje(studenti, vpisna)
		fmt.Printf("%s %s: povprečna ocena: %f -> ", student.ime, student.priimek, povprecnaOcena)
		if povprecnaOcena >= 9 {
			fmt.Println("Odličen študent!")
		} else if povprecnaOcena < 6 {
			fmt.Println("Neuspešen študent")
		} else {
			fmt.Println("Povprečen študent")
		}
	}
}
