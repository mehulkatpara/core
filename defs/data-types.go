package defs

/*
 * When PlmnId needs to be converted to string (e.g. when used in maps as key),
 * the string shall be composed of three digits "mcc" followed by "-" and two
 * or three digits "mnc", and shall match the following pattern: ^[0-9]{3}-[0-9]{2,3}$.
 *
 * Example 1: "262-01"
 * Example 2: "302-720"
 */
type PlmnId struct {
	Mcc string `json:"mcc"` // Mobile Country Code
	Mnc string `json:"mnc"` // Mobile Network Code
}
