package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3002Primary(t *testing.T){
 data:=[]byte("abc");h:=blake3.New();h.Write(data);a:=h.Sum(nil);b:=blake3.Sum256(data);same:=bytes.Equal(a,b[:]);want:=true;gotv:=same;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3002Boundary(t *testing.T){
 data:=bytes.Repeat([]byte("z"),2048);h:=blake3.New();h.Write(data);a:=h.Sum(nil);b:=blake3.Sum256(data);gotv:=bytes.Equal(a,b[:]);want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
