package types

// PermissionContext holds the permission flags for runtime operations
type PermissionContext struct {
	AllowFS  bool
	AllowNet bool
	AllowEnv bool
}
