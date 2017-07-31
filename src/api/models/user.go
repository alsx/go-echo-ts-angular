package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // Register MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// User represents user model
type User struct {
	ID            int64
	Name          string
	Email         string
	Password      string    `json:"-"`
	FacebookToken string    `db:"facebook_token" json:"-"`
	Created       time.Time `json:"-"`
}

var dbConn *sqlx.DB

// UserManager contains methods to work with User model
type UserManager struct {
	db *sqlx.DB
}

// NewUserManager returns instance of UserManager
func NewUserManager(ctx echo.Context) UserManager {
	if dbConn != nil {
		ctx.Logger().Debug("Reusing DB connection.")
		return UserManager{dbConn}
	}
	dsn, _ := ctx.Get("dsn").(string)
	var err error
	dbConn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		ctx.Logger().Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		ctx.Logger().Fatal(err)
	}
	ctx.Logger().Debug("Connected to MySQL DB.")
	return UserManager{dbConn}
}

// NewUserManager create connection pool to db.
// TODO: move to base when it will be more then one model

// Add - create new user.
func (mgr UserManager) Add(user User) (int64, error) {
	const query = "INSERT INTO `users` (`name`, `email`, `password`, `facebook_token`)  VALUES (?, ?, ?, ?)"
	result := mgr.db.MustExec(query, user.Name, user.Email, user.Password, user.FacebookToken)

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetByEmail selects user by email
func (mgr UserManager) GetByEmail(email string) (User, error) {
	const query = "SELECT * from `users` WHERE email=?"
	user := User{}
	err := mgr.db.Get(&user, query, email)
	return user, err
}
