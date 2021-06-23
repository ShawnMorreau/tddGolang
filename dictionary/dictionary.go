package dictionary

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrInvalidKey       = DictionaryErr("could not find that key in the dictionary")
	ErrKeyAlreadyExists = DictionaryErr("that key is already defined in the dictionary")
)

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrInvalidKey
	}
	return value, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrInvalidKey:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	if _, ok := d[key]; !ok {
		return ErrInvalidKey
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
