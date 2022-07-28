package client

import (
	"time"
)

//go:generate mockgen -destination mock_logger_client.go -package client github.com/tokopedia/topads-logging-client LoggerService

type LoggerService interface {
	GenerateNonce() string
	GenerateCreateTime() string
	ConvertToCreateTime(t time.Time) string
	ConvertToClientTimestamp(t time.Time) string
	PublishEvent(data WrapperEvent)
	PublishDashboardEventData(data DashboardEvent)
	PublishDisplayBrowseData(displayBrowseMessage DisplayBrowseRequestEvent)
	PublishDisplayData(data DisplayRequestEvent)
}

