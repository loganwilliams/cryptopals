package main

import "testing"

// Problem 1.1
func TestHexToBase64(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", 
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}
	for _, c := range cases {
		got := HexToBase64(c.in)
		if got != c.want {
			t.Errorf("HexToBase64(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// Problem 1.2
func TestXorHex(t *testing.T) {
	cases := []struct {
		in1, in2, want string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}
	for _, c := range cases {
		got := XorHex(c.in1, c.in2)
		if got != c.want {
			t.Errorf("TestXorHex(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}

// Problem 1.5
func TestRepeatingKeyXor(t *testing.T) {
	cases := []struct {
		in, key, want string
	}{
		{"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE", "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"},
	}

	for _, c:= range cases {
		got := RepeatingKeyXor(c.key, c.in)
		if got != c.want {
			t.Errorf("TestRepeatingKeyXor(%q, %q) == %q, want %q", c.key, c.in, got, c.want)
		}
	}
}

// Problem 1.6
func TestHammingDistance(t *testing.T) {
	cases := []struct {
		in1, in2 []byte;
		want int
	}{
		{[]byte("this is a test"), []byte("wokka wokka!!!"), 37},
	}
	for _, c:= range cases {
		got := HammingDistance(c.in1, c.in2)
		if got != c.want {
			t.Errorf("HammingDistance(%q, %q) == %v, want %v", c.in1, c.in2, got, c.want)
		}
	}
}