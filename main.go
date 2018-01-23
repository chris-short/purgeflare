package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	execute      = flag.Bool("execute", false, "Will actually purge cache; use caution because your origin will get hammered.")
	emailAsFlag  = flag.String("email", "user@example.com", "set X-Auth-Email")
	keyAsFlag    = flag.String("key", "c2547eb745079dac9320b638f5e225cf483cc5cfdda41", "set X-Auth-Key")
	cfzoneAsFlag = flag.String("cfzone", "023e105f4ecef8ad9ca31a8372d0c353", "Specify Cloudflare Zone ID")

//	emailAsEnv   = os.Getenv("AuthEmail")
//	keyAsEnv     = os.Getenv("AuthKey")
//	cfzoneAsEnv  = os.Getenv("CFZone")
)

func main() {

	flag.Parse()

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
	req.Header.Set("X-Auth-Email", *emailAsFlag)
	req.Header.Set("X-Auth-Key", *keyAsFlag)
	req.Header.Set("Content-Type", "application/json")

	// If this isn't live, just return the headers and exit.
	if !*execute {
		fmt.Println("Request Headers:")
		fmt.Printf("%+v\n", req.Header)
		fmt.Println("Add --execute to send request to Cloudflare")
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	rp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle err
	}

	fmt.Printf("Result: %v - %v", resp.StatusCode, string(rp))
}
