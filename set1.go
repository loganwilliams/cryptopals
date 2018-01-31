package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"math"
	"io/ioutil"
	"github.com/spacemonkeygo/openssl"
)

// SET 1 PROBLEM 1
// Convert a hexadecimal string to base64
func HexToBase64(msg string) (string) {
	bytes, err := hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	b64encoded := base64.StdEncoding.EncodeToString([]byte(bytes))
	
	return b64encoded
}

// SET 1 PROBLEM 2
// computer the XOR of two lists of bytes
func Xor(a, b []byte) ([]byte) {
	if len(a) != len(b) {
		panic("Lists to XOR are not the same length!")
	}

	o := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		o[i] = a[i] ^ b[i]
	}

	return o
}

func XorHex(a, b string) (string) {
	// convert hex string to byte buffers
	a_bytes, err := hex.DecodeString(a)
	if err != nil {
		panic(err)
	}

	b_bytes, err := hex.DecodeString(b)
	if err != nil {
		panic(err)
	}

	// computer XOR and re-encode to hex
	x_bytes := Xor(a_bytes, b_bytes)
	x := hex.EncodeToString(x_bytes)
	return x
}

// SET 1 PROBLEM 3
func ScoreEnglish(text string) (float64) {
	englishFrequency := map[string]float64{
		"A": .0834/13.0,
		"B": .0154/13.0,
		"C": .0273/13.0,
		"D": .0414/13.0,
		"E": .1260/13.0,
		"F": .0203/13.0,
		"G": .0192/13.0,
		"H": .0611/13.0,
		"I": .0671/13.0,
		"J": .0023/13.0,
		"K": .0087/13.0,
		"L": .0424/13.0,
		"M": .0253/13.0,
		"N": .0680/13.0,
		"O": .0770/13.0,
		"P": .0166/13.0,
		"Q": .0009/13.0,
		"R": .0568/13.0,
		"S": .0611/13.0,
		"T": .0937/13.0,
		"U": .0285/13.0,
		"V": .0106/13.0,
		"W": .0234/13.0,
		"X": .0020/13.0,
		"Y": .0204/13.0,
		"Z": .0006/13.0,
		"a": .0834/1.3,
		"b": .0154/1.3,
		"c": .0273/1.3,
		"d": .0414/1.3,
		"e": .1260/1.3,
		"f": .0203/1.3,
		"g": .0192/1.3,
		"h": .0611/1.3,
		"i": .0671/1.3,
		"j": .0023/1.3,
		"k": .0087/1.3,
		"l": .0424/1.3,
		"m": .0253/1.3,
		"n": .0680/1.3,
		"o": .0770/1.3,
		"p": .0166/1.3,
		"q": .0009/1.3,
		"r": .0568/1.3,
		"s": .0611/1.3,
		"t": .0937/1.3,
		"u": .0285/1.3,
		"v": .0106/1.3,
		"w": .0234/1.3,
		"x": .0020/1.3,
		"y": .0204/1.3,
		"z": .0006/1.3,
		" ": .15,
	}

	sequenceFrequency := make(map[string]float64)

	for i := 0; i < len(text); i++ {
		_, ok := sequenceFrequency[strings.ToUpper(string(text[i]))] 

		if ok {
			sequenceFrequency[string(text[i])] += 1.0/float64(len(text))
		} else {
			sequenceFrequency[string(text[i])] = 1.0/float64(len(text))
		}
	}

	err := 0.0

	for k := range englishFrequency {
		err += math.Abs(englishFrequency[k] - sequenceFrequency[k])
	}

	for k := range sequenceFrequency {
		_, ok := englishFrequency[k]
		if !ok {
			err += sequenceFrequency[k] * 5
		}
	}

	return err
}

func FindXor(msg string) (string, byte, float64) {
	msgBytes, err := hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}

	minScore := 10.0
	var correctText string
	var correctCipher byte

	for i := 0; i < 256; i++ {
		cipher := make([]byte, len(msgBytes))
		for j := 0; j < len(msgBytes); j++ {
			cipher[j] = byte(i)
		}

		decoded := Xor(msgBytes, cipher)
		score := ScoreEnglish(string(decoded))
		if score < minScore {
			minScore = score
			correctText = string(decoded)
			correctCipher = cipher[0]
		}

	}

	return correctText, correctCipher, minScore
}

