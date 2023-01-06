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
type BankUserDetails struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AccountNumber string `json:"accountnumber"`
}
type LoginVerification struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
