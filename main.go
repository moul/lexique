package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/peterbourgon/ff/v2/ffcli"
	"moul.io/godev"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		rootFlagSet = flag.NewFlagSet("lexique", flag.ExitOnError)
		nthFlagSet  = flag.NewFlagSet("textctl repeat", flag.ExitOnError)
		nthN        = nthFlagSet.Int("n", 0, "nth")
	)

	nth := &ffcli.Command{
		Name:       "nth",
		ShortUsage: "lexique ntx [-n POS]",
		FlagSet:    nthFlagSet,
		Exec: func(_ context.Context, args []string) error {
			lexique, err := parseLexique()
			if err != nil {
				return err
			}
			picked := lexique[*nthN]
			fmt.Println(picked.Ortho)
			return nil
		},
	}

	random := &ffcli.Command{
		Name:       "random",
		ShortUsage: "lexique random",
		Exec: func(_ context.Context, args []string) error {
			lines, err := readLines()
			if err != nil {
				return err
			}
			picked := lines[rand.Intn(len(lines))]
			entry := parseEntry(picked)
			fmt.Println(entry.Ortho)
			return nil
		},
	}

	stats := &ffcli.Command{
		Name:       "stats",
		ShortUsage: "lexique stats",
		Exec: func(_ context.Context, args []string) error {
			lexique, err := parseLexique()
			if err != nil {
				return err
			}
			fmt.Println("total:", len(lexique))
			return nil
		},
	}

	dump := &ffcli.Command{
		Name:       "dump",
		ShortUsage: "lexique dump",
		Exec: func(_ context.Context, args []string) error {
			lexique, err := parseLexique()
			if err != nil {
				return err
			}
			fmt.Println(godev.PrettyJSON(lexique))
			return nil
		},
	}

	root := &ffcli.Command{
		ShortUsage:  "lexique [flags] <subcommand>",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{nth, random, stats, dump},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
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
	Freqlemfilms2 float64
	Freqlemlivres float64
	Freqfilms2    float64
	Freqlivres    float64
	Infover       string
	Nbhomogr      int
	Nbhomoph      int
	Islem         bool
	Nblettres     int
	Nbphons       int
	Cvcv          string
	PCvcv         string
	Voisorth      int
	Voisphon      int
	Puorth        int
	Puphon        int
	Syll          string
	Nbsyll        int
	CvCv          string
	Orthrenv      string
	Phonrenv      string
	Orthosyll     string
	Cgramortho    string
	Deflem        int
	Defobs        int
	Old20         float64
	Pld20         float64
	Morphoder     string
	Nbmorph       int
}

func parseEntry(line []string) Entry {
	return Entry{
		Ortho:         line[0],
		Phon:          line[1],
		Lemme:         line[2],
		Cgram:         line[3],
		Genre:         line[4],
		Nombre:        line[5],
		Freqlemfilms2: parseFloat(line[6]),
		Freqlemlivres: parseFloat(line[7]),
		Freqfilms2:    parseFloat(line[8]),
		Freqlivres:    parseFloat(line[9]),
		Infover:       line[10],
		Nbhomogr:      parseInt(line[11]),
		Nbhomoph:      parseInt(line[12]),
		Islem:         parseBool(line[13]),
		Nblettres:     parseInt(line[14]),
		Nbphons:       parseInt(line[15]),
		Cvcv:          line[16],
		PCvcv:         line[17],
		Voisorth:      parseInt(line[18]),
		Voisphon:      parseInt(line[19]),
		Puorth:        parseInt(line[20]),
		Puphon:        parseInt(line[21]),
		Syll:          line[22],
		Nbsyll:        parseInt(line[23]),
		CvCv:          line[24],
		Orthrenv:      line[25],
		Phonrenv:      line[26],
		Orthosyll:     line[27],
		Cgramortho:    line[28],
		Deflem:        parseInt(line[29]),
		Defobs:        parseInt(line[30]),
		Old20:         parseFloat(line[31]),
		Pld20:         parseFloat(line[32]),
		Morphoder:     line[33],
		Nbmorph:       parseInt(line[34]),
	}
}

func readLines() ([][]string, error) {
	f, err := os.Open("./lexique.tsv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}

func parseLexique() ([]Entry, error) {
	lines, err := readLines()
	if err != nil {
		return nil, err
	}

	entries := make([]Entry, len(lines)-1)
	for i, line := range lines[1:] {
		entry := parseEntry(line)
		entries[i] = entry
	}
	return entries, nil
}

func parseBool(input string) bool {
	return input == "1"
}

func parseInt(input string) int {
	out, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return out
}

func parseFloat(input string) float64 {
	out, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return -1
	}
	return out
}
