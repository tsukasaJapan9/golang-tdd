package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is test"}

	assert := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	assertErr := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is test"
		assert(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("hoge")
		if err == nil {
			t.Error("expected error")
		}
		assertErr(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	assertErr := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	assertDesc := func(t *testing.T, dict Dictionary, word, wantDesc string) {
		got, err := dict.Search(word)
		if err != nil {
			t.Error("add word")
		}

		if got != wantDesc {
			t.Errorf("got:%s, want:%s", got, wantDesc)
		}
	}

	t.Run("add new word", func(t *testing.T) {
		dict := Dictionary{}

		word := "test"
		desc := "this is test"

		err := dict.Add(word, desc)
		assertErr(t, err, nil)
		assertDesc(t, dict, word, desc)
	})

	t.Run("add exist word", func(t *testing.T) {
		dict := Dictionary{"test":"this is test"}
		word := "test"
		desc := "this is test"

		err := dict.Add(word, desc)

		assertErr(t, err, ErrExist)
		assertDesc(t, dict, word, desc)
	})

}

func TestUpdate(t *testing.T) {
	assertErr := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	assertDesc := func(t *testing.T, dict Dictionary, word, wantDesc string) {
		got, err := dict.Search(word)
		if err != nil {
			t.Error("add word")
		}

		if got != wantDesc {
			t.Errorf("got:%s, want:%s", got, wantDesc)
		}
	}

	t.Run("dict update exist word", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}

		word := "test"
		desc := "this is not test"

		err := dict.Update(word, desc)
		assertErr(t, err, nil)
		assertDesc(t, dict, word, desc)
	})

	t.Run("dict update new word", func(t *testing.T) {
		dict := Dictionary{}

		word := "test"
		desc := "this is test"

		err := dict.Update(word, desc)
		assertErr(t, err, ErrNotExist)
	})

}

func TestDelete(t *testing.T) {
	assertErr := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	t.Run("delete exist word", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}

		word := "test"

		err := dict.Delete(word)
		assertErr(t, err, nil)
		_, err = dict.Search(word)
		assertErr(t, err, ErrNotFound)
	})

	t.Run("delete not exist word", func(t *testing.T) {
		dict := Dictionary{}

		word := "test"

		err := dict.Delete(word)
		assertErr(t, err, ErrNotExist)
		_, err = dict.Search(word)
		assertErr(t, err, ErrNotFound)
	})
}