package main

var WMC int = 47
var ATFD int = 5
var TCC float64 = 0.3

func calc_WMC(s Struct) int {
	wmc := 0
	for _, m := range s.Methods {
		wmc += m.Complexity
	}
	return wmc
}

func NP(s Struct) int {
	n := len(s.Methods)
	if n <= 1 {
		return 0
	}
	return n * (n - 1) / 2
}

func NDC(s Struct) int {
	ndc := 0
	for j := 0; j < len(s.Methods)-1; j++ {
		for k := j + 1; k < len(s.Methods); k++ {
			if commonAttributeAccessExists(s.Methods[j], s.Methods[k]) {
				ndc++
			}
		}
	}
	return ndc
}

func commonAttributeAccessExists(m1 Method, m2 Method) bool {
	for _, v1 := range m1.SelfVarAccessed {
		for _, v2 := range m2.SelfVarAccessed {
			if v1.right == v2.right {
				return true
			}
		}
	}

	return false
}

func calc_ATFD(s Struct) int {
	uniqList := uniqeSelectors{}

	for _, m := range s.Methods {
		for _, v := range m.OthersVarAccessed {
			if !uniqList.exists(v) {
				uniqList.add(v)
			}
		}
	}

	atfd := len(uniqList.selectors)
	return atfd
}

func calc_TCC(s Struct) float64 {
	if s.NP == 0 {
		return TCC_Null
	}
	return float64(s.NDC) / float64(s.NP)
}

func GodStruct(s Struct) bool {
	if s.WMC > WMC && s.TCC < TCC && s.ATFD > ATFD {
		return true
	}

	return false
}

func DemiGodStruct(s Struct) bool {
	if GodStruct(s) {
		return false
	}

	demiGodEligibleCounter := 0

	if s.WMC > WMC {
		demiGodEligibleCounter++
	}

	if s.TCC < TCC {
		demiGodEligibleCounter++
	}

	if s.ATFD > ATFD {
		demiGodEligibleCounter++
	}

	if demiGodEligibleCounter == 2 {
		return true
	}

	return false

	// if s.WMC > 47 && s.TCC >= 0.3 && s.TCC != TCC_Null && s.ATFD <= 5 {
	// 	return true
	// }

	// if s.WMC <= 47 && s.TCC < 0.3 && s.ATFD <= 5 {
	// 	return true
	// }

	// if s.WMC <= 47 && s.TCC >= 0.3 && s.TCC != TCC_Null && s.ATFD > 5 {
	// 	return true
	// }

	// return false
}
