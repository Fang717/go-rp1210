// Package rp1210 provides utilities for working with the RP1210 standard, with a focus on NEXIQ adapters.
// This file defines command constants for use with RP1210_SendCommand, including standard RP1210 commands
// and placeholders for NEXIQ-specific extensions.
package rp1210

// Command constants for RP1210_SendCommand, used to configure and control NEXIQ adapters.
// Values are based on the RP1210C/D standard unless noted as NEXIQ-specific.
// Constants are grouped by function for clarity and use uppercase naming for Go style.
// All commands should be verified against NEXIQ’s RP1210 SDK or driver documentation.
const (
    // General Commands
    CMD_RESET_DEVICE            = 0x0000 // Resets the NEXIQ adapter to its initial state.
    CMD_READ_VERSION            = 0x0001 // Reads DLL and API version information from the driver.
    CMD_READ_MESSAGE            = 0x0002 // Reads messages from the receive queue.
    CMD_DISALLOW_FURTHER_CONNECTIONS = 0x0013 // Prevents additional client connections to the adapter.
    CMD_SET_MAX_ERROR_MSG_SIZE  = 0x0015 // Sets the maximum size for error message buffers.
    CMD_GENERIC_DRIVER_COMMAND  = 0x8000 // Placeholder for NEXIQ-specific commands (verify with NEXIQ SDK).

    // Filtering Commands
    CMD_SET_ALL_FILTERS_TO_PASS = 0x0003 // Sets all protocol filters to pass messages.
    CMD_SET_ALL_FILTERS_TO_DISCARD = 0x000E // Sets all protocol filters to discard messages.
    CMD_SET_MESSAGE_RECEIVE     = 0x0010 // Enables or disables message reception.
    CMD_SET_J1708_FILTER_TYPE   = 0x0016 // Sets J1708 filter type (e.g., MID-based filtering).
    CMD_SET_J1939_FILTER_TYPE   = 0x0017 // Sets J1939 filter type (e.g., PGN-based filtering).
    CMD_SET_CAN_FILTER_TYPE     = 0x0018 // Sets CAN filter type (e.g., ID-based filtering).
    CMD_SET_ISO15765_FILTER_TYPE = 0x001B // Sets ISO15765 filter type (e.g., flow control filtering).

    // Message Filtering Commands
    CMD_SET_MESSAGE_FILTERING_J1939 = 0x0004 // Configures message filtering for J1939 protocol.
    CMD_SET_MESSAGE_FILTERING_CAN  = 0x0005 // Configures message filtering for CAN protocol.
    CMD_SET_MESSAGE_FILTERING_J1708 = 0x0007 // Configures message filtering for J1708 protocol.
    CMD_SET_MESSAGE_FILTERING_ISO15765 = 0x0009 // Configures message filtering for ISO15765 protocol.

    // Broadcast Commands
    CMD_SET_BROADCAST_FOR_J1939 = 0x000C // Configures broadcast messages for J1939 protocol.
    CMD_SET_BROADCAST_FOR_CAN   = 0x0011 // Configures broadcast messages for CAN protocol.
    CMD_SET_BROADCAST_FOR_J1708 = 0x0012 // Configures broadcast messages for J1708 protocol.
    CMD_SET_BROADCAST_FOR_ISO15765 = 0x001C // Configures broadcast messages for ISO15765 protocol.

    // Protocol-Specific Commands
    CMD_SET_J1708_MODE          = 0x000B // Sets J1708 operating mode (verify supported modes with NEXIQ).
    CMD_PROTECT_J1939_ADDRESS   = 0x000F // Claims and protects a J1939 address for communication.
    CMD_RELEASE_J1939_ADDRESS   = 0x0019 // Releases a claimed J1939 address.
    CMD_SET_J1939_INTERPACKET_TIME = 0x001D // Sets interpacket timing for J1939 messages.
    CMD_SET_ISO15765_FLOW_CONTROL = 0x001E // Configures flow control for ISO15765 protocol.
    CMD_CLEAR_ISO15765_FLOW_CONTROL = 0x001F // Clears flow control settings for ISO15765.
    CMD_SET_ISO15765_LINK_TYPE  = 0x0020 // Sets the link type for ISO15765 (e.g., CAN-based).
    CMD_SET_J1708_BAUD          = 0x0022 // Sets J1708 baud rate (verify supported rates with NEXIQ).
    CMD_SET_J1939_BAUD          = 0x0023 // Sets J1939 baud rate (e.g., 250k, 500k; verify with NEXIQ).
    CMD_SET_ISO15765_BAUD       = 0x0024 // Sets ISO15765 baud rate (verify with NEXIQ).
    CMD_SET_BLOCK_TIMEOUT       = 0x0025 // Sets blocking timeout for message operations.

    // Miscellaneous Commands
    CMD_ECHO_TRANSMITTED_MESSAGES = 0x000A // Enables or disables echo of transmitted messages.
    // Add more NEXIQ-specific commands as needed (e.g., from NEXIQ SDK, typically in range 0x8000–0xFFFF).
)