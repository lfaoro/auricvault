package auricvault

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// LoadEnv loads .env file variables.
func loadEnvVars() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Unable to load environment variables: %v", err)
	}
	envVars := []string{
		"AURIC_URL",
		"AURIC_CONFIGURATION",
		"AURIC_MTID",
		"AURIC_MTID_SECRET",
		"AURIC_SEGMENT",
	}
	for _, s := range envVars {
		_, yes := os.LookupEnv(s)
		if !yes {
			return fmt.Errorf("Missing required environment variable: %v", s)
		}
	}
	return nil
}

// getHMAC calculates the hmac for the X-VAULT-HMAC header.
func getHMAC(data string) string {
	_, yes := os.LookupEnv("AURIC_MTID_SECRET")
	if !yes {
		log.Fatal("Auric environment variables required.")
	}
	secret := []byte(os.Getenv("AURIC_MTID_SECRET"))
	// Create a new HMAC by defining the hash type and the key (as byte array)
	hash := hmac.New(sha512.New, secret)
	hash.Write([]byte(data))
	sha := hex.EncodeToString(hash.Sum(nil))
	return strings.ToLower(sha)
}

func setHeaders(hmacData []byte) http.Header {
	h := make(http.Header)
	h.Add("Content-Type", "application/json")
	h.Add("X-VAULT-HMAC", getHMAC(string(hmacData)))
	return h
}

func getTime() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}

func (v *Vault) doRequest() (response *Response, err error) {
	v.request.ID = 1
	v.request.Params[0].UtcTimestamp = getTime()

	body, err := json.Marshal(v.request)
	if err != nil {
		return nil, fmt.Errorf("unable to marshall the request body: %v", err)
	}
	log.Debug("request: ", string(body))
	req, err := http.NewRequest("POST", v.url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}
	req.Header = setHeaders(body)
	res, err := v.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform http call: %v", err)
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}
	log.Debug("response body: ", string(d))
	if err = json.Unmarshal(d, &response); err != nil {
		return nil, fmt.Errorf("unable parse response body: %v", err)
	}
	if response.Error != "" {
		return nil, fmt.Errorf("auric's message: %v", response.Error)
	}
	return response, nil
}
