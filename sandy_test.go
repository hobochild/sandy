package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	s := []string{"password.txt"}
	patterns := []string{""}
	reqs, err := Exec("cat", s, patterns, patterns)

	if err != nil {
		t.Errorf("Something went wrong")
	}

	if len(reqs) != 2 {
		t.Errorf("reqs count was incorrect, got: %d, want: %d.", len(reqs), 2)
	}
}

func TestInput(t *testing.T) {
	cmd := exec.Command("./sandy", "cat", "./password.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Errorf("Something went wrong")
	}
	if !strings.Contains(out.String(), "Blocked READ on ") {
		t.Errorf("Expected %s output got %s", "123", out.String())
	}
}

func TestAllowList(t *testing.T) {
	cmd := exec.Command("./sandy", "--y", "*.so", "--y", "*.txt", "cat", "./password.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Errorf("Something went wrong")
	}
	if out.String() != "123\n" {
		t.Errorf("Expected %s output got %s", "123", out.String())
	}
}

func TestBlockList(t *testing.T) {
	cmd := exec.Command("./sandy", "--y", "*.so", "--n", "*.txt", "cat", "./password.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Errorf("Something went wrong")
	}
	if !strings.Contains(out.String(), "Blocked READ on ") {
		t.Errorf("Expected %s output got %s", "123", out.String())
	}
}

func TestHelp(t *testing.T) {
	cmd := exec.Command("./sandy", "-h")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		t.Errorf("Something went wrong")
	}

	if strings.Contains(out.String(), "Usage of ./sandy:") {
		t.Errorf("Expected %s output got %s", "123", out.String())
	}
}
