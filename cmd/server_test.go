package main

import (
	"context"
	"testing"

	"achuala.in/velolimits/pbgen/velolimits"
	"achuala.in/velolimits/svc"
	"github.com/golang/mock/gomock"
)

func TestCreateNewRule(t *testing.T) {
	testCases := []struct {
		name        string
		req         *velolimits.CreateNewRuleRq
		ruleId      string
		expectedErr bool
	}{
		{
			name:        "CreateNewCountRule",
			req:         &velolimits.CreateNewRuleRq{RuleDef: &velolimits.CreateNewRuleRq_Tlr{}},
			ruleId:      "",
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ctx := context.Background()

			ruleSvc := svc.NewLimitRulesSvc(nil)
			rs, err := ruleSvc.CreateNewRule(ctx, tc.req)
			t.Log(rs, err)
		})
	}
}
