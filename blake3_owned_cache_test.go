package blake3

import (
	"testing"

)

func TestBlake3OwnedCache(t *testing.T){ c:=NEWBlake3OwnedCache();in:=[]byte("abc");c.Put("k",in);in[0]='x';if got:=string(c.Get("k"));got!="abc"{t.Fatalf("cached=%q",got)} }

func TestBlake3OwnedCacheReturnsIndependentCopy(t *testing.T){ c:=NEWBlake3OwnedCache();c.Put("k",[]byte("abc"));got:=c.Get("k");got[0]='x';if again:=string(c.Get("k"));again!="abc"{t.Fatalf("cached=%q",again)} }
