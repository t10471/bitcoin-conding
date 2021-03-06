

// Code generated by "cofing -t MsgBlock"; DO NOT EDIT.
package message

import (
	"bytes"

	"github.com/t10471/bitcoin-coding/basetype"
)

func (m *MsgBlock) Decode(b_ *bytes.Buffer) error {
	// Header
	if err := m.Header.Decode(b_); err != nil {
		return err
	} 
	// TxnCount
	{
		var err error
		m.TxnCount, err = basetype.DecodeVarInt(b_)
		if err != nil {
			return err
		}
	} 
	// Txn
	{
		m_ := int(m.TxnCount)
		m.Txn = make([]MsgTx, m_)
		for i_ := 0; i_ < m_; i_++{
			if err := m.Txn[i_].Decode(b_); err != nil {
				return err
			}
		}
	}  
	return nil
}

func (m *MsgBlock) Encode(b_ *bytes.Buffer) error {
	// Header
	if err := m.Header.Encode(b_); err != nil {
		return err
	} 
	// TxnCount
	if err := basetype.EncodeVarInt(b_, m.TxnCount); err != nil {
		return err
	} 
	// Txn
	{
		m_ := int(m.TxnCount)
		for i_ := 0; i_ < m_; i_++ {
			if err := m.Txn[i_].Encode(b_); err != nil {
				return err
			}
		}
	}  
	return nil
}
