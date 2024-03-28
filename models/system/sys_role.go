package system

import (
	"time"
)

// 系统角色
type SysRole struct {
	CreatedAt     time.Time  // 创建时间
	UpdatedAt     time.Time  // 更新时间
	DeletedAt     *time.Time `sql:"index"`
	RoleId   uint       `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	RoleName string     `json:"roleName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      *uint      `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	//DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children []SysRole `json:"children" gorm:"-"`
	//SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users         []SysUser `json:"-" gorm:"many2many:sys_user_role;"`
	DefaultRouter string    `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysRole) TableName() string {
	return "sys_roles"
}
