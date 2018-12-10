package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/jeffdupont/challenge3/go/levenshtein"
)

var leads = map[int]lead{}

type lead struct {
	first, last string
}

func main() {
	f, err := os.Open("leads.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id, _ := strconv.Atoi(record[0])
		first := record[1]
		last := record[2]

		leads[id] = lead{first, last}
	}

	firstName := "GABS"
	lastName := "DE'ELIA"

	// max tolerances
	maxFD := 3
	maxLD := 1
	maxTotal := maxFD + (maxLD - 1)

	for _, l := range leads {
		fd := levenshtein.Distance(firstName, l.first)
		ld := levenshtein.Distance(lastName, l.last)

		if fd > maxFD {
			continue
		}
		if ld > maxLD {
			continue
		}
		if (fd + ld) > maxTotal {
			continue
		}

		fmt.Println(fd, ld)
	}
}
