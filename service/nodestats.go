package service

import (
	"encoding/json"
	"net/http"
	"vigil/utils"
)

// Gives out an array on stats depending on the number of nodes present in the cluster.
// Gives a single entity for a standalone node
func GetNodesStats() (*[]NodeStats, error) {
	url := utils.GetUrlFromString(
		"http://localhost:9200/_cat/nodes?h=ip,name,heap.percent," +
			"heap.current,heap.max,ram.percent," +
			"ram.current,ram.max,node.role,master," +
			"cpu,load_1m,load_5m,load_15m," +
			"disk.used_percent,disk.used,disk.total",
	)

	h := http.Header{}
	h.Set("Accept", "application/json")

	req := http.Request{
		Method: "",
		URL:    &url,
		Header: h,
	}

	c := http.Client{}
	resp, err := c.Do(&req)
	utils.CheckError(err)

	if resp != nil {
		defer resp.Body.Close()

		var nodeStat []NodeStats
		if err := json.NewDecoder(resp.Body).Decode(&nodeStat); err != nil {
			utils.CheckError(err)
			return nil, err
		}

		return &nodeStat, nil
	}
	
	return nil, nil
}
