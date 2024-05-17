package request

type Login struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
	NIP      int    `json:"nip" validate:"required,min_len=13,max_len=15"`
}

type Register struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	NIP      int    `json:"nip" validate:"required,min_len=13,max_len=15"`
}
