package request

type SaveMedicalRecord struct {
	IdentityNumber        int    `json:"identityNumber" validate:"required"`
	Name                  string `json:"name" validate:"required"`
	BirthDate             string `json:"birthDate" validate:"required"`
	PhoneNumber           string `json:"phoneNumber" validate:"required"`
	Gender                string `json:"gender" validate:"required"`
	IdentityCardScanImage string `json:"identityCardScanImg" validate:"required,url"`
}

type GetMedicalRecordParam struct {
	IdentityNumber *int    `query:"identityNumber" validate:"omitempty"`
	Name           *string `query:"name" validate:"omitempty"`
	PhoneNumber    *string `query:"phoneNumber" validate:"omitempty"`
	CreatedAt      *string `query:"createdAt" validate:"omitempty"`
	Limit          *int    `query:"limit" validate:"omitempty"`
	Offset         *int    `query:"offset" validate:"omitempty"`
}
