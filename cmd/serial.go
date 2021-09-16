package cmd

import (
	"fmt"
	"strings"

	"go.bug.st/serial"
)

const (
	kIdAll = "ZZ"

	kCommandPowerOn             = "PON"
	kCommandPowerOff            = "POF"
	kCommandInputSelection      = "IIS"
	kCommandUpdateSettings      = "VXX"
	kCommandAutoSetup           = "OAS"
	kCommandAspect              = "VSE"
	kCommandPicture             = "VPM"
	kCommandShutter             = "OSH"
	kCommandMainPowerStatus     = "QPW"
	kCommandShutterStatus       = "QSH"
	kCommandInputSignalStatus   = "QIN"
	kCommandLampPowerModeStatus = "QVX"
	kCommandAspectModeStatus    = "QSE"
	kCommandPictureModeStatus   = "QPM"
	kCommandProjectorRunTime    = "QST"
	kCommandLampRunTime         = "Q$L"
	kCommandTemperatureStatus   = "QTM"

	kParamsNone                   = ""
	kParamsShutterOff             = "1"
	kParamsShutterOn              = "0"
	kParamsDigitalLink            = "DL1"
	kParamsDisplayPort            = "DP1"
	kParamsHdmi                   = "HD1"
	kParamsDvi                    = "DVI"
	kParamsRgb1                   = "RG1"
	kParamsRgb2                   = "RG2"
	kParamsVideo                  = "VID"
	kParamsLampPowerAuto          = kParamsLamp1 + "=+00030"
	kParamsLampPowerNormal        = kParamsLamp1 + "=+00001"
	kParamsLampPowerEco           = kParamsLamp1 + "=+00000"
	kParamsAspectAuto             = "0"
	kParamsAspect43               = "1"
	kParamsAspectWide             = "2"
	kParamsAspectReal             = "5"
	kParamsAspectFull             = "6"
	kParamsAspectHFit             = "9"
	kParamsAspectVFit             = "10"
	kParamsPictureNatural         = "NAT"
	kParamsPictureStandard        = "STD"
	kParamsPictureDynamic         = "DYN"
	kParamsPictureCinema          = "CIN"
	kParamsPictureDicomSimulation = "DIC"
	kParamsLampRunTime            = "1"
	kParamsAirIntake              = "0"
	kParamsExhaust                = "1"
	kParamsLamp1                  = "LPWI1"
)

func buildQuery(id string, cmd string, params string) string {
	if params == kParamsNone {
		return fmt.Sprintf("\x02AD%v;%v\x03", id, cmd)
	} else {
		return fmt.Sprintf("\x02AD%v;%v:%v\x03", id, cmd, params)
	}
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

	fmt.Printf("Result: %v\n", string(result[1:len(result)-1]))

	return nil
}
