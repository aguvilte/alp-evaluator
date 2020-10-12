package domain

type LicensePlate struct {
	Plate string `json:"license_plate" binding:"required"`
	Score string `json:"score" binding:"required"`
	Model string `json:"model" binding:"required"`
}
