package models

type Device struct {
	ID   string `json:"id" jwt:"id"`
	Name string `json:"name" jwt:"name"`

	LastMeterReadingsTime *string `json:"last_meter_readings_time" jwt:"-"`
}
