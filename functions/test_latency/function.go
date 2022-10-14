package test_latency

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"
)

// AnalyticsEvent is the payload of an Analytics log event.
type AnalyticsEvent struct {
	EventDimensions []EventDimensions `json:"eventDim"`
	UserDimensions  UserDimensions    `json:"userDim"`
}

// EventDimensions holds Analytics event dimensions.
type EventDimensions struct {
	Name                    string      `json:"name"`
	Date                    string      `json:"date"`
	TimestampMicros         string      `json:"timestampMicros"`
	PreviousTimestampMicros string      `json:"previousTimestampMicros"`
	Params                  interface{} `json:"params"`
}

type UserDimensions struct {
	AppInfo AppInfo `json:"appInfo"`
}

type AppInfo struct {
	AppInstanceId string `json:"appInstanceId"`
	AppPlatform   string `json:"appPlatform"`
}

func TestLatency(ctx context.Context, e AnalyticsEvent) error {
	if len(e.EventDimensions) == 0 {
		return errors.New("EventDimensions don't exist")
	}
	timestampMicros, _ := strconv.Atoi(e.EventDimensions[0].TimestampMicros)
	eventTime := time.Unix(0, int64(timestampMicros*1000))
	//log.Printf("Event time: %v", eventTime)
	log.Printf("Latency:%vs", int(time.Since(eventTime).Seconds()))
	// simulate trigger workflow grpc call
	time.Sleep(time.Second)
	return nil
}
