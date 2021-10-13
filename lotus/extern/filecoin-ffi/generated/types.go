// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#cgo pkg-config: ${SRCDIR}/../filcrypto.pc
#include "../filcrypto.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// FilBLSSignature as declared in filecoin-ffi/filcrypto.h:65
type FilBLSSignature struct {
	Inner          [96]byte
	refa2ac09ba    *C.fil_BLSSignature
	allocsa2ac09ba interface{}
}

// FilAggregateResponse as declared in filecoin-ffi/filcrypto.h:72
type FilAggregateResponse struct {
	Signature      FilBLSSignature
	refb3efa36d    *C.fil_AggregateResponse
	allocsb3efa36d interface{}
}

// FilAggregateProof as declared in filecoin-ffi/filcrypto.h:79
type FilAggregateProof struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ProofLen       uint
	ProofPtr       []byte
	ref22b6c4f6    *C.fil_AggregateProof
	allocs22b6c4f6 interface{}
}

// Fil32ByteArray as declared in filecoin-ffi/filcrypto.h:83
type Fil32ByteArray struct {
	Inner          [32]byte
	ref373ec61a    *C.fil_32ByteArray
	allocs373ec61a interface{}
}

// FilAggregationInputs as declared in filecoin-ffi/filcrypto.h:91
type FilAggregationInputs struct {
	CommR          Fil32ByteArray
	CommD          Fil32ByteArray
	SectorId       uint64
	Ticket         Fil32ByteArray
	Seed           Fil32ByteArray
	ref90b967c9    *C.fil_AggregationInputs
	allocs90b967c9 interface{}
}

// FilSealCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:100
type FilSealCommitPhase2Response struct {
	StatusCode      FCPResponseStatus
	ErrorMsg        string
	ProofPtr        []byte
	ProofLen        uint
	CommitInputsPtr []FilAggregationInputs
	CommitInputsLen uint
	ref5860b9a4     *C.fil_SealCommitPhase2Response
	allocs5860b9a4  interface{}
}

// FilClearCacheResponse as declared in filecoin-ffi/filcrypto.h:105
type FilClearCacheResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	refa9a80400    *C.fil_ClearCacheResponse
	allocsa9a80400 interface{}
}

// FilZeroSignatureResponse as declared in filecoin-ffi/filcrypto.h:112
type FilZeroSignatureResponse struct {
	Signature      FilBLSSignature
	ref835a0405    *C.fil_ZeroSignatureResponse
	allocs835a0405 interface{}
}

// FilFauxRepResponse as declared in filecoin-ffi/filcrypto.h:118
type FilFauxRepResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	Commitment     [32]byte
	refaa003f71    *C.fil_FauxRepResponse
	allocsaa003f71 interface{}
}

// FilFinalizeTicketResponse as declared in filecoin-ffi/filcrypto.h:124
type FilFinalizeTicketResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	Ticket         [32]byte
	refb370fa86    *C.fil_FinalizeTicketResponse
	allocsb370fa86 interface{}
}

// FilGenerateDataCommitmentResponse as declared in filecoin-ffi/filcrypto.h:130
type FilGenerateDataCommitmentResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	CommD          [32]byte
	ref87da7dd9    *C.fil_GenerateDataCommitmentResponse
	allocs87da7dd9 interface{}
}

// FilGenerateFallbackSectorChallengesResponse as declared in filecoin-ffi/filcrypto.h:140
type FilGenerateFallbackSectorChallengesResponse struct {
	ErrorMsg         string
	StatusCode       FCPResponseStatus
	IdsPtr           []uint64
	IdsLen           uint
	ChallengesPtr    []uint64
	ChallengesLen    uint
	ChallengesStride uint
	ref7047a3fa      *C.fil_GenerateFallbackSectorChallengesResponse
	allocs7047a3fa   interface{}
}

// FilGeneratePieceCommitmentResponse as declared in filecoin-ffi/filcrypto.h:151
type FilGeneratePieceCommitmentResponse struct {
	StatusCode      FCPResponseStatus
	ErrorMsg        string
	CommP           [32]byte
	NumBytesAligned uint64
	ref4b00fda4     *C.fil_GeneratePieceCommitmentResponse
	allocs4b00fda4  interface{}
}

