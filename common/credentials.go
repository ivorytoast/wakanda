package common

import (
	"github.com/spf13/viper"
	"log"
	"time"
	"wakanda/config"
)

var (
    Configuration config.ApplicationConfig
)

type APIKey struct {
	ID string
	Secret string
	OAuth string
	PolygonKeyID string
}

func Credentials() *APIKey {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if Configuration.Authorization.UserPrivateKey == "YOUR_PRIVATE_KEY_GOES_HERE" || Configuration.Authorization.UserPrivateKey == "" {
		log.Fatal("Before sending any requests you must put your PRIVATE KEY into the config.yml!")
	}

	log.Printf("public key is             : [%s]", Configuration.Authorization.UserPublicKey)
	log.Printf("private key is            : [%s]", Configuration.Authorization.UserPrivateKey)
	log.Printf("auth key is               : [%s]", Configuration.Authorization.UserAuthKey)
	log.Printf("polygon key is            : [%s]", Configuration.Authorization.PolygonApiKey)

	log.Printf("paper base url is         : [%s]", Configuration.Setup.PaperBaseUrl)
	log.Printf("paper data url is         : [%s]", Configuration.Setup.PaperDataUrl)
	log.Printf("paper api version is      : [%s]", Configuration.Setup.PaperApiVersion)
	log.Printf("paper request timeout is  : [%s]", Configuration.Setup.PaperClientRequestTimeout * time.Second)

	return &APIKey {
		ID: Configuration.Authorization.UserPublicKey,
		Secret: Configuration.Authorization.UserPrivateKey,
		OAuth: Configuration.Authorization.UserAuthKey,
		PolygonKeyID: Configuration.Authorization.PolygonApiKey,
	}
}