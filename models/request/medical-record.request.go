package request

type SaveMedicalRecord struct {
	IdentityNumber        int    `json:"identityNumber" validate:"required"`
	Name                  string `json:"name" validate:"required"`
	BirthDate             string `json:"birthDate" validate:"required"`
	PhoneNumber           string `json:"phoneNumber" validate:"required"`
	Gender                string `json:"gender" validate:"required"`
	IdentityCardScanImage string `json:"identityCardScanImg" validate:"required,url"`
}
