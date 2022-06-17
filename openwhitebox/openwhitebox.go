package main

import (
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/OpenWhiteBox/primitives/encoding"
	"github.com/OpenWhiteBox/primitives/matrix"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/OpenWhiteBox/AES/constructions/full"
)

const keySize = 16 * (1 + 128 + 1 + 128 + 1)

func main() {
	type info struct {
		name string
		id   int
	}
	v := info{"Nan", 33}
	fmt.Printf("%v\n", v)
	fmt.Printf("%+v", v)
	//key := "0123456789abcdeffedcba9876543210"
	text := "000000000000000000000000deadbeef"
	//生成一个key
	//GenerateKey(key)
	//用这个key加密一串字符串
	encrypt(text)
	//decrypt(text)
}

func GenerateKey(hexKey string) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Println(err)
		flag.PrintDefaults()
		return
	} else if len(key) != 16 {
		log.Println("Key must be 128 bits.")
		flag.PrintDefaults()
		return
	}

	// GenerateKey is deterministic, so we need to sample a small amount of
	// randomness to get a random white-box construction.
	seed := make([]byte, 16)
	rand.Read(seed)

	// This generates the white-box construction. inputMask and outputMask are
	// the random affine transformations on the input and output of constr.
	constr, inputMask, outputMask := full.GenerateKeys(key, seed)

	// Write the public white-box to disk.
	ioutil.WriteFile("./constr.txt", constr.Serialize(), 0777)

	// Write the private input and output mask to disk.
	buff := make([]byte, 0)
	buff = append(buff, key...)

	for _, row := range inputMask.Forwards {
		buff = append(buff, row...)
	}
	buff = append(buff, inputMask.BlockAdditive[:]...)

	for _, row := range outputMask.Forwards {
		buff = append(buff, row...)
	}
	buff = append(buff, outputMask.BlockAdditive[:]...)

	ioutil.WriteFile("./constr.key", buff, 0777)
}

func encrypt(hexBlock string) {
	block, err := hex.DecodeString(hexBlock)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		return
	} else if len(block) != 16 {
		log.Println("Block must be 128 bits.")
		return
	}

	// Read construction from disk and parse it into something usable.
	data, err := ioutil.ReadFile("./constr.txt")
	if err != nil {
		log.Fatal(err)
	}
	constr, err := full.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt block in-place, and print as hex.
	constr.Encrypt(block, block)
	fmt.Printf("%x\n", block)
}

func decrypt(hexBlock string) {
	block, err := hex.DecodeString(hexBlock)
	if err != nil {
		log.Println(err)
		flag.PrintDefaults()
		return
	} else if len(block) != 16 {
		log.Println("Block must be 128 bits.")
		flag.PrintDefaults()
		return
	}

	// Read key from disk and parse it.
	data, err := ioutil.ReadFile("./constr.key")
	if err != nil {
		log.Fatal(err)
	} else if len(data) != keySize {
		log.Fatalf("key wrong size: %v (should be %v)", len(data), keySize)
	}

	var key []byte
	inputLinear, outputLinear := matrix.Matrix{}, matrix.Matrix{}
	inputConst, outputConst := [16]byte{}, [16]byte{}

	key, data = data[:16], data[16:]
	for i := 0; i < 128; i++ {
		inputLinear, data = append(inputLinear, data[:16]), data[16:]
	}
	copy(inputConst[:], data)
	data = data[16:]
	for i := 0; i < 128; i++ {
		outputLinear, data = append(outputLinear, data[:16]), data[16:]
	}
	copy(outputConst[:], data)

	inputMask := encoding.NewBlockAffine(inputLinear, inputConst)
	outputMask := encoding.NewBlockAffine(outputLinear, outputConst)

	// Decrypt block and print as hex.
	temp := [16]byte{}
	copy(temp[:], block)

	temp = outputMask.Decode(temp)

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	c.Decrypt(temp[:], temp[:])

	temp = inputMask.Decode(temp)
	fmt.Printf("%x\n", temp)
}
