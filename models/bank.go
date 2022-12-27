package models

type Bank struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Acntnum  string `json:"acntnum"`
	Acnttype string `json:"acnttype"`
	Acntbal  string `json:"acntbal"`
}
type Banker struct {
	Bankdetails []Bank `json:"bankdetails"`
}
