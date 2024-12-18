package domain

type UserModel struct {
	id        int64
	firstname string
	lastname  string
	role      UserRoleEnum
}

func NewUserModel(id int64, firstname, lastname string, role UserRoleEnum) *UserModel {
	return &UserModel{
		id:        id,
		firstname: firstname,
		lastname:  lastname,
		role:      role,
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
