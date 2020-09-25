package models

import (
	"box/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aeternity/aepp-sdk-go/account"
	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/binary"
	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/aeternity/aepp-sdk-go/swagguard/node/models"
	"github.com/aeternity/aepp-sdk-go/transactions"
	rlp "github.com/randomshinichi/rlpae"
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"math/big"
	"strconv"
)

//var nodeURL = "https://mainnet.aeternal.io"
//var NodeURL = "http://node.aechina.io:3013"
var NodeURL = "http://localhost:3013"
var NodeURLD = "http://localhost:3113"
var CompilerURL = "http://localhost:3080"
var ContractABCAddress = "ct_chmLQ9hswEXJdvfw9f6PQF3M44MYhBfmq6jAGib36kJ9p2Lj7"
var ContractBoxAddress = "ct_dm2gdE177Ctg1xCB39x9sDBTTYMVDWPbQCRUb1fnfQwekKQUh"

//var nodeURL = nodeURL
//根据助记词返回用户
func MnemonicAccount(mnemonic string) (*account.Account, error) {

	seed, err := account.ParseMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	_, err = bip39.EntropyFromMnemonic(mnemonic)

	if err != nil {
		return nil, err
	}
	// Derive the subaccount m/44'/457'/3'/0'/1'
	key, err := account.DerivePathFromSeed(seed, 0, 0)
	if err != nil {
		return nil, err
	}

	// Deriving the aeternity Account from a BIP32 Key is a destructive process
	alice, err := account.BIP32KeyToAeKey(key)
	if err != nil {
		return nil, err
	}
	return alice, nil
}

//根据私钥返回用户
func SigningKeyHexStringAccount(signingKey string) (*account.Account, error) {
	acc, e := account.FromHexString(signingKey)
	return acc, e
}

//随机创建用户
func CreateAccount() (*account.Account, string) {
	mnemonic, signingKey, _ := CreateAccountUtils()
	acc, _ := account.FromHexString(signingKey)
	return acc, mnemonic
}

//随机创建用户
func CreateAccountUtils() (mnemonic string, signingKey string, address string) {

	//cerate mnemonic
	entropy, _ := bip39.NewEntropy(128)
	mne, _ := bip39.NewMnemonic(entropy)

	//mnemonic := "tail disagree oven fit state cube rule test economy claw nice stable"
	seed, _ := account.ParseMnemonic(mne)

	_, _ = bip39.EntropyFromMnemonic(mne)
	// Derive the subaccount m/44'/457'/3'/0'/1'
	key, _ := account.DerivePathFromSeed(seed, 0, 0)

	// Deriving the aeternity Account from a BIP32 Key is a destructive process
	alice, _ := account.BIP32KeyToAeKey(key)
	return mne, alice.SigningKeyToHexString(), alice.Address
}

//返回最新区块高度
func ApiBlocksTop() (height uint64) {
	client := naet.NewNode(NodeURL, false)
	h, _ := client.GetHeight()
	return h
}

//地址信息返回用户信息
func ApiGetAccount(address string) (account *models.Account, e error) {
	client := naet.NewNode(NodeURL, false)
	acc, e := client.GetAccount(address)
	return acc, e
}

//发起转账
func ApiSpend(account *account.Account, recipientId string, amount float64, data string) (*aeternity.TxReceipt, error) {

	accountNet, e := ApiGetAccount(account.Address)
	if e != nil {
		return nil, e
	}
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err == nil {
		if tokens/1000000000000000000 > amount {
			node := naet.NewNode(NodeURL, false)
			//_, _, ttlnoncer := transactions.GenerateTTLNoncer(node)
			ttlnoncer := transactions.NewTTLNoncer(node)

			spendTx, err := transactions.NewSpendTx(account.Address, recipientId, utils.GetRealAebalanceBigInt(amount), []byte(data), ttlnoncer)

			if err != nil {
				return nil, err
			}
			hash, err := aeternity.SignBroadcast(spendTx, account, node, "ae_mainnet")
			return hash, err
		} else {
			return nil, errors.New("tokens number insufficient")
		}
	} else {
		return nil, err
	}
}

//获取Sophia vm 当前编译版本
func ApiVersion() (v string) {
	c := naet.NewCompiler("https://compiler.aepps.com", false)
	v, _ = c.APIVersion()
	return v
}

//返回tx详细信息
func ApiThHash(th string) (tx *models.GenericSignedTx) {
	client := naet.NewNode(NodeURL, false)
	t, _ := client.GetTransactionByHash(th)
	return t
}

//获取Sophia vm 当前编译版本
func CompilerVersion() (v string) {
	c := naet.NewCompiler("https://compiler.aepps.com", false)
	v, _ = c.APIVersion()
	return v
}

//编译Sophiae
func CompileContract() (s string, e error) {

	c := naet.NewCompiler("https://compiler.aepps.com", true)

	expected, _ := ioutil.ReadFile("contract/fungible-token.aes")
	source, e := c.CompileContract(string(expected), config.Compiler.Backend)
	return source, e
}

