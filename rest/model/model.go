package model

type FeatureFlag struct {
	FlagName string `json:"flag_name"`
	IsActive bool   `json:"is_active"`
}

type FeatureFlagConfig struct {
	Percentage int  `json:"percentage"`
	True       bool `json:"true"`
	Default    bool `json:"default"`
}
