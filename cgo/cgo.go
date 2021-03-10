// WARNING: This file has automatically been generated on Wed, 10 Mar 2021 12:03:59 IST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package cgo

/*
#cgo LDFLAGS: -L${SRCDIR} -lsolana_ffi
#cgo linux LDFLAGS: -lcrypto -ldl -lm -lrt -lssl -ludev
#cgo darwin LDFLAGS: -F/Library/Frameworks -framework Security -framework CoreServices -framework IOKit -framework IOSurface -framework AppKit
#include "solana-ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// UniquePubkey function as declared in cgo/solana-ffi.h:17
func UniquePubkey() string {
	__ret := C.unique_pubkey()
	__v := packPCharString(__ret)
	return __v
}

// ProgramDerivedAddress function as declared in cgo/solana-ffi.h:19
func ProgramDerivedAddress(seedsPointer []byte, seedsSize uint32, program string) string {
	cseedsPointer, cseedsPointerAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&seedsPointer)))
	cseedsSize, cseedsSizeAllocMap := (C.size_t)(seedsSize), cgoAllocsUnknown
	program = safeString(program)
	cprogram, cprogramAllocMap := unpackPCharString(program)
	__ret := C.program_derived_address(cseedsPointer, cseedsSize, cprogram)
	runtime.KeepAlive(program)
	runtime.KeepAlive(cprogramAllocMap)
	runtime.KeepAlive(cseedsSizeAllocMap)
	runtime.KeepAlive(cseedsPointerAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// GatewayInitialize function as declared in cgo/solana-ffi.h:23
func GatewayInitialize(keypairPath string, rpcUrl string, authorityPointer []byte, selector string) string {
	keypairPath = safeString(keypairPath)
	ckeypairPath, ckeypairPathAllocMap := unpackPCharString(keypairPath)
	rpcUrl = safeString(rpcUrl)
	crpcUrl, crpcUrlAllocMap := unpackPCharString(rpcUrl)
	cauthorityPointer, cauthorityPointerAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&authorityPointer)))
	selector = safeString(selector)
	cselector, cselectorAllocMap := unpackPCharString(selector)
	__ret := C.gateway_initialize(ckeypairPath, crpcUrl, cauthorityPointer, cselector)
	runtime.KeepAlive(selector)
	runtime.KeepAlive(cselectorAllocMap)
	runtime.KeepAlive(cauthorityPointerAllocMap)
	runtime.KeepAlive(rpcUrl)
	runtime.KeepAlive(crpcUrlAllocMap)
	runtime.KeepAlive(keypairPath)
	runtime.KeepAlive(ckeypairPathAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// GatewayInitializeAccount function as declared in cgo/solana-ffi.h:28
func GatewayInitializeAccount(keypairPath string, rpcUrl string, selector string) string {
	keypairPath = safeString(keypairPath)
	ckeypairPath, ckeypairPathAllocMap := unpackPCharString(keypairPath)
	rpcUrl = safeString(rpcUrl)
	crpcUrl, crpcUrlAllocMap := unpackPCharString(rpcUrl)
	selector = safeString(selector)
	cselector, cselectorAllocMap := unpackPCharString(selector)
	__ret := C.gateway_initialize_account(ckeypairPath, crpcUrl, cselector)
	runtime.KeepAlive(selector)
	runtime.KeepAlive(cselectorAllocMap)
	runtime.KeepAlive(rpcUrl)
	runtime.KeepAlive(crpcUrlAllocMap)
	runtime.KeepAlive(keypairPath)
	runtime.KeepAlive(ckeypairPathAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// GatewayGetBurnCount function as declared in cgo/solana-ffi.h:32
func GatewayGetBurnCount(rpcUrl string) uint64 {
	rpcUrl = safeString(rpcUrl)
	crpcUrl, crpcUrlAllocMap := unpackPCharString(rpcUrl)
	__ret := C.gateway_get_burn_count(crpcUrl)
	runtime.KeepAlive(rpcUrl)
	runtime.KeepAlive(crpcUrlAllocMap)
	__v := (uint64)(__ret)
	return __v
}

// GatewayMint function as declared in cgo/solana-ffi.h:34
func GatewayMint(keypairPath string, rpcUrl string, authoritySecret string, selector string, amount uint64) string {
	keypairPath = safeString(keypairPath)
	ckeypairPath, ckeypairPathAllocMap := unpackPCharString(keypairPath)
	rpcUrl = safeString(rpcUrl)
	crpcUrl, crpcUrlAllocMap := unpackPCharString(rpcUrl)
	authoritySecret = safeString(authoritySecret)
	cauthoritySecret, cauthoritySecretAllocMap := unpackPCharString(authoritySecret)
	selector = safeString(selector)
	cselector, cselectorAllocMap := unpackPCharString(selector)
	camount, camountAllocMap := (C.ulonglong)(amount), cgoAllocsUnknown
	__ret := C.gateway_mint(ckeypairPath, crpcUrl, cauthoritySecret, cselector, camount)
	runtime.KeepAlive(camountAllocMap)
	runtime.KeepAlive(selector)
	runtime.KeepAlive(cselectorAllocMap)
	runtime.KeepAlive(authoritySecret)
	runtime.KeepAlive(cauthoritySecretAllocMap)
	runtime.KeepAlive(rpcUrl)
	runtime.KeepAlive(crpcUrlAllocMap)
	runtime.KeepAlive(keypairPath)
	runtime.KeepAlive(ckeypairPathAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// GatewayBurn function as declared in cgo/solana-ffi.h:40
func GatewayBurn(keypairPath string, rpcUrl string, selector string, burnCount uint64, burnAmount uint64, recipientLen uint32, recipientPointer []byte) string {
	keypairPath = safeString(keypairPath)
	ckeypairPath, ckeypairPathAllocMap := unpackPCharString(keypairPath)
	rpcUrl = safeString(rpcUrl)
	crpcUrl, crpcUrlAllocMap := unpackPCharString(rpcUrl)
	selector = safeString(selector)
	cselector, cselectorAllocMap := unpackPCharString(selector)
	cburnCount, cburnCountAllocMap := (C.ulonglong)(burnCount), cgoAllocsUnknown
	cburnAmount, cburnAmountAllocMap := (C.ulonglong)(burnAmount), cgoAllocsUnknown
	crecipientLen, crecipientLenAllocMap := (C.size_t)(recipientLen), cgoAllocsUnknown
	crecipientPointer, crecipientPointerAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&recipientPointer)))
	__ret := C.gateway_burn(ckeypairPath, crpcUrl, cselector, cburnCount, cburnAmount, crecipientLen, crecipientPointer)
	runtime.KeepAlive(crecipientPointerAllocMap)
	runtime.KeepAlive(crecipientLenAllocMap)
	runtime.KeepAlive(cburnAmountAllocMap)
	runtime.KeepAlive(cburnCountAllocMap)
	runtime.KeepAlive(selector)
	runtime.KeepAlive(cselectorAllocMap)
	runtime.KeepAlive(rpcUrl)
	runtime.KeepAlive(crpcUrlAllocMap)
	runtime.KeepAlive(keypairPath)
	runtime.KeepAlive(ckeypairPathAllocMap)
	__v := packPCharString(__ret)
	return __v
}