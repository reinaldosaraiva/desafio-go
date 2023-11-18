package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Quote struct {
    Bid string `json:"bid"`
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	req,err:= http.NewRequestWithContext(ctx,http.MethodGet,"http://localhost:8080/cotacao",nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var quote Quote
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		log.Fatal(err)
	}
	content := []byte("Dolar: "+quote.Bid)
	if err := ioutil.WriteFile("cotacao.txt", content, 0644); err != nil {
		log.Fatal(err)
	}

	
}