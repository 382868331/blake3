package blake3

type TYPEBlake3OwnedCache struct{ values map[string][]byte }
func NEWBlake3OwnedCache()*TYPEBlake3OwnedCache{return &TYPEBlake3OwnedCache{values:map[string][]byte{}}}
func (c *TYPEBlake3OwnedCache) Put(k string,v []byte){c.values[k]=v}
func (c *TYPEBlake3OwnedCache) Get(k string)[]byte{return c.values[k]}
