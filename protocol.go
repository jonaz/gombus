package gombus

import (
	"fmt"
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
)

func DecodeRecordFunction(dif byte) string {
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

func DeviceTypeLookup(deviceType byte) (string, error) {
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
