package search

import "hash/fnv"

type BloomFilter struct {
	bitArr []bool
	hashFn []func([]byte) uint64
}

func NewBloomFilter(bitLength, hashFnCount int) *BloomFilter {
	bf := &BloomFilter{
		bitArr: make([]bool, bitLength),
		hashFn: make([]func([]byte) uint64, hashFnCount),
	}

	for i := 0; i < hashFnCount; i++ {
		bf.hashFn[i] = bf.getHashFunc(i)
	}

	return bf
}

func (bf *BloomFilter) Add(data string) {
	for _, hashFn := range bf.hashFn {
		index := hashFn([]byte(data)) % uint64(len(bf.bitArr))
		bf.bitArr[index] = true
	}
}

func (bf *BloomFilter) Contains(data string) bool {
	for _, hashFn := range bf.hashFn {
		index := hashFn([]byte(data)) % uint64(len(bf.bitArr))
		if !bf.bitArr[index] {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) getHashFunc(i int) func([]byte) uint64 {
	return func(data []byte) uint64 {
		hashFn := fnv.New64a()
		_, err := hashFn.Write(data)
		if err != nil {
			panic("Fail to generate hash for data: " + string(data))
		}
		_, err = hashFn.Write([]byte{byte(i)})
		if err != nil {
			panic("Fail to generate hash for data: " + string(data))
		}
		return hashFn.Sum64()
	}
}
