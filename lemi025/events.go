package lemi025

import "fmt"

func (event ConfigReadEvent) MarshalSerial() (data []byte, err error) {
	if event.stationNumber != nil {
		data = []byte(fmt.Sprintf("025 %d", *event.stationNumber))
		err = nil

		return
	}

	data = nil
	err = fmt.Errorf("unable to marshal ConfigReadEvent: station number not defined")

	return
}

func (event ConfigReadEvent) UnmarshalSerial(data []byte) error {
	if len(data) != 5 {
		return fmt.Errorf(
			"invalid data length: %d. Data should have length of 5.",
			len(data),
		)
	}

	if string(data[:3]) != "025 " {
		return fmt.Errorf(
			"invalid data: %v. Data should start with \"025 \"",
			data,
		)
	}

	sn := uint8((data[4] / 10 << 4) + (data[4] % 10))
	event.stationNumber = &sn

	return nil
}
