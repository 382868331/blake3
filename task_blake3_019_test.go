package blake3_test
import("bytes";"io";"reflect";"testing";blake3 "github.com/zeebo/blake3")
var _=bytes.MinRead
var _=io.EOF
func TestTaskBlake3019Primary(t *testing.T){
 data:=[]byte("abc");got:=blake3.Sum512(data);d:=blake3.New();d.Write(data);want:=make([]byte,64);d.Digest().Read(want);gotv:=bytes.Equal(got[:],want);wantv:=true;err:=error(nil);if !reflect.DeepEqual(gotv,wantv) || err!=nil{t.Fatalf("got=%v want=%v err=%v",gotv,wantv,err)}
}
