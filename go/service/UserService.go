package service

import (
	"simpleMVC/go/dao"
	"simpleMVC/go/entity"
)

//新增user
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

//获取user集合
func GetAllUser() (userList []*entity.User, err error) {
	if err := dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

//根据id删除user
func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

//根据id查询user
func GetUserById(id int) (user *entity.User, err error) {
	user = &entity.User{}
	if err = dao.SqlSession.Debug().Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, err
}

//更新user信息
func UpdateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}

//根据用户名查询user
func GetUserByName(name string) (user *entity.User, err error) {
	user = &entity.User{}
	if err = dao.SqlSession.Debug().Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, err
}
