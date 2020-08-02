package model

import (
	"ginblog/utils/err_msg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(username string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", username).First(&cate)
	if cate.ID > 0 {
		return err_msg.ErrorCategoryNameUsed // 2001
	}
	return err_msg.SUCCESS // 200
}

// 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return err_msg.ERROR // 500
	}
	return err_msg.SUCCESS // 200
}

// 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cates []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

// 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}
