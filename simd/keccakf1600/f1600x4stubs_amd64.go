// Code generated by command: go run src.go -out ../../f1600x4_amd64.s -stubs ../../f1600x4stubs_amd64.go -pkg keccakf1600. DO NOT EDIT.

//go:build amd64 && !purego

package keccakf1600

//go:noescape
func f1600x4AVX2(state *uint64, rc *[24]uint64, turbo bool)
