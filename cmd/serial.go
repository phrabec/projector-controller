package cmd

import (
	"fmt"
	"strings"

	"go.bug.st/serial"
)

const (
	idAll = "ZZ"

	commandPowerOn             = "PON"
	commandPowerOff            = "POF"
	commandInputSelection      = "IIS"
	commandUpdateSettings      = "VXX"
	commandAutoSetup           = "OAS"
	commandAspect              = "VSE"
	commandPicture             = "VPM"
	commandShutter             = "OSH"
	commandMainPowerStatus     = "QPW"
	commandShutterStatus       = "QSH"
	commandInputSignalStatus   = "QIN"
	commandLampPowerModeStatus = "QVX"
	commandAspectModeStatus    = "QSE"
	commandPictureModeStatus   = "QPM"
	commandProjectorRunTime    = "QST"
	commandLampRunTime         = "Q$L"
	commandTemperatureStatus   = "QTM"

	paramsNone                   = ""
	paramsShutterOff             = "1"
	paramsShutterOn              = "0"
	paramsDigitalLink            = "DL1"
	paramsDisplayPort            = "DP1"
	paramsHdmi                   = "HD1"
	paramsDvi                    = "DVI"
	paramsRgb1                   = "RG1"
	paramsRgb2                   = "RG2"
	paramsVideo                  = "VID"
	paramsLampPowerAuto          = paramsLamp1 + "=+00030"
	paramsLampPowerNormal        = paramsLamp1 + "=+00001"
	paramsLampPowerEco           = paramsLamp1 + "=+00000"
	paramsAspectAuto             = "0"
	paramsAspect43               = "1"
	paramsAspectWide             = "2"
	paramsAspectReal             = "5"
	paramsAspectFull             = "6"
	paramsAspectHFit             = "9"
	paramsAspectVFit             = "10"
	paramsPictureNatural         = "NAT"
	paramsPictureStandard        = "STD"
	paramsPictureDynamic         = "DYN"
	paramsPictureCinema          = "CIN"
	paramsPictureDicomSimulation = "DIC"
	paramsLampRunTime            = "1"
	paramsAirIntake              = "0"
	paramsExhaust                = "1"
	paramsLamp1                  = "LPWI1"
)

func buildQuery(id string, cmd string, params string) string {
	if params == paramsNone {
		return fmt.Sprintf("\x02AD%v;%v\x03", id, cmd)
	}

	return fmt.Sprintf("\x02AD%v;%v:%v\x03", id, cmd, params)
}

func communicate(portName string, id string, cmd string, params string) error {
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return err
	}
	defer port.Close()

	query := buildQuery(id, cmd, params)

	_, err = port.Write([]byte(query))
	if err != nil {
		return err
	}

	buff := make([]byte, 20)
	result := make([]byte, 0)
	for {
		n, err := port.Read(buff)
		if err != nil {
			return err
		}
		if n == 0 {
			break
		}

		result = append(result, buff[:n]...)

		if strings.Contains(string(buff[:n]), "\x03") {
			break
		}
	}

	length := len(result)
	if length < 2 {
		return fmt.Errorf("unexpected response (len %d): '%v'", length, result)
	}

	text := string(result[1 : length-1]) //nolint:gosec
	fmt.Printf("Result: %v\n", text)

	return nil
}
