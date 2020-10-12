package domain

import "encoding/json"

type LicensePlate struct {
	Plate string      `json:"plate" binding:"required"`
	Score json.Number `json:"score" binding:"required"`
	Model string      `json:"model" binding:"required"`
}
