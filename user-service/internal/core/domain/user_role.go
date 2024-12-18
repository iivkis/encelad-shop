package domain

import (
	"fmt"
	"sync"
)

type UserRoleEnum interface {
	isUserRole()
	String() string
}

type implUserRole struct {
	val string
}

func (implUserRole) isUserRole() {}

func (r implUserRole) String() string {
	return r.val
}

func (r implUserRole) Eq(role UserRoleEnum) bool {
	return r.val == role.String()
}

type enumUserRole_Admin struct{ *implUserRole }
type enumUserRole_User struct{ *implUserRole }

type fieldsUserRole struct {
	Admin *enumUserRole_Admin
	User  *enumUserRole_User
}

func (f fieldsUserRole) FromString(role string) (UserRoleEnum, error) {
	switch role {
	case f.Admin.val:
		return f.Admin, nil
	case f.User.val:
		return f.User, nil
	default:
		return nil, fmt.Errorf("unknown role: %s", role)
	}
}

func (f fieldsUserRole) FromStringMust(role string) UserRoleEnum {
	r, err := f.FromString(role)
	if err != nil {
		panic(err)
	}
	return r
}

var (
	syncOnceUserRole = &sync.Once{}
	userRoleEnum     *fieldsUserRole
)

func UserRole() *fieldsUserRole {
	syncOnceUserRole.Do(func() {
		userRoleEnum = &fieldsUserRole{
			Admin: &enumUserRole_Admin{
				&implUserRole{
					val: "admin",
				},
			},
			User: &enumUserRole_User{
				&implUserRole{
					val: "user",
				},
			},
		}
	})
	return userRoleEnum
}
