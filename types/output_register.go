package types

type OutputRegister struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	CreationDate int64  `json:"creation_date"`
}
