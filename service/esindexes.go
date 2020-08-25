package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"vigil/client"
	"vigil/utils"
)

func GetIndiceInfo() []IndexInfo {

	url := utils.GetUrlFromString("http://localhost:9200/_cat/indices?h=index,health,docs.count,store.size")

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

		IndiceInfo := make([]IndexInfo, 0)
		if err := json.NewDecoder(resp.Body).Decode(&IndiceInfo); err != nil {
			utils.CheckError(err)
			return nil
		}

		return IndiceInfo
	}

	return nil
}

func GetDocumentsFromIndex(index string) {
	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*",
			},
		},
		"size": 10,
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		utils.CheckError(err)
	}

	fmt.Println(buf.String())

	es := client.ESClient

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	utils.CheckError(err)

	fmt.Println(res.String())
}
