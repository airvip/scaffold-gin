package def

var CommonCache = map[string]string{
	"user_detail" : "user:detail_%s",
}

func GetCacheString(key string) string {
	if val,ok := CommonCache[key]; ok {
      return val;
	}
	panic("cache string not exist!!!")
}





