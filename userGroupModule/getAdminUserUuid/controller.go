package getAdminUserUuid

import(
	
)

func getAdminUuid(uuid string) (string, int) {
	return getAdminUuidFromDb(uuid)
}