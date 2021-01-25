package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestHelloByName(t *testing.T) {
    got := HelloByName("Michael")
    want := "Hello, Michael"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestMain(t *testing.T) {
    main()
}

