package models

type Role string

const (
	RoleOwner      Role = "owner"
	RoleAdmin      Role = "admin"
	RoleMaintainer Role = "maintainer"
	RoleViewer     Role = "viewer"
	RoleSupport    Role = "support"
)
