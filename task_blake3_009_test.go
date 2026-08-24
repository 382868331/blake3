package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3009Primary(t *testing.T){
 h:=blake3.New();h.Write([]byte("old"));h.Reset();gotv:=bytes.Equal(h.Sum(nil),blake3.New().Sum(nil));want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3009Boundary(t *testing.T){
 h:=blake3.New();h.Write([]byte("a"));h.Reset();h.Write([]byte("b"));fresh:=blake3.New();fresh.Write([]byte("b"));gotv:=bytes.Equal(h.Sum(nil),fresh.Sum(nil));want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
