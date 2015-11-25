package core

type Channels struct {
	mask int
}

func NewChannels() (*Channels) {
	return &Channels{1}
}

func (channels *Channels) Set(channel int) {
	channels.mask = 1 << channel
}

func (channels *Channels) Enable(channel int) {
	channels.mask |= 1 << channel
}

func (channels *Channels) Toggle(channel int) {
	channels.mask ^= 1 << channel
}

func (channels *Channels) Disable(channel int) {
	channels.mask &= ^(1 << channel)
}
