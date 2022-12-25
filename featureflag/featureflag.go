package featureflag

import (
	"time"

	"github.com/go-feature-flag/config"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func InitFeatureFlag(config *config.Config) error {
	err := ffclient.Init(ffclient.Config{
		PollingInterval: 3 * time.Second,
		Retriever: &fileretriever.Retriever{
			Path: config.FeatureFlag.ConfigFilePath,
		},
		// Retriever: &httpretriever.Retriever{
		// 	URL:     "https://raw.githubusercontent.com/zikrykr/go-feature-flags/main/flag-config.json",
		// 	Timeout: 5 * time.Second,
		// },
	})
	if err != nil {
		return err
	}

	return nil
}

func IsFeatureEnabled(flagKey string) bool {
	cfg, err := config.LoadConfig()
	if err != nil {
		return false
	}
	user := ffuser.NewUser(cfg.FeatureFlag.UserDefault)
	hasFlag, _ := ffclient.BoolVariation(flagKey, user, false)
	return hasFlag
}
