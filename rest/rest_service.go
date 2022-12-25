package rest

import (
	"context"
	"errors"

	"github.com/go-feature-flag/config"
	"github.com/go-feature-flag/rest/model"
	"github.com/go-feature-flag/rest/request"
	"github.com/go-feature-flag/rest/utils"
)

type RestFeatureFlagItf interface {
	CreateFeatureFlag(ctx context.Context, data request.CreateFeatureFlagReq) (model.FeatureFlag, error)
	UpdateFeatureFlag(ctx context.Context, data request.CreateFeatureFlagReq) (model.FeatureFlag, error)
	GetFeatureFlags(ctx context.Context) ([]model.FeatureFlag, error)
	DeleteFeatureFlag(ctx context.Context, flagName string) error
}

type restFeatureFlag struct {
	Config *config.Config
}

func NewRestFeatureFlag(cfg *config.Config) RestFeatureFlagItf {
	return &restFeatureFlag{
		Config: cfg,
	}
}

func (c *restFeatureFlag) CreateFeatureFlag(ctx context.Context, data request.CreateFeatureFlagReq) (model.FeatureFlag, error) {
	var (
		result         model.FeatureFlag
		configFilePath = c.Config.FeatureFlag.ConfigFilePath
	)

	featureFlagConfig, err := utils.GetFeatureFlagConfig(configFilePath)
	if err != nil {
		return result, err
	}

	if _, ok := featureFlagConfig[data.FlagName]; ok {
		return result, errors.New("feature flag already exist")
	}

	featureFlagConfig[data.FlagName] = model.FeatureFlagConfig{
		Percentage: 100,
		True:       data.IsActive,
		Default:    false,
	}

	if err := utils.OverwriteFeatureFlagConfig(configFilePath, featureFlagConfig); err != nil {
		return result, err
	}

	result.FlagName = data.FlagName
	result.IsActive = data.IsActive

	return result, nil
}

func (c *restFeatureFlag) UpdateFeatureFlag(ctx context.Context, data request.CreateFeatureFlagReq) (model.FeatureFlag, error) {
	var (
		result         model.FeatureFlag
		configFilePath = c.Config.FeatureFlag.ConfigFilePath
	)

	featureFlagConfig, err := utils.GetFeatureFlagConfig(configFilePath)
	if err != nil {
		return result, err
	}

	if _, ok := featureFlagConfig[data.FlagName]; !ok {
		return result, errors.New("feature flag doesn't exist")
	}

	featureFlagConfig[data.FlagName] = model.FeatureFlagConfig{
		Percentage: 100,
		True:       data.IsActive,
		Default:    false,
	}

	if err := utils.OverwriteFeatureFlagConfig(configFilePath, featureFlagConfig); err != nil {
		return result, err
	}

	result.FlagName = data.FlagName
	result.IsActive = data.IsActive

	return result, nil
}

func (c *restFeatureFlag) GetFeatureFlags(ctx context.Context) ([]model.FeatureFlag, error) {
	var (
		result         []model.FeatureFlag
		configFilePath = c.Config.FeatureFlag.ConfigFilePath
	)

	featureFlagConfig, err := utils.GetFeatureFlagConfig(configFilePath)
	if err != nil {
		return result, err
	}

	featureFlagKeys := utils.GetFeatureFlagKeys(featureFlagConfig)

	for _, key := range featureFlagKeys {
		if _, ok := featureFlagConfig[key]; !ok {
			continue
		}

		featureFlagConfig := featureFlagConfig[key]

		featureFlag := model.FeatureFlag{
			FlagName: key,
			IsActive: featureFlagConfig.True,
		}

		result = append(result, featureFlag)
	}

	return result, nil
}

func (c *restFeatureFlag) DeleteFeatureFlag(ctx context.Context, flagName string) error {
	var (
		configFilePath = c.Config.FeatureFlag.ConfigFilePath
	)

	featureFlagConfig, err := utils.GetFeatureFlagConfig(configFilePath)
	if err != nil {
		return err
	}

	if _, ok := featureFlagConfig[flagName]; !ok {
		return errors.New("feature flag doesn't exist")
	}

	delete(featureFlagConfig, flagName)

	if err := utils.OverwriteFeatureFlagConfig(configFilePath, featureFlagConfig); err != nil {
		return err
	}

	return nil
}
