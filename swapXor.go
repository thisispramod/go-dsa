package main

// Binary: 0101 (5) ^  1001 (9) = 1100 (which is 12)

// The XOR ^  operator compares bits. It outputs 1 if the bits are different, and 0 if the bits are the same.\(0 \oplus 0 = 0\)

func SwapXor(a int, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
