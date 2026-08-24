package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3005Primary(t *testing.T){
 a:=make([]byte,32);b:=make([]byte,32);blake3.DeriveKey("ctx",[]byte("one"),a);blake3.DeriveKey("ctx",[]byte("two"),b);gotv:=!bytes.Equal(a,b);want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3005Boundary(t *testing.T){
 a:=make([]byte,64);b:=make([]byte,64);blake3.DeriveKey("ctx",nil,a);blake3.DeriveKey("ctx",[]byte{0},b);gotv:=!bytes.Equal(a,b);want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
