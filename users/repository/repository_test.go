package repository

import (
	"context"
	"testing"
)

func TestUserRepository_Count(t *testing.T) {
	repo := NewUserRepository()
	resultNonCondition, errNonCondition := repo.Count(context.Background(), "")
	if errNonCondition != nil{
		t.Error("Error in empty condition")
	}
	t.Log(resultNonCondition)
}