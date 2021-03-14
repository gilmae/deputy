package orgchart

import (
	"fmt"
	"sort"
)

// Role represents a role within an organisation, ex: System Administrator.
// Roles are heirarchical; Parent is the role that supervises the given role
type Role struct {
	Id     int
	Name   string
	Parent int
}

// User represents a User within an organisation.
type User struct {
	Id   int
	Name string
	Role int
}

// Organisation represents the collection of roles and users with an organisation.
type Organisation struct {
	roles    map[int]Role
	roleTree map[int][]int

	users map[int]User

	usersInRole map[int][]User
}

// NewOrganisation constructs a new Organisation and initialises data structures
func NewOrganisation() *Organisation {
	return &Organisation{
		roles:       make(map[int]Role),
		roleTree:    make(map[int][]int),
		users:       make(map[int]User),
		usersInRole: make(map[int][]User),
	}
}

// GetSubordinates queries the users and roles mappings to build a list of subordinate users,
// including subordinates of subordinates
func (o *Organisation) GetSubordinates(userId int) ([]User, error) {
	found := make(map[int]User)

	user, ok := o.users[userId]

	if !ok {
		return nil, fmt.Errorf("User not found")
	}

	o.mapSubordinates(user, found)

	subordinates := []User{}

	for _, user := range found {
		subordinates = append(subordinates, user)
	}

	sort.Slice(subordinates, func(i, j int) bool {
		return subordinates[i].Id < subordinates[j].Id
	})

	return subordinates, nil
}

// SetRoles setups the organisations roles
func (o *Organisation) SetRoles(roles []Role) {
	o.roles = make(map[int]Role)
	o.roleTree = make(map[int][]int)

	o.usersInRole = make(map[int][]User)
	for _, role := range roles {
		o.roles[role.Id] = role

		// Forgo referential integrity check for now
		if _, ok := o.roleTree[role.Parent]; !ok {
			o.roleTree[role.Parent] = []int{role.Id}
		} else {
			o.roleTree[role.Parent] = append(o.roleTree[role.Parent], role.Id)
		}

		o.usersInRole[role.Id] = make([]User, 0)
	}
	o.mapUsersToRoles()
}

// SetUsers setups the users within an organisation
func (o *Organisation) SetUsers(users []User) {
	o.users = make(map[int]User)
	for _, user := range users {
		o.users[user.Id] = user
	}
	o.mapUsersToRoles()
}

func (o *Organisation) mapSubordinates(user User, found map[int]User) {
	// Find subroles for the user's role. If the user has no subroles, no need to process further
	subRoles, ok := o.roleTree[user.Role]
	if !ok {
		return
	}

	for _, roleId := range subRoles {
		if users, ok := o.usersInRole[roleId]; ok {
			for _, u := range users {
				// If the user is not already in the foudn map, add and find their subordinates
				// If they are found, don't re-process them - avoid circular orgcharts
				if _, ok := found[u.Id]; !ok {
					found[u.Id] = u
					o.mapSubordinates(u, found)
				}
			}
		}
	}
}

func (o *Organisation) mapUsersToRoles() {
	for _, user := range o.users {
		if users, ok := o.usersInRole[user.Role]; ok {
			o.usersInRole[user.Role] = append(users, user)
		}
	}
}
