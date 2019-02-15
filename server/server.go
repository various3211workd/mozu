package server

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"../blockchain"
)

// json format
type Person struct {
	Prev		[]byte "json:'prev'"
	Data		[]byte "json:'data'"	
	Hash		[]byte "json:'hash'"
	PoW			string "json:pow"
}

/*
	handler function
*/
func handler(w http.ResponseWriter, r *http.Request) {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")

	// create json data
	data := Person{}
	data.Prev = bc.Blocks[len(bc.Blocks)-1].PrevBlockHash
	data.Data = bc.Blocks[len(bc.Blocks)-1].Data
	data.Hash = bc.Blocks[len(bc.Blocks)-1].Hash
	pow := blockchain.NewProofOfWork(bc.Blocks[len(bc.Blocks)-1])
	data.PoW = strconv.FormatBool(pow.Validate())
	
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// response
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

/*
	Run funciton
*/
func Run() {
	port := 8080

	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
