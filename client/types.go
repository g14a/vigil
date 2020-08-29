package client

type NodeInfoResponse struct {
	Name        string `json:"name"`
	ClusterName string `json:"cluster_name"`
	ClusterUUID string `json:"cluster_uuid"`
	Version     struct {
		Number                    string `json:"number"`
		BuildFavour               string `json:"build_favour"`
		BuildType                 string `json:"build_type"`
		BuildHash                 string `json:"build_hash"`
		BuildDate                 string `json:"build_date"`
		BuildSnapshot             bool   `json:"build_snapshot"`
		LuceneVersion             string `json:"lucene_version"`
		MinWireCompatibleVersion  string `json:"minimum_wire_compatible_version"`
		MinIndexCompatibleVersion string `json:"minimum_index_compatible_version"`
	} `json:"version"`
	Tagline string `json:"tagline"`
}
