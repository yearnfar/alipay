package alipay

// 生成sign
func makeSign(params map[string]string, key string) string {
	var keys []string
	var sorted []string

	for k, v := range params {
		if k != "sign" && v != "" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s=%s", k, params[k]))
	}

	str := string.Join(sorted, "&")
	str += key

	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}
