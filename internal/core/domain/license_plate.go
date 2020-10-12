package domain

type LicensePlate struct {
	Plate string  `json:"plate" binding:"required"`
	Score float64 `json:"score" binding:"required"`
}
