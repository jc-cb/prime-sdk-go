package main

import (
	"encoding/json"
	"log"
	"os"

	prime "github.com/coinbase-samples/prime-sdk-go"
)

func main() {

	credentials := &prime.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("PRIME_CREDENTIALS")), credentials); err != nil {
		log.Fatalf("unable to init credentials %v:", err)
	}

}
