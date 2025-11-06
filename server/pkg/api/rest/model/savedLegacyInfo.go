package model

import (
	"time"

	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/legacy"
)

type SavedLegacyInfo struct {
	ConnectionID string    `gorm:"primaryKey" json:"connection_id" validate:"required"`
	LegacyData   string    `gorm:"column:legacy_data" json:"legacy_data" validate:"required"`
	Status       string    `gorm:"column:status" json:"status"`
	SavedTime    time.Time `gorm:"column:saved_time" json:"saved_time"`
}

type LegacyInfoList struct {
	Servers []legacy.LegacySoftware `json:"servers" validate:"required"`
}
