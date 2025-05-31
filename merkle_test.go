package merkletree

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	content := []string{"a", "b", "c", "d", "e", "f", "g", "k", "sdf", "cxv"}
	var list []Content
	for _, v := range content {
		list = append(list, Content{X: crypto.Keccak256Hash([]byte(v)).Bytes()})
	}
	tr := New(KECCAK256)
	tr.WithOption(Option{
		Sort:              true,
		DisableLeavesHash: true,
	})
	tr.InitLeaves(list)
	log.Println(hex.EncodeToString(tr.GetRoot()))
	proofs := tr.GetProof(list[0])
	for _, proof := range proofs {
		log.Println(hex.EncodeToString(proof.Data))
	}
}

//func TestMerkleTree_WithOptions(t *testing.T) {
//
//	var list []Content
//	list = append(list, Content{X: "0x0000007adc3815c056c91dbc5ec48c4fd0950b49"})
//	list = append(list, Content{X: "0x00010a9bbde3f6ec395c179e36f1df24e59bd824"})
//	list = append(list, Content{X: "0x0004932ff86fbdb0e7cf99203111624ea57cd772"})
//	list = append(list, Content{X: "0x01291849cc904161603c1c60ff86658ad7eadb8d"})
//	list = append(list, Content{X: "0x0582f324704bf44f7e9e80a2ca68f08e91e45927"})
//	/*list = append(list, Content{x: "a"})
//	list = append(list, Content{x: "b"})
//	list = append(list, Content{x: "c"})
//	list = append(list, Content{x: "d"})*/
//
//	tr := New(SHA256)
//	tr.WithOption(Option{
//		Sort: true,
//	})
//	tr.InitLeaves(list)
//
//	//tr.AddLeaf(Content{x: "e"})
//	root := tr.GetRoot()
//	leaf := Content{X: "0x01291849cc904161603c1c60ff86658ad7eadb8d"}
//
//	root, _ = hex.DecodeString("5651b0b746e49be1307faa82cdef9a6cbd454d9b0df2318cd795e8adb877c594")
//
//	proof := tr.GetProof(leaf)
//	fmt.Println(tr.String())
//
//	t.Fatal(tr.Verify(proof, leaf, root)) // true
//
//	//t.Fatal(tr.GetHexLeaves())
//	//t.Fatal(tr.GetHexLayers())
//}
