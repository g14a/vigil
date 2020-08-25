package service

type IndexInfo struct {
	Name      string `json:"index"`
	Health    string `json:"health"`
	DocsCount string `json:"docs.count"`
	Storage   string `json:"store.size"`
}

type NodeStats struct {
	IP              string `json:"ip"`
	Name            string `json:"name"`
	HeapPercent     string `json:"heap.percent"`
	HeapCurrent     string `json:"heap.current"`
	HeapMax         string `json:"heap.max"`
	RamPercent      string `json:"ram.percent"`
	RamCurrent      string `json:"ram.current"`
	RamMax          string `json:"ram.max"`
	NodeRole        string `json:"node.role"`
	Master          string `json:"master"`
	CPUPercent      string `json:"cpu"`
	DiskUsedPercent string `json:"disk.user_percent"`
	DiskUsed        string `json:"disk.used"`
	DiskTotal       string `json:"disk.total"`
}
