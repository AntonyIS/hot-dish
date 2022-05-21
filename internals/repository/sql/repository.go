package sql

import (
	"database/sql"
	"fmt"

	"github.com/AntonyIS/go-foods/internals/domain"
)

type sqlRepository struct {
	db        *sql.DB
	tablename string
}

func dbClient() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "gofoods"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	// Ping db
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
func NewRepository(driverName, tablename string) (domain.UserRepository, error) {
	repo := &sqlRepository{}
	repo.db = dbClient()
	return repo, nil
}

func (sqlRepo sqlRepository) CreateUser(user domain.User) (*domain.User, error) {
	// perform a db.Query insert
	operation := fmt.Sprintf("INSERT INTO %v", sqlRepo.tablename)
	values := fmt.Sprintf("VALUES (%v %v %v) ", user.UserID, user.Name, user.Email)
	insert, err := sqlRepo.db.Query(operation + sqlRepo.tablename + values)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	return &user, nil
}

func (sql sqlRepository) GetUsers() (*[]domain.User, error) {
	query, err := sql.db.Query(fmt.Sprintf("SELECT * FROM %v ORDER by DESC", sql.tablename))
	user := domain.User{}
	users := []domain.User{}

	for query.Next() {
		var user_id, name, email string
		err = query.Scan(&user_id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		user.UserID = user_id
		user.Name = name
		user.Email = email

		users = append(users, user)
	}
	defer query.Close()
	defer sql.db.Close()
	return &users, nil

}

func (sql sqlRepository) GetUser(user_id string) (*domain.User, error) {
	user := domain.User{}
	query, err := sql.db.Query("SELECT * FROM Employee WHERE id=?", user_id)
	if err != nil {
		panic(err.Error())
	}

	if query.Next() {
		err = query.Scan(&user.UserID, &user.Name, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}
	return &user, nil

}

func (sql sqlRepository) UpdateUser(user domain.User) (*domain.User, error) {
	prep, err := sql.db.Prepare(fmt.Sprintf("UPDATE %v SET name=?, email=? WHERE user_id=?", sql.tablename))

	if err != nil {
		panic(err.Error())
	}
	prep.Exec(user.Name, user.Email)

	defer sql.db.Close()
	return &user, nil
}

func (sql sqlRepository) DeleteUser(user_id string) error {
	query, err := sql.db.Prepare(fmt.Sprintf("DELETE FROM %v WHERE id=?", sql.tablename))
	if err != nil {
		panic(err.Error())
	}
	query.Exec(user_id)
	defer sql.db.Close()
	return nil
}
