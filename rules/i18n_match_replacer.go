package rules

import (
	"fmt"
	"strings"
)

type i18nMatchReplacer struct{}

func (r i18nMatchReplacer) GetJavaMatchReplacer(match []string) *string {
	if len(match) == 3 {
		key := match[1]
		if len(key) > 0 {
			if len(key) > 1 {
				key = fmt.Sprintf("%s%s", strings.ToUpper(key[:1]), key[1:])
			} else {
				key = strings.ToUpper(key)
			}
		} else {
			return nil
		}
		value := fmt.Sprintf("StringValues.INSTANCE.get%s()", key)
		return &value
	}
	return nil
}

func (r i18nMatchReplacer) GetKotlinMatchReplacer(match []string) *string {
	if len(match) == 3 {
		value := fmt.Sprintf("StringValues.%s", match[1])
		return &value
	}
	return nil
}

func (r i18nMatchReplacer) GetObjcMatchReplacer(match []string) *string {
	if len(match) == 3 {
		value := fmt.Sprintf("StringValues.shared.%s", match[1])
		return &value
	}
	return nil
}

func (r i18nMatchReplacer) GetSwiftMatchReplacer(match []string) *string {
	if len(match) == 3 {
		value := fmt.Sprintf("StringValues.shared.%s", match[1])
		return &value
	}
	return nil
}
