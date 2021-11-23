package dbm

import (
	"fmt"
	"hlaing-backend-test-1/env"
	"net/mail"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var userSchema = `
CREATE TABLE %s (
    id              SERIAL,
    name            VARCHAR(80) NOT NULL,
    email           VARCHAR(250) UNIQUE,
    password        VARCHAR(250) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
`

// User represent a user record from the database
type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UserManager wraps all user database operations
type UserManager struct {
	*sqlx.DB
	*env.Envars

	TableName string
}

// NewUserManager return a new user manager
func NewUserManager(envars *env.Envars, db *sqlx.DB) *UserManager {
	return &UserManager{
		DB:        db,
		Envars:    envars,
		TableName: envars.TablePrefix + `user`,
	}
}

// AddNewUser add a new user to the system
func (um *UserManager) AddNewUser(name, email, pass string) error {
	if name == "" {
		return fmt.Errorf("name must not be empty")
	}

	if email == "" {
		return fmt.Errorf("email must not be empty")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	if pass == "" {
		return fmt.Errorf("password must not be empty")
	}

	// TODO: check password length, strength etc

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var stmt, args = squirrel.Insert(um.TableName).
		Columns("name", "email", "password").
		Values(name, email, hash).
		MustSql()

	if _, err := um.Exec(stmt, args...); err != nil {
		return err
	}

	return nil
}

// GetUser return an existing user from the system
func (um *UserManager) GetUser(email string) (User, error) {
	var user = User{}
	var stmt, args = squirrel.Select("*").From(um.TableName).
		Where(`email = ?`, email).
		MustSql()

	if err := um.Get(&user, stmt, args...); err != nil {
		return user, err
	}

	return user, nil
}

// AuthenticateUser wraps GetUser to also check password
func (um *UserManager) AuthenticateUser(email, pass string) (User, error) {
	user, err := um.GetUser(email)
	if err != nil {
		return user, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return user, err
	}

	return user, nil
}

// Setup is used for testing setup
func (um *UserManager) Setup() {
	um.MustExec(fmt.Sprintf(userSchema, um.TableName))
}

// TearDown is used for testing teardown
func (um *UserManager) TearDown() {
	um.MustExec(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, um.TableName))
}
