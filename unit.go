package gombus

import "fmt"

var MeasureUnit = map[string]string{
	"KWH":         "kWh",
	"WH":          "WH",
	"J":           "J",
	"M3":          "m^3",
	"L":           "l",
	"KG":          "kg",
	"W":           "W",
	"J_H":         "J/h",
	"M3_H":        "m^3/h",
	"M3_MIN":      "m^3/min",
	"M3_S":        "m^3/s",
	"KG_H":        "kg/h",
	"C":           "C",
	"K":           "K",
	"BAR":         "bar",
	"DATE":        "date",
	"TIME":        "time",
	"DATE_TIME":   "date time",
	"DATE_TIME_S": "date time to second",
	"SECONDS":     "seconds",
	"MINUTES":     "minutes",
	"HOURS":       "hours",
	"DAYS":        "days",
	"NONE":        "none",
	"V":           "V",
	"A":           "A",
	"HCA":         "H.C.A",
	"CURRENCY":    "Currency unit",
	"BAUD":        "Baud",
	"BIT_TIMES":   "Bittimes",
	"PERCENT":     "%",
	"DBM":         "dBm",
}

var VIFUnit = map[string]int{
	"ENERGY_WH":              0x07, // E000 0xxx
	"ENERGY_J":               0x0F, // E000 1xxx
	"VOLUME":                 0x17, // E001 0xxx
	"MASS":                   0x1F, // E001 1xxx
	"ON_TIME":                0x23, // E010 00xx
	"OPERATING_TIME":         0x27, // E010 01xx
	"POWER_W":                0x2F, // E010 1xxx
	"POWER_J_H":              0x37, // E011 0xxx
	"VOLUME_FLOW":            0x3F, // E011 1xxx
	"VOLUME_FLOW_EXT":        0x47, // E100 0xxx
	"VOLUME_FLOW_EXT_S":      0x4F, // E100 1xxx
	"MASS_FLOW":              0x57, // E101 0xxx
	"FLOW_TEMPERATURE":       0x5B, // E101 10xx
	"RETURN_TEMPERATURE":     0x5F, // E101 11xx
	"TEMPERATURE_DIFFERENCE": 0x63, // E110 00xx
	"EXTERNAL_TEMPERATURE":   0x67, // E110 01xx
	"PRESSURE":               0x6B, // E110 10xx
	"DATE":                   0x6C, // E110 1100
	"DATE_TIME_GENERAL":      0x6D, // E110 1101
	"DATE_TIME":              0x6D, // E110 1101
	"EXTENTED_TIME":          0x6D, // E110 1101
	"EXTENTED_DATE_TIME":     0x6D, // E110 1101
	"UNITS_FOR_HCA":          0x6E, // E110 1110
	"RES_THIRD_VIFE_TABLE":   0x6F, // E110 1111
	"AVG_DURATION":           0x73, // E111 00xx
	"ACTUALITY_DURATION":     0x77, // E111 01xx
	"FABRICATION_NO":         0x78, // E111 1000
	"IDENTIFICATION":         0x79, // E111 1001
	"ADDRESS":                0x7A, // E111 1010

	// NOT THE ONES FOR SPECIAL PURPOSES
	"FIRST_EXT_VIF_CODES":     0xFB, // 1111 1011
	"VARIABLE_VIF":            0xFC, // E111 1111
	"VIF_FOLLOWING":           0x7C, // E111 1100
	"SECOND_EXT_VIF_CODES":    0xFD, // 1111 1101
	"THIRD_EXT_VIF_CODES_RES": 0xEF, // 1110 1111
	"ANY_VIF":                 0x7E, // E111 1110
	"MANUFACTURER_SPEC":       0x7F, // E111 1111
}

