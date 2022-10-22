package users

type Core struct {
	UserID   int
	Name     string
	Email    string
	Password string
	Role     RoleCore
}

type RoleCore struct {
	ID       int
	RoleName string
}

type Data interface {
	InsertData(data Core) error
	IsDuplicate(data Core) error
	FindUser(Email string) (Core, error)
	SelectUser(ID int) (Core, error)
	SelectAll() ([]Core, error)
	DataDelete(string) error
}

type Bussines interface {
	Register(data Core) error
	Login(data Core) (id int, token string, err error)
	GetProfile(ID int) (Core, error)
	GetAll() ([]Core, error)
	BussDell(id string) error
}
