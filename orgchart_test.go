package orgchart

import (
	"sort"
	"testing"
)

func TestSetRoles(t *testing.T) {

	tests := [][]Role{
		[]Role{Role{Id: 1, Name: "foo", Parent: 0}},
		[]Role{
			Role{Id: 1, Name: "bar", Parent: 0},
			Role{Id: 2, Name: "baz", Parent: 1},
		},
	}

	for _, roles := range tests {
		o := NewOrganisation()
		o.SetRoles(roles)

		if len(o.Roles) != len(roles) {
			t.Errorf("Incorrect number of roles in organisation, expected %d, got %d", len(roles), len(o.Roles))
		}

		if len(o.UsersInRole) != len(roles) {
			t.Errorf("Incorrect number of roles in user to role mapping in organisation, expected %d, got %d", 1, len(o.Roles))
		}

		for _, role := range roles {
			r, ok := o.Roles[role.Id]
			if !ok {
				t.Fatalf("Could not resolve Role# %d in organisation roles", role.Id)
			}

			if r != role {
				t.Errorf("Role does not match, expected %+v, got %+v", roles[0], r)
			}

			users, ok := o.UsersInRole[role.Id]
			if !ok {
				t.Fatalf("Could not resolve Role# %d in user to role mapping", role.Id)
			}
			if len(users) != 0 {
				t.Errorf("Users in role not empty for Role#%d", role.Id)
			}
		}
	}
}

func TestRoleTree(t *testing.T) {
	roles := []Role{
		Role{Id: 1, Name: "doh", Parent: 0},
		Role{Id: 2, Name: "ray", Parent: 1},
		Role{Id: 3, Name: "me", Parent: 1},
		Role{Id: 4, Name: "fah", Parent: 2},
		Role{Id: 5, Name: "fah", Parent: 6},
	}

	o := NewOrganisation()
	o.SetRoles(roles)

	if subroles, ok := o.roleTree[1]; !ok {
		t.Fatalf("Could not resolve subroles for Role# %d", 1)
	} else {

		if len(subroles) != 2 {
			t.Errorf("Incorrect number of subroles for Role# %d, expected %d, got %d", 1, 2, len(subroles))
		}

		sort.Ints(subroles)
		if subroles[0] != 2 && subroles[1] != 3 {
			t.Errorf("Incorrect subroles, expected %+v, got %+v", []int{2, 3}, subroles)
		}
	}

	if subroles, ok := o.roleTree[2]; !ok {
		t.Fatalf("Could not resolve subroles for Role# %d", 2)
	} else {

		if len(subroles) != 1 {
			t.Errorf("Incorrect number of subroles for Role# %d, expected %d, got %d", 2, 1, len(subroles))
		}

		if subroles[0] != 4 {
			t.Errorf("Incorrect subroles, expected %+v, got %+v", []int{2, 3}, subroles)
		}
	}
}

func TestSetUsers(t *testing.T) {
	tests := [][]User{
		[]User{User{Id: 1, Name: "foo", Role: 1}},
		[]User{
			User{Id: 1, Name: "bar", Role: 1},
			User{Id: 2, Name: "baz", Role: 2},
		},
	}

	for _, users := range tests {
		o := NewOrganisation()
		o.SetUsers(users)

		if len(o.Users) != len(users) {
			t.Errorf("Incorrect number of users in organisation, expected %d, got %d", len(users), len(o.Users))
		}

		for _, user := range users {
			u, ok := o.Users[user.Id]
			if !ok {
				t.Fatalf("Could not resolve User# %d in organisation users", user.Id)
			}

			if u != user {
				t.Errorf("User does not match, expected %+v, got %+v", users[0], u)
			}
		}
	}
}
