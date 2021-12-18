package lemi025

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/sss-eda/encoding/bcd"
	lemi025 "github.com/sss-eda/lemi-025"
)

type Encoder struct {
	io.Writer
}

// NewEncoder TODO
func NewEncoder(w io.Writer) *Encoder {
	encoder := Encoder{w}

	return &encoder
}

// Encode TODO
func (encoder *Encoder) Encode(v interface{}) error {
	var buffer []byte

	switch command := v.(type) {
	case lemi025.ReadConfigCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x30
	case lemi025.ReadTimeCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x31
	case lemi025.SetTimeCommand:
		buffer = make([]byte, 8)
		buffer[0] = 0x3D
		buffer[1] = 0x32
		buffer[2] = bcd.Encode(command.Year() - 2000)
		buffer[3] = bcd.Encode(command.Month())
		buffer[4] = bcd.Encode(command.Day())
		buffer[5] = bcd.Encode(command.Hour())
		buffer[6] = bcd.Encode(command.Minute())
		buffer[7] = bcd.Encode(command.Second())
	case lemi025.SetCoefficients1Command:
		buffer = make([]byte, 4)
		buffer[0] = 0x3D
		buffer[1] = 0x33
		buffer[2] = 0x00
		buffer[3] = bcd.Encode(command.Mode())
	case lemi025.ReadCoefficients1Command:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x34
	case lemi025.SetCoefficients2Command:
		buffer = make([]byte, 84)
		buffer[0] = 0x3D
		buffer[1] = 0x35
		buffer[2] = 0xFF
		buffer[3] = 0xFF

		err := binary.Write(encoder, binary.LittleEndian, command.Ax1)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Ay1)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Az1)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Beta)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Gamma)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Xi)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Exy)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Eyz)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.Exz)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.K1x)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.K1y)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.K1z)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.K2x)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.K2y)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.K2z)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.KTF)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.KTE)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.KTF0)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.KTE0)
		if err != nil {
			return err
		}

		err = binary.Write(encoder, binary.LittleEndian, command.KVBAT)
		if err != nil {
			return err
		}
	case lemi025.ReadCoefficients2Command:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x36
	case lemi025.ReadGPSDataCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x37
	case lemi025.StopSystemCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x38
	case lemi025.StartSystemCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x39
	case lemi025.CheckFLASHCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x3A
	case lemi025.SetDACxCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x3D

		err = binary.Write(encoder, binary.LittleEndian, command.Value)
		if err != nil {
			return err
		}
	case lemi025.SetDACyCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x3E

		err = binary.Write(encoder, binary.LittleEndian, command.Value)
		if err != nil {
			return err
		}
	case lemi025.SetDACzCommand:
		buffer = make([]byte, 2)
		buffer[0] = 0x3D
		buffer[1] = 0x3F

		err = binary.Write(encoder, binary.LittleEndian, command.Value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("command type not reccognised: %v", command)
	}

	return nil
}
