package miscutils

import "strings"

func CaseInsensitiveKeyMapJoin[value any](
	primary, addOn map[string]value,
) map[string]value {
	if primary == nil && addOn == nil {
		return nil
	}

	result := map[string]value{}
	for k, v := range primary {
		result[k] = v
	}
	for k, v := range addOn {
		alreadyExists := false
		for pk := range primary {
			if strings.EqualFold(pk, k) {
				alreadyExists = true
				break
			}
		}
		if alreadyExists {
			continue
		}
		result[k] = v
	}

	return result
}
