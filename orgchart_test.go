package orgchart

import "testing"

func TestSetRoles(t *testing.T) {

	tests := [][]Role{
		[]Role{Role{Id: 1, Name: "foo", Parent: 0}},
		[]Role{
			Role{Id:1, Name:"bar", Parent:0},
			Role{Id:2, Name:"baz", Parent:1},
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
