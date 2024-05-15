package entity

type User struct {
	ID        string
	NIP       int
	Name      string
	Password  string
	CreatedAt string
}

type UserNurse struct {
	ID                  string
	NIP                 int
	Name                string
	Password            *string
	IdentityCardScanImg string
	CreatedAt           string
}

type Patient struct {
	ID                  string
	IdentityNumber      int
	PhoneNumber         string
	Name                string
	BirthDate           string
	Gender              string
	IdentityCardScanImg string
	CreatedAt           string
}

type Users struct {
	ID        string `json:"userId"`
	NIP       int    `json:"nip"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
