package sst

const (
	SstMagic   = uint64(0xe489f8a9d479536b)
	MaxKeySize = 8 * 1024
)

const (
	typeNil   = 1
	typeBytes = 2
)
