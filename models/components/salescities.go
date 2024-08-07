// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// SalesCitiesCountry - The 2-letter country code of the city: https://d.to/geo
type SalesCitiesCountry string

const (
	SalesCitiesCountryAf SalesCitiesCountry = "AF"
	SalesCitiesCountryAl SalesCitiesCountry = "AL"
	SalesCitiesCountryDz SalesCitiesCountry = "DZ"
	SalesCitiesCountryAs SalesCitiesCountry = "AS"
	SalesCitiesCountryAd SalesCitiesCountry = "AD"
	SalesCitiesCountryAo SalesCitiesCountry = "AO"
	SalesCitiesCountryAi SalesCitiesCountry = "AI"
	SalesCitiesCountryAq SalesCitiesCountry = "AQ"
	SalesCitiesCountryAg SalesCitiesCountry = "AG"
	SalesCitiesCountryAr SalesCitiesCountry = "AR"
	SalesCitiesCountryAm SalesCitiesCountry = "AM"
	SalesCitiesCountryAw SalesCitiesCountry = "AW"
	SalesCitiesCountryAu SalesCitiesCountry = "AU"
	SalesCitiesCountryAt SalesCitiesCountry = "AT"
	SalesCitiesCountryAz SalesCitiesCountry = "AZ"
	SalesCitiesCountryBs SalesCitiesCountry = "BS"
	SalesCitiesCountryBh SalesCitiesCountry = "BH"
	SalesCitiesCountryBd SalesCitiesCountry = "BD"
	SalesCitiesCountryBb SalesCitiesCountry = "BB"
	SalesCitiesCountryBy SalesCitiesCountry = "BY"
	SalesCitiesCountryBe SalesCitiesCountry = "BE"
	SalesCitiesCountryBz SalesCitiesCountry = "BZ"
	SalesCitiesCountryBj SalesCitiesCountry = "BJ"
	SalesCitiesCountryBm SalesCitiesCountry = "BM"
	SalesCitiesCountryBt SalesCitiesCountry = "BT"
	SalesCitiesCountryBo SalesCitiesCountry = "BO"
	SalesCitiesCountryBa SalesCitiesCountry = "BA"
	SalesCitiesCountryBw SalesCitiesCountry = "BW"
	SalesCitiesCountryBv SalesCitiesCountry = "BV"
	SalesCitiesCountryBr SalesCitiesCountry = "BR"
	SalesCitiesCountryIo SalesCitiesCountry = "IO"
	SalesCitiesCountryBn SalesCitiesCountry = "BN"
	SalesCitiesCountryBg SalesCitiesCountry = "BG"
	SalesCitiesCountryBf SalesCitiesCountry = "BF"
	SalesCitiesCountryBi SalesCitiesCountry = "BI"
	SalesCitiesCountryKh SalesCitiesCountry = "KH"
	SalesCitiesCountryCm SalesCitiesCountry = "CM"
	SalesCitiesCountryCa SalesCitiesCountry = "CA"
	SalesCitiesCountryCv SalesCitiesCountry = "CV"
	SalesCitiesCountryKy SalesCitiesCountry = "KY"
	SalesCitiesCountryCf SalesCitiesCountry = "CF"
	SalesCitiesCountryTd SalesCitiesCountry = "TD"
	SalesCitiesCountryCl SalesCitiesCountry = "CL"
	SalesCitiesCountryCn SalesCitiesCountry = "CN"
	SalesCitiesCountryCx SalesCitiesCountry = "CX"
	SalesCitiesCountryCc SalesCitiesCountry = "CC"
	SalesCitiesCountryCo SalesCitiesCountry = "CO"
	SalesCitiesCountryKm SalesCitiesCountry = "KM"
	SalesCitiesCountryCg SalesCitiesCountry = "CG"
	SalesCitiesCountryCd SalesCitiesCountry = "CD"
	SalesCitiesCountryCk SalesCitiesCountry = "CK"
	SalesCitiesCountryCr SalesCitiesCountry = "CR"
	SalesCitiesCountryCi SalesCitiesCountry = "CI"
	SalesCitiesCountryHr SalesCitiesCountry = "HR"
	SalesCitiesCountryCu SalesCitiesCountry = "CU"
	SalesCitiesCountryCy SalesCitiesCountry = "CY"
	SalesCitiesCountryCz SalesCitiesCountry = "CZ"
	SalesCitiesCountryDk SalesCitiesCountry = "DK"
	SalesCitiesCountryDj SalesCitiesCountry = "DJ"
	SalesCitiesCountryDm SalesCitiesCountry = "DM"
	SalesCitiesCountryDo SalesCitiesCountry = "DO"
	SalesCitiesCountryEc SalesCitiesCountry = "EC"
	SalesCitiesCountryEg SalesCitiesCountry = "EG"
	SalesCitiesCountrySv SalesCitiesCountry = "SV"
	SalesCitiesCountryGq SalesCitiesCountry = "GQ"
	SalesCitiesCountryEr SalesCitiesCountry = "ER"
	SalesCitiesCountryEe SalesCitiesCountry = "EE"
	SalesCitiesCountryEt SalesCitiesCountry = "ET"
	SalesCitiesCountryFk SalesCitiesCountry = "FK"
	SalesCitiesCountryFo SalesCitiesCountry = "FO"
	SalesCitiesCountryFj SalesCitiesCountry = "FJ"
	SalesCitiesCountryFi SalesCitiesCountry = "FI"
	SalesCitiesCountryFr SalesCitiesCountry = "FR"
	SalesCitiesCountryGf SalesCitiesCountry = "GF"
	SalesCitiesCountryPf SalesCitiesCountry = "PF"
	SalesCitiesCountryTf SalesCitiesCountry = "TF"
	SalesCitiesCountryGa SalesCitiesCountry = "GA"
	SalesCitiesCountryGm SalesCitiesCountry = "GM"
	SalesCitiesCountryGe SalesCitiesCountry = "GE"
	SalesCitiesCountryDe SalesCitiesCountry = "DE"
	SalesCitiesCountryGh SalesCitiesCountry = "GH"
	SalesCitiesCountryGi SalesCitiesCountry = "GI"
	SalesCitiesCountryGr SalesCitiesCountry = "GR"
	SalesCitiesCountryGl SalesCitiesCountry = "GL"
	SalesCitiesCountryGd SalesCitiesCountry = "GD"
	SalesCitiesCountryGp SalesCitiesCountry = "GP"
	SalesCitiesCountryGu SalesCitiesCountry = "GU"
	SalesCitiesCountryGt SalesCitiesCountry = "GT"
	SalesCitiesCountryGn SalesCitiesCountry = "GN"
	SalesCitiesCountryGw SalesCitiesCountry = "GW"
	SalesCitiesCountryGy SalesCitiesCountry = "GY"
	SalesCitiesCountryHt SalesCitiesCountry = "HT"
	SalesCitiesCountryHm SalesCitiesCountry = "HM"
	SalesCitiesCountryVa SalesCitiesCountry = "VA"
	SalesCitiesCountryHn SalesCitiesCountry = "HN"
	SalesCitiesCountryHk SalesCitiesCountry = "HK"
	SalesCitiesCountryHu SalesCitiesCountry = "HU"
	SalesCitiesCountryIs SalesCitiesCountry = "IS"
	SalesCitiesCountryIn SalesCitiesCountry = "IN"
	SalesCitiesCountryID SalesCitiesCountry = "ID"
	SalesCitiesCountryIr SalesCitiesCountry = "IR"
	SalesCitiesCountryIq SalesCitiesCountry = "IQ"
	SalesCitiesCountryIe SalesCitiesCountry = "IE"
	SalesCitiesCountryIl SalesCitiesCountry = "IL"
	SalesCitiesCountryIt SalesCitiesCountry = "IT"
	SalesCitiesCountryJm SalesCitiesCountry = "JM"
	SalesCitiesCountryJp SalesCitiesCountry = "JP"
	SalesCitiesCountryJo SalesCitiesCountry = "JO"
	SalesCitiesCountryKz SalesCitiesCountry = "KZ"
	SalesCitiesCountryKe SalesCitiesCountry = "KE"
	SalesCitiesCountryKi SalesCitiesCountry = "KI"
	SalesCitiesCountryKp SalesCitiesCountry = "KP"
	SalesCitiesCountryKr SalesCitiesCountry = "KR"
	SalesCitiesCountryKw SalesCitiesCountry = "KW"
	SalesCitiesCountryKg SalesCitiesCountry = "KG"
	SalesCitiesCountryLa SalesCitiesCountry = "LA"
	SalesCitiesCountryLv SalesCitiesCountry = "LV"
	SalesCitiesCountryLb SalesCitiesCountry = "LB"
	SalesCitiesCountryLs SalesCitiesCountry = "LS"
	SalesCitiesCountryLr SalesCitiesCountry = "LR"
	SalesCitiesCountryLy SalesCitiesCountry = "LY"
	SalesCitiesCountryLi SalesCitiesCountry = "LI"
	SalesCitiesCountryLt SalesCitiesCountry = "LT"
	SalesCitiesCountryLu SalesCitiesCountry = "LU"
	SalesCitiesCountryMo SalesCitiesCountry = "MO"
	SalesCitiesCountryMg SalesCitiesCountry = "MG"
	SalesCitiesCountryMw SalesCitiesCountry = "MW"
	SalesCitiesCountryMy SalesCitiesCountry = "MY"
	SalesCitiesCountryMv SalesCitiesCountry = "MV"
	SalesCitiesCountryMl SalesCitiesCountry = "ML"
	SalesCitiesCountryMt SalesCitiesCountry = "MT"
	SalesCitiesCountryMh SalesCitiesCountry = "MH"
	SalesCitiesCountryMq SalesCitiesCountry = "MQ"
	SalesCitiesCountryMr SalesCitiesCountry = "MR"
	SalesCitiesCountryMu SalesCitiesCountry = "MU"
	SalesCitiesCountryYt SalesCitiesCountry = "YT"
	SalesCitiesCountryMx SalesCitiesCountry = "MX"
	SalesCitiesCountryFm SalesCitiesCountry = "FM"
	SalesCitiesCountryMd SalesCitiesCountry = "MD"
	SalesCitiesCountryMc SalesCitiesCountry = "MC"
	SalesCitiesCountryMn SalesCitiesCountry = "MN"
	SalesCitiesCountryMs SalesCitiesCountry = "MS"
	SalesCitiesCountryMa SalesCitiesCountry = "MA"
	SalesCitiesCountryMz SalesCitiesCountry = "MZ"
	SalesCitiesCountryMm SalesCitiesCountry = "MM"
	SalesCitiesCountryNa SalesCitiesCountry = "NA"
	SalesCitiesCountryNr SalesCitiesCountry = "NR"
	SalesCitiesCountryNp SalesCitiesCountry = "NP"
	SalesCitiesCountryNl SalesCitiesCountry = "NL"
	SalesCitiesCountryNc SalesCitiesCountry = "NC"
	SalesCitiesCountryNz SalesCitiesCountry = "NZ"
	SalesCitiesCountryNi SalesCitiesCountry = "NI"
	SalesCitiesCountryNe SalesCitiesCountry = "NE"
	SalesCitiesCountryNg SalesCitiesCountry = "NG"
	SalesCitiesCountryNu SalesCitiesCountry = "NU"
	SalesCitiesCountryNf SalesCitiesCountry = "NF"
	SalesCitiesCountryMk SalesCitiesCountry = "MK"
	SalesCitiesCountryMp SalesCitiesCountry = "MP"
	SalesCitiesCountryNo SalesCitiesCountry = "NO"
	SalesCitiesCountryOm SalesCitiesCountry = "OM"
	SalesCitiesCountryPk SalesCitiesCountry = "PK"
	SalesCitiesCountryPw SalesCitiesCountry = "PW"
	SalesCitiesCountryPs SalesCitiesCountry = "PS"
	SalesCitiesCountryPa SalesCitiesCountry = "PA"
	SalesCitiesCountryPg SalesCitiesCountry = "PG"
	SalesCitiesCountryPy SalesCitiesCountry = "PY"
	SalesCitiesCountryPe SalesCitiesCountry = "PE"
	SalesCitiesCountryPh SalesCitiesCountry = "PH"
	SalesCitiesCountryPn SalesCitiesCountry = "PN"
	SalesCitiesCountryPl SalesCitiesCountry = "PL"
	SalesCitiesCountryPt SalesCitiesCountry = "PT"
	SalesCitiesCountryPr SalesCitiesCountry = "PR"
	SalesCitiesCountryQa SalesCitiesCountry = "QA"
	SalesCitiesCountryRe SalesCitiesCountry = "RE"
	SalesCitiesCountryRo SalesCitiesCountry = "RO"
	SalesCitiesCountryRu SalesCitiesCountry = "RU"
	SalesCitiesCountryRw SalesCitiesCountry = "RW"
	SalesCitiesCountrySh SalesCitiesCountry = "SH"
	SalesCitiesCountryKn SalesCitiesCountry = "KN"
	SalesCitiesCountryLc SalesCitiesCountry = "LC"
	SalesCitiesCountryPm SalesCitiesCountry = "PM"
	SalesCitiesCountryVc SalesCitiesCountry = "VC"
	SalesCitiesCountryWs SalesCitiesCountry = "WS"
	SalesCitiesCountrySm SalesCitiesCountry = "SM"
	SalesCitiesCountrySt SalesCitiesCountry = "ST"
	SalesCitiesCountrySa SalesCitiesCountry = "SA"
	SalesCitiesCountrySn SalesCitiesCountry = "SN"
	SalesCitiesCountrySc SalesCitiesCountry = "SC"
	SalesCitiesCountrySl SalesCitiesCountry = "SL"
	SalesCitiesCountrySg SalesCitiesCountry = "SG"
	SalesCitiesCountrySk SalesCitiesCountry = "SK"
	SalesCitiesCountrySi SalesCitiesCountry = "SI"
	SalesCitiesCountrySb SalesCitiesCountry = "SB"
	SalesCitiesCountrySo SalesCitiesCountry = "SO"
	SalesCitiesCountryZa SalesCitiesCountry = "ZA"
	SalesCitiesCountryGs SalesCitiesCountry = "GS"
	SalesCitiesCountryEs SalesCitiesCountry = "ES"
	SalesCitiesCountryLk SalesCitiesCountry = "LK"
	SalesCitiesCountrySd SalesCitiesCountry = "SD"
	SalesCitiesCountrySr SalesCitiesCountry = "SR"
	SalesCitiesCountrySj SalesCitiesCountry = "SJ"
	SalesCitiesCountrySz SalesCitiesCountry = "SZ"
	SalesCitiesCountrySe SalesCitiesCountry = "SE"
	SalesCitiesCountryCh SalesCitiesCountry = "CH"
	SalesCitiesCountrySy SalesCitiesCountry = "SY"
	SalesCitiesCountryTw SalesCitiesCountry = "TW"
	SalesCitiesCountryTj SalesCitiesCountry = "TJ"
	SalesCitiesCountryTz SalesCitiesCountry = "TZ"
	SalesCitiesCountryTh SalesCitiesCountry = "TH"
	SalesCitiesCountryTl SalesCitiesCountry = "TL"
	SalesCitiesCountryTg SalesCitiesCountry = "TG"
	SalesCitiesCountryTk SalesCitiesCountry = "TK"
	SalesCitiesCountryTo SalesCitiesCountry = "TO"
	SalesCitiesCountryTt SalesCitiesCountry = "TT"
	SalesCitiesCountryTn SalesCitiesCountry = "TN"
	SalesCitiesCountryTr SalesCitiesCountry = "TR"
	SalesCitiesCountryTm SalesCitiesCountry = "TM"
	SalesCitiesCountryTc SalesCitiesCountry = "TC"
	SalesCitiesCountryTv SalesCitiesCountry = "TV"
	SalesCitiesCountryUg SalesCitiesCountry = "UG"
	SalesCitiesCountryUa SalesCitiesCountry = "UA"
	SalesCitiesCountryAe SalesCitiesCountry = "AE"
	SalesCitiesCountryGb SalesCitiesCountry = "GB"
	SalesCitiesCountryUs SalesCitiesCountry = "US"
	SalesCitiesCountryUm SalesCitiesCountry = "UM"
	SalesCitiesCountryUy SalesCitiesCountry = "UY"
	SalesCitiesCountryUz SalesCitiesCountry = "UZ"
	SalesCitiesCountryVu SalesCitiesCountry = "VU"
	SalesCitiesCountryVe SalesCitiesCountry = "VE"
	SalesCitiesCountryVn SalesCitiesCountry = "VN"
	SalesCitiesCountryVg SalesCitiesCountry = "VG"
	SalesCitiesCountryVi SalesCitiesCountry = "VI"
	SalesCitiesCountryWf SalesCitiesCountry = "WF"
	SalesCitiesCountryEh SalesCitiesCountry = "EH"
	SalesCitiesCountryYe SalesCitiesCountry = "YE"
	SalesCitiesCountryZm SalesCitiesCountry = "ZM"
	SalesCitiesCountryZw SalesCitiesCountry = "ZW"
	SalesCitiesCountryAx SalesCitiesCountry = "AX"
	SalesCitiesCountryBq SalesCitiesCountry = "BQ"
	SalesCitiesCountryCw SalesCitiesCountry = "CW"
	SalesCitiesCountryGg SalesCitiesCountry = "GG"
	SalesCitiesCountryIm SalesCitiesCountry = "IM"
	SalesCitiesCountryJe SalesCitiesCountry = "JE"
	SalesCitiesCountryMe SalesCitiesCountry = "ME"
	SalesCitiesCountryBl SalesCitiesCountry = "BL"
	SalesCitiesCountryMf SalesCitiesCountry = "MF"
	SalesCitiesCountryRs SalesCitiesCountry = "RS"
	SalesCitiesCountrySx SalesCitiesCountry = "SX"
	SalesCitiesCountrySs SalesCitiesCountry = "SS"
	SalesCitiesCountryXk SalesCitiesCountry = "XK"
)

func (e SalesCitiesCountry) ToPointer() *SalesCitiesCountry {
	return &e
}
func (e *SalesCitiesCountry) UnmarshalJSON(data []byte) error {
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
		*e = SalesCitiesCountry(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SalesCitiesCountry: %v", v)
	}
}

type SalesCities struct {
	// The name of the city
	City string `json:"city"`
	// The 2-letter country code of the city: https://d.to/geo
	Country SalesCitiesCountry `json:"country"`
	// The number of sales from this city
	Sales float64 `json:"sales"`
	// The total amount of sales from this city
	Amount float64 `json:"amount"`
}

func (o *SalesCities) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *SalesCities) GetCountry() SalesCitiesCountry {
	if o == nil {
		return SalesCitiesCountry("")
	}
	return o.Country
}

func (o *SalesCities) GetSales() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sales
}

func (o *SalesCities) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}
