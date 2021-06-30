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
	"github.com/beego/cache"
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

var NodeUrl = "https://node.aeasy.io"
var NodeUrlDebug = "https://debug.aeasy.io"
var CompilerUrl = "https://compiler.aeasy.io"
//var CompilerUrl = "https://compiler.aeasy.io"

var TESTTNodeUrl = "https://testnet.aeternity.io"
var TESTUrlDebug = "https://testnet.aeternity.io"


var LastHeight = 0
var LockAccountSize = 0
var ConsumingTime int64 = 0
var IsCheckIng bool = false

//var NodeUrl = "https://testnet.aeternity.io"
//var NodeUrlDebug = "https://testnet.aeternity.io"
//var CompilerUrl = "https://compiler.aeasy.io"
//ct_VxetjnAkrWpCHkqkGJuda8W5Ni6ireEXPHJpACv82gLWySp5e
var ABCLockContractV3 = "ct_y7gojSY8rXW6tztE9Ftqe3kmNrqEXsREiPwGCeG3MJL38jkFo"
var BoxSwapContractV2 = "ct_2meHkLcAoZPrQj7P5WjFyJJRLJqRtv43z1QEbpcS1gHs9W8Q3g"
var OraclesContractV1 = "ct_22mvCVphg3ipN856sANq27zDkFt4tAUzeCB1w8PLrM8xoBNGvM"

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
	client := naet.NewNode(NodeUrl, false)
	h, _ := client.GetHeight()
	return h
}

//地址信息返回用户信息
func ApiGetAccount(address string) (account *models.Account, e error) {
	client := naet.NewNode(NodeUrl, false)
	acc, e := client.GetAccount(address)
	return acc, e
}

//发起转账
func ApiSpend(account *account.Account, recipientId string, amount float64, data string) (*aeternity.TxReceipt, error) {
	//获取账户
	accountNet, e := ApiGetAccount(account.Address)
	if e != nil {
		return nil, e
	}
	//格式化账户的tokens
	tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
	if err == nil {

		//判断账户余额是否大于要转账的余额
		if tokens/1000000000000000000 >= amount {
			//获取节点信息
			node := naet.NewNode(NodeUrl, false)
			//生成ttl
			ttler := transactions.CreateTTLer(node)
			noncer := transactions.CreateNoncer(node)

			ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)
			//生成转账tx
			spendTx, err := transactions.NewSpendTx(account.Address, recipientId, utils.GetRealAebalanceBigInt(amount), []byte(data), ttlNoncer)

			if err != nil {
				return nil, err
			}
			//广播转账信息
			hash, err := aeternity.SignBroadcast(spendTx, account, node, "ae_mainnet")

			//err = aeternity.WaitSynchronous(hash, config.Client.WaitBlocks, node)

			if err != nil {
				return nil, err
			}
			return hash, err
		} else {
			return nil, errors.New("tokens number insufficient")
		}
	} else {
		return nil, err
	}
}

type CallInfoResult struct {
	CallInfo CallInfo `json:"call_info"`
	Reason   string   `json:"reason"`
}
type CallInfo struct {
	ReturnType  string `json:"return_type"`
	ReturnValue string `json:"return_value"`
}


//正常调用合约
func CallContractFunction(account *account.Account, ctID string, function string, args []string) (s interface{}, e error) {
	//获取节点信息
	n := naet.NewNode(NodeUrl, false)
	//获取编译器信息
	c := naet.NewCompiler(CompilerUrl, false)
	//创建上下文
	ctx := aeternity.NewContext(account, n)
	//关联编译器
	ctx.SetCompiler(c)
	//创建合约
	contract := aeternity.NewContract(ctx)
	var source []byte
	if ctID == ABCLockContractV3 {
		source, _ = ioutil.ReadFile("contract/ABCLockContractV3.aes")
	} else {
		source, _ = ioutil.ReadFile("contract/AEX9Contract.aes")
	}
	//调用合约代码
	callReceipt, err := contract.Call(ctID, string(source), function, args)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(NodeUrl + "/v2/transactions/" + callReceipt.Hash + "/info")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//获取合约调用信息
	//response := Get(NodeURL + "/v2/transactions/" + callReceipt.Hash + "/info")
	//解析jSON
	var callInfoResult CallInfoResult
	err = json.Unmarshal(body, &callInfoResult)
	if err != nil {
		return nil, err
	}
	//解析结果
	decodeResult, err := c.DecodeCallResult(callInfoResult.CallInfo.ReturnType, callInfoResult.CallInfo.ReturnValue, function, string(source))
	if err != nil {
		return nil, err
	}
	//返回结果
	return decodeResult, err
}

//存放解析结果的缓存
//var cacheResultMap = make(map[string]interface{})

