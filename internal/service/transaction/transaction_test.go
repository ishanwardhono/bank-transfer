package transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ishanwardhono/transfer-system/internal/entity/dto"
	"github.com/ishanwardhono/transfer-system/internal/entity/model"
	"github.com/ishanwardhono/transfer-system/pkg/logger"
	mockaccountrepo "github.com/ishanwardhono/transfer-system/test/mock/repository/account"
	mockdbtrxrepo "github.com/ishanwardhono/transfer-system/test/mock/repository/dbtrx"
	mocktransactionrepo "github.com/ishanwardhono/transfer-system/test/mock/repository/transaction"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"
)

func init() {
	logger.Init("debug")
}

func Test_service_Transfer(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountRepo := mockaccountrepo.NewMockRepository(ctrl)
	mockTransactionRepo := mocktransactionrepo.NewMockRepository(ctrl)
	mockDbTrx := mockdbtrxrepo.NewMockRepository(ctrl)

	// Test data
	sourceAccountID := int64(123)
	destAccountID := int64(456)
	amount := decimal.NewFromInt(100)

	sourceAccount := model.Account{
		ID:        sourceAccountID,
		Balance:   decimal.NewFromInt(1000),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	destinationAccount := model.Account{
		ID:        destAccountID,
		Balance:   decimal.NewFromInt(500),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	sourceBalanceAfter := sourceAccount.Balance.Sub(amount)    // 900
	destBalanceAfter := destinationAccount.Balance.Add(amount) // 600
	mockTx := &sqlx.Tx{}

	tests := []struct {
		name    string
		req     dto.TransferRequest
		wantErr bool
		mockFn  func()
	}{
		{
			name: "error getting source account",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(model.Account{}, errors.New("source account error"))
			},
		},
		{
			name: "error getting destination account",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(model.Account{}, errors.New("destination account error"))
			},
		},
		{
			name: "insufficient balance in source account",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               decimal.NewFromInt(2000), // More than source balance
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
			},
		},
		{
			name: "error beginning transaction",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
				mockDbTrx.EXPECT().BeginTx(ctx).Return(nil, errors.New("transaction error"))
			},
		},
		{
			name: "error inserting transaction",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
				mockDbTrx.EXPECT().BeginTx(ctx).Return(mockTx, nil)

				// Expect model with any reference number
				mockTransactionRepo.EXPECT().TxInsertTransaction(ctx, mockTx, gomock.Any()).Return(errors.New("insert error"))
				mockDbTrx.EXPECT().RollbackTx(ctx, mockTx)
			},
		},
		{
			name: "error updating source account balance",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
				mockDbTrx.EXPECT().BeginTx(ctx).Return(mockTx, nil)
				mockTransactionRepo.EXPECT().TxInsertTransaction(ctx, mockTx, gomock.Any()).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, sourceAccountID, sourceBalanceAfter).Return(errors.New("update error"))
				mockDbTrx.EXPECT().RollbackTx(ctx, mockTx)
			},
		},
		{
			name: "error updating destination account balance",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
				mockDbTrx.EXPECT().BeginTx(ctx).Return(mockTx, nil)
				mockTransactionRepo.EXPECT().TxInsertTransaction(ctx, mockTx, gomock.Any()).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, sourceAccountID, sourceBalanceAfter).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, destAccountID, destBalanceAfter).Return(errors.New("update error"))
				mockDbTrx.EXPECT().RollbackTx(ctx, mockTx)
			},
		},
		{
			name: "error committing transaction",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: true,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
				mockDbTrx.EXPECT().BeginTx(ctx).Return(mockTx, nil)
				mockTransactionRepo.EXPECT().TxInsertTransaction(ctx, mockTx, gomock.Any()).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, sourceAccountID, sourceBalanceAfter).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, destAccountID, destBalanceAfter).Return(nil)
				mockDbTrx.EXPECT().CommitTx(mockTx).Return(errors.New("commit error"))
				mockDbTrx.EXPECT().RollbackTx(ctx, mockTx)
			},
		},
		{
			name: "success",
			req: dto.TransferRequest{
				SourceAccountID:      sourceAccountID,
				DestinationAccountID: destAccountID,
				Amount:               amount,
			},
			wantErr: false,
			mockFn: func() {
				mockAccountRepo.EXPECT().GetAccount(ctx, sourceAccountID).Return(sourceAccount, nil)
				mockAccountRepo.EXPECT().GetAccount(ctx, destAccountID).Return(destinationAccount, nil)
				mockDbTrx.EXPECT().BeginTx(ctx).Return(mockTx, nil)
				mockTransactionRepo.EXPECT().TxInsertTransaction(ctx, mockTx, gomock.Any()).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, sourceAccountID, sourceBalanceAfter).Return(nil)
				mockAccountRepo.EXPECT().TxUpdateBalance(ctx, mockTx, destAccountID, destBalanceAfter).Return(nil)
				mockDbTrx.EXPECT().CommitTx(mockTx).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn()
			}
			s := NewService(mockDbTrx, mockTransactionRepo, mockAccountRepo)
			if err := s.Transfer(ctx, tt.req); (err != nil) != tt.wantErr {
				t.Errorf("service.Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
