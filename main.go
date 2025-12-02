package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ela-lab/ps_dn5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Izpis in obdelava ocen študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen za pozitivno oceno",
				Value: 3,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša dovoljena ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja dovoljena ocena",
				Value: 10,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			fmt.Println("Nastavitve:")
			fmt.Println("  Zahtevano št. ocen za pozitivno:", stOcen)
			fmt.Println("  Dovoljene ocene:", minOcena, "-", maxOcena)
			fmt.Println()

			// Poženemo logiko
			runRedovalnica(minOcena, maxOcena)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runRedovalnica(minOcena int, maxOcena int) {

	var student1 redovalnica.Student
	var student2 redovalnica.Student

	student1.Ime = "Ela"
	student1.Priimek = "Roš"
	student1.Ocene = []int{9, 10, 8, 6, 7, 7}

	student2.Ime = "Ana"
	student2.Priimek = "Novak"
	student2.Ocene = []int{6, 7, 8}

	slovar := make(map[string]redovalnica.Student)

	slovar["012"] = student1
	slovar["011"] = student2
	slovar["003"] = redovalnica.Student{"Janez", "Novak", []int{10, 6, 6, 6, 7, 8}}

	ocena := 10
	if ocena >= minOcena && ocena <= maxOcena {
		redovalnica.DodajOceno(slovar, "011", ocena)
	} else {
		fmt.Println("Ocena", ocena, "je izven dovoljenega območja.")
	}

	redovalnica.IzpisRedovalnice(slovar)
	redovalnica.IzpisiKoncniUspeh(slovar)
}
