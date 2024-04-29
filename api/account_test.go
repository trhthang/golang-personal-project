package api

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/trhthang/simplebank/db/mock"
	db "github.com/trhthang/simplebank/db/sqlc"
	"github.com/trhthang/simplebank/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	//build stubs
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID))
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
