package db_test

import (
	"context"
	"github.com/stretchr/testify/require"
	db "simplebank/db/sqlc"
	"simplebank/util"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account db.Account) db.Entry {
	arg := db.CreateEntryParams{
		Amount:    util.RandomMoney(),
		AccountID: account.ID,
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.ID, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	entries, err := testQueries.ListEntries(context.Background(), db.ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	})
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entries)
		require.Equal(t, entry.AccountID, account.ID)
	}
}
