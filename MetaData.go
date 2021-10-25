package ElasticSearchClient

type metaData struct {
	Name string `json:"name"`
	Decimals string `json:"decimals"`
	Description string `json:"description"`
	Image string `json:"image"`
	Properties interface{}
}

type Property1 struct {
	Size string `json:"size"`
	Color string `json:"color"`
}

type Property2 struct {
	Weight string `json:"weight"`
	Length string `json:"length"`
}