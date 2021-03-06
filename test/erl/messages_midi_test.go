package erl

// Basic imports
import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ut-proj/midiserver/pkg/erl"
	"github.com/ut-proj/midiserver/pkg/erl/messages"
	"github.com/ut-proj/midiserver/pkg/erl/packets"
	"github.com/ut-proj/midiserver/pkg/types"
)

const (
	Bb     uint8 = 34
	Volume uint8 = 40
)

type MidiMessageTestSuite struct {
	suite.Suite
	batch  interface{}
	device interface{}
	noteOn interface{}
}

func (suite *MidiMessageTestSuite) SetupTest() {
	batchBytes := []byte{0x38, 0x33, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x34, 0x36, 0x44, 0x36, 0x39, 0x36, 0x34, 0x36, 0x39, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x35, 0x36, 0x32, 0x36, 0x31, 0x37, 0x34, 0x36, 0x33, 0x36, 0x38, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x34, 0x36, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x33, 0x30, 0x39, 0x36, 0x39, 0x35, 0x37, 0x39, 0x43, 0x41, 0x35, 0x33, 0x34, 0x42, 0x41, 0x30, 0x42, 0x34, 0x41, 0x46, 0x41, 0x43, 0x46, 0x43, 0x45, 0x44, 0x37, 0x30, 0x39, 0x38, 0x36, 0x34, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x38, 0x36, 0x44, 0x36, 0x35, 0x37, 0x33, 0x37, 0x33, 0x36, 0x31, 0x36, 0x37, 0x36, 0x35, 0x37, 0x33, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x36, 0x36, 0x34, 0x36, 0x35, 0x37, 0x36, 0x36, 0x39, 0x36, 0x33, 0x36, 0x35, 0x36, 0x31, 0x30, 0x30, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x37, 0x36, 0x33, 0x36, 0x38, 0x36, 0x31, 0x36, 0x45, 0x36, 0x45, 0x36, 0x35, 0x36, 0x43, 0x36, 0x31, 0x30, 0x30, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x37, 0x36, 0x45, 0x36, 0x46, 0x37, 0x34, 0x36, 0x35, 0x35, 0x46, 0x36, 0x46, 0x36, 0x45, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x35, 0x37, 0x30, 0x36, 0x39, 0x37, 0x34, 0x36, 0x33, 0x36, 0x38, 0x36, 0x31, 0x32, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x38, 0x37, 0x36, 0x36, 0x35, 0x36, 0x43, 0x36, 0x46, 0x36, 0x33, 0x36, 0x39, 0x37, 0x34, 0x37, 0x39, 0x36, 0x31, 0x32, 0x38, 0x36, 0x41, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x38, 0x36, 0x45, 0x36, 0x46, 0x37, 0x34, 0x36, 0x35, 0x35, 0x46, 0x36, 0x46, 0x36, 0x36, 0x36, 0x36, 0x36, 0x31, 0x32, 0x32, 0x36, 0x41, 0x36, 0x41, 0xa}
	deviceBytes := []byte{0x38, 0x33, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x34, 0x36, 0x44, 0x36, 0x39, 0x36, 0x34, 0x36, 0x39, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x35, 0x36, 0x32, 0x36, 0x31, 0x37, 0x34, 0x36, 0x33, 0x36, 0x38, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x34, 0x36, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x31, 0x31, 0x46, 0x46, 0x31, 0x33, 0x35, 0x43, 0x37, 0x38, 0x44, 0x35, 0x34, 0x31, 0x35, 0x43, 0x38, 0x38, 0x31, 0x38, 0x43, 0x44, 0x45, 0x37, 0x32, 0x32, 0x35, 0x32, 0x46, 0x46, 0x30, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x38, 0x36, 0x44, 0x36, 0x35, 0x37, 0x33, 0x37, 0x33, 0x36, 0x31, 0x36, 0x37, 0x36, 0x35, 0x37, 0x33, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x36, 0x36, 0x34, 0x36, 0x35, 0x37, 0x36, 0x36, 0x39, 0x36, 0x33, 0x36, 0x35, 0x36, 0x31, 0x30, 0x30, 0x36, 0x41, 0x36, 0x41, 0xa}
	noteOnBytes := []byte{0x38, 0x33, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x34, 0x36, 0x44, 0x36, 0x39, 0x36, 0x34, 0x36, 0x39, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x35, 0x36, 0x32, 0x36, 0x31, 0x37, 0x34, 0x36, 0x33, 0x36, 0x38, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x34, 0x36, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x44, 0x45, 0x39, 0x35, 0x30, 0x37, 0x37, 0x39, 0x45, 0x36, 0x30, 0x41, 0x34, 0x33, 0x39, 0x41, 0x42, 0x43, 0x38, 0x33, 0x33, 0x32, 0x37, 0x41, 0x44, 0x46, 0x37, 0x30, 0x44, 0x39, 0x36, 0x31, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x38, 0x36, 0x44, 0x36, 0x35, 0x37, 0x33, 0x37, 0x33, 0x36, 0x31, 0x36, 0x37, 0x36, 0x35, 0x37, 0x33, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x37, 0x36, 0x45, 0x36, 0x46, 0x37, 0x34, 0x36, 0x35, 0x35, 0x46, 0x36, 0x46, 0x36, 0x45, 0x36, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x35, 0x37, 0x30, 0x36, 0x39, 0x37, 0x34, 0x36, 0x33, 0x36, 0x38, 0x36, 0x31, 0x32, 0x32, 0x36, 0x38, 0x30, 0x32, 0x36, 0x34, 0x30, 0x30, 0x30, 0x38, 0x37, 0x36, 0x36, 0x35, 0x36, 0x43, 0x36, 0x46, 0x36, 0x33, 0x36, 0x39, 0x37, 0x34, 0x37, 0x39, 0x36, 0x31, 0x32, 0x38, 0x36, 0x41, 0x36, 0x41, 0x36, 0x41, 0xa}

	opts := &erl.Opts{IsHexEncoded: true}
	bPkt, _ := packets.NewPacket(batchBytes, opts)
	suite.batch, _ = bPkt.Term()
	dPkt, _ := packets.NewPacket(deviceBytes, opts)
	suite.device, _ = dPkt.Term()
	nPkt, _ := packets.NewPacket(noteOnBytes, opts)
	suite.noteOn, _ = nPkt.Term()
}

