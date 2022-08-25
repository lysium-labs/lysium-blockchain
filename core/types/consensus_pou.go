// Copyright 2022 The lysium-dev Authors
// This file is part of the lysium-dev library.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// txJSON is the JSON representation of transactions.
type proofOfUsageStruct struct {
	Type hexutil.Uint64 `json:"type"`

	// Common transaction fields:
	Nonce    *hexutil.Uint64 `json:"nonce"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Gas      *hexutil.Uint64 `json:"gas"`
	Value    *hexutil.Big    `json:"value"`
	Data     *hexutil.Bytes  `json:"input"`
	V        *hexutil.Big    `json:"v"`
	R        *hexutil.Big    `json:"r"`
	S        *hexutil.Big    `json:"s"`
	To       *common.Address `json:"to"`

	// Access list transaction fields:
	ChainID    *hexutil.Big `json:"chainId,omitempty"`
	AccessList *AccessList  `json:"accessList,omitempty"`

	// Only used for encoding:
	Hash common.Hash `json:"hash"`
}

// MarshalJSON marshals as JSON with a hash.
func (t *Transaction) generateProofOfUsage() ([]byte, error) {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = t.Hash()
	enc.Type = hexutil.Uint64(t.Type())

	// Other fields are set conditionally depending on tx type.
	switch tx := t.inner.(type) {
	case *LegacyTx:
		enc.Nonce = (*hexutil.Uint64)(&tx.Nonce)
		enc.Gas = (*hexutil.Uint64)(&tx.Gas)
		enc.GasPrice = (*hexutil.Big)(tx.GasPrice)
		enc.Value = (*hexutil.Big)(tx.Value)
		enc.Data = (*hexutil.Bytes)(&tx.Data)
		enc.To = t.To()
		enc.V = (*hexutil.Big)(tx.V)
		enc.R = (*hexutil.Big)(tx.R)
		enc.S = (*hexutil.Big)(tx.S)
	case *AccessListTx:
		enc.ChainID = (*hexutil.Big)(tx.ChainID)
		enc.AccessList = &tx.AccessList
		enc.Nonce = (*hexutil.Uint64)(&tx.Nonce)
		enc.Gas = (*hexutil.Uint64)(&tx.Gas)
		enc.GasPrice = (*hexutil.Big)(tx.GasPrice)
		enc.Value = (*hexutil.Big)(tx.Value)
		enc.Data = (*hexutil.Bytes)(&tx.Data)
		enc.To = t.To()
		enc.V = (*hexutil.Big)(tx.V)
		enc.R = (*hexutil.Big)(tx.R)
		enc.S = (*hexutil.Big)(tx.S)
	}
	return json.Marshal(&enc)
}

func (w *trezorDriver) Open(device io.ReadWriter, passphrase string) error {
	w.device, w.failure = device, nil

	// If phase 1 is requested, init the connection and wait for user callback
	if passphrase == "" && !w.passphrasewait {
		// If we're already waiting for a PIN entry, insta-return
		if w.pinwait {
			return ErrTrezorPINNeeded
		}
		// Initialize a connection to the device
		features := new(trezor.Features)
		if _, err := w.trezorExchange(&trezor.Initialize{}, features); err != nil {
			return err
		}
		w.version = [3]uint32{features.GetMajorVersion(), features.GetMinorVersion(), features.GetPatchVersion()}
		w.label = features.GetLabel()

		// Do a manual ping, forcing the device to ask for its PIN and Passphrase
		askPin := true
		askPassphrase := true
		res, err := w.trezorExchange(&trezor.Ping{PinProtection: &askPin, PassphraseProtection: &askPassphrase}, new(trezor.PinMatrixRequest), new(trezor.PassphraseRequest), new(trezor.Success))
		if err != nil {
			return err
		}
		// Only return the PIN request if the device wasn't unlocked until now
		switch res {
		case 0:
			w.pinwait = true
			return ErrTrezorPINNeeded
		case 1:
			w.pinwait = false
			w.passphrasewait = true
			return ErrTrezorPassphraseNeeded
		case 2:
			return nil // responded with trezor.Success
		}
	}
	// Phase 2 requested with actual PIN entry
	if w.pinwait {
		w.pinwait = false
		res, err := w.trezorExchange(&trezor.PinMatrixAck{Pin: &passphrase}, new(trezor.Success), new(trezor.PassphraseRequest))
		if err != nil {
			w.failure = err
			return err
		}
		if res == 1 {
			w.passphrasewait = true
			return ErrTrezorPassphraseNeeded
		}
	} else if w.passphrasewait {
		w.passphrasewait = false
		if _, err := w.trezorExchange(&trezor.PassphraseAck{Passphrase: &passphrase}, new(trezor.Success)); err != nil {
			w.failure = err
			return err
		}
	}

	return nil
}
