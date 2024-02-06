package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestNewSuccess(t *testing.T) {
	// When
	name := "test"
	a, err := item.New(name)
	if err != nil {
		t.Errorf("validation error: %s", err)
	}

	// Then
	if a.GetId().UUID() == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, a.GetId().UUID())
	}
	if a.GetName() != name {
		t.Errorf("expected %s, got %s", name, a.GetName())
	}
	if a.IsDeleted() != false {
		t.Errorf("expected %t, got %t", false, a.IsDeleted())
	}
}

func TestNewFail(t *testing.T) {
	// When
	name := ""
	_, err := item.New(name)
	if err == nil {
		t.Errorf("validation error : %s", err)
	}
}

func TestChangeName(t *testing.T) {
	// Given
	name := "test"
	a, err := item.New(name)
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.ChangeName("test2")

	// Then
	if a.GetName() != "test2" {
		t.Errorf("expected %s, got %s", "test2", a.GetName())
	}
}

func TestDelete(t *testing.T) {
	// When
	name := "test"
	a, err := item.New(name)
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.Delete()

	// Then
	if a.IsDeleted() != true {
		t.Errorf("expected %t, got %t", true, a.IsDeleted())
	}
}
