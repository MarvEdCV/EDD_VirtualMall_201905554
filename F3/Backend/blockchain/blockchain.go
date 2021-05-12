package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//"bytes"
//"crypto/sha256"

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) AsociarHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Esto unirá la información relevante de nuestro bloque anterior con los nuevos bloques
	hash := sha256.Sum256(info)
	//Esto realiza el hash actual
	b.Hash = hash[:]
	//Si esto (^) no tiene sentido, se puede buscar los valores predeterminados de los segmentos
}

func CrearBloque(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash} //crea un block
	block.AsociarHash()
	return block
}

func (chain *BlockChain) AgregarBloque(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CrearBloque(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Inicio() *Block {
	return CrearBloque("BloqueInicial", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Inicio()}}
}

func (b *BlockChain) Imprimir(chain *BlockChain) {
	for _, lock := range chain.blocks {
		fmt.Printf("Hash anterior: %x\n", lock.PrevHash)
		fmt.Printf("data: %s\n", lock.Data)
		fmt.Printf("hash: %x\n", lock.Hash)
		fmt.Println("")
	}
}
