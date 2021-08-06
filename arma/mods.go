package arma

import (
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/endigma/garma/common"
)

var (
	ActivePreset Preset
)

var (
	ErrNoPreset = errors.New("no presets")
)

func GetPresets() []Preset {
	var presets []Preset

	files, err := ioutil.ReadDir("./ignore/presets")
	common.CheckErr(err, "failed to read presets")

	for _, f := range files {
		presets = append(presets, Preset{
			Name:     f.Name(),
			Filename: f.Name(),
		})
	}

	return presets
}

// func SavePreset()

func GetMods() []Mod {
	// Open the file
	csvfile, err := os.Open("mods.csv")
	common.CheckErr(err, "failed to parse csv")

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	rgx, _ := regexp.Compile(`\d+$`)

	var mods []Mod

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		common.CheckErr(err, "failed to parse csv")

		mods = append(mods, Mod{
			Name:       record[0],
			Enabled:    false,
			WorkshopID: rgx.FindString(record[1]),
		})
	}

	return mods
}
