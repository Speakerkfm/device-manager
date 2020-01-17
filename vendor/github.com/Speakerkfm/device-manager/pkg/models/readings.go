package models

import "time"

type DeviceReadings struct {
	MeterReadingsTime string  `json:"meter_readings_time"`
	Temperature       float64 `json:"temperature"`

	Created time.Time `json:"-"`
}
