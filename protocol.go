package gombus

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/sirupsen/logrus"
)

const (
	DATA_RECORD_DIF_MASK_INST        = 0x00
	DATA_RECORD_DIF_MASK_MIN         = 0x10
	DATA_RECORD_DIF_MASK_TYPE_INT32  = 0x04
	DATA_RECORD_DIF_MASK_DATA        = 0x0F
	DATA_RECORD_DIF_MASK_FUNCTION    = 0x30
	DATA_RECORD_DIF_MASK_STORAGE_NO  = 0x40
	DATA_RECORD_DIF_MASK_EXTENSION   = 0x80
	DATA_RECORD_DIF_MASK_NON_DATA    = 0xF0
	DATA_RECORD_DIFE_MASK_STORAGE_NO = 0x0F
	DATA_RECORD_DIFE_MASK_TARIFF     = 0x30
	DATA_RECORD_DIFE_MASK_DEVICE     = 0x40
	DATA_RECORD_DIFE_MASK_EXTENSION  = 0x80

	DIB_DIF_WITHOUT_EXTENSION     = 0x7F
	DIB_DIF_EXTENSION_BIT         = 0x80
	DIB_VIF_WITHOUT_EXTENSION     = 0x7F
	DIB_VIF_EXTENSION_BIT         = 0x80
	DIB_DIF_MANUFACTURER_SPECIFIC = 0x0F
	DIB_DIF_MORE_RECORDS_FOLLOW   = 0x1F
	DIB_DIF_IDLE_FILLER           = 0x2F

	VARIABLE_DATA_MEDIUM_OTHER         = 0x00
	VARIABLE_DATA_MEDIUM_OIL           = 0x01
	VARIABLE_DATA_MEDIUM_ELECTRICITY   = 0x02
	VARIABLE_DATA_MEDIUM_GAS           = 0x03
	VARIABLE_DATA_MEDIUM_HEAT_OUT      = 0x04
	VARIABLE_DATA_MEDIUM_STEAM         = 0x05
	VARIABLE_DATA_MEDIUM_HOT_WATER     = 0x06
	VARIABLE_DATA_MEDIUM_WATER         = 0x07
	VARIABLE_DATA_MEDIUM_HEAT_COST     = 0x08
	VARIABLE_DATA_MEDIUM_COMPR_AIR     = 0x09
	VARIABLE_DATA_MEDIUM_COOL_OUT      = 0x0A
	VARIABLE_DATA_MEDIUM_COOL_IN       = 0x0B
	VARIABLE_DATA_MEDIUM_HEAT_IN       = 0x0C
	VARIABLE_DATA_MEDIUM_HEAT_COOL     = 0x0D
	VARIABLE_DATA_MEDIUM_BUS           = 0x0E
	VARIABLE_DATA_MEDIUM_UNKNOWN       = 0x0F
	VARIABLE_DATA_MEDIUM_IRRIGATION    = 0x10
	VARIABLE_DATA_MEDIUM_WATER_LOGGER  = 0x11
	VARIABLE_DATA_MEDIUM_GAS_LOGGER    = 0x12
	VARIABLE_DATA_MEDIUM_GAS_CONV      = 0x13
	VARIABLE_DATA_MEDIUM_COLORIFIC     = 0x14
	VARIABLE_DATA_MEDIUM_BOIL_WATER    = 0x15
	VARIABLE_DATA_MEDIUM_COLD_WATER    = 0x16
	VARIABLE_DATA_MEDIUM_DUAL_WATER    = 0x17
	VARIABLE_DATA_MEDIUM_PRESSURE      = 0x18
	VARIABLE_DATA_MEDIUM_ADC           = 0x19
	VARIABLE_DATA_MEDIUM_SMOKE         = 0x1A
	VARIABLE_DATA_MEDIUM_ROOM_SENSOR   = 0x1B
	VARIABLE_DATA_MEDIUM_GAS_DETECTOR  = 0x1C
	VARIABLE_DATA_MEDIUM_BREAKER_E     = 0x20
	VARIABLE_DATA_MEDIUM_VALVE         = 0x21
	VARIABLE_DATA_MEDIUM_CUSTOMER_UNIT = 0x25
	VARIABLE_DATA_MEDIUM_WASTE_WATER   = 0x28
	VARIABLE_DATA_MEDIUM_GARBAGE       = 0x29
	VARIABLE_DATA_MEDIUM_VOC           = 0x2B
	VARIABLE_DATA_MEDIUM_SERVICE_UNIT  = 0x30
	VARIABLE_DATA_MEDIUM_RC_SYSTEM     = 0x36
	VARIABLE_DATA_MEDIUM_RC_METER      = 0x37

	CONTROL_MASK_FCB = 0x20
	CONTROL_MASK_FCV = 0x10
)

