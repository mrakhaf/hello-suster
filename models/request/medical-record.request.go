package request

type SavePatient struct {
	IdentityNumber        int    `json:"identityNumber" validate:"required,min_len=16,max_len=16"`
	Name                  string `json:"name" validate:"required,min=3,max=30"`
	BirthDate             string `json:"birthDate" validate:"required,datetime"`
	PhoneNumber           string `json:"phoneNumber" validate:"required,min=10,max=15,startswith=+62"`
	Gender                string `json:"gender" validate:"required,oneof=male female"`
	IdentityCardScanImage string `json:"identityCardScanImg" validate:"required,url"`
}

type SaveMedicalRecord struct {
	IdentityNumber int    `json:"identityNumber" validate:"required"`
	Symptoms       string `json:"symptoms" validate:"required"`
	Medications    string `json:"medications" validate:"required"`
}

type GetPatientsParam struct {
	IdentityNumber *int    `query:"identityNumber" validate:"omitempty"`
	Name           *string `query:"name" validate:"omitempty"`
	PhoneNumber    *string `query:"phoneNumber" validate:"omitempty"`
	CreatedAt      *string `query:"createdAt" validate:"omitempty"`
	Limit          *int    `query:"limit" validate:"omitempty"`
	Offset         *int    `query:"offset" validate:"omitempty"`
}

type GetMedicalRecordsParam struct {
	IdentityNumber *int    `query:"identityDetail.identityNumber" validate:"omitempty"`
	UserId         *string `query:"createdBy.userId" validate:"omitempty"`
	Nip            *string `query:"createdBy.nip" validate:"omitempty"`
	Limit          *int    `query:"limit" validate:"omitempty"`
	Offset         *int    `query:"offset" validate:"omitempty"`
	CreatedAt      *string `query:"createdAt" validate:"omitempty"`
}
