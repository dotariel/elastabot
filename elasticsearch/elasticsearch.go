package elasticsearch

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Connection is a
type Connection struct {
	*elastic.Client
	URL string
}

type TermQueryOptions struct {
	Index     string
	Sort      string
	Ascending bool
	Query     string
}

// Connect returns a Connection to an Elasticsearch endpoint
func Connect(url string) (*Connection, error) {
	c, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}

	return &Connection{Client: c, URL: url}, nil
}

// Ping queries the ES cluster as a simple health check
func (con *Connection) Ping() bool {
	ctx := context.Background()

	_, code, err := con.Client.Ping(con.URL).Do(ctx)
	if err != nil {
		log.Error(err)
		return false
	}

	return code == 200
}

func (con *Connection) TermQuery(option TermQueryOptions) interface{} {
	return nil
}

// TODO: Function to write "silence" document to ES

// TODO: TermQuery

// def findNewestAlert(es, index, recentMinutes, name):
//   alert = None
//   details = None
//   res = es.search(index=index, size=1, sort='alert_time:desc', q='alert_time:[now-' + str(recentMinutes) + 'm TO now] AND !rule_name:Deadman* AND rule_name:' + name , _source_include=['rule_name', 'match_body'])
//   if res['hits']['hits']:
//     alert = res['hits']['hits'][0]['_source']['rule_name']
//     details = res['hits']['hits'][0]['_source']['match_body']
// 	return alert, details

// if res['hits']['hits']:
//     alert = res['hits']['hits'][0]['_source']['rule_name']
//     details = res['hits']['hits'][0]['_source']['match_body']
// 	return alert, details

// FindNewestAlert returns the
// func (con *Connection) FindNewestAlert(index string) bool {
// 	query := elastic.NewTermQuery("rule_name", "*")
// 	searchResult, err := con.Search().Index(index).Query(query).Do(context.Background())

// 	// Handle this
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(searchResult)

// 	return false
// }
