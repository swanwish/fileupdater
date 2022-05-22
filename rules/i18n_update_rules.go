package rules

import (
	"github.com/swanwish/fileupdaterdemo/common"
)

var matchReplacer = i18nMatchReplacer{}

var (
	I18nJavaUpdateRule = UpdateRule{
		Suffixes:         []string{common.ExtensionJava},
		Pattern:          "StringProvider\\.getValue\\(\"(.*?)\",\\s*?\"(.*?)\"\\)",
		GetMatchReplacer: matchReplacer.GetJavaMatchReplacer,
	}
	I18nKotlinUpdateRule = UpdateRule{
		Suffixes:         []string{common.ExtensionKotlin},
		Pattern:          "StringProvider\\.getValue\\(\"(.*?)\",\\s*?\"(.*?)\"\\)",
		GetMatchReplacer: matchReplacer.GetKotlinMatchReplacer,
	}
	I18nObjcUpdateRule = UpdateRule{
		Suffixes:         []string{common.ExtensionObjc},
		Pattern:          "\\[\\[StringProvider sharedInstance\\]\\s*?getValue:@\"(.*?)\"\\s*?defaultValue:@\"(.*?)\"\\]",
		GetMatchReplacer: matchReplacer.GetObjcMatchReplacer,
	}
	I18nSwiftUpdateReule = UpdateRule{
		Suffixes:         []string{common.ExtensionSwift},
		Pattern:          "StringProvider\\.sharedInstance\\(\\)\\.getValue\\(\"(.*?)\",\\s*?defaultValue: \"(.*?)\"\\)",
		GetMatchReplacer: matchReplacer.GetSwiftMatchReplacer,
	}
)
