package domain

import "golang.org/x/crypto/bcrypt"

type UserModel struct {
	id             int64
	firstname      string
	lastname       string
	role           UserRoleEnum
	hashedPassword string
}

func NewUserModel(id int64, firstname, lastname string, role UserRoleEnum, hashedPassword string) *UserModel {
	return &UserModel{
		id:             id,
		firstname:      firstname,
		lastname:       lastname,
		role:           role,
		hashedPassword: hashedPassword,
	}
}

func (m *UserModel) SetID(id int64) {
	m.id = id
}

func (m *UserModel) GetID() int64 {
	return m.id
}

func (m *UserModel) SetFirstname(firstname string) {
	m.firstname = firstname
}

func (m *UserModel) GetFirstname() string {
	return m.firstname
}

func (m *UserModel) SetLastname(lastname string) {
	m.lastname = lastname
}

func (m *UserModel) GetLastname() string {
	return m.lastname
}

func (m *UserModel) GetFullname() string {
	return m.firstname + " " + m.lastname
}

func (m *UserModel) SetRole(role UserRoleEnum) {
	m.role = role
}

func (m *UserModel) GetRole() UserRoleEnum {
	return m.role
}

func (*UserModel) PasswordToHash(password string) ([]byte, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *UserModel) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(m.hashedPassword),
		[]byte(password),
	) == nil
}
