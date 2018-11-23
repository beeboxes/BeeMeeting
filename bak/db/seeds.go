package db

import (
	"beeboxes.com/BeeMeeting/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {

	//uis := make()
	//for _, ui := range uis {
	//
	//}
}

func Seed() {
	db, err := gorm.Open("mysql", "root:@/beemeeting_development?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	roleGroups := make(map[string][]string)
	roleGroups["员工管理"] = []string{
		"添加/编辑员工",
		"删除员工",
	}

	roleGroups["角色管理"] = []string{
		"创建/编辑角色",
		"删除角色",
	}

	roleGroups["会议管理"] = []string{
		"添加/编辑会议",
		"删除会议",
	}

	roleGroups["参会人管理"] = []string{
		"导入参会人",
		"现场采集",
		"补签",
		"导出参会人信息",
		"删除参会人",
	}
	/*roleGroups[models.PermissionGroup{Name: "员工管理"}] = []models.Permission{
		models.Permission{Name: "添加/编辑员工"},
		models.Permission{Name: "删除员工"},
	}

	roleGroups[models.PermissionGroup{Name: "角色管理"}] = []models.Permission{
		models.Permission{Name: "创建/编辑角色"},
		models.Permission{Name: "删除角色"},
	}

	roleGroups[models.PermissionGroup{Name: "会议管理"}] = []models.Permission{
		models.Permission{Name: "添加/编辑会议"},
		models.Permission{Name: "删除会议"},
	}

	roleGroups[models.PermissionGroup{Name: "参会人管理"}] = []models.Permission{
		models.Permission{Name: "导入参会人"},
		models.Permission{Name: "现场采集"},
		models.Permission{Name: "补签"},
		models.Permission{Name: "导出参会人信息"},
		models.Permission{Name: "删除参会人"},
	}*/
	db.AutoMigrate(&models.PermissionGroup{}, &models.Permission{})
	for group, roles := range roleGroups {
		g:=models.PermissionGroup{Name: group}
		ret := db.Create(&g)
		fmt.Println("创建分组： ", ret)
		for _, role := range roles {
			ret := db.Create(&models.Permission{Name: role, PermissionGroupID: g.Model.ID})
			fmt.Println("创建权限： ", ret)
		}
	}
}
