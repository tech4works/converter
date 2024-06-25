package converter

func CanToBytes(a any) bool {
	_, err := ToBytesWithErr(a)
	return err == nil
}

func ToBytes(a any) []byte {
	bs, err := ToBytesWithErr(a)
	if err != nil {
		panic(err)
	}
	return bs
}

func ToBytesWithErr(a any) ([]byte, error) {
	s, err := ToStringWithErr(a)
	return []byte(s), err
}
