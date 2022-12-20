package mpeg

// TODO: should not be named mpeg_*? just iti_*?

// ANSI/SCTE 128-1 2020
// https://wagtail-prod-storage.s3.amazonaws.com/documents/ANSI_SCTE-128-1-2020-1586877225672.pdf

import (
	"github.com/wader/fq/format"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/interp"
	"github.com/wader/fq/pkg/scalar"
)

var ituTT35MpegCCDataGroup decode.Group

func init() {
	interp.RegisterFormat(
		format.MPEG_ITU_T35,
		&decode.Format{
			Description: "ITU-T.35 closed captioning data",
			Dependencies: []decode.Dependency{
				{Groups: []*decode.Group{format.MPEG_CC_Data}, Out: &ituTT35MpegCCDataGroup},
			},
			DecodeFn: mpegItuT35SDecode,
		})
}

// TODO: sym some shorter string name?
var itutt35CountryCodesMap = scalar.UintMapSymStr{
	0x00: "Japan",
	0x01: "Albania",
	0x02: "Algeria",
	0x03: "American Samoa",
	0x04: "Germany",
	0x05: "Anguilla",
	0x06: "Antigua and Barbuda",
	0x07: "Argentina",
	0x08: "Ascension (see Saint Helena)",
	0x09: "Australia",
	0x0a: "Austria",
	0x0b: "Bahamas",
	0x0c: "Bahrain",
	0x0d: "Bangladesh",
	0x0e: "Barbados",
	0x0f: "Belgium",
	0x10: "Belize",
	0x11: "Benin",
	0x12: "Bermuda",
	0x13: "Bhutan",
	0x14: "Bolivia",
	0x15: "Botswana",
	0x16: "Brazil",
	0x17: "British Antarctic Territory",
	0x18: "British Indian Ocean Territory (Diego Garcia)",
	0x19: "British Virgin Islands",
	0x1a: "Brunei Darussalam",
	0x1b: "Bulgaria",
	0x1c: "Myanmar",
	0x1d: "Burundi",
	0x1e: "Belarus",
	0x1f: "Cameroon",
	0x20: "Canada",
	0x21: "Cape Verde",
	0x22: "Cayman Islands",
	0x23: "Central African Rep.",
	0x24: "Chad",
	0x25: "Chile",
	0x26: "China",
	0x27: "Colombia",
	0x28: "Comoros",
	0x29: "Congo",
	0x2a: "Cook Islands",
	0x2b: "Costa Rica",
	0x2c: "Cuba",
	0x2d: "Cyprus",
	0x2e: "Czech Rep.",
	0x2f: "Cambodia",
	0x30: "Dem. People's Rep. of Korea",
	0x31: "Denmark",
	0x32: "Djibouti",
	0x33: "Dominican Rep.",
	0x34: "Dominica",
	0x35: "Ecuador",
	0x36: "Egypt",
	0x37: "El Salvador",
	0x38: "Equatorial Guinea",
	0x39: "Ethiopia",
	0x3a: "Falkland Islands (Malvinas)",
	0x3b: "Fiji",
	0x3c: "Finland",
	0x3d: "France",
	0x3e: "French Polynesia",
	0x3f: "(Available)",
	0x40: "Gabon",
	0x41: "Gambia",
	0x42: "Germany",
	0x43: "Angola",
	0x44: "Ghana",
	0x45: "Gibraltar",
	0x46: "Greece",
	0x47: "Grenada",
	0x48: "Guam",
	0x49: "Guatemala",
	0x4a: "Guernsey",
	0x4b: "Guinea",
	0x4c: "Guinea-Bissau",
	0x4d: "Guyana",
	0x4e: "Haiti",
	0x4f: "Honduras",
	0x50: "Hong Kong: China",
	0x51: "Hungary",
	0x52: "Iceland",
	0x53: "India",
	0x54: "Indonesia",
	0x55: "Iran (Islamic Republic of)",
	0x56: "Iraq",
	0x57: "Ireland",
	0x58: "Israel",
	0x59: "Italy",
	0x5a: "Côte d'Ivoire",
	0x5b: "Jamaica",
	0x5c: "Afghanistan",
	0x5d: "Jersey",
	0x5e: "Jordan",
	0x5f: "Kenya",
	0x60: "Kiribati",
	0x61: "Korea (Rep. of)",
	0x62: "Kuwait",
	0x63: "Lao P.D.R.",
	0x64: "Lebanon",
	0x65: "Lesotho",
	0x66: "Liberia",
	0x67: "Libya",
	0x68: "Liechtenstein",
	0x69: "Luxembourg",
	0x6a: "Macao: China",
	0x6b: "Madagascar",
	0x6c: "Malaysia",
	0x6d: "Malawi",
	0x6e: "Maldives",
	0x6f: "Mali",
	0x70: "Malta",
	0x71: "Mauritania",
	0x72: "Mauritius",
	0x73: "Mexico",
	0x74: "Monaco",
	0x75: "Mongolia",
	0x76: "Montserrat",
	0x77: "Morocco",
	0x78: "Mozambique",
	0x79: "Nauru",
	0x7a: "Nepal",
	0x7b: "Netherlands",
	0x7c: "Curaçao",
	0x7d: "New Caledonia",
	0x7e: "New Zealand",
	0x7f: "Nicaragua",
	0x80: "Niger",
	0x81: "Nigeria",
	0x82: "Norway",
	0x83: "Oman",
	0x84: "Pakistan",
	0x85: "Panama",
	0x86: "Papua New Guinea",
	0x87: "Paraguay",
	0x88: "Peru",
	0x89: "Philippines",
	0x8a: "Poland",
	0x8b: "Portugal",
	0x8c: "Puerto Rico",
	0x8d: "Qatar",
	0x8e: "Romania",
	0x8f: "Rwanda",
	0x90: "Saint Kitts and Nevis",
	0x91: "Saint Croix",
	0x92: "Saint Helena: Ascension and Tristan da Cuhna",
	0x93: "Saint Lucia",
	0x94: "San Marino",
	0x95: "Saint Thomas",
	0x96: "Sao Tome and Principe",
	0x97: "Saint Vincent and the Grenadines",
	0x98: "Saudi Arabia",
	0x99: "Senegal",
	0x9a: "Seychelles",
	0x9b: "Sierra Leone",
	0x9c: "Singapore",
	0x9d: "Solomon Islands",
	0x9e: "Somalia",
	0x9f: "South Africa",
	0xa0: "Spain",
	0xa1: "Sri Lanka",
	0xa2: "Sudan",
	0xa3: "Suriname",
	0xa4: "Swaziland",
	0xa5: "Sweden",
	0xa6: "Switzerland",
	0xa7: "Syrian Arab Republic",
	0xa8: "Tanzania",
	0xa9: "Thailand",
	0xaa: "Togo",
	0xab: "Tonga",
	0xac: "Trinidad and Tobago",
	0xad: "Tunisia",
	0xae: "Turkey",
	0xaf: "Turks and Caicos Islands",
	0xb0: "Tuvalu",
	0xb1: "Uganda",
	0xb2: "Ukraine",
	0xb3: "United Arab Emirates",
	0xb4: "United Kingdom",
	0xb5: "United States",
	0xb6: "Burkina Faso",
	0xb7: "Uruguay",
	0xb8: "Russian Federation",
	0xb9: "Vanuatu",
	0xba: "Vatican",
	0xbb: "Venezuela",
	0xbc: "Viet Nam",
	0xbd: "Wallis and Futuna",
	0xbe: "Samoa",
	0xbf: "Yemen",
	0xc0: "Yemen",
	0xc1: "Serbia",
	0xc2: "Dem. Rep. of the Congo",
	0xc3: "Zambia",
	0xc4: "Zimbabwe",
	0xc5: "Slovakia",
	0xc6: "Slovenia",
	0xc7: "Lithuania",
	0xc8: "Montenegro",
	0xc9: "(Available)",
	0xca: "(Available)",
	0xcb: "(Available)",
	0xcc: "(Available)",
	0xcd: "(Available)",
	0xce: "(Available)",
	0xcf: "(Available)",
	0xff: "Escape code to extension list",
}

