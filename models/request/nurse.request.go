package request

type RegisterNurse struct {
	Name                  string `json:"name" validate:"required,min=5,max=50"`
	NIP                   int    `json:"nip" validate:"required,int_len=13"`
	IdentityCardScanImage string `json:"identityCardScanImg" validate:"required,url"`
}

type AccessNurse struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
}
