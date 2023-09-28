package constant

import "errors"

type ScheduleStatus string

const (
	ScheduleStatusCreated    ScheduleStatus = "created"
	ScheduleStatusProcessing ScheduleStatus = "processing"
	ScheduleStatusSuccess    ScheduleStatus = "success"
	ScheduleStatusFailed     ScheduleStatus = "failed"
)

var ScheduleStatuses = []ScheduleStatus{
	ScheduleStatusCreated,
	ScheduleStatusProcessing,
	ScheduleStatusSuccess,
	ScheduleStatusFailed,
}

func ParseScheduleStatus(str string) (ScheduleStatus, error) {
	for _, t := range ScheduleStatuses {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
