// Code generated by https://github.com/henrymbaldwin/solana-anchor-go. DO NOT EDIT.
// Code generated by https://github.com/henrymbaldwin/solana-anchor-go. DO NOT EDIT.

package contract_bindings_solana

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(PublicKey ag_solanago.PublicKey) {
	ProgramID = PublicKey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "ContractBindingsSolana"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_TransferOwnership = ag_binary.TypeID([8]byte{65, 177, 215, 73, 53, 45, 99, 47})

	Instruction_UpdateSingleUpdateFeeInLamports = ag_binary.TypeID([8]byte{1, 154, 255, 107, 3, 43, 213, 151})

	Instruction_UpdateStorkEvmPublicKey = ag_binary.TypeID([8]byte{153, 177, 89, 97, 168, 213, 163, 107})

	Instruction_UpdateStorkSolPublicKey = ag_binary.TypeID([8]byte{152, 181, 100, 254, 176, 84, 135, 126})

	Instruction_UpdateTemporalNumericValueEvm = ag_binary.TypeID([8]byte{201, 67, 98, 156, 22, 183, 29, 81})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_TransferOwnership:
		return "TransferOwnership"
	case Instruction_UpdateSingleUpdateFeeInLamports:
		return "UpdateSingleUpdateFeeInLamports"
	case Instruction_UpdateStorkEvmPublicKey:
		return "UpdateStorkEvmPublicKey"
	case Instruction_UpdateStorkSolPublicKey:
		return "UpdateStorkSolPublicKey"
	case Instruction_UpdateTemporalNumericValueEvm:
		return "UpdateTemporalNumericValueEvm"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			Name: "initialize", Type: (*Initialize)(nil),
		},
		{
			Name: "transfer_ownership", Type: (*TransferOwnership)(nil),
		},
		{
			Name: "update_single_update_fee_in_lamports", Type: (*UpdateSingleUpdateFeeInLamports)(nil),
		},
		{
			Name: "update_stork_evm_public_key", Type: (*UpdateStorkEvmPublicKey)(nil),
		},
		{
			Name: "update_stork_sol_public_key", Type: (*UpdateStorkSolPublicKey)(nil),
		},
		{
			Name: "update_temporal_numeric_value_evm", Type: (*UpdateTemporalNumericValueEvm)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := decodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func decodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}

func DecodeInstructions(message *ag_solanago.Message) (instructions []*Instruction, err error) {
	for _, ins := range message.Instructions {
		var programID ag_solanago.PublicKey
		if programID, err = message.Program(ins.ProgramIDIndex); err != nil {
			return
		}
		if !programID.Equals(ProgramID) {
			continue
		}
		var accounts []*ag_solanago.AccountMeta
		if accounts, err = ins.ResolveInstructionAccounts(message); err != nil {
			return
		}
		var insDecoded *Instruction
		if insDecoded, err = decodeInstruction(accounts, ins.Data); err != nil {
			return
		}
		instructions = append(instructions, insDecoded)
	}
	return
}
