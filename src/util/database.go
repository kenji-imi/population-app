package util

import (
	"encoding/csv"
	"fmt"
	"github.com/kenji-imi/population-app/src/xo"
	"io"
	"os"
	"strconv"
)

//util.WriteCSVData(db, "data/population.csv")
func WriteCSVData(db xo.XODB, filename string) {
	var err error
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない！
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if len(rec) != 9 {
			fmt.Printf("invalid format. %d != 9, %v\n", len(rec), rec)
			continue
		}
		//fmt.Println(rec)
		//for i, v := range rec {
		//	fmt.Printf("%d: %v\n", i, v)
		//}

		id, err := strconv.Atoi(rec[0])
		eraYear, err := strconv.Atoi(rec[4])
		year, err := strconv.Atoi(rec[5])
		population, err := strconv.Atoi(rec[6])
		malePopulation, err := strconv.Atoi(rec[7])
		femalePopulation, err := strconv.Atoi(rec[8])

		m := xo.Population{
			ID:               id,
			PrefCode:         rec[1],
			PrefName:         rec[2],
			EraName:          rec[3],
			EraYear:          eraYear,
			Year:             year,
			Population:       population,
			MalePopulation:   malePopulation,
			FemalePopulation: femalePopulation,
		}
		if err := m.Insert(db); err != nil {
			panic(err)
		}
	}
}
