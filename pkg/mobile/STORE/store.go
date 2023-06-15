package store

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"

	// "github.com/golang-jwt/jwt/v4"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	"golang.org/x/crypto/bcrypt"
)

type Store interface{
	CreateMobile(mobile Model.Mobile) error  
	GetAllMobiles()  ([]Model.Mobile, error)
	GetByID( id int ) Model.Mobile
	DeleteByID( id int) error
	UpdateByID( id int, brand string) error
} 

type store struct {
	session *gocql.Session
}
var mobileStore Store

func NewStore () Store {
	Session := db.Session
	return &store{
		session : Session,
	}
}

func GetStore() Store {
	if mobileStore == nil{
		mobileStore = NewStore()
	}
	return mobileStore 
}

func (s *store) CreateMobile(mobile Model.Mobile) error {
	query := `INSERT INTO mobile_app.mobile(id,brand,model,year,price) VALUES (?,?,?,?,?)`
	err := s.session.Query(query,mobile.ID,mobile.Brand,mobile.Model,mobile.Year,mobile.Price).Exec()
	if err != nil{
		log.Printf("ERROR: %s", err.Error())
	}
	return err
}

func (s *store) GetAllMobiles() ([]Model.Mobile, error) {
	query := `SELECT * FROM mobile_app.mobile`
	iter := s.session.Query(query).Iter()
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

func (s *store) GetByID( id int ) Model.Mobile{
	query := `SELECT * FROM mobile_app.mobile WHERE id=?`
	var mobile Model.Mobile
	s.session.Query(query,id).Scan(&mobile.ID, &mobile.Brand, &mobile.Model, &mobile.Year, &mobile.Price)
	return mobile
}

func (s *store) DeleteByID( id int) error {
	query := `DELETE FROM mobile_app.mobile WHERE id=?`
	err := s.session.Query(query,id).Exec()
	if err != nil {
		log.Printf("ERROR: fail edit document, %s", err.Error())
	}

	return err
}

func (s *store) UpdateByID( id int, brand string) error {
	query := `UPDATE mobile_app.mobile SET  brand=? WHERE id=?`
	err := s.session.Query(query, brand, id).Exec()
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

func CheckEmail( session *gocql.Session, email string) Model.User {
    fmt.Println("email",email)
	query := `SELECT * FROM mobile_app.user WHERE Email=?`
	var user Model.User
    session.Query(query,email).Scan(&user.Email,&user.Password)
	fmt.Println("userdb",user)
	return user

}