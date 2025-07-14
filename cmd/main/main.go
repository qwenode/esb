package main

import (
    "context"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/qwenode/esb"
    "log"
)

func main() {
    client, err := elasticsearch.NewTypedClient(
        elasticsearch.Config{
            Addresses: []string{"http://10.1.1.6:9200"},
            Username:  "elastic",
            Password:  "Elastic_eQ4N7i",
        },
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Println(client)
    // client.Indices.Create("xx").Do(context.Background())
    // client.Index("xx").Document(
    //     map[string]any{
    //         "hello": "world",
    //     },
    // ).Do(context.Background())
    do, err := client.Search().Index("xx").Query(
        esb.NewQuery(esb.Term("hello", "world1")),
    
    ).Do(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    log.Println(do)
    for _, hit := range do.Hits.Hits {
        log.Println(string(hit.Source_))
    }
}
