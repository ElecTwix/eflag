package eflag_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ElecTwix/eflag"
	"github.com/ElecTwix/eflag/flag"
)

func TestEflagParser(t *testing.T) {
	handler := eflag.New()
	testVal := 15
	err := handler.AddFlag(flag.New("test", testVal))
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"./binary", "-test"}

	flags, err := handler.ParseRaw(args[1:])
	if err != nil {
		t.Fatal(err)
	}

	if len(flags) != 1 {
		t.Fatal("return flags need to be 1")
	}
}

func TestDataChange(t *testing.T) {
	handler := eflag.New()
	testVal, expectedVal := 10, 15
	err := handler.AddFlag(flag.New("test", testVal).AddInput(reflect.Int))
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"./binary", "-test", fmt.Sprint(expectedVal)}

	flags, err := handler.ParseRaw(args[1:])
	if err != nil {
		t.Fatal(err)
	}

	if len(flags) != 1 {
		t.Fatal("return flags need to be 1")
	}

	flag := flags[0]

	if flag.Data == testVal {
		t.Fatalf("data is not changed data: %v", testVal)
	}

	if flag.Data != expectedVal {
		t.Fatalf("expected %v got %v", expectedVal, flag.Data)
	}
}

func TestDefualtValue(t *testing.T) {
	handler := eflag.New()
	testVal := 10
	err := handler.AddFlag(flag.New("test", testVal))
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"./binary", "-test"}

	flags, err := handler.ParseRaw(args[1:])
	if err != nil {
		t.Fatal(err)
	}

	if len(flags) != 1 {
		t.Fatal("return flags need to be 1")
	}

	flag := flags[0]

	if flag.Data != testVal {
		t.Fatalf("expected %v got %v", testVal, flag.Data)
	}
}

func TestWrongValue(t *testing.T) {
	handler := eflag.New()
	testVal := 10
	err := handler.AddFlag(flag.New("test", testVal).AddInput(reflect.Int))
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"./binary", "-test", "1.1"}

	_, err = handler.ParseRaw(args[1:])
	if err == nil {
		t.Fatal("should give error when parse float to int")
	}
}
