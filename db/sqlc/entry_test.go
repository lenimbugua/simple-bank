package db

import (
	"context"
	"testing"
	"time"

	"github.com/lenimbugua/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomAmount(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	require.WithinDuration(t, time.Now(), entry.CreatedAt, 1*time.Second)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, 1*time.Second)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
}
