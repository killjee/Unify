create table UserGroupMetaData (
	uuid INT64,
	name VARCHAR,
	adminUserUUID INT64,
	parentGroupUUID INT64,
	PRIMARY KEY (uuid)
);

create table UserGroupMembers(
	groupUuid INT64,
	userUuid INT64,
	PRIMARY KEY (groupUuid, userUuid),
	FOREIGN KEY (groupUuid) REFRENCES UserGroupMetaData(uuid)
);