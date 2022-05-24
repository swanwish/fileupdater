package rules

import (
	"github.com/swanwish/fileupdater/common"
)

var matchReplacer = i18nMatchReplacer{}

var (
	I18nJavaUpdateRule = UpdateRule{
		Exts:             []string{common.ExtensionJava},
		Pattern:          "StringProvider\\.getValue\\(\"(.*?)\",\\s*?\"(.*?)\"\\)",
		GetMatchReplacer: matchReplacer.GetJavaMatchReplacer,
	}
	I18nKotlinUpdateRule = UpdateRule{
		Exts:             []string{common.ExtensionKotlin},
		Pattern:          "StringProvider\\.getValue\\(\"(.*?)\",\\s*?\"(.*?)\"\\)",
		GetMatchReplacer: matchReplacer.GetKotlinMatchReplacer,
	}
	I18nObjcUpdateRule = UpdateRule{
		Exts:             []string{common.ExtensionObjc},
		Pattern:          "\\[\\[StringProvider sharedInstance\\]\\s*?getValue:@\"(.*?)\"\\s*?defaultValue:@\"(.*?)\"\\]",
		GetMatchReplacer: matchReplacer.GetObjcMatchReplacer,
	}
	I18nSwiftUpdateReule = UpdateRule{
		Exts:             []string{common.ExtensionSwift},
		Pattern:          "StringProvider\\.sharedInstance\\(\\)\\.getValue\\(\"(.*?)\",\\s*?defaultValue: \"(.*?)\"\\)",
		GetMatchReplacer: matchReplacer.GetSwiftMatchReplacer,
	}
)
