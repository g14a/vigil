package service

type IndexInfo struct {
	Name      string `json:"index"`
	Health    string `json:"health"`
	DocsCount string `json:"docs.count"`
	Storage   string `json:"store.size"`
}
