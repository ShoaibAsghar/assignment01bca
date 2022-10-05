package function

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"strconv"
)

func calculateHash(stringToHash string) string {
	// fmt.Printf("string Received: %s \n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

type block struct {
	transaction  string
	nonce        int
	previousHash string
}

func NewBlock(transaction string, nonce int, previousHash string) *block {
	var block1 block
	block1.transaction = transaction
	block1.nonce = nonce
	block1.previousHash = previousHash

	// str1 := strconv.Itoa(block1.nonce)
	// str2 := block1.transaction + str1 //+ block1.previousHash

	// block1.previousHash = calculateHash(str2)
	// fmt.Println("HASH: %x\n ", block1.previousHash)

	return &block1

}

func DisplayBlocks(li1 *list.List) {
	var blocknumber int
	blocknumber++
	var str3 string
	var str4 int
	var str5 string
	var str6 string
	// var str7 string
	for i := li1.Front(); i != nil; i = i.Next() {
		fmt.Println("---------------------------------------------------------")
		fmt.Println("-----------------Block ", blocknumber, " ------------------------------")
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Transaction -> ", i.Value.(*block).transaction)
		str3 = i.Value.(*block).transaction
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Nonce -> ", i.Value.(*block).nonce)
		str4 = i.Value.(*block).nonce
		str5 = strconv.Itoa(str4)
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Prevoius Hash -> ", i.Value.(*block).previousHash)
		// str7 = i.Value.(*block).previousHash
		fmt.Println("---------------------------------------------------------")
		str6 = str3 + str5 //+ str7
		fmt.Println("Current Hash -> ", calculateHash(str6))
		fmt.Println("---------------------------------------------------------")
		fmt.Println("---------------------------------------------------------")

		blocknumber++

	}
}

func ChangeBlock(blockNum int, blockTrans string, blockNonc int, li2 *list.List) {
	var num int
	num = 0
	for i := li2.Front(); i != nil; i = i.Next() {
		if num == blockNum {
			i.Value.(*block).transaction = blockTrans
			i.Value.(*block).nonce = blockNonc
		} else {
			num++
		}

	}

}

func verifyChain(li3 *list.List) {
	var preHash string
	var Hashcurrent string

	for i := li3.Front(); i != nil; i = i.Next() {
		preHash = strconv.Itoa(i.Value.(*block).nonce) + i.Value.(*block).transaction
		Hashcurrent = calculateHash(preHash)
		if i != li3.Front() {
			if i.Value.(*block).previousHash != Hashcurrent {
				fmt.Println("Changes are made in a Blockchain!")
				fmt.Println("Because Current Hash is not match with Prevoius Hash!")
				return
			}
		}

	}

}
