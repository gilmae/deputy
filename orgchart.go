package orgchart

type Role struct {
	Id int
	Name string
	Parent int
}

type User struct {
	Id int
	Name string
	Role int
}

type Organisation struct {
	Roles map[int]Role
	Users map[int]User

	UsersInRole map[int][]User
	
}

func NewOrganisation() *Organisation{
	return &Organisation{
		Roles: make(map[int]Role),
		Users: make(map[int]User),
		UsersInRole: make(map[int][]User),
	}
}

func (o *Organisation) SetRoles(roles []Role) {
	o.Roles = make(map[int]Role)
	o.UsersInRole = make(map[int][]User)
	for _, role := range roles {
		o.Roles[role.Id] = role
		// TODO Check if users exist and remap
		o.UsersInRole[role.Id] = make([]User, 0)
	}
}