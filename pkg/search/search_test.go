package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {
	ch := All(context.Background(), "Go", []string{"test.txt"})
	results, err := <-ch
	if !err {
		t.Errorf("error: %v", err)
	}
	log.Println("result: ", results)
}

func TestAll_singleOne(t *testing.T) {
	ch := All(context.Background(), "golang", []string{"test.txt"})
	results, err := <-ch
	if !err {
		t.Errorf("error: %v", err)
	}
	log.Println("result: ", results)
}

func TestAll_notFound(t *testing.T) {
	ch := All(context.Background(), "alif", []string{"test.txt"})
	results, err := <-ch
	if err != false {
		t.Errorf("error: %v", err)
	}
	log.Println("result: ", results)
}

func TestAny_positiv(t *testing.T) {
	ch := Any(context.Background(), "golang", []string{"test.txt"})
	results, err := <-ch
	if !err {
		t.Errorf("error: %v", err)
	}
	log.Println("result: ", results)
}

func TestAny_negative(t *testing.T) {
	ch := Any(context.Background(), "alif", []string{"test.txt"})
	results, err := <-ch
	if err == true {
		t.Errorf("error: %v", err)
	}
	log.Println("result: ", results)
}

func TestAny_multiSearch(t *testing.T) {
	ch := Any(context.Background(), "Go", []string{"test.txt"})
	results, err := <-ch
	if !err {
		t.Errorf("error: %v", err)
	}
	log.Println("result: ", results)
}