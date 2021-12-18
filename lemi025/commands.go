package lemi025

import "io"

type ReadConfigCommand struct{}

func (command *ReadConfigCommand) MarshalLEMI025() ([]byte, error) {
	return []byte{0x3D, 0x30}, nil
}

func (command *ReadConfigCommand) UnmarshalLEMI025([]byte) error {
	return []byte{0x3D, 0x30}, nil
}

type ReadTimeCommand struct{}

func (command *ReadTimeCommand) MarshalLEMI025() ([]byte, error) {
	return []byte{0x3D, 0x31}, nil
}

func (command *ReadTimeCommand) UnmarshalLEMI025([]byte) error {

}

type SetTimeCommand struct{}

func (command *SetTimeCommand) MarshalLEMI025() ([]byte, error) {

}

// ReadConfig TODO
func ReadConfig(w io.Writer) error {
	buffer = make([]byte, 2)

	buffer[0] = 0x3D
	buffer[1] = 0x30

	_, err := w.Write(buffer)
	return err
}
