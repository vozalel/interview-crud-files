package unit_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/internal/entity/mock"
	"github.com/vozalel/interview-crud-files/internal/usecase"
	"testing"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

var (
	adminID          = 1
	permissionDenyID = 2
	simpleUserID     = 3
)

var (
	testData = "testData"
)

func datasourceCrudUC(t *testing.T) (
	entity.IDatasourceUC,
	*mock.MockIManagerACL,
	*mock.MockIManagerDatasource) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	aclManager := mock.NewMockIManagerACL(mockCtl)
	datasourceManager := mock.NewMockIManagerDatasource(mockCtl)

	uc := usecase.New(aclManager, datasourceManager)
	return uc, aclManager, datasourceManager
}

func TestCreate(t *testing.T) {
	datasourceUC, aclManager, _ := datasourceCrudUC(t)

	tests := []test{
		{
			name: "permission deny",
			mock: func() {
				aclManager.EXPECT().GetUserPerformACL(context.Background(), &entity.User{
					Name: "nonPermission",
					ID:   &permissionDenyID,
				}).Return(entity.PerformACL{
					Create: false,
					List:   false,
				}, nil)
			},
			res: nil,
			err: usecase.ErrorPermissionDenied,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			err := datasourceUC.CreateDataSource(
				context.Background(),
				&entity.User{
					Name: "nonPermission",
					ID:   &permissionDenyID,
				},
				&entity.Datasource{
					Name: "test permission deny",
					Data: &testData,
				},
			)

			require.ErrorIs(t, err, tc.err)
		})
	}
}
