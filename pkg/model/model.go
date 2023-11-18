package model

import "time"

type CacheBreak bool

const BreakCache CacheBreak = true

type GlobalAlert struct {
	AlertKey       string      `json:"AlertKey"`
	AlertHtml      string      `json:"AlertHtml"`
	AlertTimestamp time.Time   `json:"AlertTimestamp"`
	AlertLink      string      `json:"AlertLink"`
	AlertLevel     int32       `json:"AlertLevel"`
	AlertType      int32       `json:"AlertType"`
	StreamInfo     *StreamInfo `json:"StreamInfo"`
}

type StreamInfo struct {
	ChannelName string `json:"ChannelName"`
}