// TODO: sym some shorter string name?
var atsc1UserDataTypeCodeMap = scalar.UintMap{
	0x03: scalar.Uint{Sym: "dtvcc", Description: "DTVCC Captions (EIA-608, EIA-708)"},
	0x04: scalar.Uint{Sym: "additional_608_data", Description: "Additional 608 Data"},
	0x05: scalar.Uint{Sym: "luma_pam", Description: "Luma PAM"},
	0x06: scalar.Uint{Sym: "afd", Description: "AFD"},
	// TODO: more and there seems to be conflicts like for 0x7 and 0x8?
}

// 8.1.1. SEI T.35 Construct for Auxiliary Data Syntax
func mpegItuT35SDecode(d *decode.D) any {
	d.FieldU8("country_code", itutt35CountryCodesMap)
	d.FieldU16("provider_code") // TODO: per country code tables?
	userIdentifier := d.FieldUTF8("user_identifier", 4)

	switch userIdentifier {
	case "GA94":
		d.FieldStruct("user_structure", func(d *decode.D) {
			d.FieldU8("user_data_type_code", atsc1UserDataTypeCodeMap)
			d.FieldFormat("user_data_type_structure", &ituTT35MpegCCDataGroup, nil)
		})
	default:
		d.FieldRawLen("user_structure", d.BitsLeft())

	}

	return nil
}
