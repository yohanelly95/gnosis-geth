package params

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

//go:embed chainspecs
var chainspecs embed.FS

func readChainSpec(filename string) *ChainConfig {
	f, err := chainspecs.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not open chainspec for %s: %v", filename, err))
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	spec := &ChainConfig{}
	err = decoder.Decode(&spec)
	if err != nil {
		panic(fmt.Sprintf("Could not parse chainspec for %s: %v", filename, err))
	}
	return spec
}

var (
	GnosisGenesisHash    = common.HexToHash("0x4f1dd23188aab3a76b463e4af801b52b1248ef073c648cbdc4c9333d3da79756")
	ChiadoGenesisHash    = common.HexToHash("0xada44fd8d2ecab8b08f256af07ad3e777f17fb434f8f8e678b312f576212ba9a")
	MetalayerGenesisHash = common.HexToHash("0x56d44e60411eae7f4f4544f5c1cb70178dfb37e0434f6573e5cd138101d906ab")

	GnosisGenesisStateRoot    = common.HexToHash("0x40cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133b")
	ChiadoGenesisStateRoot    = common.HexToHash("0x9ec3eaf4e6188dfbdd6ade76eaa88289b57c63c9a2cde8d35291d5a29e143d31")
	MetalayerGenesisStateRoot = common.HexToHash("0xda8db002b6a9dade245b8269545a86b51f96cee2579b8fd394adf18c2ca7ec99")

	GnosisChainConfig    = readChainSpec("chainspecs/gnosis.json")
	ChiadoChainConfig    = readChainSpec("chainspecs/chiado.json")
	MetalayerChainConfig = readChainSpec("chainspecs/metalayer.json")

	GnosisForkBlock = uint64(25349536)
)
