package Models

import (
	Mysql "SimpleApi/databses"
	"errors"
)

func GetPermissionList(page int, pageSize int) []Permission {
	var permissions []Permission
	Mysql.DB.Scopes(Mysql.Paginate(page, pageSize)).Find(&permissions)

	return permissions
}

type MenuTree struct {
	Name     string      `json:"name"`
	Parent   string      `json:"parent"`
	Status   string      `json:"status"`
	Uid      string      `json:"uid"`
	Type     string      `json:"type"`
	Url      string      `json:"url"`
	Icon     string      `json:"icon"`
	Describe string      `json:"describe"`
	Children []*MenuTree `json:"children"`
}

func GetAllPermissions() ([]interface{}, error) {
	permTree, err := GetMenuTree()
	if err != nil {
		return permTree, errors.New("get permissions error")
	}
	return permTree, nil
}

//构建树形结构
func GetMenuTree() (dataList []interface{}, err error) {
	var parentList []Permission
	//获取父节点

	Mysql.DB.Where("type = ?", "01").Find(&parentList)

	for _, v := range parentList {
		parent := MenuTree{v.Name, v.Parent, v.Status, v.Uid, v.Type, v.Url, v.Icon, v.Describe, []*MenuTree{}}
		var childrenList []Permission

		Mysql.DB.Where("parent = ?", v.ID).Find(&childrenList)

		for _, c := range childrenList {
			child := MenuTree{c.Name, c.Parent, c.Status, c.Uid, c.Type, c.Url, c.Icon, c.Describe, []*MenuTree{}}
			parent.Children = append(parent.Children, &child)
		}
		dataList = append(dataList, parent)
	}
	return dataList, nil
}
