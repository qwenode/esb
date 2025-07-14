package main

import (
    "encoding/json"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
    "github.com/qwenode/esb"
    "log"
)

func main() {
    client, err := elasticsearch.NewTypedClient(elasticsearch.Config{})
    if err != nil {
        log.Fatal(err)
    }
    log.Println(client)
    client.Search().Index("xx").Query(
        &types.Query{
            Terms: &types.TermsQuery{
                TermsQuery: map[string]types.TermsQueryField{
                    "aa": []string{"xxx"},
                },
            },
            Bool: &types.BoolQuery{
                Filter: []types.Query{
                    {
                        Term: map[string]types.TermQuery{
                            "abc": {
                                Value: "ww",
                            },
                        },
                    },
                },
                Must: []types.Query{
                    {
                        Terms: &types.TermsQuery{
                            TermsQuery: map[string]types.TermsQueryField{
                                "field": []string{"xxx"},
                            },
                        },
                    },
                },
            },
        },
    )
    query := esb.NewQuery(
        esb.Bool(
            esb.Filter(
                esb.Term("field", "aa"),
                esb.Term("field", "abc"),
            
            ),
            esb.Filter(esb.NumberRange("create").Gt(11).Build()),
        
        ),
    )
    marshal, err := json.Marshal(query)
    log.Println(string(marshal))
}
