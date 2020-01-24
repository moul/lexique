package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	err := generate()
	if err != nil {
		panic(err)
	}
}

type Entry struct {
	// FIXME: use other types (ints, bools, etc)
	Ortho         string
	Phon          string
	Lemme         string
	Cgram         string
	Genre         string
	Nombre        string
	Freqlemfilms2 string
	Freqlemlivres string
	Freqfilms2    string
	Freqlivres    string
	Infover       string
	Nbhomogr      string
	Nbhomoph      string
	Islem         string
	Nblettres     string
	Nbphons       string
	Cvcv          string
	PCvcv         string
	Voisorth      string
	Voisphon      string
	Puorth        string
	Puphon        string
	Syll          string
	Nbsyll        string
	CvCv          string
	Orthrenv      string
	Phonrenv      string
	Orthosyll     string
	Cgramortho    string
	Deflem        string
	Defobs        string
	Old20         string
	Pld20         string
	Morphoder     string
	Nbmorph       string
}

func generate() error {
	f, err := os.Open("./lexique.tsv")
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	entries := make([]Entry, len(lines)-1)
	for i, line := range lines[1:] {
		entry := Entry{
			Ortho:         line[0],
			Phon:          line[1],
			Lemme:         line[2],
			Cgram:         line[3],
			Genre:         line[4],
			Nombre:        line[5],
			Freqlemfilms2: line[6],
			Freqlemlivres: line[7],
			Freqfilms2:    line[8],
			Freqlivres:    line[9],
			Infover:       line[10],
			Nbhomogr:      line[11],
			Nbhomoph:      line[12],
			Islem:         line[13],
			Nblettres:     line[14],
			Nbphons:       line[15],
			Cvcv:          line[16],
			PCvcv:         line[17],
			Voisorth:      line[18],
			Voisphon:      line[19],
			Puorth:        line[20],
			Puphon:        line[21],
			Syll:          line[22],
			Nbsyll:        line[23],
			CvCv:          line[24],
			Orthrenv:      line[25],
			Phonrenv:      line[26],
			Orthosyll:     line[27],
			Cgramortho:    line[28],
			Deflem:        line[29],
			Defobs:        line[30],
			Old20:         line[31],
			Pld20:         line[32],
			Morphoder:     line[33],
			Nbmorph:       line[34],
		}
		entries[i] = entry
	}

	fmt.Println(len(entries))

	return nil
}
