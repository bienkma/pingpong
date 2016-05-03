package main

import (
	"encoding/json"
	"github.com/bienkma/pingpong/lib"
	"github.com/olivere/elastic"
	"os"
	"strconv"
	"time"
)

type Configuration struct {
	HostName       string
	NodeSameRack   string
	NodeDiffRack   string
	NodeDiffRegion string
	ElasticURL     []string
}

type ndp struct {
	HostName       string `json: "hostname"`
	CostSameRack   int64  `json: "same_rack"`
	CostDiffRack   int64  `json: "diff_rack"`
	CostDiffRegion int64  `json: "diff_region"`
	Timestamp      string `json: "timestamp"`
}

func main() {
	// Define configuration file
	file, err_opening := os.Open("config.json")
	if err_opening != nil {
		pingpong.Log("config.json file not found")
		panic(err_opening)
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		pingpong.Log("config.json file wrong format!..")
		panic(err)
	}
	// close define Configuration
	// Create connection
	client, err := elastic.NewClient(
		elastic.SetURL(configuration.ElasticURL[0]),
	)
	if err != nil {
		pingpong.Log("Can't connection elasticsearch!...")
		panic(err)
	}

	for {
		timestamp := time.Now().Format(time.RFC3339Nano)
		// create doc for management host
		doc := ndp{
			HostName:       configuration.HostName,
			CostSameRack:   pingpong.Ping(configuration.NodeSameRack),
			CostDiffRack:   pingpong.Ping(configuration.NodeDiffRack),
			CostDiffRegion: pingpong.Ping(configuration.NodeDiffRegion),
			Timestamp:      timestamp,
		}
		// Add document to the index
		t2 := strconv.FormatInt(time.Now().UnixNano(), 10)
		_, err = client.Index().
			Index("ndp").
			Type(configuration.HostName).
			Id(t2).
			BodyJson(doc).
			Do()
		if err != nil {
			pingpong.Log("Can't create index in elasticsearch. Maybe duplicate index or problem connection elastic server!.. ")
			panic(err)
		}
	}
}
