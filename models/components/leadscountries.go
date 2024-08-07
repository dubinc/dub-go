// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// LeadsCountriesCountry - The 2-letter ISO 3166-1 country code for the country associated with the location of the user. Learn more: https://d.to/geo
type LeadsCountriesCountry string

const (
	LeadsCountriesCountryAf LeadsCountriesCountry = "AF"
	LeadsCountriesCountryAl LeadsCountriesCountry = "AL"
	LeadsCountriesCountryDz LeadsCountriesCountry = "DZ"
	LeadsCountriesCountryAs LeadsCountriesCountry = "AS"
	LeadsCountriesCountryAd LeadsCountriesCountry = "AD"
	LeadsCountriesCountryAo LeadsCountriesCountry = "AO"
	LeadsCountriesCountryAi LeadsCountriesCountry = "AI"
	LeadsCountriesCountryAq LeadsCountriesCountry = "AQ"
	LeadsCountriesCountryAg LeadsCountriesCountry = "AG"
	LeadsCountriesCountryAr LeadsCountriesCountry = "AR"
	LeadsCountriesCountryAm LeadsCountriesCountry = "AM"
	LeadsCountriesCountryAw LeadsCountriesCountry = "AW"
	LeadsCountriesCountryAu LeadsCountriesCountry = "AU"
	LeadsCountriesCountryAt LeadsCountriesCountry = "AT"
	LeadsCountriesCountryAz LeadsCountriesCountry = "AZ"
	LeadsCountriesCountryBs LeadsCountriesCountry = "BS"
	LeadsCountriesCountryBh LeadsCountriesCountry = "BH"
	LeadsCountriesCountryBd LeadsCountriesCountry = "BD"
	LeadsCountriesCountryBb LeadsCountriesCountry = "BB"
	LeadsCountriesCountryBy LeadsCountriesCountry = "BY"
	LeadsCountriesCountryBe LeadsCountriesCountry = "BE"
	LeadsCountriesCountryBz LeadsCountriesCountry = "BZ"
	LeadsCountriesCountryBj LeadsCountriesCountry = "BJ"
	LeadsCountriesCountryBm LeadsCountriesCountry = "BM"
	LeadsCountriesCountryBt LeadsCountriesCountry = "BT"
	LeadsCountriesCountryBo LeadsCountriesCountry = "BO"
	LeadsCountriesCountryBa LeadsCountriesCountry = "BA"
	LeadsCountriesCountryBw LeadsCountriesCountry = "BW"
	LeadsCountriesCountryBv LeadsCountriesCountry = "BV"
	LeadsCountriesCountryBr LeadsCountriesCountry = "BR"
	LeadsCountriesCountryIo LeadsCountriesCountry = "IO"
	LeadsCountriesCountryBn LeadsCountriesCountry = "BN"
	LeadsCountriesCountryBg LeadsCountriesCountry = "BG"
	LeadsCountriesCountryBf LeadsCountriesCountry = "BF"
	LeadsCountriesCountryBi LeadsCountriesCountry = "BI"
	LeadsCountriesCountryKh LeadsCountriesCountry = "KH"
	LeadsCountriesCountryCm LeadsCountriesCountry = "CM"
	LeadsCountriesCountryCa LeadsCountriesCountry = "CA"
	LeadsCountriesCountryCv LeadsCountriesCountry = "CV"
	LeadsCountriesCountryKy LeadsCountriesCountry = "KY"
	LeadsCountriesCountryCf LeadsCountriesCountry = "CF"
	LeadsCountriesCountryTd LeadsCountriesCountry = "TD"
	LeadsCountriesCountryCl LeadsCountriesCountry = "CL"
	LeadsCountriesCountryCn LeadsCountriesCountry = "CN"
	LeadsCountriesCountryCx LeadsCountriesCountry = "CX"
	LeadsCountriesCountryCc LeadsCountriesCountry = "CC"
	LeadsCountriesCountryCo LeadsCountriesCountry = "CO"
	LeadsCountriesCountryKm LeadsCountriesCountry = "KM"
	LeadsCountriesCountryCg LeadsCountriesCountry = "CG"
	LeadsCountriesCountryCd LeadsCountriesCountry = "CD"
	LeadsCountriesCountryCk LeadsCountriesCountry = "CK"
	LeadsCountriesCountryCr LeadsCountriesCountry = "CR"
	LeadsCountriesCountryCi LeadsCountriesCountry = "CI"
	LeadsCountriesCountryHr LeadsCountriesCountry = "HR"
	LeadsCountriesCountryCu LeadsCountriesCountry = "CU"
	LeadsCountriesCountryCy LeadsCountriesCountry = "CY"
	LeadsCountriesCountryCz LeadsCountriesCountry = "CZ"
	LeadsCountriesCountryDk LeadsCountriesCountry = "DK"
	LeadsCountriesCountryDj LeadsCountriesCountry = "DJ"
	LeadsCountriesCountryDm LeadsCountriesCountry = "DM"
	LeadsCountriesCountryDo LeadsCountriesCountry = "DO"
	LeadsCountriesCountryEc LeadsCountriesCountry = "EC"
	LeadsCountriesCountryEg LeadsCountriesCountry = "EG"
	LeadsCountriesCountrySv LeadsCountriesCountry = "SV"
	LeadsCountriesCountryGq LeadsCountriesCountry = "GQ"
	LeadsCountriesCountryEr LeadsCountriesCountry = "ER"
	LeadsCountriesCountryEe LeadsCountriesCountry = "EE"
	LeadsCountriesCountryEt LeadsCountriesCountry = "ET"
	LeadsCountriesCountryFk LeadsCountriesCountry = "FK"
	LeadsCountriesCountryFo LeadsCountriesCountry = "FO"
	LeadsCountriesCountryFj LeadsCountriesCountry = "FJ"
	LeadsCountriesCountryFi LeadsCountriesCountry = "FI"
	LeadsCountriesCountryFr LeadsCountriesCountry = "FR"
	LeadsCountriesCountryGf LeadsCountriesCountry = "GF"
	LeadsCountriesCountryPf LeadsCountriesCountry = "PF"
	LeadsCountriesCountryTf LeadsCountriesCountry = "TF"
	LeadsCountriesCountryGa LeadsCountriesCountry = "GA"
	LeadsCountriesCountryGm LeadsCountriesCountry = "GM"
	LeadsCountriesCountryGe LeadsCountriesCountry = "GE"
	LeadsCountriesCountryDe LeadsCountriesCountry = "DE"
	LeadsCountriesCountryGh LeadsCountriesCountry = "GH"
	LeadsCountriesCountryGi LeadsCountriesCountry = "GI"
	LeadsCountriesCountryGr LeadsCountriesCountry = "GR"
	LeadsCountriesCountryGl LeadsCountriesCountry = "GL"
	LeadsCountriesCountryGd LeadsCountriesCountry = "GD"
	LeadsCountriesCountryGp LeadsCountriesCountry = "GP"
	LeadsCountriesCountryGu LeadsCountriesCountry = "GU"
	LeadsCountriesCountryGt LeadsCountriesCountry = "GT"
	LeadsCountriesCountryGn LeadsCountriesCountry = "GN"
	LeadsCountriesCountryGw LeadsCountriesCountry = "GW"
	LeadsCountriesCountryGy LeadsCountriesCountry = "GY"
	LeadsCountriesCountryHt LeadsCountriesCountry = "HT"
	LeadsCountriesCountryHm LeadsCountriesCountry = "HM"
	LeadsCountriesCountryVa LeadsCountriesCountry = "VA"
	LeadsCountriesCountryHn LeadsCountriesCountry = "HN"
	LeadsCountriesCountryHk LeadsCountriesCountry = "HK"
	LeadsCountriesCountryHu LeadsCountriesCountry = "HU"
	LeadsCountriesCountryIs LeadsCountriesCountry = "IS"
	LeadsCountriesCountryIn LeadsCountriesCountry = "IN"
	LeadsCountriesCountryID LeadsCountriesCountry = "ID"
	LeadsCountriesCountryIr LeadsCountriesCountry = "IR"
	LeadsCountriesCountryIq LeadsCountriesCountry = "IQ"
	LeadsCountriesCountryIe LeadsCountriesCountry = "IE"
	LeadsCountriesCountryIl LeadsCountriesCountry = "IL"
	LeadsCountriesCountryIt LeadsCountriesCountry = "IT"
	LeadsCountriesCountryJm LeadsCountriesCountry = "JM"
	LeadsCountriesCountryJp LeadsCountriesCountry = "JP"
	LeadsCountriesCountryJo LeadsCountriesCountry = "JO"
	LeadsCountriesCountryKz LeadsCountriesCountry = "KZ"
	LeadsCountriesCountryKe LeadsCountriesCountry = "KE"
	LeadsCountriesCountryKi LeadsCountriesCountry = "KI"
	LeadsCountriesCountryKp LeadsCountriesCountry = "KP"
	LeadsCountriesCountryKr LeadsCountriesCountry = "KR"
	LeadsCountriesCountryKw LeadsCountriesCountry = "KW"
	LeadsCountriesCountryKg LeadsCountriesCountry = "KG"
	LeadsCountriesCountryLa LeadsCountriesCountry = "LA"
	LeadsCountriesCountryLv LeadsCountriesCountry = "LV"
	LeadsCountriesCountryLb LeadsCountriesCountry = "LB"
	LeadsCountriesCountryLs LeadsCountriesCountry = "LS"
	LeadsCountriesCountryLr LeadsCountriesCountry = "LR"
	LeadsCountriesCountryLy LeadsCountriesCountry = "LY"
	LeadsCountriesCountryLi LeadsCountriesCountry = "LI"
	LeadsCountriesCountryLt LeadsCountriesCountry = "LT"
	LeadsCountriesCountryLu LeadsCountriesCountry = "LU"
	LeadsCountriesCountryMo LeadsCountriesCountry = "MO"
	LeadsCountriesCountryMg LeadsCountriesCountry = "MG"
	LeadsCountriesCountryMw LeadsCountriesCountry = "MW"
	LeadsCountriesCountryMy LeadsCountriesCountry = "MY"
	LeadsCountriesCountryMv LeadsCountriesCountry = "MV"
	LeadsCountriesCountryMl LeadsCountriesCountry = "ML"
	LeadsCountriesCountryMt LeadsCountriesCountry = "MT"
	LeadsCountriesCountryMh LeadsCountriesCountry = "MH"
	LeadsCountriesCountryMq LeadsCountriesCountry = "MQ"
	LeadsCountriesCountryMr LeadsCountriesCountry = "MR"
	LeadsCountriesCountryMu LeadsCountriesCountry = "MU"
	LeadsCountriesCountryYt LeadsCountriesCountry = "YT"
	LeadsCountriesCountryMx LeadsCountriesCountry = "MX"
	LeadsCountriesCountryFm LeadsCountriesCountry = "FM"
	LeadsCountriesCountryMd LeadsCountriesCountry = "MD"
	LeadsCountriesCountryMc LeadsCountriesCountry = "MC"
	LeadsCountriesCountryMn LeadsCountriesCountry = "MN"
	LeadsCountriesCountryMs LeadsCountriesCountry = "MS"
	LeadsCountriesCountryMa LeadsCountriesCountry = "MA"
	LeadsCountriesCountryMz LeadsCountriesCountry = "MZ"
	LeadsCountriesCountryMm LeadsCountriesCountry = "MM"
	LeadsCountriesCountryNa LeadsCountriesCountry = "NA"
	LeadsCountriesCountryNr LeadsCountriesCountry = "NR"
	LeadsCountriesCountryNp LeadsCountriesCountry = "NP"
	LeadsCountriesCountryNl LeadsCountriesCountry = "NL"
	LeadsCountriesCountryNc LeadsCountriesCountry = "NC"
	LeadsCountriesCountryNz LeadsCountriesCountry = "NZ"
	LeadsCountriesCountryNi LeadsCountriesCountry = "NI"
	LeadsCountriesCountryNe LeadsCountriesCountry = "NE"
	LeadsCountriesCountryNg LeadsCountriesCountry = "NG"
	LeadsCountriesCountryNu LeadsCountriesCountry = "NU"
	LeadsCountriesCountryNf LeadsCountriesCountry = "NF"
	LeadsCountriesCountryMk LeadsCountriesCountry = "MK"
	LeadsCountriesCountryMp LeadsCountriesCountry = "MP"
	LeadsCountriesCountryNo LeadsCountriesCountry = "NO"
	LeadsCountriesCountryOm LeadsCountriesCountry = "OM"
	LeadsCountriesCountryPk LeadsCountriesCountry = "PK"
	LeadsCountriesCountryPw LeadsCountriesCountry = "PW"
	LeadsCountriesCountryPs LeadsCountriesCountry = "PS"
	LeadsCountriesCountryPa LeadsCountriesCountry = "PA"
	LeadsCountriesCountryPg LeadsCountriesCountry = "PG"
	LeadsCountriesCountryPy LeadsCountriesCountry = "PY"
	LeadsCountriesCountryPe LeadsCountriesCountry = "PE"
	LeadsCountriesCountryPh LeadsCountriesCountry = "PH"
	LeadsCountriesCountryPn LeadsCountriesCountry = "PN"
	LeadsCountriesCountryPl LeadsCountriesCountry = "PL"
	LeadsCountriesCountryPt LeadsCountriesCountry = "PT"
	LeadsCountriesCountryPr LeadsCountriesCountry = "PR"
	LeadsCountriesCountryQa LeadsCountriesCountry = "QA"
	LeadsCountriesCountryRe LeadsCountriesCountry = "RE"
	LeadsCountriesCountryRo LeadsCountriesCountry = "RO"
	LeadsCountriesCountryRu LeadsCountriesCountry = "RU"
	LeadsCountriesCountryRw LeadsCountriesCountry = "RW"
	LeadsCountriesCountrySh LeadsCountriesCountry = "SH"
	LeadsCountriesCountryKn LeadsCountriesCountry = "KN"
	LeadsCountriesCountryLc LeadsCountriesCountry = "LC"
	LeadsCountriesCountryPm LeadsCountriesCountry = "PM"
	LeadsCountriesCountryVc LeadsCountriesCountry = "VC"
	LeadsCountriesCountryWs LeadsCountriesCountry = "WS"
	LeadsCountriesCountrySm LeadsCountriesCountry = "SM"
	LeadsCountriesCountrySt LeadsCountriesCountry = "ST"
	LeadsCountriesCountrySa LeadsCountriesCountry = "SA"
	LeadsCountriesCountrySn LeadsCountriesCountry = "SN"
	LeadsCountriesCountrySc LeadsCountriesCountry = "SC"
	LeadsCountriesCountrySl LeadsCountriesCountry = "SL"
	LeadsCountriesCountrySg LeadsCountriesCountry = "SG"
	LeadsCountriesCountrySk LeadsCountriesCountry = "SK"
	LeadsCountriesCountrySi LeadsCountriesCountry = "SI"
	LeadsCountriesCountrySb LeadsCountriesCountry = "SB"
	LeadsCountriesCountrySo LeadsCountriesCountry = "SO"
	LeadsCountriesCountryZa LeadsCountriesCountry = "ZA"
	LeadsCountriesCountryGs LeadsCountriesCountry = "GS"
	LeadsCountriesCountryEs LeadsCountriesCountry = "ES"
	LeadsCountriesCountryLk LeadsCountriesCountry = "LK"
	LeadsCountriesCountrySd LeadsCountriesCountry = "SD"
	LeadsCountriesCountrySr LeadsCountriesCountry = "SR"
	LeadsCountriesCountrySj LeadsCountriesCountry = "SJ"
	LeadsCountriesCountrySz LeadsCountriesCountry = "SZ"
	LeadsCountriesCountrySe LeadsCountriesCountry = "SE"
	LeadsCountriesCountryCh LeadsCountriesCountry = "CH"
	LeadsCountriesCountrySy LeadsCountriesCountry = "SY"
	LeadsCountriesCountryTw LeadsCountriesCountry = "TW"
	LeadsCountriesCountryTj LeadsCountriesCountry = "TJ"
	LeadsCountriesCountryTz LeadsCountriesCountry = "TZ"
	LeadsCountriesCountryTh LeadsCountriesCountry = "TH"
	LeadsCountriesCountryTl LeadsCountriesCountry = "TL"
	LeadsCountriesCountryTg LeadsCountriesCountry = "TG"
	LeadsCountriesCountryTk LeadsCountriesCountry = "TK"
	LeadsCountriesCountryTo LeadsCountriesCountry = "TO"
	LeadsCountriesCountryTt LeadsCountriesCountry = "TT"
	LeadsCountriesCountryTn LeadsCountriesCountry = "TN"
	LeadsCountriesCountryTr LeadsCountriesCountry = "TR"
	LeadsCountriesCountryTm LeadsCountriesCountry = "TM"
	LeadsCountriesCountryTc LeadsCountriesCountry = "TC"
	LeadsCountriesCountryTv LeadsCountriesCountry = "TV"
	LeadsCountriesCountryUg LeadsCountriesCountry = "UG"
	LeadsCountriesCountryUa LeadsCountriesCountry = "UA"
	LeadsCountriesCountryAe LeadsCountriesCountry = "AE"
	LeadsCountriesCountryGb LeadsCountriesCountry = "GB"
	LeadsCountriesCountryUs LeadsCountriesCountry = "US"
	LeadsCountriesCountryUm LeadsCountriesCountry = "UM"
	LeadsCountriesCountryUy LeadsCountriesCountry = "UY"
	LeadsCountriesCountryUz LeadsCountriesCountry = "UZ"
	LeadsCountriesCountryVu LeadsCountriesCountry = "VU"
	LeadsCountriesCountryVe LeadsCountriesCountry = "VE"
	LeadsCountriesCountryVn LeadsCountriesCountry = "VN"
	LeadsCountriesCountryVg LeadsCountriesCountry = "VG"
	LeadsCountriesCountryVi LeadsCountriesCountry = "VI"
	LeadsCountriesCountryWf LeadsCountriesCountry = "WF"
	LeadsCountriesCountryEh LeadsCountriesCountry = "EH"
	LeadsCountriesCountryYe LeadsCountriesCountry = "YE"
	LeadsCountriesCountryZm LeadsCountriesCountry = "ZM"
	LeadsCountriesCountryZw LeadsCountriesCountry = "ZW"
	LeadsCountriesCountryAx LeadsCountriesCountry = "AX"
	LeadsCountriesCountryBq LeadsCountriesCountry = "BQ"
	LeadsCountriesCountryCw LeadsCountriesCountry = "CW"
	LeadsCountriesCountryGg LeadsCountriesCountry = "GG"
	LeadsCountriesCountryIm LeadsCountriesCountry = "IM"
	LeadsCountriesCountryJe LeadsCountriesCountry = "JE"
	LeadsCountriesCountryMe LeadsCountriesCountry = "ME"
	LeadsCountriesCountryBl LeadsCountriesCountry = "BL"
	LeadsCountriesCountryMf LeadsCountriesCountry = "MF"
	LeadsCountriesCountryRs LeadsCountriesCountry = "RS"
	LeadsCountriesCountrySx LeadsCountriesCountry = "SX"
	LeadsCountriesCountrySs LeadsCountriesCountry = "SS"
	LeadsCountriesCountryXk LeadsCountriesCountry = "XK"
)

