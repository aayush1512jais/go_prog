package model

type Medicine struct {
	MedicineID int    `json:"medicine_id"`
	Name       string `json:"name,omitempty"`
	Company    string `json:"company,omitempty"`
	IsExpired  bool   `json:"is_expired,omitempty"`
}
