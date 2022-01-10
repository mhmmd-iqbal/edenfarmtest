package main

import (
	"testing"
)

func TestSatu(t *testing.T){
    input := []int{1,2,3,8,9,3,2,1}

    got := mainProgram(input)
    want := 3

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestDua(t *testing.T){
    input := []int{1,2,1,2,2,1}

    got := mainProgram(input)
    want := 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestTiga(t *testing.T){
    input := []int{7,1,2,9,7,2,1}

    got := mainProgram(input)
    want := 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
