package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3013Primary(t *testing.T){
 h:=blake3.New();h.Write([]byte("abc"));prefix:=[]byte("pre");result:=h.Sum(prefix);wantPrefix:=[]byte("pre");gotv:=bytes.Equal(result[:3],wantPrefix)&&len(result)==35;want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3013Boundary(t *testing.T){
 h:=blake3.New();b:=make([]byte,2,2);copy(b,"xy");result:=h.Sum(b);gotv:=string(result[:2])=="xy"&&len(result)==34;want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want) || err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