func (suite *MidiMessageTestSuite) TestConvertDevice() {
	converted, err := messages.Convert(suite.device)
	suite.NoError(err)
	suite.Equal("11ff135c-78d5-415c-8818-cde72252ff02", converted.Id())
	suite.Equal(types.MidiDeviceType(), converted.Calls()[0].Op)
	suite.Equal(uint8(0), converted.Calls()[0].Args.Device)
}

func (suite *MidiMessageTestSuite) TestConvertNoteOn() {
	converted, err := messages.Convert(suite.noteOn)
	suite.NoError(err)
	suite.Equal("de950779-e60a-439a-bc83-327adf70d961", converted.Id())
	suite.Equal(types.MidiNoteOnType(), converted.Calls()[0].Op)
	suite.Equal(Bb, converted.Calls()[0].Args.NoteOn.Pitch)
	suite.Equal(Volume, converted.Calls()[0].Args.NoteOn.Velocity)
}

func (suite *MidiMessageTestSuite) TestConvertBatch() {
	converted, err := messages.Convert(suite.batch)
	suite.NoError(err)
	suite.Equal("30969579-ca53-4ba0-b4af-acfced709864", converted.Id())
	suite.Equal(4, len(converted.Calls()))
	suite.Equal(1, converted.Calls()[0].Id)
	suite.Equal(types.MidiDeviceType(), converted.Calls()[0].Op)
	suite.Equal(uint8(0), converted.Calls()[0].Args.Device)
	suite.Equal(2, converted.Calls()[1].Id)
	suite.Equal(types.MidiChannelType(), converted.Calls()[1].Op)
	suite.Equal(uint8(0), converted.Calls()[1].Args.Channel)
	suite.Equal(3, converted.Calls()[2].Id)
	suite.Equal(types.MidiNoteOnType(), converted.Calls()[2].Op)
	suite.Equal(Bb, converted.Calls()[2].Args.NoteOn.Pitch)
	suite.Equal(Volume, converted.Calls()[2].Args.NoteOn.Velocity)
	suite.Equal(4, converted.Calls()[3].Id)
	suite.Equal(types.MidiNoteOffType(), converted.Calls()[3].Op)
	suite.Equal(Bb, converted.Calls()[3].Args.NoteOff)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMidiMessageTestSuite(t *testing.T) {
	suite.Run(t, new(MidiMessageTestSuite))
}