func decodeRecordFunction(dif byte) string {
	switch dif & DATA_RECORD_DIF_MASK_FUNCTION {
	case 0x00:
		return "Instantaneous value"
	case 0x10:
		return "Maximum value"
	case 0x20:
		return "Minimum value"
	case 0x30:
		return "Value during error state"
	default:
		return "Unknown"
	}
}

func deviceTypeLookup(deviceType byte) (string, error) {
	switch deviceType {
	case VARIABLE_DATA_MEDIUM_OTHER:
		return "Other", nil
	case VARIABLE_DATA_MEDIUM_OIL:
		return "Oild", nil
	case VARIABLE_DATA_MEDIUM_ELECTRICITY:
		return "Electricity", nil
	case VARIABLE_DATA_MEDIUM_GAS:
		return "Gas", nil
	case VARIABLE_DATA_MEDIUM_HEAT_OUT:
		return "Heat: Outlet", nil
	case VARIABLE_DATA_MEDIUM_STEAM:
		return "Steam", nil
	case VARIABLE_DATA_MEDIUM_HOT_WATER:
		return "Warm water (30-90°C)", nil
	case VARIABLE_DATA_MEDIUM_WATER:
		return "Warm", nil
	case VARIABLE_DATA_MEDIUM_HEAT_COST:
		return "Heat Cost Allocator", nil
	case VARIABLE_DATA_MEDIUM_COMPR_AIR:
		return "Compressed Air", nil
	case VARIABLE_DATA_MEDIUM_COOL_OUT:
		return "Cooling load meter: Outlet", nil
	case VARIABLE_DATA_MEDIUM_COOL_IN:
		return "Cooling load meter: Inlet", nil
	case VARIABLE_DATA_MEDIUM_HEAT_IN:
		return "Heat: Inlet", nil
	case VARIABLE_DATA_MEDIUM_HEAT_COOL:
		return "Heat / Cooling load meter", nil
	case VARIABLE_DATA_MEDIUM_BUS:
		return "Bus / System", nil
	case VARIABLE_DATA_MEDIUM_UNKNOWN:
		return "Unknown Device type", nil
	case VARIABLE_DATA_MEDIUM_IRRIGATION:
		return "Irrigation Water", nil
	case VARIABLE_DATA_MEDIUM_WATER_LOGGER:
		return "Water Logger", nil
	case VARIABLE_DATA_MEDIUM_GAS_LOGGER:
		return "Gas Logger", nil
	case VARIABLE_DATA_MEDIUM_GAS_CONV:
		return "Gas Converter", nil
	case VARIABLE_DATA_MEDIUM_COLORIFIC:
		return "Calorific value", nil
	case VARIABLE_DATA_MEDIUM_BOIL_WATER:
		return "Hot water (>90°C)", nil
	case VARIABLE_DATA_MEDIUM_COLD_WATER:
		return "Cold water", nil
	case VARIABLE_DATA_MEDIUM_DUAL_WATER:
		return "Dual water", nil
	case VARIABLE_DATA_MEDIUM_PRESSURE:
		return "Pressure", nil
	case VARIABLE_DATA_MEDIUM_ADC:
		return "A/D Converter", nil
	case VARIABLE_DATA_MEDIUM_SMOKE:
		return "Smoke Detector", nil
	case VARIABLE_DATA_MEDIUM_ROOM_SENSOR:
		return "Ambient Sensor", nil
	case VARIABLE_DATA_MEDIUM_GAS_DETECTOR:
		return "Gas Detector", nil
	case VARIABLE_DATA_MEDIUM_BREAKER_E:
		return "Breaker: Electricity", nil
	case VARIABLE_DATA_MEDIUM_VALVE:
		return "Valve: Gas or Water", nil
	case VARIABLE_DATA_MEDIUM_CUSTOMER_UNIT:
		return "Customer Unit: Display Device", nil
	case VARIABLE_DATA_MEDIUM_WASTE_WATER:
		return "Waste Water", nil
	case VARIABLE_DATA_MEDIUM_GARBAGE:
		return "Garbage", nil
	case VARIABLE_DATA_MEDIUM_VOC:
		return "VOC Sensor", nil
	case VARIABLE_DATA_MEDIUM_SERVICE_UNIT:
		return "Service Unit", nil
	case VARIABLE_DATA_MEDIUM_RC_SYSTEM:
		return "Radio Converter: System", nil
	case VARIABLE_DATA_MEDIUM_RC_METER:
		return "Radio Converter: Meter", nil
	case 0x22, 0x23, 0x24, 0x26, 0x27, 0x2A, 0x2C, 0x2D, 0x2E, 0x2F, 0x31, 0x32, 0x33, 0x34, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E, 0x3F:
		return "Reserved", nil
	}

	return "", fmt.Errorf("Unknown medium (0x%.2x)", deviceType)
}

