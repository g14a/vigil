package service

import (
	"encoding/json"
	"net/http"
	"vigil/client"
	"vigil/utils"
)

func GetDistinctIndexes() []IndexInfo {
	_, err := client.ClusterInfo(client.ESClient)
	utils.CheckError(err)

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

}
