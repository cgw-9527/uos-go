package msig

import uos "github.com/lialvin/uos-go"

type ProposalRow struct {
	ProposalName       uos.Name              `json:"proposal_name"`
	RequestedApprovals []uos.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []uos.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  uos.HexBytes          `json:"packed_transaction"`
}
