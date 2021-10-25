package ElasticSearchClient

import (
	"log"
	"testing"
)

const NODE1URL = "http://15.164.215.206:9200"
const NODE2URL = "http://13.209.41.122:9200"
const NODE3URL = "http://54.180.87.65:9200"

const tenant = "hanwha"
const contractId = "ShoesNFT"
const metaMapping = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":1
	},
      "mappings" : {
         "properties" : {
            "decimals" : {
               "type" : "text",
               "fields" : {
                  "keyword" : {
                     "type" : "keyword",
                     "ignore_above" : 256
                  }
               }
            },
            "image" : {
               "fields" : {
                  "keyword" : {
                     "type" : "keyword",
                     "ignore_above" : 256
                  }
               },
               "type" : "text"
            },
            "description" : {
               "type" : "text",
               "fields" : {
                  "keyword" : {
                     "type" : "keyword",
                     "ignore_above" : 256
                  }
               }
            },
            "size" : {
               "type" : "text",
               "fields" : {
                  "keyword" : {
                     "type" : "keyword",
                     "ignore_above" : 256
                  }
               }
            },
            "name" : {
               "fields" : {
                  "keyword" : {
                     "ignore_above" : 256,
                     "type" : "keyword"
                  }
               },
               "type" : "text"
            }
         }
      }
}
`
const testMapping = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":2
	},
	"mappings":{
		"properties":{
			"user":{
				"type":"keyword"
			},
			"message":{
				"type":"text",
				"store": true,
				"fielddata": true
			},
			"tags":{
				"type":"keyword"
			},
			"location":{
				"type":"geo_point"
			},
			"suggest_field":{
				"type":"completion"
			}
		}
	}
}
`

func TestNewElasticSearchClient(t *testing.T) {
	_, err := NewElasticSearchClient(NODE1URL, NODE2URL, NODE3URL)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateIndex(t *testing.T) {
	esclient, err := NewElasticSearchClient(NODE1URL, NODE2URL, NODE3URL)
	if err != nil {
		t.Error(err)
	}

	err = esclient.CreateIndex(NFT)
	if err != nil {
		t.Error(err)
	}
}

func TestAddDocument(t *testing.T) {
	esclient, err := NewElasticSearchClient(NODE1URL, NODE2URL, NODE3URL)
	if err != nil {
		t.Error(err)
	}

	tokenId := "2"

	body := metaData{
		Name:        "test2",
		Decimals:    "18",
		Description: "test",
		Image:       "cidxxxx",
		Properties: Property2{
			Weight: "10",
			Length: "5",
		},
	}
	//body2 := `{
	//	Name : "test3",
	//	Decimals: "18",
	//	Description: "test",
	//	Image: "cidxxxx",
	//	Color: "large",
	//}

	err = esclient.AddDocument(tenant, contractId, tokenId, body)
	if err != nil {
		return
	}
}

func TestGetDocument(t *testing.T) {
	esclient, err := NewElasticSearchClient(NODE1URL, NODE2URL, NODE3URL)
	if err != nil {
		t.Error(err)
	}

	tokenId := "2"

	meta, err := esclient.GetDocument(tenant, contractId, tokenId)
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v", meta)
}

func TestGetIndexMapping(t *testing.T) {
	esclient, err := NewElasticSearchClient(NODE1URL, NODE2URL, NODE3URL)
	if err != nil {
		t.Error(err)
	}

	err = esclient.GetIndexMapping(NFT)
	if err != nil {
		t.Error(err)
	}
}
