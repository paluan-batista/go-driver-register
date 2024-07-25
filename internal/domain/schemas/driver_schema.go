package schemas

type CreateDriver struct {
	Name      string `json:"name"`
	LicenseID string `json:"license_id"`
}
