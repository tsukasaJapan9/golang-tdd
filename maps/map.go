package main

const (
	ErrNotFound = DictErr("word is not found")
	ErrExist    = DictErr("word is exist")
	ErrNotExist = DictErr("word is not exist")
)

type DictErr string
func (e DictErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	str, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return str, nil
}

func (d Dictionary) Add(word string, desc string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = desc
	case nil:
		return ErrExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, desc string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		d[word] = desc
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}