// SET 1 PROBLEM 5
func RepeatingKeyXor(key, msg string) (string) {
	keyBytes := []byte(key)
	msgBytes := []byte(msg)
	repeatedKey := make([]byte, len(msgBytes))
	for i := 0; i < len(msgBytes); i++ {
		repeatedKey[i] = keyBytes[i % len(keyBytes)]
	}

	xorBytes := Xor(repeatedKey, msgBytes)

	return hex.EncodeToString(xorBytes)
}

// SET 1 PROBLEM 6
func HammingDistance(a, b []byte) (int) {
	c := Xor(a, b)
	sum := 0
	for i:= 0; i < len(c); i++ {
		// count the non-zero bits by incrementally shifting to ther right
		for j:=0; j < 8; j++ {
			sum += int((c[i] >> byte(j)) & byte(1))
		}
	}
	return sum
}

func DetermineKeySize(msg []byte) (int) {
	minDistance := 1000.0
	bestKeysize := 0
	for keysize := 2; keysize <= 40; keysize++ {
		nreps := int(math.Min(math.Floor(float64(len(msg))/float64(keysize)), 10))
		distance := float64(HammingDistance(msg[0:keysize*(nreps-1)], msg[keysize:keysize*nreps]))/float64(keysize)
		if distance < minDistance {
			bestKeysize = keysize
			minDistance = distance
		}
	}

	return bestKeysize
}

func RepetitionScore(msg []byte) (float64) {
	repCount := 0

	// count the number of repeated two byte sequences
	for i := 0; i < len(msg)-1; i++ {
		for j := i+1; j < len(msg)-1; j++ {
			if msg[i] == msg[j] && msg[i+1] == msg[j+1] {
				repCount += 1
			}
		}
	}

	return float64(repCount) / float64(len(msg))
}

func DetermineRepeatingKey(msg []byte, keysize int) ([]byte, []byte) {
	// slice msg into keysize different pieces
	slicedMsg := make([][]byte, keysize)
	decodedMsg := make([]byte, len(msg))
	key := make([]byte, keysize)

	for i := 0; i < keysize; i++ {
		slicedMsg[i] = make([]byte, len(msg)/keysize)
		for j := 0; j < len(slicedMsg[i]); j++ {
			slicedMsg[i][j] = msg[j*keysize + i]
		}

		text, cipher, _ := FindXor(hex.EncodeToString(slicedMsg[i]))
		key[i] = cipher

		for j := 0; j < len(slicedMsg[i]); j++ {
			decodedMsg[j*keysize + i] = text[j]
		}
	}

	return decodedMsg, key
}

func main() {
	text, _, score := FindXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(text, score)

	dat, err := ioutil.ReadFile("testcases/set1/4.txt")
	if err != nil {
		panic(err)
	}
    
    minScore := 10.0
    testSequences := strings.Fields(string(dat))
    var bestText string
    for i := 0; i < len(testSequences); i++ {
    	text, _, score := FindXor(testSequences[170])
    	if score < minScore {
    		minScore = score
    		bestText = text
    	}
    }
    fmt.Println(bestText, minScore)

	dat, err = ioutil.ReadFile("testcases/set1/6.txt")
	if err != nil {
		panic(err)
	}
	dat,_ = base64.StdEncoding.DecodeString(string(dat))
	keysize := DetermineKeySize(dat)
	fmt.Println(keysize)

	decodedMsg, repeatingKey := DetermineRepeatingKey(dat, keysize)
	fmt.Println(string(decodedMsg))
	fmt.Println(string(repeatingKey))

	// CHALLENGE 7

	cipher, err := openssl.GetCipherByName("aes-128-ecb")
	
	ch7, err := ioutil.ReadFile("testcases/set1/7.txt")
	if err != nil {
		panic(err)
	}

	dat,_ = base64.StdEncoding.DecodeString(string(ch7))

	dCtx, err := openssl.NewDecryptionCipherCtx(cipher, nil, []byte("YELLOW SUBMARINE"), nil)

	if err != nil {
		panic(err)
	}

	plaintext, _ :=  dCtx.DecryptUpdate(dat)
	fmt.Println(string(plaintext))

	// CHALLENGE 8

	ch8, _ := ioutil.ReadFile("testcases/set1/8.txt")
	blocks := strings.Split(string(ch8), "\n")

	likelyIndex := 0
	likelyBlock := blocks[0]
	maxScore := 0.0

	for i, block := range blocks {
		if len(block) > 0 {
			dat, _ = hex.DecodeString(block)
			repScore := RepetitionScore(dat)
			if repScore > maxScore {
				likelyIndex = i
				likelyBlock = block
				maxScore = repScore
			}
		}
	}

	fmt.Println(likelyIndex, likelyBlock)


}