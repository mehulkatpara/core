/*
The file holds all the regex patterns used by the library.
Any regex pattern must be added here then used in the library.
*/
package validation

import "regexp"

/*
	3GPP TS 29.571 V17.7.0:
*/

// Common Patterns
var domainNameMatcher = regexp.MustCompile("^(?:[_a-z0-9](?:[_a-z0-9-]{0,61}[a-z0-9])?\\.)+(?:[a-z](?:[a-z0-9-]{0,61}[a-z0-9])?)?$")

// Table 5.3.2-1: Simple Data Types Regex
var amfIdPattern = regexp.MustCompile("^[A-Fa-f0-9]{6}$")
var amfRegionIdPattern = regexp.MustCompile("^[A-Fa-f0-9]{2}$")
var amfSetPattern = regexp.MustCompile("^[0-3][A-Fa-f0-9]{2}$")
var cagIdPattern = regexp.MustCompile("^[A-Fa-f0-9]{8}$")
var wildcardDnnPattern = regexp.MustCompile("^[*]$")
var gpsiPattern = regexp.MustCompile("^(msisdn-[0-9]{5,15}|extid-.+@.+|.+)$")
var groupIdPattern = regexp.MustCompile("^[A-Fa-f0-9]{8}-[0-9]{3}-[0-9]{2,3}-([A-Fa-f0-9][A-Fa-f0-9]){1,10}$")
var externalGroupIdPattern = regexp.MustCompile("^extgroupid-[^@]+@[^@]+$")
var peiPattern = regexp.MustCompile("^(imei-[0-9]{15}|imeisv-[0-9]{16}|mac((-[0-9a-fA-F]{2}){6})(-untrusted)?|eui((-[0-9a-fA-F]{2}){8})|.+)$")
var supiPattern = regexp.MustCompile("^(imsi-[0-9]{5,15}|nai-.+| gci-.+|gli-.+|.+)$")
var supiOrSuciPattern = regexp.MustCompile("^(imsi-[0-9]{5,15}|nai-.+|gli-.+|gci-.+|suci-(0-[0-9]{3}-[0-9]{2,3}|[1-7]-.+)-[0-9]{1,4}-(0-0-.*|[a-fA-F1-9]-([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])-[a-fA-F0-9]+)|.+)$")

// Table 5.4.2-1: Simple Data Types
var mccPattern = regexp.MustCompile("^[0-9]{3}$")
var mncPattern = regexp.MustCompile("^[0-9]{2,3}$")
var tacPattern = regexp.MustCompile("^[A-Fa-f0-9]{2,3}$")
var eutraCellIdPattern = regexp.MustCompile("^[A-Fa-f0-9]{7}$")
var nrCellIdPattern = regexp.MustCompile("^[A-Fa-f0-9]{9}$")
var n3IwfIdPattern = regexp.MustCompile("^[A-Fa-f0-9]+$")
var wAgfIdPattern = regexp.MustCompile("^[A-Fa-f0-9]+$")
var tngfIdPattern = regexp.MustCompile("^[A-Fa-f0-9]+$")
var ngeNbIdPattern = regexp.MustCompile("^('MacroNGeNB-[A-Fa-f0-9]{5}|LMacroNGeNB-[A-Fa-f0-9]{6}|SMacroNGeNB-[A-Fa-f0-9]{5})$")
var nidPattern = regexp.MustCompile("^[A-Fa-f0-9]{11}$")
var nfSetIdPattern = regexp.MustCompile("^([A-Za-z0-9\\-]*[A-Za-z0-9])$")
var nfServiceSetIdPattern = regexp.MustCompile("^([A-Za-z0-9\\-]*[A-Za-z0-9])$")
var typeAllocationCodePattern = regexp.MustCompile("^[0-9]{8}$")
var fiveGPrukIdPattern = regexp.MustCompile("^rid[0-9]{1,4}\\.pid[0-9a-fA-F]+\\@prose-cp\\.5gc\\.mnc[0-9]{2,3}\\.mcc[0-9]{3}\\.3gppnetwork\\.org$")