func decodeUnit(vif byte, vife []byte) Unit {
	var code int

	logrus.Tracef("vif raw is: % x\n", vif)
	logrus.Tracef("vife raw is: % x\n", vife)
	if vif == 0xFB {
		code = int(vife[0])&DIB_VIF_WITHOUT_EXTENSION | 0x200
	} else if vif == 0xFD {
		code = int(vife[0])&DIB_VIF_WITHOUT_EXTENSION | 0x100
		// } else if vif == 0x7C {  // handled in frame.go
	} else if vif == 0xFC {
		code := vife[0] & DIB_VIF_WITHOUT_EXTENSION
		var factor float64

		if 0x70 <= code && code <= 0x77 {
			factor = math.Pow10((int(vife[0]) & 0x07) - 6)
		} else if 0x78 <= code && code <= 0x7B {
			factor = math.Pow10((int(vife[0]) & 0x03) - 3)
		} else if code == 0x7D {
			factor = 1000
		}

		return Unit{
			Exp: factor,
			//Unit:        vib.Custom, //TODO here
			Type:        VIFUnit["VARIABLE_VIF"],
			VIFUnitDesc: "",
		}
	} else {
		code = int(vif) & DIB_VIF_WITHOUT_EXTENSION
	}

	logrus.Tracef("vif code is: % x\n", code)
	return unitTable[code]
}
func decodeStorageNumber(dif int, dife []byte) int {
	bitIndex := 0
	result := 0

	result |= dif & DATA_RECORD_DIF_MASK_STORAGE_NO >> 6
	bitIndex++

	size := len(dife)
	for i := 0; i < size; i++ {
		result |= int(dife[i]&DATA_RECORD_DIFE_MASK_STORAGE_NO) << bitIndex
		bitIndex += 4
	}

	return result
}

func decodeTariff(dif int, dife []byte) int {
	bitIndex := 0
	result := 0
	size := len(dife)
	for i := 0; i < size; i++ {
		result |= int(dife[i]&DATA_RECORD_DIFE_MASK_TARIFF>>4) << bitIndex
		bitIndex += 2
	}

	return result
}

func decodeDevice(dif int, dife []byte) int {
	bitIndex := 0
	result := 0

	size := len(dife)
	for i := 0; i < size; i++ {
		result |= int(dife[i]&DATA_RECORD_DIFE_MASK_DEVICE>>6) << bitIndex
		bitIndex++
	}

	return result
}

func decodeASCII(data []byte) string {
	s := ""
	for i := len(data) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%c", data[i])
	}
	return s
}

func uintToBCD(value uint64, size int) []byte {
	buf := make([]byte, size)
	if value > 0 {
		remainder := value
		for pos := size - 1; pos >= 0 && remainder > 0; pos-- {
			tail := byte(remainder % 100)
			hi, lo := tail/10, tail%10
			buf[size-1-pos] = hi<<4 + lo
			remainder /= 100
		}
	}
	return buf
}

func bcdToInt(bcd []byte) int {
	var i = 0
	size := len(bcd)
	for k := range bcd {
		r0 := bcd[size-1-k] & 0xf
		r1 := bcd[size-1-k] >> 4 & 0xf
		r := r1*10 + r0
		i = i*100 + int(r)
	}
	return i
}

// if data fields is 0100 == 32 bit integer.
func int32ToInt(data []byte) (int, error) {
	if len(data) != 4 {
		return 0.0, fmt.Errorf("wrong data length")
	}
	i := binary.LittleEndian.Uint32(data)
	return int(i), nil
}

func int24ToInt(b []byte) int {
	_ = b[2] // bounds check hint to compiler; see golang.org/issue/14808
	return int(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16)
}

func checkKthBitSet(n, k int) bool {
	if n&(1<<(k)) == 0 {
		return false
	}
	return true
}

// setBit set bits at pos. example 0000 pos 1 will be 0010.
func setBit(n byte, pos uint) byte {
	n |= (1 << pos)
	return n
}

func setBitFromMask(b, mask byte) byte {
	return b | mask
}

func calcCheckSum(data []byte) byte {
	var res byte
	for _, v := range data {
		res += v
	}
	return res & 0xff
}
