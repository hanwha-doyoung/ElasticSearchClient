package ElasticSearchClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type ESClient struct {
	*elastic.Client
}

// 아마 read from config
const NFT = "nft"

// load의 기능이 필요한가?
// new는 최초 서비스 기동 시 되고, 그 뒤로는 dial을 해서 있는 client를 가져와 사용하는 것 아닌가?
func NewElasticSearchClient(url ...string) (*ESClient, error) {
	client, err := elastic.NewClient(
		// TODO read username, password from config
		elastic.SetBasicAuth("elastic", "qwer1234!@#$"),
		elastic.SetURL(url...),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new client error %v", err))

	}
	return &ESClient{
		client,
	}, nil
}

func (esc *ESClient) CreateIndex(index string) error {
	service, err := esc.Client.CreateIndex(index).Do(context.TODO())
	if err != nil {
		return errors.New(fmt.Sprintf("create index error %v", err))

	}
	if !service.Acknowledged {
		return errors.New(fmt.Sprintf("index.acknowledged got false"))
	}
	return nil
}

// tenant from ctx?
func (esc *ESClient) AddDocument(tenant string, contractId string, tokenId string, doc metaData) error {
	jsonBody, err := json.Marshal(&doc)
	if err != nil {
		return errors.New(fmt.Sprintf("marshal error %v", err))
	}
	var bodyMap map[string]interface{}
	err = json.Unmarshal(jsonBody, &bodyMap)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal error %v", err))
	}

	docId := tenant + contractId + tokenId
	_, err = esc.Index().Index(NFT).Id(docId).BodyJson(bodyMap).Do(context.TODO())
	if err != nil {
		return errors.New(fmt.Sprintf("add document error %v", err))
	}

	return nil
}

func (esc *ESClient) GetDocument(tenant string, contractId string, tokenId string) (*metaData, error) {
	docId := tenant + contractId + tokenId
	res, err := esc.Get().Index(NFT).Id(docId).Do(context.TODO())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get document error %v", err))
	}

	var meta metaData
	err = json.Unmarshal(res.Source, &meta)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unmarshal error %v", err))
	}

	return &meta, nil
}

// index info 필요한가?
// function name -> GetIndexInfo?
func (esc *ESClient) GetIndexMapping(contractAddress string) error {
	res, err := esc.Client.IndexGet().Index(contractAddress).Do(context.TODO())
	if err != nil {
		return errors.New(fmt.Sprintf("get document error %v", err))
	}
	fmt.Printf("%v\n", res[contractAddress])
	fmt.Printf("%v\n", res[contractAddress].Mappings)
	fmt.Printf("%v\n", res[contractAddress].Settings)
	fmt.Printf("%v\n", res[contractAddress].Aliases)
	fmt.Printf("%v\n", res[contractAddress].Warmers)
	return nil
}
