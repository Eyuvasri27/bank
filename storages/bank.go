package storages

import (
	"echobank/models"
	"fmt"
)

func (c *cursor) AccountDetails(bank *models.BankUserDetails) error {
	sqlstatement := "INSERT INTO bankuserdetails(id,username,password,accountnumber)values($1,$2,$3,$4)"
	_, err := c.Db.Query(sqlstatement, bank.Id, bank.Username, bank.Password, bank.AccountNumber)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func (c *cursor) Verification(username, password string) (*models.BankUserDetails, error) {
	sqlstatement := "SELECT * FROM bankuserdetails where username = $1 and password = $2"
	var v models.BankUserDetails
	res := c.Db.QueryRow(sqlstatement, username, password)
	err := res.Scan(&v.Id, &v.Username, &v.Password, &v.AccountNumber)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (c *cursor) Createbank(b *models.Bank) error {
	sqlstatement := "INSERT INTO bank(name,acntnum,acnttype,acntbal)VALUES($1,$2,$3,$4)"
	_, err := c.Db.Query(sqlstatement, b.Name, b.Acntnum, b.Acnttype, b.Acntbal)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (c *cursor) Listbank() (*models.Banker, error) {
	sqlStatement := "SELECT id, name, acntnum,acnttype,acntbal FROM bank order by id"
	rows, err := c.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)

	}
	defer rows.Close()

	var result models.Banker
	for rows.Next() {

		var dummy models.Bank
		err2 := rows.Scan(&dummy.Id, &dummy.Name, &dummy.Acntnum, &dummy.Acnttype, &dummy.Acntbal)

		if err2 != nil {
			return nil, err2
		}
		result.Bankdetails = append(result.Bankdetails, dummy)

	}
	return &result, nil
}

func (c *cursor) Listbyidbank(id int) (*models.Bank, error) {

	sqlStatement := "SELECT * from bank where id = $1"
	fmt.Println(sqlStatement)
	var result1 models.Bank
	res := c.Db.QueryRow(sqlStatement, id)
	err := res.Scan(&result1.Id, &result1.Name, &result1.Acntnum, &result1.Acnttype, &result1.Acntbal)
	if err != nil {
		fmt.Println(err)
	}
	return &result1, nil
}

func (c *cursor) Putbank(id int, u *models.Bank) error {

	sqlStatement := "UPDATE bank SET name=$1,acntnum=$2,acnttype=$3,acntbal=$4,id"
	_, err := c.Db.Query(sqlStatement, u.Name, u.Acntnum, u.Acnttype, u.Acntbal, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *cursor) Deletebank(id int) error {
	sqlStatement := "DELETE FROM bank WHERE id = $1"
	_, err := c.Db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
