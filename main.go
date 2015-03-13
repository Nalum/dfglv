package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/nalum/dfglv/structs"
	"io/ioutil"
	"strconv"
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

	if err != nil {
		panic(err)
	}

	legendsString := string(legendsRaw)
	legendsStringConv := strconv.Quote(legendsString)
	fmt.Println(legendsStringConv[100000:100500])
	legendsByte := []byte(legendsStringConv)

	err = xml.Unmarshal(legendsByte, &legends)

	if err != nil {
		panic(err)
	}

	fmt.Println(legends)
}