var callCache, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"1","EmbedExpiry":"12000"}`)

//获取合约数据try-run
func CallStaticContractFunction(address string, ctID string, function string, args []string) (s interface{}, functionEncode string, e error) {

	var nodeUrl =""
	var compilerUrl =""
	//var nodeUrlDebug =""
	if ctID != OraclesContractV1{
		nodeUrl = NodeUrl
		compilerUrl = CompilerUrl
		//nodeUrlDebug = NodeUrlDebug
	}else{
		//nodeUrl = TESTTNodeUrl
		//compilerUrl = CompilerUrl
		//nodeUrlDebug = TESTUrlDebug
		nodeUrl = NodeUrl
		compilerUrl = CompilerUrl
		//nodeUrlDebug = NodeUrlDebug
	}

	node := naet.NewNode(nodeUrl, false)
	compile := naet.NewCompiler(compilerUrl, false)
	var source []byte
	if ctID == ABCLockContractV3 {
		source, _ = ioutil.ReadFile("contract/ABCLockContractV3.aes")
	}else if ctID == BoxSwapContractV2 {
		source, _ = ioutil.ReadFile("contract/BoxSwapContractV2.aes")
	}else if ctID == OraclesContractV1 {
		source, _ = ioutil.ReadFile("contract/OraclesLottery.aes")
	} else {
		source, _ = ioutil.ReadFile("contract/AEX9Contract.aes")
	}

	var callData = ""


	if callCache.IsExist(utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))) {
		callData = callCache.Get(utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))).(string)
	} else {
		data, err := compile.EncodeCalldata(string(source), function, args)
		if err != nil {
			return nil, function, err
		}
		callData = data
		_ = callCache.Put(utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args)), callData, 1000*time.Hour)
	}


	callTx, err := transactions.NewContractCallTx(address, ctID, big.NewInt(0),  big.NewInt(10000000000),config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, transactions.NewTTLNoncer(node))
	if err != nil {

		return nil, function, err
	}
	w := &bytes.Buffer{}
	err = callTx.EncodeRLP(w)
	if err != nil {

		println(callTx.CallData)
		return nil, function, err
	}

	txStr := binary.Encode(binary.PrefixTransaction, w.Bytes())

	body := "{\"accounts\":[{\"pub_key\":\"" + address + "\",\"amount\":100000000000000000000000000000000000}],\"txs\":[{\"tx\":\"" + txStr + "\"}]}"


	//response := utils.PostBody(nodeUrlDebug+"/v2/debug/transactions/dry-run", body, "application/json")
	response := utils.PostBody(NodeUrl+"/v3/dry-run", body, "application/json")
	var tryRun TryRun
	err = json.Unmarshal([]byte(response), &tryRun)
	if err != nil {

		return nil, function, err
	}

	//
	//if v, ok := cacheResultMap[utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))+"#"+tryRun.Results[0].CallObj.ReturnValue]; ok {
	//	return v, function, err
	//} else {
	//	decodeResult, err := compile.DecodeCallResult(tryRun.Results[0].CallObj.ReturnType, tryRun.Results[0].CallObj.ReturnValue, function, string(source))
	//	cacheResultMap[utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))+"#"+tryRun.Results[0].CallObj.ReturnValue] = decodeResult
	//	return decodeResult, function, err
	//}

	decodeResult, err := compile.DecodeCallResult(tryRun.Results[0].CallObj.ReturnType, tryRun.Results[0].CallObj.ReturnValue, function, string(source))
	//cacheResultMap[utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))+"#"+tryRun.Results[0].CallObj.ReturnValue] = decodeResult
	return decodeResult, function, err
}

var tokenCache, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"1","EmbedExpiry":"12000"}`)

//获取代币余额调用
func TokenBalanceFunction(address string, ctID string, t string, function string, args []string) (s interface{}, functionEncode string, e error) {
	node := naet.NewNode(NodeUrl, false)
	compile := naet.NewCompiler(CompilerUrl, false)
	var source []byte
	if t == "full" {
		source, _ = ioutil.ReadFile("contract/AEX9Contract.aes")
	} else if t == "basic" {
		source, _ = ioutil.ReadFile("contract/AEX9BasicContract.aes")
	}

	var callData = ""

	if tokenCache.IsExist(utils.Md5V(function + "#" + address + "#" + ctID + "#" + fmt.Sprintf("%s", args))) {
		callData = tokenCache.Get(utils.Md5V(function + "#" + address + "#" + ctID + "#" + fmt.Sprintf("%s", args))).(string)
	} else {
		data, err := compile.EncodeCalldata(string(source), function, args)
		if err != nil {
			return nil, function, err
		}
		callData = data
		_ = tokenCache.Put(utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args)), callData, 1000*time.Hour)
	}

	callTx, err := transactions.NewContractCallTx(address, ctID, big.NewInt(0), config.Client.Contracts.GasLimit, config.Client.GasPrice, config.Client.Contracts.ABIVersion, callData, transactions.NewTTLNoncer(node))
	if err != nil {
		return nil, function, err
	}

	w := &bytes.Buffer{}
	err = callTx.EncodeRLP(w)
	if err != nil {
		println(callTx.CallData)
		return nil, function, err
	}

	txStr := binary.Encode(binary.PrefixTransaction, w.Bytes())

	body := "{\"accounts\":[{\"pub_key\":\"" + address + "\",\"amount\":100000000000000000000000000000000000}],\"txs\":[{\"tx\":\"" + txStr + "\"}]}"

	response := utils.PostBody(NodeUrl+"/v3/dry-run", body, "application/json")
	var tryRun TryRun
	err = json.Unmarshal([]byte(response), &tryRun)
	if err != nil {
		return nil, function, err
	}

	//if v, ok := cacheResultMap[utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))+"#"+tryRun.Results[0].CallObj.ReturnValue]; ok {
	//	return v, function, err
	//} else {
	//	decodeResult, err := compile.DecodeCallResult(tryRun.Results[0].CallObj.ReturnType, tryRun.Results[0].CallObj.ReturnValue, function, string(source))
	//	cacheResultMap[utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))+"#"+tryRun.Results[0].CallObj.ReturnValue] = decodeResult
	//	return decodeResult, function, err
	//}
	decodeResult, err := compile.DecodeCallResult(tryRun.Results[0].CallObj.ReturnType, tryRun.Results[0].CallObj.ReturnValue, function, string(source))
	//cacheResultMap[utils.Md5V(function+"#"+address+"#"+ctID+"#"+fmt.Sprintf("%s", args))+"#"+tryRun.Results[0].CallObj.ReturnValue] = decodeResult
	return decodeResult, function, err

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
