package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/jeffdupont/challenge3/jeffdupont/go/levenshtein"
	"github.com/spf13/viper"
)

var leads = map[int]lead{}
var maxFD, maxLD int

type lead struct {
	id          int
	first, last string
}

func init() {
	loadLeads()

	// max tolerances
	maxFD = 3
	maxLD = 1

	setConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		setConfig()
	})
}

func setConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err == nil {
		maxLD = viper.GetInt("max_distance_last_name")
	}
}

func main() {
	first := "GABS"
	last := "DE'ELIA"

	result := search(first, last)
	fmt.Println(result)

	for {
	}
}

func loadLeads() {
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

		leads[id] = lead{id, first, last}
	}
}

func search(first, last string) []lead {
	result := []lead{}

	for _, l := range leads {
		fd := levenshtein.Distance(first, l.first)
		ld := levenshtein.Distance(last, l.last)

		if fd > maxFD {
			continue
		}
		if ld > maxLD {
			continue
		}

		if (fd + ld) > (maxFD + (maxLD - 1)) {
			continue
		}

		result = append(result, l)
	}
	return result
}
