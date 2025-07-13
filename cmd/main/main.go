package main

import (
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types"
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
    
}
