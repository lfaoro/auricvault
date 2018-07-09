package auricvault

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SetDebug() {
	log.SetLevel(logrus.DebugLevel)
}

func init() {
	if err := loadEnvVars(); err != nil {
		log.Fatal(err)
	}
}

// AuricVault API Implementation
//
// https://docs.auricvault.com/api-docs/
//

type Vault struct {
	url        string
	mtid       string
	mtidSecret string
	client     *http.Client
	request    request
}

type request struct {
	ID     int      `json:"id,omitempty"`
	Method string   `json:"method,omitempty"`
	Params []Params `json:"params,omitempty"`
}

type Params struct {
	ConfigurationID string    `json:"configurationId,omitempty"`
	Last4           string    `json:"last4,omitempty"`
	Mtid            string    `json:"mtid,omitempty"`
	PlaintextValue  string    `json:"plaintextValue,omitempty"`
	Retention       Retention `json:"retention,omitempty"`
	Segment         string    `json:"segment,omitempty"`
	UtcTimestamp    string    `json:"utcTimestamp,omitempty"`
	Token           string    `json:"token,omitempty"`
}

type Response struct {
	ID     int    `json:"id,omitempty"`
	Result Result `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Result struct {
	Version             string `json:"version,omitempty"`
	LastActionSucceeded int    `json:"lastActionSucceeded,omitempty"`
	Token               string `json:"token,omitempty"`
	PlaintextValue      string `json:"plaintextValue,omitempty"`
	ElapsedTime         string `json:"elapsedTime,omitempty"`
	TokenExists         string `json:"tokenExists,omitempty"`
	TokenCreatedDate    string `json:"tokenCreatedDate,omitempty"`
	LastAccessedDate    string `json:"lastAccessedDate,omitempty"`
	Segment             string `json:"segment,omitempty"`
	Retention           string `json:"retention,omitempty"`
	IsVaultEncrypted    string `json:"isVaultEncrypted,omitempty"`
}

// Retention enumerates the data retention options.
type Retention string

const (
	// BigYear data is kept approximately 14 months (14 * 31 days).
	BigYear Retention = "big-year"
	// Forever data is never deleted.
	Forever Retention = "forever"
)

func New(retention Retention) *Vault {
	return &Vault{
		url:        os.Getenv("AURIC_URL"),
		mtidSecret: os.Getenv("AURIC_MTID_SECRET"),
		client:     &http.Client{},
		request: request{
			ID:     0,
			Method: "",
			// Default params, used in all calls
			Params: []Params{
				{
					ConfigurationID: os.Getenv("AURIC_CONFIGURATION"),
					Mtid:            os.Getenv("AURIC_MTID"),
					Segment:         os.Getenv("AURIC_SEGMENT"),
					Retention:       retention,
				},
			},
		},
	}
}
