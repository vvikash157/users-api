package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GlobalUse struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewGlobalUse(DB *gorm.DB, log *logrus.Logger) *GlobalUse {
	s := &GlobalUse{
		db:     DB,
		logger: log,
	}
	return s
}
