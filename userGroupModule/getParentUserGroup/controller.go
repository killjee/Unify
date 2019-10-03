package getParentUserGroup

import(
	
)

func getParentUserGroup(uuid string) (string, int) {
	return getParentUserGroupFromDb(uuid)
}