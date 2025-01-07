// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lighthouse

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MemoryClose is the `MemoryClose` instruction.
type MemoryClose struct {
	MemoryId   *uint8
	MemoryBump *uint8

	// [0] = [] programId
	// ··········· Lighthouse program
	//
	// [1] = [WRITE, SIGNER] payer
	// ··········· Payer account
	//
	// [2] = [WRITE] memory
	// ··········· Memory account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMemoryCloseInstructionBuilder creates a new `MemoryClose` instruction builder.
func NewMemoryCloseInstructionBuilder() *MemoryClose {
	nd := &MemoryClose{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMemoryId sets the "memoryId" parameter.
func (inst *MemoryClose) SetMemoryId(memoryId uint8) *MemoryClose {
	inst.MemoryId = &memoryId
	return inst
}

// SetMemoryBump sets the "memoryBump" parameter.
func (inst *MemoryClose) SetMemoryBump(memoryBump uint8) *MemoryClose {
	inst.MemoryBump = &memoryBump
	return inst
}

// SetProgramIdAccount sets the "programId" account.
// Lighthouse program
func (inst *MemoryClose) SetProgramIdAccount(programId ag_solanago.PublicKey) *MemoryClose {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(programId)
	return inst
}

// GetProgramIdAccount gets the "programId" account.
// Lighthouse program
func (inst *MemoryClose) GetProgramIdAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
// Payer account
func (inst *MemoryClose) SetPayerAccount(payer ag_solanago.PublicKey) *MemoryClose {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer account
func (inst *MemoryClose) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMemoryAccount sets the "memory" account.
// Memory account
func (inst *MemoryClose) SetMemoryAccount(memory ag_solanago.PublicKey) *MemoryClose {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(memory).WRITE()
	return inst
}

// GetMemoryAccount gets the "memory" account.
// Memory account
func (inst *MemoryClose) GetMemoryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst MemoryClose) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MemoryClose,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MemoryClose) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MemoryClose) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MemoryId == nil {
			return errors.New("MemoryId parameter is not set")
		}
		if inst.MemoryBump == nil {
			return errors.New("MemoryBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ProgramId is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Memory is not set")
		}
	}
	return nil
}

func (inst *MemoryClose) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MemoryClose")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  MemoryId", *inst.MemoryId))
						paramsBranch.Child(ag_format.Param("MemoryBump", *inst.MemoryBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("programId", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   memory", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj MemoryClose) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MemoryId` param:
	err = encoder.Encode(obj.MemoryId)
	if err != nil {
		return err
	}
	// Serialize `MemoryBump` param:
	err = encoder.Encode(obj.MemoryBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MemoryClose) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MemoryId`:
	err = decoder.Decode(&obj.MemoryId)
	if err != nil {
		return err
	}
	// Deserialize `MemoryBump`:
	err = decoder.Decode(&obj.MemoryBump)
	if err != nil {
		return err
	}
	return nil
}

// NewMemoryCloseInstruction declares a new MemoryClose instruction with the provided parameters and accounts.
func NewMemoryCloseInstruction(
	// Parameters:
	memoryId uint8,
	memoryBump uint8,
	// Accounts:
	programId ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	memory ag_solanago.PublicKey) *MemoryClose {
	return NewMemoryCloseInstructionBuilder().
		SetMemoryId(memoryId).
		SetMemoryBump(memoryBump).
		SetProgramIdAccount(programId).
		SetPayerAccount(payer).
		SetMemoryAccount(memory)
}
