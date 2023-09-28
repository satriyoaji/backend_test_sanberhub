package constant

import "errors"

type CampaignStatus string

const (
	CampaignStatusSaved       CampaignStatus = "saved"
	CampaignStatusBroadcasted CampaignStatus = "broadcasted"
	CampaignStatusScheduled   CampaignStatus = "scheduled"
)

var CampaignStatuses = []CampaignStatus{
	CampaignStatusSaved,
	CampaignStatusBroadcasted,
	CampaignStatusScheduled,
}

func ParseCampaignStatus(str string) (CampaignStatus, error) {
	for _, t := range CampaignStatuses {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
