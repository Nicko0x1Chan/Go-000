package service

import "Week02/dao/userDao"

func DoService() (int, error) {
	return userDao.DoDao(), nil
}
