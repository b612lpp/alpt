package repository

import (
	"gb-go-architecture/lesson-1/shop_new/models"
	"gb-go-architecture/lesson-1/shop_new/repository"
	"testing"
	"time"
)

func TestCreateItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}

	result, err = db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}

	input = &models.Item{
		Name:  "someName2",
		Price: 20,
	}
	expected = &models.Item{
		ID:    2,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err = db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}
}

//из-за того что база живет в рантайме не тестируем поля время
func TestListItem(t *testing.T) {
	input := repository.ItemFilter{
		PriceLeft:  10,
		PriceRight: 15,
		Limit:      5,
		Offset:     0,
	}
	expected := []*models.Item{{1, "name", 100, time.Time{}, time.Time{}}}
	result, err := Repository.ListItems(input)
	if err != nil {
		t.Error("unexp error %s", err)
	}
	if result[0].ID != expected[0].ID {
		t.Error("wrong ID was waiting %d got %d")
	}
	if result[0].Name != expected[0].Name {
		t.Error("wrong Name was waiting %s got %s")
	}

}

func TestGetItem(t *testing.T) {
	input := 1
	expected := &models.Item{1, "name", 100, time.Time{}, time.Time{}}
	result, err := repository.NewMapDB().GetItem(input)
	if err != nil {
		t.Error("unexp error %s", err)
	}
	if expected != result {
		t.Error("wrong item expected %v got %v", expected, result)
	}

}

func TestUpdateItem(t *testing.T) {
	input := models.Item{99, "newname", 200, time.Time{}, time.Time{}}
	expected := models.Item{99, "newname", 200, time.Time{}, time.Time{}}
	result := repository.NewMapDB().UpdateItem(input)
	if result != expected {
		t.Error("was waiting %v got %v ", input, result)
	}
}

func TestDeleteItem(t testing.T) {
	input := 1
	//expected:= nil
	result := repository.NewMapDB().DeleteItem(input)
	if result != nil && repository.NewMapDB().GetItem(input) != nil {
		t.Error("element was not deleted %v", repository.NewMapDB().GetItem(input))
	}

}
