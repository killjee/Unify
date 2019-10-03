package createUserGroup

import(
	
)	

func createGroup(uuid string, name string, adminUserUuid string, parentGroupUuid string) bool {
	return addGroupToDb(uuid, name, adminUserUuid, parentGroupUuid)
}