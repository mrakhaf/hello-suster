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
