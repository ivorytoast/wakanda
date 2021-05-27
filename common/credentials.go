package common

import (
	"fmt"
	"os"
	"sync"
)

var (
	once sync.Once
	key *APIKey
)

/*
Polygon refers to the outside market data company.
Account must be setup to use Polygon -- is not by default
 */
const (
	EnvApiKeyID = "APCA_API_KEY_ID"
	EnvApiSecretKey = "APCA_API_SECRET_KEY"
	EnvApiOAuth = "APCA_API_OAUTH"
	EnvPolygonKeyID = "POLY_API_KEY_ID"
)

type APIKey struct {
	ID string
	Secret string
	OAuth string
	PolygonKeyID string
}

func Credentials() *APIKey {
	s := os.Getenv(EnvPolygonKeyID)

	var polygonKeyID string
	if s != "" {
		polygonKeyID = s
	} else {
		polygonKeyID = os.Getenv(EnvApiKeyID)
	}

	fmt.Println("polygonKeyID: " + polygonKeyID)

	return &APIKey {
		ID: os.Getenv(EnvApiKeyID),
		PolygonKeyID: polygonKeyID,
		Secret: os.Getenv(EnvApiSecretKey),
		OAuth: os.Getenv(EnvApiOAuth),
	}
}

func TestCredentials() *APIKey {
	s := os.Getenv(EnvPolygonKeyID)

	var polygonKeyID string
	if s != "" {
		polygonKeyID = s
	} else {
		polygonKeyID = os.Getenv(EnvApiKeyID)
	}

	fmt.Println("polygonKeyID: " + polygonKeyID)

	return &APIKey{
		ID: "PKX3UF3428XKJKNL3G6U",
		PolygonKeyID: polygonKeyID,
		Secret: "Dwm841cYo68pCKr4R5STBwRRKvAlZM7JTEibVD4z",
		OAuth: os.Getenv(EnvApiOAuth),
	}
}