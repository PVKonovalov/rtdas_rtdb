package rtdb

type Quality uint32

const QualityNumberOfBits = 16

const (
	QualityGood Quality = 0

	QualityIec60870Overflow           Quality = 0x0001
	QualityIec60870Reserved2          Quality = 0x0002
	QualityIec60870Reserved4          Quality = 0x0004
	QualityIec60870ElapsedTimeInvalid Quality = 0x0008
	QualityIec60870Blocked            Quality = 0x0010
	QualityIec60870Substituted        Quality = 0x0020
	QualityIec60870NonTopical         Quality = 0x0040
	QualityIec60870Invalid            Quality = 0x0080
	QualityNoDataSource               Quality = 0x0100
	QualitySubstitutedManual          Quality = 0x0200
	QualitySubstitutedCalc            Quality = 0x0400
	QualityDeleteFromScan             Quality = 0x0800
	QualityLimitOverride              Quality = 0x1000
	QualityLimitExceeded              Quality = 0x2000
	QualityAlarmInhibit               Quality = 0x4000
	QualityTestMode                   Quality = 0x8000
)

func (q Quality) String() string {
	qualityCode := []string{"", "OV", "R2", "R4", "EI", "BL", "SB", "NT", "IV", "NS", "SM", "SC", "DS", "LO", "LE", "AI", "TM"}
	qualityString := ""

	needToAddSpace := false
	if q != 0 {
		for i := 0; i < QualityNumberOfBits; i++ {
			if q&(0x0001<<i) == 0x0001<<i {
				if needToAddSpace {
					qualityString += " "
				}
				qualityString += qualityCode[i+1]
				needToAddSpace = true
			}
		}
	}
	return qualityString
}
