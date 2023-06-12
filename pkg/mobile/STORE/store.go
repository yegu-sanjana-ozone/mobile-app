package store

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	"golang.org/x/crypto/bcrypt"
)

func CreateMobile(session *gocql.Session,mobile Model.Mobile) error {
	query := `INSERT INTO mobile_app.mobile(id,brand,model,year,price) VALUES (?,?,?,?,?)`
	err := session.Query(query,mobile.ID,mobile.Brand,mobile.Model,mobile.Year,mobile.Price).Exec()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
	}
	return err
}

func GetAllMobiles(session *gocql.Session) ([]Model.Mobile, error) {
	query := `SELECT * FROM mobile_app.mobile`
	iter := session.Query(query).Iter()
	defer iter.Close()

	var mobiles []Model.Mobile

	for {
		var mobile Model.Mobile
		if !iter.Scan(&mobile.ID, &mobile.Brand, &mobile.Model, &mobile.Year, &mobile.Price) {
			break
		}
		mobiles = append(mobiles, mobile)
	}

	if err := iter.Close(); err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil, err
	}

	return mobiles, nil
}

func GetByID(session *gocql.Session, id int ) Model.Mobile{
	query := `SELECT * FROM mobile_app.mobile WHERE id=?`
	var mobile Model.Mobile
	session.Query(query,id).Scan(&mobile.ID, &mobile.Brand, &mobile.Model, &mobile.Year, &mobile.Price)
	return mobile
}

func DeleteByID(session *gocql.Session, id int) error {
	query := `DELETE FROM mobile_app.mobile WHERE id=?`
	err := session.Query(query,id).Exec()
	if err != nil {
		log.Printf("ERROR: fail edit document, %s", err.Error())
	}

	return err
}

func UpdateByID(session *gocql.Session, id int, brand string) error {
	query := `UPDATE mobile_app.mobile SET  brand=? WHERE id=?`
	err := session.Query(query, brand, id).Exec()
	if err != nil {
		log.Printf("ERROR: failed to edit document, %s", err.Error())
	}

	return err
}

func CreateUser(session *gocql.Session,user Model.User)error {
	query := `INSERT INTO mobile_app.user(Email,Password) VALUES (?,?)`
	err := session.Query(query,user.Email,user.Password).Exec()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
	}
	return err 
}

func ValidateUser(session *gocql.Session, user Model.User) error {
	var email, hashpwd string
	query := `SELECT * FROM mobile_app.user WHERE email = ?`

	err:= session.Query(query,user.Email).Scan(&email,&hashpwd)
	if err!= nil {
		if err == gocql.ErrNotFound {
            return fmt.Errorf("user not found")
        }
        return err
    }
	err = bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(user.Password))
	if err != nil {
		return fmt.Errorf("invalid email or password")
	}

	return nil

}