func CompileContractInit(account *account.Account, name string, number string) (s string, e error) {
	n := naet.NewNode(NodeURL, false)
	c := naet.NewCompiler("https://compiler.aepps.com", true)
	ctx := aeternity.NewContext(account, n)
	ctx.SetCompiler(c)
	contract := aeternity.NewContract(ctx)
	expected, _ := ioutil.ReadFile("contract/fungible-token.aes")
	ctID, _, err := contract.Deploy(string(expected), "init", []string{"\"" + name + "\"", "18", "\"" + name + "\"", "Some(" + number + ")"}, config.CompilerBackendFATE)
	if err != nil {
		return "", err
	}

	_, err = n.GetContractByID(ctID)
	if err != nil {
		return "", err
	}
	return ctID, err

}

type CallInfoResult struct {

	CallInfo CallInfo `json:"call_info"`
	Reason string `json:"reason"`
}
type CallInfo struct {
	ReturnType  string `json:"return_type"`
	ReturnValue string `json:"return_value"`
}

func Is1AE(address string) bool {
	accountNet, err := ApiGetAccount(address)
	if err != nil {
		return false
	}
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err != nil {
		return false
	}
	if tokens/1000000000000000000 < 1 {
		return false
	}
	return true
}

func CallContractFunction(account *account.Account, ctID string, function string, args []string, amount float64) (s interface{}, functionEncode string, e error) {
	n := naet.NewNode(NodeURL, false)
	//c := naet.NewCompiler("https://compiler.aepps.com", false)
	c := naet.NewCompiler(CompilerURL, false)
	ctx := aeternity.NewContext(account, n)
	ctx.SetCompiler(c)

	var callData = function
	println("123",account.Address)
	if v, ok := cacheCallMap["CALL#"+function+"#"+account.Address+"#"+ctID+"#"+fmt.Sprintf("%s", args)]; ok {
		callData = v

	} else {
		var source []byte
		if ctID == ContractBoxAddress{
			source, _ = ioutil.ReadFile("contract/BoxContract.aes")
		}else{
			source, _ = ioutil.ReadFile("contract/AbcContract.aes")
		}

		callData, _ = ctx.Compiler().EncodeCalldata(string(source), function, args, config.CompilerBackendFATE)
		cacheCallMap["CALL#"+function+"#"+account.Address+"#"+ctID+"#"+fmt.Sprintf("%s", args)] = callData
		println("456")
		println(function)
	}

	callTx, err := transactions.NewContractCallTx(ctx.SenderAccount(), ctID, utils.GetRealAebalanceBigInt(amount), config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, ctx.TTLNoncer())
	if err != nil {
		return nil, function, err
	}
	callReceipt, err := ctx.SignBroadcast(callTx, 1)

	if err != nil {
		return nil, function, err
	}

	return callReceipt.Hash, function, err
	//return decodeResult, err
}

var cacheCallMap = make(map[string]string)
var cacheResultlMap = make(map[string]interface{})

func CallStaticContractFunction(address string, ctID string, function string, args []string) (s interface{}, functionEncode string, e error) {
	node := naet.NewNode(NodeURL, false)
	compile := naet.NewCompiler(CompilerURL, false)
	source, _ := ioutil.ReadFile("contract/BoxContract.aes")
	var callData = function
	if v, ok := cacheCallMap[function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args)]; ok {
		callData = v
	} else {
		callData, _ = compile.EncodeCalldata(string(source), function, args, config.CompilerBackendFATE)
		cacheCallMap[function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args)] = callData
	}

	println(callData)
	println(fmt.Sprintf("%s", args))
	callTx, err := transactions.NewContractCallTx(address, ctID, big.NewInt(0), config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, transactions.NewTTLNoncer(node))
	if err != nil {
		println(err.Error())
		return nil, function, err
	}
	w := &bytes.Buffer{}
	err = rlp.Encode(w, callTx)
	if err != nil {
		return nil, function, err
	}
	txStr := binary.Encode(binary.PrefixTransaction, w.Bytes())

	body := "{\"accounts\":[{\"pub_key\":\"" + address + "\",\"amount\":100000000000000000000000000000000000}],\"txs\":[{\"tx\":\"" + txStr + "\"}]}"

	response := utils.PostBody(NodeURLD+"/v2/debug/transactions/dry-run", body, "application/json")
	var tryRun TryRun
	err = json.Unmarshal([]byte(response), &tryRun)
	if err != nil {
		return nil, function, err
	}

	if v, ok := cacheResultlMap[tryRun.Results[0].CallObj.ReturnValue]; ok {
		return v, function, err
	} else {
		decodeResult, err := compile.DecodeCallResult(tryRun.Results[0].CallObj.ReturnType, tryRun.Results[0].CallObj.ReturnValue, function, string(source), config.Compiler.Backend)
		cacheResultlMap[tryRun.Results[0].CallObj.ReturnValue] = decodeResult
		return decodeResult, function, err
	}

}

type TryRun struct {
	Results []Results `json:"results"`
}
type CallObj struct {
	CallerID    string        `json:"caller_id"`
	CallerNonce int           `json:"caller_nonce"`
	ContractID  string        `json:"contract_id"`
	GasPrice    int           `json:"gas_price"`
	GasUsed     int           `json:"gas_used"`
	Height      int           `json:"height"`
	Log         []interface{} `json:"log"`
	ReturnType  string        `json:"return_type"`
	ReturnValue string        `json:"return_value"`
}
type Results struct {
	CallObj CallObj `json:"call_obj"`
	Result  string  `json:"result"`
	Type    string  `json:"type"`
}
