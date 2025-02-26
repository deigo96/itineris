package constant

type Role int

const (
	PPK Role = iota + 1
	Staff
)

var roleName = map[Role]string{
	PPK:   "Pejabat Pembina Kepegawaian",
	Staff: "Staff",
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
	return Staff
}
