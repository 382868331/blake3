package blake3

type TYPEBlake3OwnedCache struct{ values map[string][]byte }
func NEWBlake3OwnedCache()*TYPEBlake3OwnedCache{return &TYPEBlake3OwnedCache{values:map[string][]byte{}}}
func (c *TYPEBlake3OwnedCache) Put(k string,v []byte){c.values[k]=append([]byte(nil),v...)}
func (c *TYPEBlake3OwnedCache) Get(k string)[]byte{
	v,ok:=c.values[k];if !ok{return nil}
	return append([]byte(nil),v...)
}
