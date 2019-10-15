package addUserToUserGroup

import (

)

func validateGroupUuid(groupUuid string) bool {
	return validateGroupUuidInDb(groupUuid)
}

func addUsers(groupUuid string, userList []string) bool {
	return addUsersToDb(groupUuid, userList)
}