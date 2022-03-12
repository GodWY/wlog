package filter

type Filter struct {
	GaKeys  []string
	N9eKeys []string
}

// WithGaFilter 添加GA过滤字段
func (f *Filter) WithGaFilter(key string) *Filter {
	f.GaKeys = append(f.GaKeys, key)
	return f
}

// WithN9eFilter 过滤n9e字段
func (f *Filter) WithN9eFilter(key string) *Filter {
	f.N9eKeys = append(f.N9eKeys, key)
	return f
}