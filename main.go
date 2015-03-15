package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/foobaz/ansi"
	"github.com/nalum/dfglv/structs"
	"io/ioutil"
)

var showVersion bool
var legendsFile string = "/home/nalum/d2/r2/legends/00125/01/01/data.xml"
var historyFile string = "/home/nalum/d2/r2/history/00125/01/01/data.xml"
var configFile string

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
	flag.StringVar(&legendsFile, "legends", legendsFile, "Set the location of the config file.")
	flag.StringVar(&historyFile, "history", historyFile, "Set the location of the config file.")
	flag.StringVar(&configFile, "config", configFile, "Set the location of the config file.")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println("Version: 0.1.0")
		return
	}

	var legends structs.Legends
	legendsRaw, err := ioutil.ReadFile(legendsFile)
	var legendsBytesProcessed []byte

	if err != nil {
		panic(err)
	}

	for i := range legendsRaw {
		if legendsRaw[i] > 127 {
			for j := range ansi.CP437toUTF8[legendsRaw[i]-128] {
				legendsBytesProcessed = append(legendsBytesProcessed, ansi.CP437toUTF8[legendsRaw[i]-128][j])
			}
		} else {
			legendsBytesProcessed = append(legendsBytesProcessed, legendsRaw[i])
		}
	}

	err = xml.Unmarshal(legendsBytesProcessed, &legends)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Artifacts: %v\n", len(legends.Artifacts.Artifacts))
	fmt.Printf("Entities: %v\n", len(legends.Entities.Entities))
	fmt.Printf("Entity Populations: %v\n", len(legends.EntityPopulations.EntityPopulations))
	fmt.Printf("Historical Eras: %v\n", len(legends.HistoricalEras.HistoricalEras))
	fmt.Printf("Regions: %v\n", len(legends.Regions.Regions))
	fmt.Printf("Sites: %v\n", len(legends.Sites.Sites))
}
