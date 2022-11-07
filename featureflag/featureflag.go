package featureflag

import (
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

const (
	configFile  = "./flag-config.json"
	userDefault = "user-default-key"
)

func InitFeatureFlag() error {
	err := ffclient.Init(ffclient.Config{
		PollingInterval: 3 * time.Second,
		Retriever: &fileretriever.Retriever{
			Path: configFile,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func IsFeatureEnabled(flagKey string) bool {
	user := ffuser.NewUser(userDefault)
	hasFlag, _ := ffclient.BoolVariation(flagKey, user, false)
	return hasFlag
}
