package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tadejparis/ps-dn5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	// testiranje

	var R = make(map[string]redovalnica.Student)

	R["63220001"] = redovalnica.Student{Ime: "Eva", Priimek: "Novak", Ocene: []int{8, 9, 7}}
	R["63220002"] = redovalnica.Student{Ime: "Janez", Priimek: "Kovač", Ocene: []int{10, 8, 9}}
	R["63220003"] = redovalnica.Student{Ime: "Bob", Priimek: "Primer", Ocene: []int{7, 6, 8}}

	// cmd

	cmd := &cli.Command{
		Name:  "ps-dn5",
		Usage: "Uporabi funkcije iz paketa redovalnica",
		Commands: []*cli.Command{
			{
				Name:  "izpisi-ocene",
				Usage: "izpiše vse ocene",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "stOcen",
						Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
						Value: 3,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.IzpisVsehOcen(R)
					return nil
				},
			},
			{
				Name:  "izpisi-uspeh",
				Usage: "izpiše končne uspehe študentov",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.IzpisiKoncniUspeh(R, cmd.Int("stOcen"))
					return nil
				},
			},
			{
				Name:  "dodaj-oceno",
				Usage: "doda oceno v redovalnici za študenta",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "minOcena",
						Usage: "najmanjša pozitivna ocena",
						Value: 5,
					},
					&cli.IntFlag{
						Name:  "maxOcena",
						Usage: "največja možna ocena",
						Value: 10,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					vpisna_st := cmd.Args().First()
					ocena, conv_err := strconv.Atoi(cmd.Args().Get(1))
					if conv_err != nil {
						fmt.Println("ocena mora biti število!")
					}
					if ocena < cmd.Int("minOcena") || ocena > cmd.Int("maxOcena") {
						fmt.Println("Prevelika ali premajhna ocena")
					}
					redovalnica.DodajOceno(R, vpisna_st, ocena)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
