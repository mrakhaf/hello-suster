package request

type RegisterNurse struct {
	Name                  string `json:"name" validate:"required,min=5,max=50"`
	NIP                   int    `json:"nip" validate:"required,min_len=13,max_len=15"`
	IdentityCardScanImage string `json:"identityCardScanImg" validate:"required,url"`
}

type AccessNurse struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type GetUsers struct {
	UserId    *string `query:"userId" validate:"omitempty"`
	Limit     *int    `query:"limit" validate:"omitempty"`
	Offset    *int    `query:"offset" validate:"omitempty"`
	Name      *string `query:"name" validate:"omitempty"`
	Nip       *string `query:"nip" validate:"omitempty"`
	Role      *string `query:"role" validate:"omitempty,oneof=it nurse"`
	CreatedAt *string `query:"createdAt" validate:"omitempty,oneof=asc desc"`
}

type EditNurse struct {
	Name string `json:"name" validate:"required,min=5,max=50"`
	NIP  int    `json:"nip" validate:"required,min_len=13,max_len=15"`
}
