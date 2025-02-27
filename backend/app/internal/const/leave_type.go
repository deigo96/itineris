package constant

type Status string

const (
	PENDING  = "pending"
	APPROVED = "approved"
	REJECTED = "rejected"
)

var StatusInt = map[Status]int{
	PENDING:  1,
	APPROVED: 2,
	REJECTED: 3,
}

func (s Status) Int() int {
	return StatusInt[s]
}

func (s Status) String() string {
	return string(s)
}

func GetStatus(s int) Status {
	for k, v := range StatusInt {
		if v == s {
			return k
		}
	}
	return PENDING
}
