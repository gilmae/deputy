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
	roleTree map[int][]int

	Users map[int]User

	UsersInRole map[int][]User
	
}

func NewOrganisation() *Organisation{
	return &Organisation{
		Roles: make(map[int]Role),
		roleTree: make(map[int][]int),
		Users: make(map[int]User),
		UsersInRole: make(map[int][]User),
	}
}

func (o *Organisation) SetRoles(roles []Role) {
	o.Roles = make(map[int]Role)
	o.roleTree = make(map[int][]int)

	o.UsersInRole = make(map[int][]User)
	for _, role := range roles {
		o.Roles[role.Id] = role
		
		// Forgo referential integrity check for now
		if _,ok := o.roleTree[role.Parent]; !ok {
			o.roleTree[role.Parent] = []int {role.Id}
		} else {
			o.roleTree[role.Parent] = append(o.roleTree[role.Parent], role.Id)
		}

		o.UsersInRole[role.Id] = make([]User, 0)
	}
}

func (o *Organisation) SetUsers(users []User) {
	o.Users = make(map[int]User)
	for _, user := range users {
		o.Users[user.Id] = user
	}
}
