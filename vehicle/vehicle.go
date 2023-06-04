package vehicle

type Capacity struct {
	Value int    `json:"value" binding:"required"`
	Unit  string `json:"unit" binding:"required"`
}

type Vehicle struct {
	Vin          string   `json:"vin" binding:"required"`
	Manufacturer string   `json:"manufacturer" binding:"required"`
	Model        string   `json:"model" binding:"required"`
	Year         int      `json:"year" binding:"required"`
	Color        string   `json:"color" binding:"required"`
	Capacity     Capacity `json:"capacity" binding:"required"`
	LicensePlate string   `json:"license_plate"`
}

type Update struct {
	Manufacturer string   `json:"manufacturer" binding:"required"`
	Model        string   `json:"model" binding:"required"`
	Year         int      `json:"year" binding:"required"`
	Color        string   `json:"color" binding:"required"`
	Capacity     Capacity `json:"capacity" binding:"required"`
	LicensePlate string   `json:"license_plate"`
}
