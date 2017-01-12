package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sysnotifsPurge/helpers"
)

// ratio in percent between available size and total size
const ratioMin = 20

// retrieve size
const statsURL = "https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com/_cluster/stats"

// retrive list of index name
const _statsURL = "https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com/_stats"

// delete index by name
const indexURL = "https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com/"

// struct to map json object Indices from _stats request
type indicesStruct struct {
	Indices map[string]struct {
		Primaries string `json:"primaries"`
	} `json:"indices"`
}

// struct to map json object Nodes from stats request
type nodesStruct struct {
	Nodes struct {
		Fs struct {
			TotalInBytes     int `json:"total_in_bytes"`
			AvailableInBytes int `json:"available_in_bytes"`
		} `json:"fs"`
	} `json:"nodes"`
}

//CheckRatio function : request size to calculate ratio and send ratio in the response
func CheckRatio(w http.ResponseWriter, r *http.Request) {

	fmt.Println("CheckRatio")

	var ratio = 0
	var data nodesStruct

	// request size while ratio is less tha ratioMin
	for ratio < ratioMin {

		// send request to amazone to retrieve total size and available size
		resp, err := http.Get(statsURL)
		if err != nil {
			// handle error
			helpers.SetTextResponse(w, 400, "error")
			return
		}

		helpers.GetBody(resp, &data)

		resp.Body.Close()

		// calculate the ratio in percent
		ratio = int(float64(data.Nodes.Fs.AvailableInBytes) / float64(data.Nodes.Fs.TotalInBytes) * 100)

		if ratio < ratioMin {
			// try to purge if ratio less than ratioMin
			purgeIndex(searchOldIndex())

		}
	}

	helpers.SetTextResponse(w, 200, strconv.Itoa(ratio))

}

//searchOldIndex: search the oldest index by date
func searchOldIndex() string {

	// send request to list index
	resp, err := http.Get(_statsURL)
	if err != nil {
		// handle error

	}
	defer resp.Body.Close()

	var data indicesStruct

	helpers.GetBody(resp, &data)

	// retrieve the name of the indices, ignore kibana indice
	var keys []string
	for key := range data.Indices {
		if !strings.Contains(key, "kibana") {
			keys = append(keys, key)
		}

	}
	sort.Strings(keys)

	fmt.Println("old index " + keys[0])

	// return the oldest indicex
	return keys[0]
	//	fmt.Printf("Results: %v\n", data.Indices)

}

// purgeIndex: delete index by name
func purgeIndex(indexName string) {

	fmt.Println("purgeIndex " + indexName)

	// send request to delete an index
	req, err := http.NewRequest("DELETE", indexURL+indexName, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error

	}
	defer resp.Body.Close()

}
