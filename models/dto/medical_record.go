package dto

type PatientResponse struct {
	IdentityNumber int    `json:"identityNumber"`
	PhoneNumber    string `json:"phoneNumber"`
	Name           string `json:"name"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	CreatedAt      string `json:"createdAt"`
}

type MedicalRecordResponse struct {
	IdentityDetail IdentityDetail `json:"identityDetail"`
	Symptoms       string         `json:"symptoms"`
	Medications    string         `json:"medications"`
	CreatedAt      string         `json:"createdAt"`
	CreatedBy      CreatedBy      `json:"createdBy"`
}

type IdentityDetail struct {
	IdentityNumber      int    `json:"identityNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	Name                string `json:"name"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

type CreatedBy struct {
	UserId string `json:"userId"`
	Nip    int    `json:"nip"`
	Name   string `json:"name"`
}