// FilVanillaProof as declared in filecoin-ffi/filcrypto.h:156
type FilVanillaProof struct {
	ProofLen       uint
	ProofPtr       []byte
	refb3e7638c    *C.fil_VanillaProof
	allocsb3e7638c interface{}
}

// FilGenerateSingleVanillaProofResponse as declared in filecoin-ffi/filcrypto.h:162
type FilGenerateSingleVanillaProofResponse struct {
	ErrorMsg       string
	VanillaProof   FilVanillaProof
	StatusCode     FCPResponseStatus
	reff9d21b04    *C.fil_GenerateSingleVanillaProofResponse
	allocsf9d21b04 interface{}
}

// FilPoStProof as declared in filecoin-ffi/filcrypto.h:168
type FilPoStProof struct {
	RegisteredProof FilRegisteredPoStProof
	ProofLen        uint
	ProofPtr        []byte
	ref3451bfa      *C.fil_PoStProof
	allocs3451bfa   interface{}
}

// FilGenerateWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:177
type FilGenerateWindowPoStResponse struct {
	ErrorMsg         string
	ProofsLen        uint
	ProofsPtr        []FilPoStProof
	FaultySectorsLen uint
	FaultySectorsPtr []uint64
	StatusCode       FCPResponseStatus
	ref2a5f3ba8      *C.fil_GenerateWindowPoStResponse
	allocs2a5f3ba8   interface{}
}

// FilGenerateWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:184
type FilGenerateWinningPoStResponse struct {
	ErrorMsg       string
	ProofsLen      uint
	ProofsPtr      []FilPoStProof
	StatusCode     FCPResponseStatus
	ref1405b8ec    *C.fil_GenerateWinningPoStResponse
	allocs1405b8ec interface{}
}

// FilGenerateWinningPoStSectorChallenge as declared in filecoin-ffi/filcrypto.h:191
type FilGenerateWinningPoStSectorChallenge struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	IdsPtr         []uint64
	IdsLen         uint
	ref69d2a405    *C.fil_GenerateWinningPoStSectorChallenge
	allocs69d2a405 interface{}
}

// FilGpuDeviceResponse as declared in filecoin-ffi/filcrypto.h:198
type FilGpuDeviceResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	DevicesLen     uint
	DevicesPtr     []string
	ref58f92915    *C.fil_GpuDeviceResponse
	allocs58f92915 interface{}
}

// FilBLSDigest as declared in filecoin-ffi/filcrypto.h:202
type FilBLSDigest struct {
	Inner          [96]byte
	ref215fc78c    *C.fil_BLSDigest
	allocs215fc78c interface{}
}

// FilHashResponse as declared in filecoin-ffi/filcrypto.h:209
type FilHashResponse struct {
	Digest         FilBLSDigest
	refc52a22ef    *C.fil_HashResponse
	allocsc52a22ef interface{}
}

// FilInitLogFdResponse as declared in filecoin-ffi/filcrypto.h:214
type FilInitLogFdResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref3c1a0a08    *C.fil_InitLogFdResponse
	allocs3c1a0a08 interface{}
}

// FilBLSPrivateKey as declared in filecoin-ffi/filcrypto.h:218
type FilBLSPrivateKey struct {
	Inner          [32]byte
	ref2f77fe3a    *C.fil_BLSPrivateKey
	allocs2f77fe3a interface{}
}

// FilPrivateKeyGenerateResponse as declared in filecoin-ffi/filcrypto.h:225
type FilPrivateKeyGenerateResponse struct {
	PrivateKey    FilBLSPrivateKey
	ref2dba09f    *C.fil_PrivateKeyGenerateResponse
	allocs2dba09f interface{}
}

// FilBLSPublicKey as declared in filecoin-ffi/filcrypto.h:229
type FilBLSPublicKey struct {
	Inner          [48]byte
	ref6d0cab13    *C.fil_BLSPublicKey
	allocs6d0cab13 interface{}
}

// FilPrivateKeyPublicKeyResponse as declared in filecoin-ffi/filcrypto.h:236
type FilPrivateKeyPublicKeyResponse struct {
	PublicKey      FilBLSPublicKey
	refee14e59d    *C.fil_PrivateKeyPublicKeyResponse
	allocsee14e59d interface{}
}

// FilPrivateKeySignResponse as declared in filecoin-ffi/filcrypto.h:243
type FilPrivateKeySignResponse struct {
	Signature      FilBLSSignature
	refcdf97b28    *C.fil_PrivateKeySignResponse
	allocscdf97b28 interface{}
}

// FilSealCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:250
type FilSealCommitPhase1Response struct {
	StatusCode                FCPResponseStatus
	ErrorMsg                  string
	SealCommitPhase1OutputPtr []byte
	SealCommitPhase1OutputLen uint
	ref61ed8561               *C.fil_SealCommitPhase1Response
	allocs61ed8561            interface{}
}

// FilSealPreCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:257
type FilSealPreCommitPhase1Response struct {
	ErrorMsg                     string
	StatusCode                   FCPResponseStatus
	SealPreCommitPhase1OutputPtr []byte
	SealPreCommitPhase1OutputLen uint
	ref132bbfd8                  *C.fil_SealPreCommitPhase1Response
	allocs132bbfd8               interface{}
}

// FilSealPreCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:265
type FilSealPreCommitPhase2Response struct {
	ErrorMsg        string
	StatusCode      FCPResponseStatus
	RegisteredProof FilRegisteredSealProof
	CommD           [32]byte
	CommR           [32]byte
	ref2aa6831d     *C.fil_SealPreCommitPhase2Response
	allocs2aa6831d  interface{}
}

// FilStringResponse as declared in filecoin-ffi/filcrypto.h:274
type FilStringResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	StringVal      string
	ref4f413043    *C.fil_StringResponse
	allocs4f413043 interface{}
}

// FilUnsealRangeResponse as declared in filecoin-ffi/filcrypto.h:279
type FilUnsealRangeResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref61e219c9    *C.fil_UnsealRangeResponse
	allocs61e219c9 interface{}
}

// FilVerifyAggregateSealProofResponse as declared in filecoin-ffi/filcrypto.h:285
type FilVerifyAggregateSealProofResponse struct {
	StatusCode    FCPResponseStatus
	ErrorMsg      string
	IsValid       bool
	ref66180e0    *C.fil_VerifyAggregateSealProofResponse
	allocs66180e0 interface{}
}

// FilVerifySealResponse as declared in filecoin-ffi/filcrypto.h:291
type FilVerifySealResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refd4397079    *C.fil_VerifySealResponse
	allocsd4397079 interface{}
}

// FilVerifyWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:297
type FilVerifyWindowPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	ref34c4d49f    *C.fil_VerifyWindowPoStResponse
	allocs34c4d49f interface{}
}

// FilVerifyWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:303
type FilVerifyWinningPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refaca6860c    *C.fil_VerifyWinningPoStResponse
	allocsaca6860c interface{}
}

// FilWriteWithAlignmentResponse as declared in filecoin-ffi/filcrypto.h:311
type FilWriteWithAlignmentResponse struct {
	CommP                 [32]byte
	ErrorMsg              string
	LeftAlignmentUnpadded uint64
	StatusCode            FCPResponseStatus
	TotalWriteUnpadded    uint64
	refa330e79            *C.fil_WriteWithAlignmentResponse
	allocsa330e79         interface{}
}

// FilWriteWithoutAlignmentResponse as declared in filecoin-ffi/filcrypto.h:318
type FilWriteWithoutAlignmentResponse struct {
	CommP              [32]byte
	ErrorMsg           string
	StatusCode         FCPResponseStatus
	TotalWriteUnpadded uint64
	refc8e1ed8         *C.fil_WriteWithoutAlignmentResponse
	allocsc8e1ed8      interface{}
}

// FilPublicPieceInfo as declared in filecoin-ffi/filcrypto.h:323
type FilPublicPieceInfo struct {
	NumBytes       uint64
	CommP          [32]byte
	refd00025ac    *C.fil_PublicPieceInfo
	allocsd00025ac interface{}
}

// FilPrivateReplicaInfo as declared in filecoin-ffi/filcrypto.h:331
type FilPrivateReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CacheDirPath    string
	CommR           [32]byte
	ReplicaPath     string
	SectorId        uint64
	ref81a31e9b     *C.fil_PrivateReplicaInfo
	allocs81a31e9b  interface{}
}

// FilPublicReplicaInfo as declared in filecoin-ffi/filcrypto.h:337
type FilPublicReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CommR           [32]byte
	SectorId        uint64
	ref81b617c2     *C.fil_PublicReplicaInfo
	allocs81b617c2  interface{}
}
