package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/foobaz/ansi"
	"github.com/nalum/dfglv/structs"
)

var showVersion bool
var httpPort int
var legendsFile string = "/home/nalum/d2/r2/legends/00125/01/01/data.xml"
var historyFile string = "/home/nalum/d2/r2/history/00125/01/01/data.xml"
var legends structs.Legends

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
	flag.BoolVar(&showVersion, "v", false, "Show the version of dfglv and check if there is a new version available.")
	flag.IntVar(&httpPort, "port", 8080, "This is the http port to use.")
	flag.IntVar(&httpPort, "p", 8080, "This is the http port to use.")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println("Dwarf Fortress GO Legends View")
		fmt.Println("Version         : 0.1.0")
		fmt.Println("Source          : github.com/nalum/dfglv")
		fmt.Println("Author          : Luke Mallon <nalum@lukemallon.com>")
		fmt.Println("Current Release : 0.1.0") // TODO: check the current release
		return
	}

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

	http.HandleFunc("/", root)
	http.Handle("/artifacts/", legends.Artifacts)
	http.Handle("/entities/", legends.Entities)
	http.Handle("/entity-populations/", legends.EntityPopulations)
	http.Handle("/historical-eras/", legends.HistoricalEras)
	http.Handle("/historical-event-collections/", legends.HistoricalEventCollections)
	http.Handle("/historical-events/", legends.HistoricalEvents)
	http.Handle("/historical-figures/", legends.HistoricalFigures)
	http.Handle("/regions/", legends.Regions)
	http.Handle("/sites/", legends.Sites)
	http.Handle("/underground-regions/", legends.UndergroundRegions)
	http.Handle("/world-constructions/", legends.WorldConstructions)
	http.Handle("/legends/", legends)
	log.Printf("Starting HTTP Server on Port %v", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", httpPort), nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "Artifacts: <a href=\"/artifacts\">%v</a><br>\n", len(legends.Artifacts.Artifacts))
	fmt.Fprintf(w, "Entities: <a href=\"/entities\">%v</a><br>\n", len(legends.Entities.Entities))
	fmt.Fprintf(w, "Entity Populations: <a href=\"/entity-populations\">%v</a><br>\n", len(legends.EntityPopulations.EntityPopulations))
	fmt.Fprintf(w, "Historical Eras: <a href=\"/historical-eras\">%v</a><br>\n", len(legends.HistoricalEras.HistoricalEras))
	fmt.Fprintf(w, "Historical Event Collections: <a href=\"/historical-event-collections\">%v</a><br>\n", len(legends.HistoricalEventCollections.HistoricalEventCollections))
	fmt.Fprintf(w, "Historical Events: <a href=\"/historical-events\">%v</a><br>\n", len(legends.HistoricalEvents.HistoricalEvents))
	fmt.Fprintf(w, "Historical Figures: <a href=\"/historical-figures\">%v</a><br>\n", len(legends.HistoricalFigures.HistoricalFigures))
	fmt.Fprintf(w, "Regions: <a href=\"/regions\">%v</a><br>\n", len(legends.Regions.Regions))
	fmt.Fprintf(w, "Sites: <a href=\"/sites\">%v</a><br>\n", len(legends.Sites.Sites))
	fmt.Fprintf(w, "Underground Regions: <a href=\"/underground-regions\">%v</a><br>\n", len(legends.UndergroundRegions.UndergroundRegions))
	fmt.Fprintf(w, "World Constructions: <a href=\"/world-constructions\">%v</a><br>\n", len(legends.WorldConstructions.WorldConstructions))
	log.Println(r.URL)
}
