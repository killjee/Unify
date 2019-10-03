create table UserGroupMetaData (
	uuid VARCHAR NOT NULL,
	name VARCHAR NOT NULL,
	adminUserUuid VArCHAR NOT NULL,
	parentGroupUuid VArCHAR,
	PRIMARY KEY (uuid)
);

create table UserGroupMembers(
	groupUuid VArCHAR NOT NULL REFERENCES UserGroupMetaData(uuid),
	userUuid VArCHAR NOT NULL,
	PRIMARY KEY (groupUuid, userUuid)
);