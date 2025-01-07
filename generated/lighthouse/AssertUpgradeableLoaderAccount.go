// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lighthouse

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AssertUpgradeableLoaderAccount is the `AssertUpgradeableLoaderAccount` instruction.
type AssertUpgradeableLoaderAccount struct {
	LogLevel  *LogLevel
	Assertion *UpgradeableLoaderStateAssertion

	// [0] = [] targetAccount
	// ··········· Target account to be asserted
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAssertUpgradeableLoaderAccountInstructionBuilder creates a new `AssertUpgradeableLoaderAccount` instruction builder.
func NewAssertUpgradeableLoaderAccountInstructionBuilder() *AssertUpgradeableLoaderAccount {
	nd := &AssertUpgradeableLoaderAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetLogLevel sets the "logLevel" parameter.
func (inst *AssertUpgradeableLoaderAccount) SetLogLevel(logLevel LogLevel) *AssertUpgradeableLoaderAccount {
	inst.LogLevel = &logLevel
	return inst
}

// SetAssertion sets the "assertion" parameter.
func (inst *AssertUpgradeableLoaderAccount) SetAssertion(assertion UpgradeableLoaderStateAssertion) *AssertUpgradeableLoaderAccount {
	inst.Assertion = &assertion
	return inst
}

// SetTargetAccountAccount sets the "targetAccount" account.
// Target account to be asserted
func (inst *AssertUpgradeableLoaderAccount) SetTargetAccountAccount(targetAccount ag_solanago.PublicKey) *AssertUpgradeableLoaderAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(targetAccount)
	return inst
}

// GetTargetAccountAccount gets the "targetAccount" account.
// Target account to be asserted
func (inst *AssertUpgradeableLoaderAccount) GetTargetAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

func (inst AssertUpgradeableLoaderAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AssertUpgradeableLoaderAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AssertUpgradeableLoaderAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AssertUpgradeableLoaderAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LogLevel == nil {
			return errors.New("LogLevel parameter is not set")
		}
		if inst.Assertion == nil {
			return errors.New("Assertion parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TargetAccount is not set")
		}
	}
	return nil
}

func (inst *AssertUpgradeableLoaderAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AssertUpgradeableLoaderAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" LogLevel", *inst.LogLevel))
						paramsBranch.Child(ag_format.Param("Assertion", *inst.Assertion))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("target", inst.AccountMetaSlice.Get(0)))
					})
				})
		})
}

func (obj AssertUpgradeableLoaderAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LogLevel` param:
	err = encoder.Encode(obj.LogLevel)
	if err != nil {
		return err
	}
	// Serialize `Assertion` param:
	err = encoder.Encode(obj.Assertion)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AssertUpgradeableLoaderAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LogLevel`:
	err = decoder.Decode(&obj.LogLevel)
	if err != nil {
		return err
	}
	// Deserialize `Assertion`:
	err = decoder.Decode(&obj.Assertion)
	if err != nil {
		return err
	}
	return nil
}

// NewAssertUpgradeableLoaderAccountInstruction declares a new AssertUpgradeableLoaderAccount instruction with the provided parameters and accounts.
func NewAssertUpgradeableLoaderAccountInstruction(
	// Parameters:
	logLevel LogLevel,
	assertion UpgradeableLoaderStateAssertion,
	// Accounts:
	targetAccount ag_solanago.PublicKey) *AssertUpgradeableLoaderAccount {
	return NewAssertUpgradeableLoaderAccountInstructionBuilder().
		SetLogLevel(logLevel).
		SetAssertion(assertion).
		SetTargetAccountAccount(targetAccount)
}
