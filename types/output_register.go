package types

type OutputRegister struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	CreationDate int64  `json:"creation_date"`
}
