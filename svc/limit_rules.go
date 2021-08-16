package svc

import (
	"context"

	"achuala.in/velolimits/pbgen/velolimits"
	"gorm.io/gorm"
)

type LimitRulesSvc struct {
	velolimits.UnimplementedVelocityLimitServer
	db *gorm.DB
}

func NewLimitRulesSvc(db *gorm.DB) *LimitRulesSvc {
	return &LimitRulesSvc{db: db}
}

func (s *LimitRulesSvc) CreateNewRule(context.Context, *velolimits.CreateNewRuleRq) (*velolimits.CreateNewRuleRs, error) {
	return nil, nil
}
