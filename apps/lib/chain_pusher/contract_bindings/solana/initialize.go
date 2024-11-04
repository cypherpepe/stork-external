// Code generated by https://github.com/henrymbaldwin/solana-anchor-go. DO NOT EDIT.

package contract_bindings_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize is the `initialize` instruction.
type Initialize struct {
	StorkSolPublicKey         *ag_solanago.PublicKey
	StorkEvmPublicKey         *[20]uint8
	SingleUpdateFeeInLamports *uint64

	// [0] = [WRITE] config
	//
	// [1] = [WRITE, SIGNER] owner
	//
	// [2] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	return nd
}

// SetStorkSolPublicKey sets the "stork_sol_public_key" parameter.
func (inst *Initialize) SetStorkSolPublicKey(stork_sol_public_key ag_solanago.PublicKey) *Initialize {
	inst.StorkSolPublicKey = &stork_sol_public_key
	return inst
}

// SetStorkEvmPublicKey sets the "stork_evm_public_key" parameter.
func (inst *Initialize) SetStorkEvmPublicKey(stork_evm_public_key [20]uint8) *Initialize {
	inst.StorkEvmPublicKey = &stork_evm_public_key
	return inst
}

// SetSingleUpdateFeeInLamports sets the "single_update_fee_in_lamports" parameter.
func (inst *Initialize) SetSingleUpdateFeeInLamports(single_update_fee_in_lamports uint64) *Initialize {
	inst.SingleUpdateFeeInLamports = &single_update_fee_in_lamports
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Initialize) SetConfigAccount(config ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

func (inst *Initialize) findFindConfigAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: stork_config
	seeds = append(seeds, []byte{byte(0x73), byte(0x74), byte(0x6f), byte(0x72), byte(0x6b), byte(0x5f), byte(0x63), byte(0x6f), byte(0x6e), byte(0x66), byte(0x69), byte(0x67)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindConfigAddressWithBumpSeed calculates Config account address with given seeds and a known bump seed.
func (inst *Initialize) FindConfigAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindConfigAddress(bumpSeed)
	return
}

func (inst *Initialize) MustFindConfigAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindConfigAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindConfigAddress finds Config account address with given seeds.
func (inst *Initialize) FindConfigAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindConfigAddress(0)
	return
}

func (inst *Initialize) MustFindConfigAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindConfigAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetConfigAccount gets the "config" account.
func (inst *Initialize) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *Initialize) SetOwnerAccount(owner ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *Initialize) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.StorkSolPublicKey == nil {
			return errors.New("StorkSolPublicKey parameter is not set")
		}
		if inst.StorkEvmPublicKey == nil {
			return errors.New("StorkEvmPublicKey parameter is not set")
		}
		if inst.SingleUpdateFeeInLamports == nil {
			return errors.New("SingleUpdateFeeInLamports parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("            StorkSolPublicKey", *inst.StorkSolPublicKey))
						paramsBranch.Child(ag_format.Param("            StorkEvmPublicKey", *inst.StorkEvmPublicKey))
						paramsBranch.Child(ag_format.Param("    SingleUpdateFeeInLamports", *inst.SingleUpdateFeeInLamports))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("system_program", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StorkSolPublicKey` param:
	err = encoder.Encode(obj.StorkSolPublicKey)
	if err != nil {
		return err
	}
	// Serialize `StorkEvmPublicKey` param:
	err = encoder.Encode(obj.StorkEvmPublicKey)
	if err != nil {
		return err
	}
	// Serialize `SingleUpdateFeeInLamports` param:
	err = encoder.Encode(obj.SingleUpdateFeeInLamports)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StorkSolPublicKey`:
	err = decoder.Decode(&obj.StorkSolPublicKey)
	if err != nil {
		return err
	}
	// Deserialize `StorkEvmPublicKey`:
	err = decoder.Decode(&obj.StorkEvmPublicKey)
	if err != nil {
		return err
	}
	// Deserialize `SingleUpdateFeeInLamports`:
	err = decoder.Decode(&obj.SingleUpdateFeeInLamports)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	stork_sol_public_key ag_solanago.PublicKey,
	stork_evm_public_key [20]uint8,
	single_update_fee_in_lamports uint64,
	// Accounts:
	config ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetStorkSolPublicKey(stork_sol_public_key).
		SetStorkEvmPublicKey(stork_evm_public_key).
		SetSingleUpdateFeeInLamports(single_update_fee_in_lamports).
		SetConfigAccount(config).
		SetOwnerAccount(owner).
		SetSystemProgramAccount(systemProgram)
}
