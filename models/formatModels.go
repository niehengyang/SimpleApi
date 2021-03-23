package Models

// 用户
type UserForm struct {
	User
	Roles []int `json:"Roles"`
}

//角色
type RoleForm struct {
	Role
	Permissions []int `json:"Permissions"`
}

//权限
type PermissionForm struct {
	Permission
}

//文件上传
type FileForm struct {
	Path       string
	Name       string
	OriginName string
	HashName   string
}

//obj特殊格式
type ObjType struct {
	Img []string
	Obj string
	Mtl string
}
