package blake3_test
import("bytes";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
func TestTaskBlake3008Primary(t *testing.T){
 a:=blake3.New();b:=blake3.New();a.WriteString("hello");b.Write([]byte("hello"));gotv:=bytes.Equal(a.Sum(nil),b.Sum(nil));want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}

func TestTaskBlake3008Boundary(t *testing.T){
 a:=blake3.New();b:=blake3.New();a.WriteString("\u4e16\u754c");b.Write([]byte("\u4e16\u754c"));gotv:=bytes.Equal(a.Sum(nil),b.Sum(nil));want:=true;err:=error(nil);got:=gotv;if !reflect.DeepEqual(got,want)||err!=nil{t.Fatalf("got=%v want=%v err=%v",got,want,err)}
}
