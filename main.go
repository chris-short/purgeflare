package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func main() {

	type Payload struct {
		PurgeEverything bool `json:"purge_everything"`
	}

	data := Payload{
	// fill struct
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("DELETE", "https://api.cloudflare.com/client/v4/zones/89a96b8b091a805efea9e0eeee8e6e1b/purge_cache", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("X-Auth-Email", "user@example.com")
	req.Header.Set("X-Auth-Key", "c2547eb745079dac9320b638f5e225cf483cc5cfdda41")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
