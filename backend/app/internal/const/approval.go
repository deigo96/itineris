package constant

type Approval int

const (
	APPROVE = iota + 1
	REJECT
)

var approvalName = map[Approval]string{
	APPROVE: APPROVED,
	REJECT:  REJECTED,
}

func (a Approval) String() string {
	return approvalName[a]
}

func GetApprovalStatus(status string) Approval {
	for k, v := range approvalName {
		if v == status {
			return k
		}
	}
	return 0
}