var VIFUnitExt = map[string]int{
	// Currency Units
	"CURRENCY_CREDIT": 0x03, // E000 00nn Credit of 10 nn-3 of the nominal ...
	"CURRENCY_DEBIT":  0x07, // E000 01nn Debit of 10 nn-3 of the nominal ...

	// Enhanced Identification
	"ACCESS_NUMBER":    0x08, // E000 1000 Access Number (transmission count)
	"MEDIUM":           0x09, // E000 1001 Medium (as in fixed header)
	"MANUFACTURER":     0x0A, // E000 1010 Manufacturer (as in fixed header)
	"PARAMETER_SET_ID": 0x0B, // E000 1011 Parameter set identification Enha ...
	"MODEL_VERSION":    0x0C, // E000 1100 Model / Version
	"HARDWARE_VERSION": 0x0D, // E000 1101 Hardware version //
	"FIRMWARE_VERSION": 0x0E, // E000 1110 Firmware version //
	"SOFTWARE_VERSION": 0x0F, // E000 1111 Software version //

	// Implementation of all TC294 WG1 requirements (improved selection ..)
	"CUSTOMER_LOCATION":           0x10, // E001 0000 Customer location
	"CUSTOMER":                    0x11, // E001 0001 Customer
	"ACCESS_CODE_USER":            0x12, // E001 0010 Access Code User
	"ACCESS_CODE_OPERATOR":        0x13, // E001 0011 Access Code Operator
	"ACCESS_CODE_SYSTEM_OPERATOR": 0x14, // E001 0100 Access Code System Operator
	"ACCESS_CODE_DEVELOPER":       0x15, // E001 0101 Access Code Developer
	"PASSWORD":                    0x16, // E001 0110 Password
	"ERROR_FLAGS":                 0x17, // E001 0111 Error flags (binary)
	"ERROR_MASKS":                 0x18, // E001 1000 Error mask
	"RESERVED":                    0x19, // E001 1001 Reserved
	"DIGITAL_OUTPUT":              0x1A, // E001 1010 Digital Output (binary)
	"DIGITAL_INPUT":               0x1B, // E001 1011 Digital Input (binary)
	"BAUDRATE":                    0x1C, // E001 1100 Baudrate [Baud]
	"RESPONSE_DELAY":              0x1D, // E001 1101 response delay time
	"RETRY":                       0x1E, // E001 1110 Retry
	"RESERVED_2":                  0x1F, // E001 1111 Reserved

	// Enhanced storage management
	"FIRST_STORAGE_NR":       0x20, // E010 0000 First storage
	"LAST_STORAGE_NR":        0x21, // E010 0001 Last storage
	"SIZE_OF_STORAGE_BLOCK":  0x22, // E010 0010 Size of storage block
	"RESERVED_3":             0x23, // E010 0011 Reserved
	"STORAGE_INTERVAL":       0x27, // E010 01nn Storage interval
	"STORAGE_INTERVAL_MONTH": 0x28, // E010 1000 Storage interval month(s)
	"STORAGE_INTERVAL_YEARS": 0x29, // E010 1001 Storage interval year(s)

	// E010 1010 Reserved
	// E010 1011 Reserved
	"DURATION_SINCE_LAST_READOUT": 0x2F, // E010 11nn Duration since last ...

	//  Enhanced tarif management
	"START_OF_TARIFF":        0x30, // E011 0000 Start (date/time) of tariff
	"DURATION_OF_TARIFF":     0x3,  // E011 00nn Duration of tariff
	"PERIOD_OF_TARIFF":       0x37, // E011 01nn Period of tariff
	"PERIOD_OF_TARIFF_MONTH": 0x38, // E011 1000 Period of tariff months(s)
	"PERIOD_OF_TARIFF_YEARS": 0x39, // E011 1001 Period of tariff year(s)
	"DIMENSIONLESS":          0x3A, // E011 1010 dimensionless / no VIF

	// E011 1011 Reserved
	// E011 11xx Reserved
	// Electrical units
	"VOLTS":                          0x4F, // E100 nnnn 10 nnnn-9 Volts
	"AMPERE":                         0x5F, // E101 nnnn 10 nnnn-12 A
	"RESET_COUNTER":                  0x60, // E110 0000 Reset counter
	"CUMULATION_COUNTER":             0x61, // E110 0001 Cumulation counter
	"CONTROL_SIGNAL":                 0x62, // E110 0010 Control signal
	"DAY_OF_WEEK":                    0x63, // E110 0011 Day of week
	"WEEK_NUMBER":                    0x64, // E110 0100 Week number
	"TIME_POINT_OF_DAY_CHANGE":       0x65, // E110 0101 Time point of day ...
	"STATE_OF_PARAMETER_ACTIVATION":  0x66, // E110 0110 State of parameter
	"SPECIAL_SUPPLIER_INFORMATION":   0x67, // E110 0111 Special supplier ...
	"DURATION_SINCE_LAST_CUMULATION": 0x6B, // E110 10pp Duration since last
	"OPERATING_TIME_BATTERY":         0x6F, // E110 11pp Operating time battery
	"DATEAND_TIME_OF_BATTERY_CHANGE": 0x70, // E111 0000 Date and time of bat...
	// E111 0001 to E111 1111 Reserved

	"RSSI": 0x71, // E111 0001 RSSI
}

var VIFUnitSecExt = map[string]int{
	"RELATIVE_HUMIDITY": 0x1A,
}

type VIF struct {
	Exp         float64
	Unit        string
	Type        int
	VIFUnitDesc string
}

func (vif VIF) Value(v float64) float64 {
	return v * vif.Exp
}

