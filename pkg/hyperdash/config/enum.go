// Code generated by generate-enum.go DO NOT EDIT.
package config

// UnitType type
type UnitType string

func (t UnitType) String() string {
	return string(t)
}

// UnitType enums
const (

	// Misc

	UnitTypeNone               UnitType = "none"
	UnitTypeString             UnitType = "string"
	UnitTypeShort              UnitType = "short"
	UnitTypePercent0100        UnitType = "percent"
	UnitTypePercent0010        UnitType = "percentunit"
	UnitTypeHumidityH          UnitType = "humidity"
	UnitTypeDecibel            UnitType = "dB"
	UnitTypeHexadecimal0x      UnitType = "hex0x"
	UnitTypeHexadecimal        UnitType = "hex"
	UnitTypeScientificnotation UnitType = "sci"
	UnitTypeLocaleformat       UnitType = "locale"
	UnitTypePixels             UnitType = "pixel"

	// Acceleration

	UnitTypeMeterssec UnitType = "accMS2"
	UnitTypeFeetsec   UnitType = "accFS2"
	UnitTypeGunit     UnitType = "accG"

	// Angle

	UnitTypeDegrees    UnitType = "degree"
	UnitTypeRadians    UnitType = "radian"
	UnitTypeGradian    UnitType = "grad"
	UnitTypeArcMinutes UnitType = "arcmin"
	UnitTypeArcSeconds UnitType = "arcsec"

	// Area

	UnitTypeSquareMetersm UnitType = "areaM2"
	UnitTypeSquareFeetft  UnitType = "areaF2"
	UnitTypeSquareMilesmi UnitType = "areaMI2"

	// Computation

	UnitTypeFLOPs  UnitType = "flops"
	UnitTypeMFLOPs UnitType = "mflops"
	UnitTypeGFLOPs UnitType = "gflops"
	UnitTypeTFLOPs UnitType = "tflops"
	UnitTypePFLOPs UnitType = "pflops"
	UnitTypeEFLOPs UnitType = "eflops"
	UnitTypeZFLOPs UnitType = "zflops"
	UnitTypeYFLOPs UnitType = "yflops"

	// Concentration

	UnitTypePartspermillionppm               UnitType = "ppm"
	UnitTypePartsperbillionppb               UnitType = "conppb"
	UnitTypeNanogrampercubicmeterngm         UnitType = "conngm3"
	UnitTypeNanogrampernormalcubicmeterngNm  UnitType = "conngNm3"
	UnitTypeMicrogrampercubicmetergm         UnitType = "conμgm3"
	UnitTypeMicrogrampernormalcubicmetergNm  UnitType = "conμgNm3"
	UnitTypeMilligrampercubicmetermgm        UnitType = "conmgm3"
	UnitTypeMilligrampernormalcubicmetermgNm UnitType = "conmgNm3"
	UnitTypeGrampercubicmetergm              UnitType = "congm3"
	UnitTypeGrampernormalcubicmetergNm       UnitType = "congNm3"
	UnitTypeMilligramsperdecilitremgdL       UnitType = "conmgdL"
	UnitTypeMillimolesperlitremmolL          UnitType = "conmmolL"

	// Currency

	UnitTypeDollars            UnitType = "currencyUSD"
	UnitTypePounds             UnitType = "currencyGBP"
	UnitTypeEuro               UnitType = "currencyEUR"
	UnitTypeYen                UnitType = "currencyJPY"
	UnitTypeRubles             UnitType = "currencyRUB"
	UnitTypeHryvnias           UnitType = "currencyUAH"
	UnitTypeRealR              UnitType = "currencyBRL"
	UnitTypeDanishKronekr      UnitType = "currencyDKK"
	UnitTypeIcelandicKrnakr    UnitType = "currencyISK"
	UnitTypeNorwegianKronekr   UnitType = "currencyNOK"
	UnitTypeSwedishKronakr     UnitType = "currencySEK"
	UnitTypeCzechkorunaczk     UnitType = "currencyCZK"
	UnitTypeSwissfrancCHF      UnitType = "currencyCHF"
	UnitTypePolishZotyPLN      UnitType = "currencyPLN"
	UnitTypeBitcoin            UnitType = "currencyBTC"
	UnitTypeMilliBitcoin       UnitType = "currencymBTC"
	UnitTypeMicroBitcoin       UnitType = "currencyμBTC"
	UnitTypeSouthAfricanRandR  UnitType = "currencyZAR"
	UnitTypeIndianRupee        UnitType = "currencyINR"
	UnitTypeSouthKoreanWon     UnitType = "currencyKRW"
	UnitTypeIndonesianRupiahRp UnitType = "currencyIDR"
	UnitTypePhilippinePesoPHP  UnitType = "currencyPHP"
	UnitTypeVietnameseDongVND  UnitType = "currencyVND"

	// Data

	UnitTypeBytesIEC  UnitType = "bytes"
	UnitTypeBytesSI   UnitType = "decbytes"
	UnitTypeBitsIEC   UnitType = "bits"
	UnitTypeBitsSI    UnitType = "decbits"
	UnitTypeKibibytes UnitType = "kbytes"
	UnitTypeKilobytes UnitType = "deckbytes"
	UnitTypeMebibytes UnitType = "mbytes"
	UnitTypeMegabytes UnitType = "decmbytes"
	UnitTypeGibibytes UnitType = "gbytes"
	UnitTypeGigabytes UnitType = "decgbytes"
	UnitTypeTebibytes UnitType = "tbytes"
	UnitTypeTerabytes UnitType = "dectbytes"
	UnitTypePebibytes UnitType = "pbytes"
	UnitTypePetabytes UnitType = "decpbytes"

	// Data rate

	UnitTypePacketssec   UnitType = "pps"
	UnitTypeBytessecIEC  UnitType = "binBps"
	UnitTypeBytessecSI   UnitType = "Bps"
	UnitTypeBitssecIEC   UnitType = "binbps"
	UnitTypeBitssecSI    UnitType = "bps"
	UnitTypeKibibytessec UnitType = "KiBs"
	UnitTypeKibibitssec  UnitType = "Kibits"
	UnitTypeKilobytessec UnitType = "KBs"
	UnitTypeKilobitssec  UnitType = "Kbits"
	UnitTypeMibibytessec UnitType = "MiBs"
	UnitTypeMibibitssec  UnitType = "Mibits"
	UnitTypeMegabytessec UnitType = "MBs"
	UnitTypeMegabitssec  UnitType = "Mbits"
	UnitTypeGibibytessec UnitType = "GiBs"
	UnitTypeGibibitssec  UnitType = "Gibits"
	UnitTypeGigabytessec UnitType = "GBs"
	UnitTypeGigabitssec  UnitType = "Gbits"
	UnitTypeTebibytessec UnitType = "TiBs"
	UnitTypeTebibitssec  UnitType = "Tibits"
	UnitTypeTerabytessec UnitType = "TBs"
	UnitTypeTerabitssec  UnitType = "Tbits"
	UnitTypePetibytessec UnitType = "PiBs"
	UnitTypePetibitssec  UnitType = "Pibits"
	UnitTypePetabytessec UnitType = "PBs"
	UnitTypePetabitssec  UnitType = "Pbits"

	// Date & time

	UnitTypeDatetimeISO                UnitType = "dateTimeAsIso"
	UnitTypeDatetimeISONodateiftoday   UnitType = "dateTimeAsIsoNoDateIfToday"
	UnitTypeDatetimeUS                 UnitType = "dateTimeAsUS"
	UnitTypeDatetimeUSNodateiftoday    UnitType = "dateTimeAsUSNoDateIfToday"
	UnitTypeDatetimelocal              UnitType = "dateTimeAsLocal"
	UnitTypeDatetimelocalNodateiftoday UnitType = "dateTimeAsLocalNoDateIfToday"
	UnitTypeDatetimedefault            UnitType = "dateTimeAsSystem"
	UnitTypeFromNow                    UnitType = "dateTimeFromNow"

	// Energy

	UnitTypeWattW                      UnitType = "watt"
	UnitTypeKilowattkW                 UnitType = "kwatt"
	UnitTypeMegawattMW                 UnitType = "megwatt"
	UnitTypeGigawattGW                 UnitType = "gwatt"
	UnitTypeMilliwattmW                UnitType = "mwatt"
	UnitTypeWattpersquaremeterWm       UnitType = "Wm2"
	UnitTypeVoltampereVA               UnitType = "voltamp"
	UnitTypeKilovoltamperekVA          UnitType = "kvoltamp"
	UnitTypeVoltamperereactivevar      UnitType = "voltampreact"
	UnitTypeKilovoltamperereactivekvar UnitType = "kvoltampreact"
	UnitTypeWatthourWh                 UnitType = "watth"
	UnitTypeWatthourperKilogramWhkg    UnitType = "watthperkg"
	UnitTypeKilowatthourkWh            UnitType = "kwatth"
	UnitTypeKilowattminkWm             UnitType = "kwattm"
	UnitTypeAmperehourAh               UnitType = "amph"
	UnitTypeKiloamperehourkAh          UnitType = "kamph"
	UnitTypeMilliamperehourmAh         UnitType = "mamph"
	UnitTypeJouleJ                     UnitType = "joule"
	UnitTypeElectronvolteV             UnitType = "ev"
	UnitTypeAmpereA                    UnitType = "amp"
	UnitTypeKiloamperekA               UnitType = "kamp"
	UnitTypeMilliamperemA              UnitType = "mamp"
	UnitTypeVoltV                      UnitType = "volt"
	UnitTypeKilovoltkV                 UnitType = "kvolt"
	UnitTypeMillivoltmV                UnitType = "mvolt"
	UnitTypeDecibelmilliwattdBm        UnitType = "dBm"
	UnitTypeOhm                        UnitType = "ohm"
	UnitTypeKiloohmk                   UnitType = "kohm"
	UnitTypeMegaohmM                   UnitType = "Mohm"
	UnitTypeFaradF                     UnitType = "farad"
	UnitTypeMicrofaradF                UnitType = "µfarad"
	UnitTypeNanofaradnF                UnitType = "nfarad"
	UnitTypePicofaradpF                UnitType = "pfarad"
	UnitTypeFemtofaradfF               UnitType = "ffarad"
	UnitTypeHenryH                     UnitType = "henry"
	UnitTypeMillihenrymH               UnitType = "mhenry"
	UnitTypeMicrohenryH                UnitType = "µhenry"
	UnitTypeLumensLm                   UnitType = "lumens"

	// Flow

	UnitTypeGallonsmingpm      UnitType = "flowgpm"
	UnitTypeCubicmetersseccms  UnitType = "flowcms"
	UnitTypeCubicfeetseccfs    UnitType = "flowcfs"
	UnitTypeCubicfeetmincfm    UnitType = "flowcfm"
	UnitTypeLitrehour          UnitType = "litreh"
	UnitTypeLitreminLmin       UnitType = "flowlpm"
	UnitTypeMilliLitreminmLmin UnitType = "flowmlpm"
	UnitTypeLuxlx              UnitType = "lux"

	// Force

	UnitTypeNewtonmetersNm      UnitType = "forceNm"
	UnitTypeKilonewtonmeterskNm UnitType = "forcekNm"
	UnitTypeNewtonsN            UnitType = "forceN"
	UnitTypeKilonewtonskN       UnitType = "forcekN"

	// Hash rate

	UnitTypeHashessec     UnitType = "Hs"
	UnitTypeKilohashessec UnitType = "KHs"
	UnitTypeMegahashessec UnitType = "MHs"
	UnitTypeGigahashessec UnitType = "GHs"
	UnitTypeTerahashessec UnitType = "THs"
	UnitTypePetahashessec UnitType = "PHs"
	UnitTypeExahashessec  UnitType = "EHs"

	// Mass

	UnitTypeMilligrammg UnitType = "massmg"
	UnitTypeGramg       UnitType = "massg"
	UnitTypeKilogramkg  UnitType = "masskg"
	UnitTypeMetrictont  UnitType = "masst"

	// Length

	UnitTypeMillimetermm UnitType = "lengthmm"
	UnitTypeFeetft       UnitType = "lengthft"
	UnitTypeMeterm       UnitType = "lengthm"
	UnitTypeKilometerkm  UnitType = "lengthkm"
	UnitTypeMilemi       UnitType = "lengthmi"

	// Pressure

	UnitTypeMillibars       UnitType = "pressurembar"
	UnitTypeBars            UnitType = "pressurebar"
	UnitTypeKilobars        UnitType = "pressurekbar"
	UnitTypePascals         UnitType = "pressurepa"
	UnitTypeHectopascals    UnitType = "pressurehpa"
	UnitTypeKilopascals     UnitType = "pressurekpa"
	UnitTypeInchesofmercury UnitType = "pressurehg"
	UnitTypePSI             UnitType = "pressurepsi"

	// Radiation

	UnitTypeBecquerelBq          UnitType = "radbq"
	UnitTypeCurieCi              UnitType = "radci"
	UnitTypeGrayGy               UnitType = "radgy"
	UnitTypeRad                  UnitType = "radrad"
	UnitTypeSievertSv            UnitType = "radsv"
	UnitTypeMilliSievertmSv      UnitType = "radmsv"
	UnitTypeMicroSievertSv       UnitType = "radusv"
	UnitTypeRem                  UnitType = "radrem"
	UnitTypeExposureCkg          UnitType = "radexpckg"
	UnitTypeRoentgenR            UnitType = "radr"
	UnitTypeSieverthourSvh       UnitType = "radsvh"
	UnitTypeMilliSieverthourmSvh UnitType = "radmsvh"
	UnitTypeMicroSieverthourSvh  UnitType = "radusvh"

	// Rotational Speed

	UnitTypeRevolutionsperminuterpm UnitType = "rotrpm"
	UnitTypeHertzHz                 UnitType = "rothz"
	UnitTypeRadianspersecondrads    UnitType = "rotrads"
	UnitTypeDegreesperseconds       UnitType = "rotdegs"

	// Temperature

	UnitTypeCelsiusC    UnitType = "celsius"
	UnitTypeFahrenheitF UnitType = "fahrenheit"
	UnitTypeKelvinK     UnitType = "kelvin"

	// Time

	UnitTypeHertz1s         UnitType = "hertz"
	UnitTypeNanosecondsns   UnitType = "ns"
	UnitTypeMicrosecondss   UnitType = "µs"
	UnitTypeMillisecondsms  UnitType = "ms"
	UnitTypeSecondss        UnitType = "s"
	UnitTypeMinutesm        UnitType = "m"
	UnitTypeHoursh          UnitType = "h"
	UnitTypeDaysd           UnitType = "d"
	UnitTypeDurationms      UnitType = "dtdurationms"
	UnitTypeDurations       UnitType = "dtdurations"
	UnitTypeDurationhhmmss  UnitType = "dthms"
	UnitTypeDurationdhhmmss UnitType = "dtdhms"
	UnitTypeTimetickss100   UnitType = "timeticks"
	UnitTypeClockms         UnitType = "clockms"
	UnitTypeClocks          UnitType = "clocks"

	// Throughput

	UnitTypeCountsseccps   UnitType = "cps"
	UnitTypeOpssecops      UnitType = "ops"
	UnitTypeRequestssecrps UnitType = "reqps"
	UnitTypeReadssecrps    UnitType = "rps"
	UnitTypeWritessecwps   UnitType = "wps"
	UnitTypeIOopsseciops   UnitType = "iops"
	UnitTypeCountsmincpm   UnitType = "cpm"
	UnitTypeOpsminopm      UnitType = "opm"
	UnitTypeReadsminrpm    UnitType = "rpm"
	UnitTypeWritesminwpm   UnitType = "wpm"

	// Velocity

	UnitTypeMeterssecondms    UnitType = "velocityms"
	UnitTypeKilometershourkmh UnitType = "velocitykmh"
	UnitTypeMileshourmph      UnitType = "velocitymph"
	UnitTypeKnotkn            UnitType = "velocityknot"

	// Volume

	UnitTypeMillilitremL     UnitType = "mlitre"
	UnitTypeLitreL           UnitType = "litre"
	UnitTypeCubicmeter       UnitType = "m3"
	UnitTypeNormalcubicmeter UnitType = "Nm3"
	UnitTypeCubicdecimeter   UnitType = "dm3"
	UnitTypeGallons          UnitType = "gallons"
)