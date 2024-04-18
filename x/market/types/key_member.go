package types

const (
	// MemberKeyPrefix is the prefix to retrieve all Member
	MemberKeyPrefix = "Member/value/"
)

// MemberKey returns the store key to retrieve a Member from the index fields
func MemberKey(
	denomA string,
	denomB string,
) []byte {
	var key []byte

	denomABytes := []byte(denomA)
	key = append(key, denomABytes...)
	key = append(key, []byte("/")...)

	denomBBytes := []byte(denomB)
	key = append(key, denomBBytes...)
	key = append(key, []byte("/")...)

	return key
}
