package account

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
	mockaccountrepo "github.com/ishanwardhono/transfer-system/test/mock/repository/account"
	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"
)

func init() {
	logger.Init("debug")
}

func Test_service_Register(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mockaccountrepo.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		req     dto.RegisterAccountRequest
		wantErr bool
		mockFn  func()
	}{
		{
			name: "error inserting account to database",
			req: dto.RegisterAccountRequest{
				AccountID:      123,
				InitialBalance: decimal.NewFromInt(1000),
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().InsertAccount(ctx, model.Account{
					ID:      123,
					Balance: decimal.NewFromInt(1000),
				}).Return(errors.New("mock error"))
			},
		},
		{
			name: "success",
			req: dto.RegisterAccountRequest{
				AccountID:      123,
				InitialBalance: decimal.NewFromInt(1000),
			},
			wantErr: false,
			mockFn: func() {
				mockAccountRepo.EXPECT().InsertAccount(ctx, model.Account{
					ID:      123,
					Balance: decimal.NewFromInt(1000),
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn()
			}
			s := NewService(mockAccountRepo)
			if err := s.Register(ctx, tt.req); (err != nil) != tt.wantErr {
				t.Errorf("service.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetById(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mockaccountrepo.NewMockRepository(ctrl)

	// Create test account with ID 123 and Balance 1000
	testAccount := model.Account{
		ID:        123,
		Balance:   decimal.NewFromInt(1000),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name      string
		accountID int64
		want      dto.GetAccountByIdResponse
		wantErr   bool
		mockFn    func()
	}{
		{
			name:      "error getting account from database",
			accountID: 123,
			want:      dto.GetAccountByIdResponse{},
			wantErr:   true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, int64(123)).Return(model.Account{}, errors.New("mock error"))
			},
		},
		{
			name:      "success",
			accountID: 123,
			want: dto.GetAccountByIdResponse{
				AccountID: 123,
				Balance:   decimal.NewFromInt(1000),
			},
			wantErr: false,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, int64(123)).Return(testAccount, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn()
			}
			s := NewService(mockAccountRepo)
			got, err := s.GetById(ctx, tt.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
