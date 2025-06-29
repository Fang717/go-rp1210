// Package rp1210 provides utilities for working with the RP1210 standard, with a focus on NEXIQ adapters.
// It includes error code mappings for standard RP1210 errors and NEXIQ-specific extensions.
package rp1210

// RP1210 provides utilities for working with maps the RP1210 error codes to their human-readable descriptions.
// The map is read-only and safe for concurrent access. Do not modify this map at runtime.
// Error codes are grouped by category for clarity, based on the RP1210C standard and NEXIQ-specific extensions.
var RP1210Errors = map[uint16]string{
    // Client-related errors (RP1210 standard)
    128: "ERR_DLL_NOT_INITIALIZED",
    129: "ERR_INVALID_CLIENT_ID",
    130: "ERR_CLIENT_ALREADY_CONNECTED",
    131: "ERR_CLIENT_AREA_FULL",
    148: "ERR_CLIENT_DISCONNECTED",
    149: "ERR_CONNECT_NOT_ALLOWED",
    156: "ERR_MULTIPLE_CLIENTS_CONNECTED",

    // Memory-related errors (RP1210 standard)
    132: "ERR_FREE_MEMORY",
    133: "ERR_NOT_ENOUGH_MEMORY",

    // Device-related errors (RP1210 standard)
    134: "ERR_INVALID_DEVICE",
    135: "ERR_DEVICE_IN_USE",
    142: "ERR_HARDWARE_NOT_RESPONDING",
    162: "ERR_HARDWARE_STATUS_CHANGE",
    207: "ERR_DEVICE_NOT_SUPPORTED",
    230: "ERR_COMM_DEVICE_IN_USE",
    453: "ERR_ADAPTER_NOT_RESPONDING",
    602: "ERR_HARDWARE_NOT_SUPPORTED",

    // Protocol-related errors (RP1210 standard)
    136: "ERR_INVALID_PROTOCOL",
    224: "ERR_CANNOT_SET_CAN_BAUDRATE",
    454: "ERR_CAN_BAUD_SET_NONSTANDARD",
    456: "ERR_J1708_BAUD_SET_NONSTANDARD",
    457: "ERR_J1939_BAUD_SET_NONSTANDARD",
    458: "ERR_ISO15765_BAUD_SET_NONSTANDARD",
    603: "ERR_CANNOT_DETERMINE_BAUD_RATE",

    // Queue-related errors (RP1210 standard)
    137: "ERR_TX_QUEUE_FULL",
    138: "ERR_TX_QUEUE_CORRUPT",
    139: "ERR_RX_QUEUE_FULL",
    140: "ERR_RX_QUEUE_CORRUPT",
    222: "ERR_COMMAND_QUEUE_IS_FULL",

    // Command-related errors (RP1210 standard)
    141: "ERR_MESSAGE_TOO_LONG",
    143: "ERR_COMMAND_NOT_SUPPORTED",
    144: "ERR_INVALID_COMMAND",
    145: "ERR_TXMESSAGE_STATUS",
    159: "ERR_MESSAGE_NOT_SENT",
    600: "ERR_INVALID_IOCTL_ID",
    601: "ERR_NULL_PARAMETER",

    // Address-related errors (RP1210 standard)
    146: "ERR_ADDRESS_CLAIM_FAILED",
    147: "ERR_CANNOT_SET_PRIORITY",
    152: "ERR_COULD_NOT_TX_ADDRESS_CLAIMED",
    153: "ERR_ADDRESS_LOST",
    157: "ERR_ADDRESS_NEVER_CLAIMED",
    225: "ERR_CANNOT_CLAIM_BROADCAST_ADDRESS",
    226: "ERR_OUT_OF_ADDRESS_RESOURCES",
    227: "ERR_ADDRESS_RELEASE_FAILED",

    // Configuration-related errors (RP1210 standard)
    202: "ERR_INI_FILE_NOT_IN_WIN_DIR",
    204: "ERR_INI_SECTION_NOT_FOUND",
    205: "ERR_INI_KEY_NOT_FOUND",
    206: "ERR_INVALID_KEY_STRING",

    // Miscellaneous errors (RP1210 standard)
    150: "ERR_CHANGE_MODE_FAILED",
    151: "ERR_BUS_OFF",
    154: "ERR_CODE_NOT_FOUND",
    155: "ERR_BLOCK_NOT_ALLOWED",
    158: "ERR_WINDOW_HANDLE_REQUIRED",
    160: "ERR_MAX_NOTIFY_EXCEEDED",
    161: "ERR_MAX_FILTERS_EXCEEDED",
    208: "ERR_INVALID_PORT_PARAM",
    220: "ERR_OS_NOT_SUPPORTED",
    441: "ERR_DATA_LINK_CONFLICT",
    455: "ERR_MULTIPLE_CONNECTIONS_NOT_ALLOWED",

    // NEXIQ-specific errors (non-standard, from NEXIQ RP1210 API or SDK)
    3000: "API_RP1210_INI_NOT_FOUND",
    3001: "API_VENDOR_INI_NOT_FOUND",
    3002: "API_INVALID_VENDOR_FILE_NAME",
    3003: "API_VENDOR_DLL_FILE_NOT_FOUND",
    3004: "API_NOT_INITIALIZED",
    3005: "API_DLL_NOT_LOADED",
    3006: "API_FUNCTIONS_NOT_INITIALIZED",
    3007: "API_NO_VENDORS_FOUND",
    3008: "API_NO_DEVICES_FOUND",
    3009: "API_NO_PROTOCOLS_FOUND",
    3010: "API_FUNCTION_NOT_SUPPORTED",
    // Add more NEXIQ-specific or other vendor-specific codes as needed
}

// GetErrorDescription returns the human-readable description for a given RP1210 error code.
// If the code is not found in RP1210Errors, it returns "UNKNOWN_ERROR".
// This function is safe for concurrent use.
func GetErrorDescription(code uint16) string {
    if desc, ok := RP1210Errors[code]; ok {
        return desc
    }
    return "UNKNOWN_ERROR"
}