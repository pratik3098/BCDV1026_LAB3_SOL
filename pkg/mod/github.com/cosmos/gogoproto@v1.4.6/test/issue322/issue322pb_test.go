// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue322/issue322.proto

package issue322

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	github_com_cosmos_gogoproto_jsonpb "github.com/cosmos/gogoproto/jsonpb"
	github_com_cosmos_gogoproto_proto "github.com/cosmos/gogoproto/proto"
	proto "github.com/cosmos/gogoproto/proto"
	go_parser "go/parser"
	math "math"
	math_rand "math/rand"
	testing "testing"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestOneofTestProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedOneofTest(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &OneofTest{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cosmos_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestOneofTestMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedOneofTest(popr, false)
	size := p.Size()
	dAtA := make([]byte, size)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(dAtA)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &OneofTest{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestOneofTestJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedOneofTest(popr, true)
	marshaler := github_com_cosmos_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &OneofTest{}
	err = github_com_cosmos_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestOneofTestProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedOneofTest(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.MarshalTextString(p)
	msg := &OneofTest{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestOneofTestProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedOneofTest(popr, true)
	dAtA := github_com_cosmos_gogoproto_proto.CompactTextString(p)
	msg := &OneofTest{}
	if err := github_com_cosmos_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestOneofTestCompare(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedOneofTest(popr, false)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OneofTest{}
	if err := github_com_cosmos_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if c := p.Compare(msg); c != 0 {
		t.Fatalf("%#v !Compare %#v, since %d", msg, p, c)
	}
	p2 := NewPopulatedOneofTest(popr, false)
	c := p.Compare(p2)
	c2 := p2.Compare(p)
	if c != (-1 * c2) {
		t.Errorf("p.Compare(p2) = %d", c)
		t.Errorf("p2.Compare(p) = %d", c2)
		t.Errorf("p = %#v", p)
		t.Errorf("p2 = %#v", p2)
	}
}
func TestOneofTestGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedOneofTest(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}
func TestOneofTestSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedOneofTest(popr, true)
	size2 := github_com_cosmos_gogoproto_proto.Size(p)
	dAtA, err := github_com_cosmos_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(dAtA) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(dAtA))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_cosmos_gogoproto_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

//These tests are generated by github.com/cosmos/gogoproto/plugin/testgen