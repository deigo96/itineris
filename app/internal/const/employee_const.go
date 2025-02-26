package constant

type Role int

const (
	SUPER Role = iota + 1
	ADMIN
	USER
)

var roleName = map[Role]string{
	SUPER: "super",
	ADMIN: "admin",
	USER:  "user",
}

func (r Role) String() string {
	return roleName[r]
}

func GetRole(role string) Role {
	for k, v := range roleName {
		if v == role {
			return k
		}
	}
	return USER
}