var VIFTable = map[int]VIF{
	// E000 0nnn    Energy Wh (0.001Wh to 10000Wh)
	0x00: {Exp: 1.0e-3, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x01: {Exp: 1.0e-2, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x02: {Exp: 1.0e-1, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x03: {Exp: 1.0e0, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x04: {Exp: 1.0e1, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x05: {Exp: 1.0e2, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x06: {Exp: 1.0e3, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},
	0x07: {Exp: 1.0e4, Unit: MeasureUnit["WH"], Type: VIFUnit["ENERGY_WH"]},

	// E000 1nnn    Energy  J (0.001kJ to 10000kJ)
	0x08: {Exp: 1.0e0, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x09: {Exp: 1.0e1, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x0A: {Exp: 1.0e2, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x0B: {Exp: 1.0e3, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x0C: {Exp: 1.0e4, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x0D: {Exp: 1.0e5, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x0E: {Exp: 1.0e6, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},
	0x0F: {Exp: 1.0e7, Unit: MeasureUnit["J"], Type: VIFUnit["ENERGY_J"]},

	// E001 0nnn    Volume m^3 (0.001l to 10000l)
	0x10: {Exp: 1.0e-6, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x11: {Exp: 1.0e-5, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x12: {Exp: 1.0e-4, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x13: {Exp: 1.0e-3, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x14: {Exp: 1.0e-2, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x15: {Exp: 1.0e-1, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x16: {Exp: 1.0e0, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},
	0x17: {Exp: 1.0e1, Unit: MeasureUnit["M3"], Type: VIFUnit["VOLUME"]},

	// E001 1nnn    Mass kg (0.001kg to 10000kg)
	0x18: {Exp: 1.0e-3, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x19: {Exp: 1.0e-2, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x1A: {Exp: 1.0e-1, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x1B: {Exp: 1.0e0, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x1C: {Exp: 1.0e1, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x1D: {Exp: 1.0e2, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x1E: {Exp: 1.0e3, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},
	0x1F: {Exp: 1.0e4, Unit: MeasureUnit["KG"], Type: VIFUnit["MASS"]},

	// E010 00nn    On Time s
	0x20: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ON_TIME"]},     // seconds
	0x21: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ON_TIME"]},    // minutes
	0x22: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ON_TIME"]},  // hours
	0x23: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ON_TIME"]}, // days

	// E010 01nn    Operating Time s
	0x24: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["OPERATING_TIME"]},     // sec
	0x25: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["OPERATING_TIME"]},    // min
	0x26: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["OPERATING_TIME"]},  // hours
	0x27: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["OPERATING_TIME"]}, // days

	// E010 1nnn    Power W (0.001W to 10000W)
	0x28: {Exp: 1.0e-3, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x29: {Exp: 1.0e-2, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x2A: {Exp: 1.0e-1, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x2B: {Exp: 1.0e0, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x2C: {Exp: 1.0e1, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x2D: {Exp: 1.0e2, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x2E: {Exp: 1.0e3, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},
	0x2F: {Exp: 1.0e4, Unit: MeasureUnit["W"], Type: VIFUnit["POWER_W"]},

	// E011 0nnn    Power J/h (0.001kJ/h to 10000kJ/h)
	0x30: {Exp: 1.0e0, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x31: {Exp: 1.0e1, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x32: {Exp: 1.0e2, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x33: {Exp: 1.0e3, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x34: {Exp: 1.0e4, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x35: {Exp: 1.0e5, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x36: {Exp: 1.0e6, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},
	0x37: {Exp: 1.0e7, Unit: MeasureUnit["J_H"], Type: VIFUnit["POWER_J_H"]},

	// E011 1nnn    Volume Flow m3/h (0.001l/h to 10000l/h)
	0x38: {Exp: 1.0e-6, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x39: {Exp: 1.0e-5, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x3A: {Exp: 1.0e-4, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x3B: {Exp: 1.0e-3, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x3C: {Exp: 1.0e-2, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x3D: {Exp: 1.0e-1, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x3E: {Exp: 1.0e0, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},
	0x3F: {Exp: 1.0e1, Unit: MeasureUnit["M3_H"], Type: VIFUnit["VOLUME_FLOW"]},

	// E100 0nnn     Volume Flow ext.  m^3/min (0.0001l/min to 1000l/min)
	0x40: {Exp: 1.0e-7, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x41: {Exp: 1.0e-6, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x42: {Exp: 1.0e-5, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x43: {Exp: 1.0e-4, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x44: {Exp: 1.0e-3, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x45: {Exp: 1.0e-2, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x46: {Exp: 1.0e-1, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},
	0x47: {Exp: 1.0e0, Unit: MeasureUnit["M3_MIN"], Type: VIFUnit["VOLUME_FLOW_EXT"]},

	// E100 1nnn     Volume Flow ext.  m^3/s (0.001ml/s to 10000ml/s)
	0x48: {Exp: 1.0e-9, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x49: {Exp: 1.0e-8, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x4A: {Exp: 1.0e-7, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x4B: {Exp: 1.0e-6, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x4C: {Exp: 1.0e-5, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x4D: {Exp: 1.0e-4, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x4E: {Exp: 1.0e-3, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},
	0x4F: {Exp: 1.0e-2, Unit: MeasureUnit["M3_S"], Type: VIFUnit["VOLUME_FLOW_EXT_S"]},

	// E101 0nnn     Mass flow kg/h (0.001kg/h to 10000kg/h)
	0x50: {Exp: 1.0e-3, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x51: {Exp: 1.0e-2, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x52: {Exp: 1.0e-1, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x53: {Exp: 1.0e0, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x54: {Exp: 1.0e1, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x55: {Exp: 1.0e2, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x56: {Exp: 1.0e3, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},
	0x57: {Exp: 1.0e4, Unit: MeasureUnit["KG_H"], Type: VIFUnit["MASS_FLOW"]},

	// E101 10nn     Flow Temperature degC (0.001degC to 1degC)
	0x58: {Exp: 1.0e-3, Unit: MeasureUnit["C"], Type: VIFUnit["FLOW_TEMPERATURE"]},
	0x59: {Exp: 1.0e-2, Unit: MeasureUnit["C"], Type: VIFUnit["FLOW_TEMPERATURE"]},
	0x5A: {Exp: 1.0e-1, Unit: MeasureUnit["C"], Type: VIFUnit["FLOW_TEMPERATURE"]},
	0x5B: {Exp: 1.0e0, Unit: MeasureUnit["C"], Type: VIFUnit["FLOW_TEMPERATURE"]},

	// E101 11nn Return Temperature degC (0.001degC to 1degC)
	0x5C: {Exp: 1.0e-3, Unit: MeasureUnit["C"], Type: VIFUnit["RETURN_TEMPERATURE"]},
	0x5D: {Exp: 1.0e-2, Unit: MeasureUnit["C"], Type: VIFUnit["RETURN_TEMPERATURE"]},
	0x5E: {Exp: 1.0e-1, Unit: MeasureUnit["C"], Type: VIFUnit["RETURN_TEMPERATURE"]},
	0x5F: {Exp: 1.0e0, Unit: MeasureUnit["C"], Type: VIFUnit["RETURN_TEMPERATURE"]},

	// E110 00nn    Temperature Difference  K   (mK to  K)
	0x60: {Exp: 1.0e-3, Unit: MeasureUnit["K"], Type: VIFUnit["TEMPERATURE_DIFFERENCE"]},
	0x61: {Exp: 1.0e-2, Unit: MeasureUnit["K"], Type: VIFUnit["TEMPERATURE_DIFFERENCE"]},
	0x62: {Exp: 1.0e-1, Unit: MeasureUnit["K"], Type: VIFUnit["TEMPERATURE_DIFFERENCE"]},
	0x63: {Exp: 1.0e0, Unit: MeasureUnit["K"], Type: VIFUnit["TEMPERATURE_DIFFERENCE"]},

	// E110 01nn     External Temperature degC (0.001degC to 1degC)
	0x64: {Exp: 1.0e-3, Unit: MeasureUnit["C"], Type: VIFUnit["EXTERNAL_TEMPERATURE"]},
	0x65: {Exp: 1.0e-2, Unit: MeasureUnit["C"], Type: VIFUnit["EXTERNAL_TEMPERATURE"]},
	0x66: {Exp: 1.0e-1, Unit: MeasureUnit["C"], Type: VIFUnit["EXTERNAL_TEMPERATURE"]},
	0x67: {Exp: 1.0e0, Unit: MeasureUnit["C"], Type: VIFUnit["EXTERNAL_TEMPERATURE"]},

	// E110 10nn     Pressure bar (1mbar to 1000mbar)
	0x68: {Exp: 1.0e-3, Unit: MeasureUnit["BAR"], Type: VIFUnit["PRESSURE"]},
	0x69: {Exp: 1.0e-2, Unit: MeasureUnit["BAR"], Type: VIFUnit["PRESSURE"]},
	0x6A: {Exp: 1.0e-1, Unit: MeasureUnit["BAR"], Type: VIFUnit["PRESSURE"]},
	0x6B: {Exp: 1.0e0, Unit: MeasureUnit["BAR"], Type: VIFUnit["PRESSURE"]},

	// E110 110n     Time Point
	0x6C: {Exp: 1.0e0, Unit: MeasureUnit["DATE"], Type: VIFUnit["DATE"]},           // Exp G
	0x6D: {Exp: 1.0e0, Unit: MeasureUnit["DATE_TIME"], Type: VIFUnit["DATE_TIME"]}, // Exp F

	// E110 1110     Units for H.C.A. dimensionless
	0x6E: {Exp: 1.0e0, Unit: MeasureUnit["HCA"], Type: VIFUnit["UNITS_FOR_HCA"]},

	// E110 1111     Reserved
	0x6F: {Exp: 0.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["RES_THIRD_VIFE_TABLE"]},

	// E111 00nn     Averaging Duration s
	0x70: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["AVG_DURATION"]},     // seconds
	0x71: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["AVG_DURATION"]},    // minutes
	0x72: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["AVG_DURATION"]},  // hours
	0x73: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["AVG_DURATION"]}, // days

	// E111 01nn     Actuality Duration s
	0x74: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ACTUALITY_DURATION"]},
	0x75: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ACTUALITY_DURATION"]},
	0x76: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ACTUALITY_DURATION"]},
	0x77: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnit["ACTUALITY_DURATION"]},

	// Fabrication No
	0x78: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["FABRICATION_NO"]},

	// E111 1001 (Enhanced) Identification
	0x79: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["IDENTIFICATION"]},

	// E111 1010 Bus Address
	0x7A: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["ADDRESS"]},

	// Unknown VIF: 7Ch
	0x7C: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["ANY_VIF"]},

	// Any VIF: 7Eh
	0x7E: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["ANY_VIF"]},

	// Manufacturer specific: 7Fh
	0x7F: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["MANUFACTURER_SPEC"]},

	// Any VIF: 7Eh
	0xFE: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["ANY_VIF"]},

	// Manufacturer specific: FFh
	0xFF: {Exp: 1.0, Unit: MeasureUnit["NONE"], Type: VIFUnit["MANUFACTURER_SPEC"]},

	// Main VIFE-Code Extension table (following VIF=FDh for primary VIF)
	// See 8.4.4 a, only some of them are here. Using range 0x100 - 0x1FF

	// E000 00nn Credit of 10nn-3 of the nominal local legal currency units
	0x100: {Exp: 1.0e-3, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_CREDIT"]},
	0x101: {Exp: 1.0e-2, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_CREDIT"]},
	0x102: {Exp: 1.0e-1, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_CREDIT"]},
	0x103: {Exp: 1.0e0, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_CREDIT"]},

	// E000 01nn Debit of 10nn-3 of the nominal local legal currency units
	0x104: {Exp: 1.0e-3, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_DEBIT"]},
	0x105: {Exp: 1.0e-2, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_DEBIT"]},
	0x106: {Exp: 1.0e-1, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_DEBIT"]},
	0x107: {Exp: 1.0e0, Unit: MeasureUnit["CURRENCY"], Type: VIFUnitExt["CURRENCY_DEBIT"]},

	// E000 1000 Access Number (transmission count)
	0x108: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ACCESS_NUMBER"]},

	// E000 1001 Medium (as in fixed header)
	0x109: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["MEDIUM"]},

	// E000 1010 Manufacturer (as in fixed header)
	0x10A: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["MANUFACTURER"]},

	// E000 1011 Parameter set identification
	0x10B: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["PARAMETER_SET_ID"]},

	// E000 1100 Model / Version
	0x10C: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["MODEL_VERSION"]},

	// E000 1101 Hardware version //
	0x10D: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["HARDWARE_VERSION"]},

	// E000 1110 Firmware version //
	0x10E: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["FIRMWARE_VERSION"]},

	// E000 1111 Software version //
	0x10F: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["SOFTWARE_VERSION"]},

	// E001 0000 Customer location
	0x110: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["CUSTOMER_LOCATION"]},

	// E001 0001 Customer
	0x111: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["CUSTOMER"]},

	// E001 0010 Access Code User
	0x112: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ACCESS_CODE_USER"]},

	// E001 0011 Access Code Operator
	0x113: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ACCESS_CODE_OPERATOR"]},

	// E001 0100 Access Code System Operator
	0x114: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ACCESS_CODE_SYSTEM_OPERATOR"]},

	// E001 0101 Access Code Developer
	0x115: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ACCESS_CODE_DEVELOPER"]},

	// E001 0110 Password
	0x116: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["PASSWORD"]},

	// E001 0111 Error flags (binary)
	0x117: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ERROR_FLAGS"]},

	// E001 1000 Error mask
	0x118: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["ERROR_MASKS"]},

	// E001 1001 Reserved
	0x119: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},

	// E001 1010 Digital Output (binary)
	0x11A: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["DIGITAL_OUTPUT"]},

	// E001 1011 Digital Input (binary)
	0x11B: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["DIGITAL_INPUT"]},

	// E001 1100 Baudrate [Baud]
	0x11C: {Exp: 1.0e0, Unit: MeasureUnit["BAUD"], Type: VIFUnitExt["BAUDRATE"]},

	// E001 1101 Response delay time [bittimes]
	0x11D: {Exp: 1.0e0, Unit: MeasureUnit["BIT_TIMES"], Type: VIFUnitExt["RESPONSE_DELAY"]},

	// E001 1110 Retry
	0x11E: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RETRY"]},

	// E001 1111 Reserved
	0x11F: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED_2"]},

	// E010 0000 First storage // for cyclic storage
	0x120: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["FIRST_STORAGE_NR"]},

	// E010 0001 Last storage // for cyclic storage
	0x121: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["LAST_STORAGE_NR"]},

	// E010 0010 Size of storage block
	0x122: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["SIZE_OF_STORAGE_BLOCK"]},

	// E010 0011 Reserved
	0x123: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED_3"]},

	// E010 01nn Storage interval [sec(s)..day(s)]
	0x124: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},
	0x125: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},
	0x126: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},
	0x127: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},
	0x128: {Exp: 2629743.83, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},
	0x129: {Exp: 31556926.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},

	// E010 1010 Reserved
	0x12A: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},

	// E010 1011 Reserved
	0x12B: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},

	// E010 11nn Duration since last readout [sec(s)..day(s)]
	0x12C: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_READOUT"]},     // seconds
	0x12D: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_READOUT"]},    // minutes
	0x12E: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_READOUT"]},  // hours
	0x12F: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_READOUT"]}, // days

	// E011 0000 Start (date/time) of tariff
	// The information about usage of data Exp F (date and time) or data Exp G (date) can
	// be derived from the datafield (0010b: Exp G / 0100: Exp F).
	0x130: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]}, // ????

	// E011 00nn Duration of tariff (nn=01 ..11: min to days)
	0x131: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},    // minute(s)
	0x132: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]},  // hour(s)
	0x133: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["STORAGE_INTERVAL"]}, // day(s)

	// E011 01nn Period of tariff [sec(s) to day(s)]
	0x134: {Exp: 1.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["PERIOD_OF_TARIFF"]},        // seconds
	0x135: {Exp: 60.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["PERIOD_OF_TARIFF"]},       // minutes
	0x136: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["PERIOD_OF_TARIFF"]},     // hours
	0x137: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["PERIOD_OF_TARIFF"]},    // days
	0x138: {Exp: 2629743.83, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["PERIOD_OF_TARIFF"]}, // month(s)
	0x139: {Exp: 31556926.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["PERIOD_OF_TARIFF"]}, // year(s)

	// E011 1010 dimensionless / no VIF
	0x13A: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["DIMENSIONLESS"]},

	// E011 1011 Reserved
	0x13B: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},

	// E011 11xx Reserved
	0x13C: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x13D: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x13E: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x13F: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},

	// E100 nnnn   Volts electrical units
	0x140: {Exp: 1.0e-9, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x141: {Exp: 1.0e-8, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x142: {Exp: 1.0e-7, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x143: {Exp: 1.0e-6, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x144: {Exp: 1.0e-5, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x145: {Exp: 1.0e-4, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x146: {Exp: 1.0e-3, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x147: {Exp: 1.0e-2, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x148: {Exp: 1.0e-1, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x149: {Exp: 1.0e0, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x14A: {Exp: 1.0e1, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x14B: {Exp: 1.0e2, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x14C: {Exp: 1.0e3, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x14D: {Exp: 1.0e4, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x14E: {Exp: 1.0e5, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},
	0x14F: {Exp: 1.0e6, Unit: MeasureUnit["V"], Type: VIFUnitExt["VOLTS"]},

	// E101 nnnn   A
	0x150: {Exp: 1.0e-12, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x151: {Exp: 1.0e-11, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x152: {Exp: 1.0e-10, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x153: {Exp: 1.0e-9, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x154: {Exp: 1.0e-8, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x155: {Exp: 1.0e-7, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x156: {Exp: 1.0e-6, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x157: {Exp: 1.0e-5, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x158: {Exp: 1.0e-4, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x159: {Exp: 1.0e-3, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x15A: {Exp: 1.0e-2, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x15B: {Exp: 1.0e-1, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x15C: {Exp: 1.0e0, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x15D: {Exp: 1.0e1, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x15E: {Exp: 1.0e2, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},
	0x15F: {Exp: 1.0e3, Unit: MeasureUnit["A"], Type: VIFUnitExt["AMPERE"]},

	// E110 0000 Reset counter
	0x160: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESET_COUNTER"]},

	// E110 0001 Accumulation counter
	0x161: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["CUMULATION_COUNTER"]},

	// E110 0010 Control signal
	0x162: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["CONTROL_SIGNAL"]},

	// E110 0011 Day of week
	0x163: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["DAY_OF_WEEK"]},

	// E110 0100 Week number
	0x164: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["WEEK_NUMBER"]},

	// E110 0101 Time point of day change
	0x165: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["TIME_POINT_OF_DAY_CHANGE"]},

	// E110 0110 State of parameter activation
	0x166: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["STATE_OF_PARAMETER_ACTIVATION"]},

	// E110 0111 Special supplier information
	0x167: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["SPECIAL_SUPPLIER_INFORMATION"]},

	// E110 10pp Duration since last accumulation [hour(s)..years(s)]
	0x168: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_CUMULATION"]},     // hours
	0x169: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_CUMULATION"]},    // days
	0x16A: {Exp: 2629743.83, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_CUMULATION"]}, // month(s)
	0x16B: {Exp: 31556926.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["DURATION_SINCE_LAST_CUMULATION"]}, // year(s)

	// E110 11pp Operating time battery [hour(s)..years(s)]
	0x16C: {Exp: 3600.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["OPERATING_TIME_BATTERY"]},     // hours
	0x16D: {Exp: 86400.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["OPERATING_TIME_BATTERY"]},    // days
	0x16E: {Exp: 2629743.83, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["OPERATING_TIME_BATTERY"]}, // month(s)
	0x16F: {Exp: 31556926.0, Unit: MeasureUnit["SECONDS"], Type: VIFUnitExt["OPERATING_TIME_BATTERY"]}, // year(s)

	// E111 0000 Date and time of battery change
	0x170: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["DATEAND_TIME_OF_BATTERY_CHANGE"]},

	// E111 0001-1111 Reserved
	0x171: {Exp: 1.0e0, Unit: MeasureUnit["DBM"], Type: VIFUnitExt["RSSI"]},
	0x172: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x173: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x174: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x175: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x176: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x177: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x178: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x179: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x17A: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x17B: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x17C: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x17D: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x17E: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},
	0x17F: {Exp: 1.0e0, Unit: MeasureUnit["NONE"], Type: VIFUnitExt["RESERVED"]},

	// Alternate VIFE-Code Extension table (following VIF=0FBh for primary VIF)
	// See 8.4.4 b, only some of them are here. Using range 0x200 - 0x2FF

	// E000 000n Energy 10(n-1) MWh 0.1MWh to 1MWh
	0x200: {Exp: 1.0e5, Unit: MeasureUnit["WH"], VIFUnitDesc: "Energy"},
	0x201: {Exp: 1.0e6, Unit: MeasureUnit["WH"], VIFUnitDesc: "Energy"},

	// E000 001n Reserved
	0x202: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x203: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E000 01nn Reserved
	0x204: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x205: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x206: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x207: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E000 100n Energy 10(n-1) GJ 0.1GJ to 1GJ
	0x208: {Exp: 1.0e8, Unit: "Reserved", VIFUnitDesc: "Energy"},
	0x209: {Exp: 1.0e9, Unit: "Reserved", VIFUnitDesc: "Energy"},

	// E000 101n Reserved
	0x20A: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x20B: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E000 11nn Reserved
	0x20C: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x20D: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x20E: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x20F: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E001 000n Volume 10(n+2) m3 100m3 to 1000m3
	0x210: {Exp: 1.0e2, Unit: MeasureUnit["M3"], VIFUnitDesc: "Volume"},
	0x211: {Exp: 1.0e3, Unit: MeasureUnit["M3"], VIFUnitDesc: "Volume"},

	// E001 001n Reserved
	0x212: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x213: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E001 01nn Reserved
	0x214: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x215: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x216: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x217: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E001 100n Mass 10(n+2) t 100t to 1000t
	0x218: {Exp: 1.0e5, Unit: MeasureUnit["KG"], VIFUnitDesc: "Mass"},
	0x219: {Exp: 1.0e6, Unit: MeasureUnit["KG"], VIFUnitDesc: "Mass"},

	// E001 1010 to E010 0000 Reserved
	0x21A: {Exp: 1.0e-1, Unit: MeasureUnit["PERCENT"], Type: VIFUnitSecExt["RELATIVE_HUMIDITY"]},
	0x21B: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x21C: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x21D: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x21E: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x21F: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x220: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E010 0001 Volume 0,1 feet^3
	0x221: {Exp: 1.0e-1, Unit: "feet^3", VIFUnitDesc: "Volume"},

	// E010 001n Volume 0,1-1 american gallon
	0x222: {Exp: 1.0e-1, Unit: "American gallon", VIFUnitDesc: "Volume"},
	0x223: {Exp: 1.0e-0, Unit: "American gallon", VIFUnitDesc: "Volume"},

	// E010 0100    Volume flow 0,001 american gallon/min
	0x224: {Exp: 1.0e-3, Unit: "American gallon/min", VIFUnitDesc: "Volume flow"},

	// E010 0101 Volume flow 1 american gallon/min
	0x225: {Exp: 1.0e0, Unit: "American gallon/min", VIFUnitDesc: "Volume flow"},

	// E010 0110 Volume flow 1 american gallon/h
	0x226: {Exp: 1.0e0, Unit: "American gallon/h", VIFUnitDesc: "Volume flow"},

	// E010 0111 Reserved
	0x227: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E010 100n Power 10(n-1) MW 0.1MW to 1MW
	0x228: {Exp: 1.0e5, Unit: MeasureUnit["W"], VIFUnitDesc: "Power"},
	0x229: {Exp: 1.0e6, Unit: MeasureUnit["W"], VIFUnitDesc: "Power"},

	// E010 101n Reserved
	0x22A: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x22B: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E010 11nn Reserved
	0x22C: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x22D: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x22E: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x22F: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E011 000n Power 10(n-1) GJ/h 0.1GJ/h to 1GJ/h
	0x230: {Exp: 1.0e8, Unit: MeasureUnit["J"], VIFUnitDesc: "Power"},
	0x231: {Exp: 1.0e9, Unit: MeasureUnit["J"], VIFUnitDesc: "Power"},

	// E011 0010 to E101 0111 Reserved
	0x232: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x233: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x234: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x235: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x236: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x237: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x238: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x239: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x23A: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x23B: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x23C: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x23D: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x23E: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x23F: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x240: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x241: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x242: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x243: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x244: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x245: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x246: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x247: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x248: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x249: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x24A: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x24B: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x24C: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x24D: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x24E: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x24F: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x250: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x251: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x252: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x253: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x254: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x255: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x256: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x257: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E101 10nn Flow Temperature 10(nn-3) degF 0.001degF to 1degF
	0x258: {Exp: 1.0e-3, Unit: "degF", VIFUnitDesc: "Flow temperature"},
	0x259: {Exp: 1.0e-2, Unit: "degF", VIFUnitDesc: "Flow temperature"},
	0x25A: {Exp: 1.0e-1, Unit: "degF", VIFUnitDesc: "Flow temperature"},
	0x25B: {Exp: 1.0e0, Unit: "degF", VIFUnitDesc: "Flow temperature"},

	// E101 11nn Return Temperature 10(nn-3) degF 0.001degF to 1degF
	0x25C: {Exp: 1.0e-3, Unit: "degF", VIFUnitDesc: "Return temperature"},
	0x25D: {Exp: 1.0e-2, Unit: "degF", VIFUnitDesc: "Return temperature"},
	0x25E: {Exp: 1.0e-1, Unit: "degF", VIFUnitDesc: "Return temperature"},
	0x25F: {Exp: 1.0e0, Unit: "degF", VIFUnitDesc: "Return temperature"},

	// E110 00nn Temperature Difference 10(nn-3) degF 0.001degF to 1degF
	0x260: {Exp: 1.0e-3, Unit: "degF", VIFUnitDesc: "Temperature difference"},
	0x261: {Exp: 1.0e-2, Unit: "degF", VIFUnitDesc: "Temperature difference"},
	0x262: {Exp: 1.0e-1, Unit: "degF", VIFUnitDesc: "Temperature difference"},
	0x263: {Exp: 1.0e0, Unit: "degF", VIFUnitDesc: "Temperature difference"},

	// E110 01nn External Temperature 10(nn-3) degF 0.001degF to 1degF
	0x264: {Exp: 1.0e-3, Unit: "degF", VIFUnitDesc: "External temperature"},
	0x265: {Exp: 1.0e-2, Unit: "degF", VIFUnitDesc: "External temperature"},
	0x266: {Exp: 1.0e-1, Unit: "degF", VIFUnitDesc: "External temperature"},
	0x267: {Exp: 1.0e0, Unit: "degF", VIFUnitDesc: "External temperature"},

	// E110 1nnn Reserved
	0x268: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x269: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x26A: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x26B: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x26C: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x26D: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x26E: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},
	0x26F: {Exp: 1.0e0, Unit: "Reserved", VIFUnitDesc: "Reserved"},

	// E111 00nn Cold / Warm Temperature Limit 10(nn-3) degF 0.001degF to 1degF
	0x270: {Exp: 1.0e-3, Unit: "degF", VIFUnitDesc: "Cold / Warm Temperature Limit"},
	0x271: {Exp: 1.0e-2, Unit: "degF", VIFUnitDesc: "Cold / Warm Temperature Limit"},
	0x272: {Exp: 1.0e-1, Unit: "degF", VIFUnitDesc: "Cold / Warm Temperature Limit"},
	0x273: {Exp: 1.0e0, Unit: "degF", VIFUnitDesc: "Cold / Warm Temperature Limit"},

	// E111 01nn Cold / Warm Temperature Limit 10(nn-3) degC 0.001degC to 1degC
	0x274: {Exp: 1.0e-3, Unit: MeasureUnit["C"], VIFUnitDesc: "Cold / Warm Temperature Limit"},
	0x275: {Exp: 1.0e-2, Unit: MeasureUnit["C"], VIFUnitDesc: "Cold / Warm Temperature Limit"},
	0x276: {Exp: 1.0e-1, Unit: MeasureUnit["C"], VIFUnitDesc: "Cold / Warm Temperature Limit"},
	0x277: {Exp: 1.0e0, Unit: MeasureUnit["C"], VIFUnitDesc: "Cold / Warm Temperature Limit"},

	// E111 1nnn cumul. count max power 10(nnn-3) W 0.001W to 10000W
	0x279: {Exp: 1.0e-3, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x278: {Exp: 1.0e-3, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x27A: {Exp: 1.0e-1, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x27B: {Exp: 1.0e0, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x27C: {Exp: 1.0e1, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x27D: {Exp: 1.0e2, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x27E: {Exp: 1.0e3, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
	0x27F: {Exp: 1.0e4, Unit: MeasureUnit["W"], VIFUnitDesc: "Cumul count max power"},
}

var TelegramDateMasks = map[string]int{
	"DATE":          0x02, // "Actual Date",            	0010 Type G
	"DATE_TIME":     0x04, // "Actual Date and Time",    	0100 Type F
	"EXT_TIME":      0x03, // "Extended Date",           	0011 Type J
	"EXT_DATE_TIME": 0x60, // "Extended Date and Time", 	0110 Type I
}

var DateCalculatorMasks = map[string]int{
	"SECOND_MASK":       0x3F, // 0011 1111
	"MINUTE_MASK":       0x3F, // 0011 1111
	"HOUR_MASK":         0x1F, // 0001 1111
	"DAY_MASK":          0x1F, // 0001 1111
	"MONTH_MASK":        0x0F, // 0000 1111
	"YEAR_MASK":         0xE0, // 1110 0000
	"YEAR_MASK_2":       0xF0, // 1111 0000
	"HUNDRED_YEAR_MASK": 0xC0, // 1100 0000
	"WEEK_DAY":          0xE0, // 1110 0000
	"WEEK":              0x3F, // 0011 1111
	"TIME_INVALID":      0x80, // 1000 0000
	"SUMMERTIME":        0x40, // 0100 0000
	"LEAP_YEAR":         0x80, // 1000 0000
	"DIF_SUMMERTIME":    0xC0, // 1100 0000
}

type DateCalculator struct{}

func (dc *DateCalculator) GetTimeWithSeconds(second int, minute int, hour int) string {
	return fmt.Sprintf("%s%02d", dc.GetTime(minute, hour), dc.GetSeconds(second))
}

func (dc *DateCalculator) GetTime(minute int, hour int) string {
	return fmt.Sprintf("%02d%02d", dc.GetHour(hour), dc.GetMinutes(minute))
}

func (dc *DateCalculator) GetDate(day int, month int) string {
	return fmt.Sprintf("%04d-%02d-%02d",
		dc.GetYear(day, month, 0, false),
		dc.GetMonth(month),
		dc.GetDay(day),
	)
}

func (dc *DateCalculator) GetSeconds(second int) int {
	return second & DateCalculatorMasks["SECOND_MASK"]
}

func (dc *DateCalculator) GetMinutes(minute int) int {
	return minute & DateCalculatorMasks["MINUTE_MASK"]
}

func (dc *DateCalculator) GetHour(hour int) int {
	return hour & DateCalculatorMasks["HOUR_MASK"]
}

func (dc *DateCalculator) GetDay(day int) int {
	return day & DateCalculatorMasks["DAY_MASK"]
}

func (dc *DateCalculator) GetMonth(month int) int {
	return month & DateCalculatorMasks["MONTH_MASK"]
}

func (dc *DateCalculator) GetYear(v1 int, v2 int, hundredYearValue int, calcHundredYear bool) int {
	year1 := v1 & DateCalculatorMasks["YEAR_MASK"]
	year2 := v2 & DateCalculatorMasks["YEAR_MASK_2"]
	hundredYear := 1

	// we move the bits of year1 value 4 bits to the right
	// and concat (or) them with year2. Afterwards we have
	// to move the result one bit to the right so that it
	// is at the right position (0xxx xxxx).
	year := (year2 | (year1 >> 4)) >> 1

	// to be compatible with older meters it is recommended to interpret the
	// years 0 to 80 as 2000 to 2080. Only year values in between 0 and 99
	// should be used

	// another option is to calculate the hundred year value (in new meters)
	// from a third value the hundred year is generated and calculated
	// the year is then calculated according to following formula:
	// year = 1900 + 100 * hundredYear + year
	if calcHundredYear {
		hundredYear = (hundredYearValue & DateCalculatorMasks["HUNDRED_YEAR_MASK"]) >> 6
		year = 1900 + (100 * hundredYear) + year
	} else if year < 81 {
		year = 2000 + year
	} else {
		year = 1900 + year
	}

	return year
}
