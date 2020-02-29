package roles

type Role int

const (
	USER Role = 0x1
	TEACHER Role = 0x1 << 100
	AUTHORITY Role = 0x1 << 200
	OPERATOR Role = 0x1 << 300
	ADMIN Role = 0x1 << 400
)

func (r Role) IsUser() bool {
	return r&USER != 0
}

func (r Role) IsOperator() bool {
	return r&OPERATOR != 0
}

func (r Role) IsAdmin() bool {
	return r&ADMIN != 0
}

func (r Role) IsTeacher() bool {
	return r&TEACHER != 0
}

func (r Role) IsAuthority() bool {
	return r&AUTHORITY != 0
}