package request

type CreateFeatureFlagReq struct {
	FlagName string `json:"flag_name"`
	IsActive bool   `json:"is_active"`
}
