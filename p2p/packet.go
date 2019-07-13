package p2p

import (
	"github.com/lialvin/uos-go"
)

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *uos.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *uos.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}
