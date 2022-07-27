package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T){
	cases := []Config{
		{
			Title: "test1",
			Version: "0.0.1-test1",
		},
	}
	for _, c := range cases {
		fmt.Printf("It works, version %s", c.Version)
	}
}