func (e LeadsCountriesCountry) ToPointer() *LeadsCountriesCountry {
	return &e
}
func (e *LeadsCountriesCountry) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AF":
		fallthrough
	case "AL":
		fallthrough
	case "DZ":
		fallthrough
	case "AS":
		fallthrough
	case "AD":
		fallthrough
	case "AO":
		fallthrough
	case "AI":
		fallthrough
	case "AQ":
		fallthrough
	case "AG":
		fallthrough
	case "AR":
		fallthrough
	case "AM":
		fallthrough
	case "AW":
		fallthrough
	case "AU":
		fallthrough
	case "AT":
		fallthrough
	case "AZ":
		fallthrough
	case "BS":
		fallthrough
	case "BH":
		fallthrough
	case "BD":
		fallthrough
	case "BB":
		fallthrough
	case "BY":
		fallthrough
	case "BE":
		fallthrough
	case "BZ":
		fallthrough
	case "BJ":
		fallthrough
	case "BM":
		fallthrough
	case "BT":
		fallthrough
	case "BO":
		fallthrough
	case "BA":
		fallthrough
	case "BW":
		fallthrough
	case "BV":
		fallthrough
	case "BR":
		fallthrough
	case "IO":
		fallthrough
	case "BN":
		fallthrough
	case "BG":
		fallthrough
	case "BF":
		fallthrough
	case "BI":
		fallthrough
	case "KH":
		fallthrough
	case "CM":
		fallthrough
	case "CA":
		fallthrough
	case "CV":
		fallthrough
	case "KY":
		fallthrough
	case "CF":
		fallthrough
	case "TD":
		fallthrough
	case "CL":
		fallthrough
	case "CN":
		fallthrough
	case "CX":
		fallthrough
	case "CC":
		fallthrough
	case "CO":
		fallthrough
	case "KM":
		fallthrough
	case "CG":
		fallthrough
	case "CD":
		fallthrough
	case "CK":
		fallthrough
	case "CR":
		fallthrough
	case "CI":
		fallthrough
	case "HR":
		fallthrough
	case "CU":
		fallthrough
	case "CY":
		fallthrough
	case "CZ":
		fallthrough
	case "DK":
		fallthrough
	case "DJ":
		fallthrough
	case "DM":
		fallthrough
	case "DO":
		fallthrough
	case "EC":
		fallthrough
	case "EG":
		fallthrough
	case "SV":
		fallthrough
	case "GQ":
		fallthrough
	case "ER":
		fallthrough
	case "EE":
		fallthrough
	case "ET":
		fallthrough
	case "FK":
		fallthrough
	case "FO":
		fallthrough
	case "FJ":
		fallthrough
	case "FI":
		fallthrough
	case "FR":
		fallthrough
	case "GF":
		fallthrough
	case "PF":
		fallthrough
	case "TF":
		fallthrough
	case "GA":
		fallthrough
	case "GM":
		fallthrough
	case "GE":
		fallthrough
	case "DE":
		fallthrough
	case "GH":
		fallthrough
	case "GI":
		fallthrough
	case "GR":
		fallthrough
	case "GL":
		fallthrough
	case "GD":
		fallthrough
	case "GP":
		fallthrough
	case "GU":
		fallthrough
	case "GT":
		fallthrough
	case "GN":
		fallthrough
	case "GW":
		fallthrough
	case "GY":
		fallthrough
	case "HT":
		fallthrough
	case "HM":
		fallthrough
	case "VA":
		fallthrough
	case "HN":
		fallthrough
	case "HK":
		fallthrough
	case "HU":
		fallthrough
	case "IS":
		fallthrough
	case "IN":
		fallthrough
	case "ID":
		fallthrough
	case "IR":
		fallthrough
	case "IQ":
		fallthrough
	case "IE":
		fallthrough
	case "IL":
		fallthrough
	case "IT":
		fallthrough
	case "JM":
		fallthrough
	case "JP":
		fallthrough
	case "JO":
		fallthrough
	case "KZ":
		fallthrough
	case "KE":
		fallthrough
	case "KI":
		fallthrough
	case "KP":
		fallthrough
	case "KR":
		fallthrough
	case "KW":
		fallthrough
	case "KG":
		fallthrough
	case "LA":
		fallthrough
	case "LV":
		fallthrough
	case "LB":
		fallthrough
	case "LS":
		fallthrough
	case "LR":
		fallthrough
	case "LY":
		fallthrough
	case "LI":
		fallthrough
	case "LT":
		fallthrough
	case "LU":
		fallthrough
	case "MO":
		fallthrough
	case "MG":
		fallthrough
	case "MW":
		fallthrough
	case "MY":
		fallthrough
	case "MV":
		fallthrough
	case "ML":
		fallthrough
	case "MT":
		fallthrough
	case "MH":
		fallthrough
	case "MQ":
		fallthrough
	case "MR":
		fallthrough
	case "MU":
		fallthrough
	case "YT":
		fallthrough
	case "MX":
		fallthrough
	case "FM":
		fallthrough
	case "MD":
		fallthrough
	case "MC":
		fallthrough
	case "MN":
		fallthrough
	case "MS":
		fallthrough
	case "MA":
		fallthrough
	case "MZ":
		fallthrough
	case "MM":
		fallthrough
	case "NA":
		fallthrough
	case "NR":
		fallthrough
	case "NP":
		fallthrough
	case "NL":
		fallthrough
	case "NC":
		fallthrough
	case "NZ":
		fallthrough
	case "NI":
		fallthrough
	case "NE":
		fallthrough
	case "NG":
		fallthrough
	case "NU":
		fallthrough
	case "NF":
		fallthrough
	case "MK":
		fallthrough
	case "MP":
		fallthrough
	case "NO":
		fallthrough
	case "OM":
		fallthrough
	case "PK":
		fallthrough
	case "PW":
		fallthrough
	case "PS":
		fallthrough
	case "PA":
		fallthrough
	case "PG":
		fallthrough
	case "PY":
		fallthrough
	case "PE":
		fallthrough
	case "PH":
		fallthrough
	case "PN":
		fallthrough
	case "PL":
		fallthrough
	case "PT":
		fallthrough
	case "PR":
		fallthrough
	case "QA":
		fallthrough
	case "RE":
		fallthrough
	case "RO":
		fallthrough
	case "RU":
		fallthrough
	case "RW":
		fallthrough
	case "SH":
		fallthrough
	case "KN":
		fallthrough
	case "LC":
		fallthrough
	case "PM":
		fallthrough
	case "VC":
		fallthrough
	case "WS":
		fallthrough
	case "SM":
		fallthrough
	case "ST":
		fallthrough
	case "SA":
		fallthrough
	case "SN":
		fallthrough
	case "SC":
		fallthrough
	case "SL":
		fallthrough
	case "SG":
		fallthrough
	case "SK":
		fallthrough
	case "SI":
		fallthrough
	case "SB":
		fallthrough
	case "SO":
		fallthrough
	case "ZA":
		fallthrough
	case "GS":
		fallthrough
	case "ES":
		fallthrough
	case "LK":
		fallthrough
	case "SD":
		fallthrough
	case "SR":
		fallthrough
	case "SJ":
		fallthrough
	case "SZ":
		fallthrough
	case "SE":
		fallthrough
	case "CH":
		fallthrough
	case "SY":
		fallthrough
	case "TW":
		fallthrough
	case "TJ":
		fallthrough
	case "TZ":
		fallthrough
	case "TH":
		fallthrough
	case "TL":
		fallthrough
	case "TG":
		fallthrough
	case "TK":
		fallthrough
	case "TO":
		fallthrough
	case "TT":
		fallthrough
	case "TN":
		fallthrough
	case "TR":
		fallthrough
	case "TM":
		fallthrough
	case "TC":
		fallthrough
	case "TV":
		fallthrough
	case "UG":
		fallthrough
	case "UA":
		fallthrough
	case "AE":
		fallthrough
	case "GB":
		fallthrough
	case "US":
		fallthrough
	case "UM":
		fallthrough
	case "UY":
		fallthrough
	case "UZ":
		fallthrough
	case "VU":
		fallthrough
	case "VE":
		fallthrough
	case "VN":
		fallthrough
	case "VG":
		fallthrough
	case "VI":
		fallthrough
	case "WF":
		fallthrough
	case "EH":
		fallthrough
	case "YE":
		fallthrough
	case "ZM":
		fallthrough
	case "ZW":
		fallthrough
	case "AX":
		fallthrough
	case "BQ":
		fallthrough
	case "CW":
		fallthrough
	case "GG":
		fallthrough
	case "IM":
		fallthrough
	case "JE":
		fallthrough
	case "ME":
		fallthrough
	case "BL":
		fallthrough
	case "MF":
		fallthrough
	case "RS":
		fallthrough
	case "SX":
		fallthrough
	case "SS":
		fallthrough
	case "XK":
		*e = LeadsCountriesCountry(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LeadsCountriesCountry: %v", v)
	}
}

type LeadsCountries struct {
	// The 2-letter ISO 3166-1 country code for the country associated with the location of the user. Learn more: https://d.to/geo
	Country LeadsCountriesCountry `json:"country"`
	// The number of leads from this country
	Leads float64 `json:"leads"`
}

func (o *LeadsCountries) GetCountry() LeadsCountriesCountry {
	if o == nil {
		return LeadsCountriesCountry("")
	}
	return o.Country
}

func (o *LeadsCountries) GetLeads() float64 {
	if o == nil {
		return 0.0
	}
	return o.Leads
}
