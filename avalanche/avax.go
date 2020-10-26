package avalanche

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/anaskhan96/base58check"
	"github.com/btcsuite/btcutil/base58"
)

//	"github.com/joho/godotenv"

// Client ...
type Client struct {
	NodeURL string
	Port    string
}

/*({
	JSONRPC: (string) (len=3) "2.0",
	Result: (map[string]interface {}) (len=1) {
	(string) (len=6) "nodeID": (string) (len=40) "NodeID-6asQ8tThdWTXCQstmhgibpE1aLF3DX2Kp"
	},
	Error: (*jsonrpc.RPCError)(<nil>),
	ID: (int) 0
})*/

//GetNodeID ...
func (c *Client) GetNodeID() (string, error) {
	fmtp("-----------------== GetNodeID()")
	JSONRPCMethod := "info.getNodeID"
	endpoint := "/ext/info"
	resField := "nodeID"
	params := Params{}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}
	val, ok := (respOut.Result[resField]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField)
	}
	return val, nil
}

//IsBootstrapped ...
func (c *Client) IsBootstrapped() (bool, error) {
	fmtp("-----------------== IsBootstrapped()")
	JSONRPCMethod := "info.isBootstrapped"
	endpoint := "/ext/info"
	resField := "isBootstrapped"
	params := Params{Chain: "X"}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return false, errors.New(respOut.Mesg)
	}
	val, ok := (respOut.Result[resField]).(bool)
	if !ok {
		return false, errors.New("err@ parsing result")
	}
	return val, nil
}

//GetPeers ...
func (c *Client) GetPeers() ([]Peer, error) {
	fmtp("-----------------== GetPeers()")
	JSONRPCMethod := "info.peers"
	endpoint := "/ext/info"
	resField1 := "peers"
	params := Params{}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return nil, errors.New(respOut.Mesg)
	}
	peersRaw, ok := (respOut.Result[resField1]).([]map[string]interface{})
	if !ok {
		return nil, errors.New("could not parse peersRaw")
	}

	peers := make([]Peer, 0)
	for i := range peersRaw {
		fmtp("peer of index", i, peersRaw[i])
		ip, ok := (peersRaw[i]["ip"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing ip")
		}

		publicIP, ok := (peersRaw[i]["publicIP"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing publicIP")
		}

		id, ok := (peersRaw[i]["id"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing id")
		}

		version, ok := (peersRaw[i]["version"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing version")
		}

		lastSent, ok := (peersRaw[i]["lastSent"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing lastSent")
		}

		lastReceived, ok := (peersRaw[i]["lastReceived"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing lastReceived")
		}
		fmtp("ip:", ip, ", publicIP:", publicIP, ", id:", id, ", version:", version, ", lastSent:", lastSent, ", lastReceived:", lastReceived)

		peers = append(peers, Peer{ip, publicIP, id, version, lastSent, lastReceived})
	}

	return peers, nil
}

//ImportKey ...
func (c *Client) ImportKey(chain string, privateKey string, username string, password string) (string, error) {
	fmtp("-----------------== ImportKey()")
	var JSONRPCMethod string
	var endpoint string
	if chain == "X" {
		JSONRPCMethod = "avm.importKey"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return "", errors.New("chain is not valid")
	}
	resField1 := "address"
	params := Params{
		PrivateKey: privateKey,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}
	val, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField1)
	}
	return val, nil
}

//ImportAVAXlocalNework ...
func (c *Client) ImportAVAXlocalNework(username string, password string) error {
	fmtp("-----------------== ImportAVAXlocalNework()")
	JSONRPCMethod := "platform.importKey"
	endpoint := "/ext/platform"
	//resField1 := "xyz"
	params := Params{
		PrivateKey: "PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN",
		Username:   username,
		Password:   password,
	}

	/* curl --location --request POST 'localhost:9650/ext/platform' --header 'Content-Type: application/json' --data-raw
	 */
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return errors.New(respOut.Mesg)
	}
	// peersRaw, ok := (respOut.Result[resField1]).([]map[string]interface{})
	// if !ok {
	// 	return errors.New("could not parse peersRaw")
	// }
	return nil
}

//CreateUser ...
func (c *Client) CreateUser(username string, password string) (bool, error) {
	fmtp("-----------------== CreateUser()")
	JSONRPCMethod := "keystore.createUser"
	endpoint := "/ext/keystore"
	resField := "success"
	params := Params{
		Username: username,
		Password: password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return false, errors.New(respOut.Mesg)
	}
	val, ok := (respOut.Result[resField]).(bool)
	if !ok {
		return false, errors.New("cannot parse result")
	}
	return val, nil
}

//CreateAddress ...
func (c *Client) CreateAddress(chain string, username string, password string) (string, error) {
	fmtp("-----------------== CreateAddress()")
	var JSONRPCMethod string
	var endpoint string
	if chain == "X" {
		JSONRPCMethod = "avm.createAddress"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return "", errors.New("chain is not valid")
	}
	resField := "address"
	params := Params{
		Username: username,
		Password: password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}
	val, ok := (respOut.Result[resField]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField)
	}
	return val, nil
}

//CreateSubnet ...
func (c *Client) CreateSubnet(controlKeys []string, threshold int, username string, password string) (string, string, error) {
	fmtp("-----------------== CreateSubnet()")
	JSONRPCMethod := "platform.createSubnet"
	endpoint := "/ext/P"
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		ControlKeys: controlKeys,
		Threshold:   threshold,
		Username:    username,
		Password:    password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse txID")
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse changeAddr")
	}
	return txID, changeAddrM, nil
}

//GetSubnets ...
func (c *Client) GetSubnets(username string, password string) (string, []string, string, error) {
	fmtp("-----------------== GetSubnets()")
	JSONRPCMethod := "platform.getSubnets"
	endpoint := "/ext/P"
	resField1 := "subnets"
	resField1a := "id"
	resField1b := "controlKeys"
	resField1c := "threshold"

	params := Params{}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", nil, "", errors.New(respOut.Mesg)
	}
	subnets, ok := (respOut.Result[resField1]).([]map[string]interface{})
	if !ok {
		return "", nil, "", errors.New("could not parse subnets")
	}

	id, ok := (subnets[0][resField1a]).(string)
	if !ok {
		return "", nil, "", errors.New("could not parse " + resField1a)
	}
	fmtp("id:", id)

	ctrlKeys, ok := (subnets[0][resField1b]).([]string)
	if !ok {
		return "", nil, "", errors.New("could not parse " + resField1b)
	}
	fmtp("ctrlKeys:", ctrlKeys)

	threshold, ok := (subnets[0][resField1c]).(string)
	if !ok {
		return "", nil, "", errors.New("could not parse " + resField1c)
	}
	fmtp("threshold:", threshold)

	return id, ctrlKeys, threshold, nil
}

//AddValidator ...
func (c *Client) AddValidator(chain string, nodeID string, startTime string, endTime string, stakeAmount int, rewardAddress string, changeAddr string, delegationFeeRate int, username string, password string) (string, string, error) {
	fmtp("-----------------== AddValidator()")
	JSONRPCMethod := "platform.addValidator"
	endpoint := "/ext/P"
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		NodeID:            nodeID,
		StartTime:         startTime,
		EndTime:           endTime,
		StakeAmount:       stakeAmount,
		RewardAddress:     chain + "-" + rewardAddress,
		ChangeAddr:        chain + "-" + changeAddr,
		DelegationFeeRate: delegationFeeRate,
		Username:          username,
		Password:          password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse txID")
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse changeAddrM")
	}
	return txID, changeAddrM, nil
}

//AddSubnetValidator ...
func (c *Client) AddSubnetValidator(nodeID string, subnetID string, startTime string, endTime string, weight int, changeAddr string, username string, password string) (string, string, error) {
	fmtp("-----------------== AddSubnetValidator()")
	JSONRPCMethod := "platform.addSubnetValidator"
	endpoint := "/ext/P"
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		NodeID:     nodeID,
		SubnetID:   subnetID,
		StartTime:  startTime,
		EndTime:    endTime,
		Weight:     weight,
		ChangeAddr: changeAddr,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse txID")
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse changeAddrM")
	}
	return txID, changeAddrM, nil
}

//GetPendingValidators ...
func (c *Client) GetPendingValidators(subnetID string) (string, string, string, string, string, error) {
	fmtp("-----------------== GetPendingValidators()")
	JSONRPCMethod := "platform.getPendingValidators"
	endpoint := "/ext/P"
	resField1 := "nodeID"
	resField2 := "startTime"
	resField3 := "endTime"
	resField4m := "stakeAmount"
	resField4s := "weight"

	params := Params{
		SubnetID: subnetID,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", "", "", "", errors.New(respOut.Mesg)
	}
	nodeID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", "", "", "", errors.New("could not parse nodeID")
	}
	startTime, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", "", "", "", errors.New("could not parse startTime")
	}
	endTime, ok := (respOut.Result[resField3]).(string)
	if !ok {
		return "", "", "", "", "", errors.New("could not parse endTime")
	}

	stakeAmount := ""
	weight := ""
	if subnetID == "" {
		stakeAmount, ok = (respOut.Result[resField4m]).(string)
		if !ok {
			return "", "", "", "", "", errors.New("could not parse stakeAmount")
		}
	} else {
		weight, ok = (respOut.Result[resField4s]).(string)
		if !ok {
			return "", "", "", "", "", errors.New("could not parse weight")
		}
	}
	return nodeID, startTime, endTime, stakeAmount, weight, nil
}

//ListAddresses ...
func (c *Client) ListAddresses(username string, password string, chain string) ([]string, error) {
	fmtp("-----------------== ListAddresses()")
	JSONRPCMethod := "avm.listAddresses"
	endpoint := "/ext/bc/" + chain
	resField1 := "addresses"
	params := Params{
		Username: username,
		Password: password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)

	if respOut.Mesg != "ok" {
		//x := make([]interface{}, 1)
		//x := []interface{(respOut.Mesg).(interface{})}
		return nil, errors.New(respOut.Mesg)
	}
	itfs, ok := (respOut.Result[resField1]).([]interface{})
	if !ok {
		return nil, errors.New("could not parse []interface{}")
	}

	//fmtp("itfs[1]", itfs[1])
	len1 := len(itfs)
	if len1 < 1 {
		return nil, errors.New("none")
	}

	fmtp("itfs[0]", itfs[0])
	fixed := make([]string, len1)
	for i := range itfs {
		fixed[i], ok = (itfs[i]).(string)
		if !ok {
			fmtp("err@ cannot parse address")
			fixed[i] = "non string"
		}
	}
	return fixed, nil
}

//GetBalance ...
func (c *Client) GetBalance(chain string, addr string, assetID string) (int, []OutputIdxTxID, error) {
	fmtp("-----------------== GetBalance()")
	JSONRPCMethod := "avm.getBalance"
	endpoint := "/ext/bc/" + chain
	resField1 := "balance"
	resField2 := "utxoIDs"
	params := Params{Address: chain + "-" + addr, AssetID: assetID}
	//fmtp("chain:", chain)
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return -1, nil, errors.New("invokeRPC")
	}
	balanceStr, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return -1, nil, errors.New("balance cannot be parsed as a string")
	}

	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		fmtp("err@ balance is not integer!", err)
		balance = -1
		return balance, nil, errors.New("balance is not integer")
	}

	utxoIDs, ok := (respOut.Result[resField2]).([]interface{})
	if !ok {
		return -1, nil, errors.New("utxoIDs cannot be parsed")
	}

	Dump1("utxoIDs:", utxoIDs)
	//a := respOut.Result[resField2]
	utxoIDsStr := make([]OutputIdxTxID, len(utxoIDs))
	for i := range utxoIDs {
		item, ok := (utxoIDs[i]).(map[string]interface{})
		if !ok {
			fmtp("err@ parsing utxoID @ index", i)
			return -1, nil, errors.New("parsing utxoID")
		}

		outIndexJSONNumber, ok := item["outputIndex"].(json.Number)
		if !ok {
			return -1, nil, errors.New("outIndexJSONNumber cannot be parsed")
		}

		txID, ok := item["txID"].(string)
		if !ok {
			return -1, nil, errors.New("err@ txID cannot be parsed")
		}

		f, err := outIndexJSONNumber.Int64()
		if err != nil {
			utxoIDsStr[i] = OutputIdxTxID{
				OutputIndex: -1,
				TxID:        txID,
			}
			continue
		}
		utxoIDsStr[i] = OutputIdxTxID{
			OutputIndex: f,
			TxID:        txID,
		}
	}
	fmtp("balance:", balance)
	Dump1("utxoIDsStr:", utxoIDsStr)
	return balance, utxoIDsStr, nil
}

//GetTxStatus ...
func (c *Client) GetTxStatus(chain string, txID string) (string, error) {
	fmtp("-----------------== getTxStatus()")
	var JSONRPCMethod string
	var endpoint string
	if chain == "X" {
		JSONRPCMethod = "avm.getTxStatus"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.getTxStatus"
		endpoint = "/ext/P"
	} else {
		return "", errors.New("chain is not valid")
	}

	resField1 := "status"
	params := Params{TxID: txID}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}

	status, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField1)
	}
	return status, nil
}

// FixAddrSlice ...
func FixAddrSlice(chain string, addrs []string) []string {
	fmtp("-----------== FixAddrSlice")
	fmtp("addrs", addrs)

	fixed := make([]string, len(addrs))
	for i := range addrs {
		fixed[i] = chain + "-" + addrs[i]
	}
	return fixed
}

// FixMinterSet ...
func FixMinterSet(chain string, minterSets []MinterSet) []MinterSet {
	fmtp("-----------== FixMinterSet")
	fmtp("minterSets", minterSets)

	fixedMinterSets := make([]MinterSet, len(minterSets))
	for i := range minterSets {
		minters := minterSets[i].Minters
		fixedMinters := make([]string, len(minters))

		for j := range minters {
			fixedMinters[j] = chain + "-" + minters[j]
		}
		fixedMinterSets[i] = MinterSet{
			Minters:   fixedMinters,
			Threshold: minterSets[i].Threshold,
		}
	}
	return fixedMinterSets
}

// FixAddrBal ...
func FixAddrBal(chain string, addrBals []AddrBal) []AddrBal {
	fmtp("-----------== FixAddrBal")
	fmtp("addrBals", addrBals)

	fixed := make([]AddrBal, len(addrBals))
	for i := range addrBals {
		fixed[i].Address = chain + "-" + addrBals[i].Address
		fixed[i].Amount = addrBals[i].Amount
	}
	return fixed
}

//CreateFixedCapAsset ...
func (c *Client) CreateFixedCapAsset(chain string, name string, symbol string, denomination int, initialHolders []AddrBal, from []string, changeAddr string, username string, password string) (string, string, error) {
	fmtp("-----------------== CreateFixedCapAsset()")
	JSONRPCMethod := "avm.createFixedCapAsset"
	//chain := initialHolders[0].Address[:1]
	endpoint := "/ext/bc/" + chain
	resField1 := "assetID"
	resField2 := "changeAddr"
	params := Params{
		Name:           name,
		Symbol:         symbol,
		Denomination:   denomination,
		InitialHolders: FixAddrBal(chain, initialHolders),
		From:           from,
		ChangeAddr:     chain + "-" + changeAddr,
		Username:       username,
		Password:       password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	assetID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse assetID")
	}

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse changeAddr")
	}
	return assetID, changeAddrM, nil
}

//CreateVariableCapAsset ...
func (c *Client) CreateVariableCapAsset(chain string, name string, symbol string, denomination int, minterSets []MinterSet, from []string, changeAddr string, username string, password string) (string, string, error) {
	fmtp("-----------------== CreateVariableCapAsset()")
	JSONRPCMethod := "avm.createVariableCapAsset"
	endpoint := "/ext/bc/" + chain
	resField1 := "assetID"
	resField2 := "changeAddr"
	params := Params{
		Name:         name,
		Symbol:       symbol,
		Denomination: denomination,
		MinterSets:   minterSets,
		From:         from,
		ChangeAddr:   chain + "-" + changeAddr,
		Username:     username,
		Password:     password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	assetID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse assetID")
	}

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse changeAddr")
	}
	return assetID, changeAddrM, nil
}

// MintAsset ...
func (c *Client) MintAsset(chain string, amount int, assetID string, addrTo string, from []string, changeAddr string, username string, password string) (string, string, error) {
	fmtp("-----------------== MintAsset()")
	JSONRPCMethod := "avm.mint"
	endpoint := "/ext/bc/" + chain
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		Amount:     amount,
		AssetID:    assetID,
		To:         chain + "-" + addrTo,
		From:       from,
		ChangeAddr: changeAddr,
		Username:   username,
		Password:   password,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField1)
	}
	fmtp("txID:", txID)

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField2)
	}
	fmtp("changeAddrM:", changeAddrM)
	return txID, changeAddrM, nil
}

//SendAsset ...
func (c *Client) SendAsset(chain string, changeAddr string, addrTo string, from []string, memo string, assetID string, amount int, username string, password string) (string, error) {
	fmtp("-----------------== SendAsset()")
	JSONRPCMethod := "avm.send"

	endpoint := "/ext/bc/" + chain
	resField1 := "txID"
	params := Params{
		AssetID:    assetID,
		To:         chain + "-" + addrTo,
		From:       from,
		ChangeAddr: chain + "-" + changeAddr,
		Amount:     amount,
		Memo:       memo,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField1)
	}
	return txID, nil
}

//WalletSend ...
func (c *Client) WalletSend(chain string, changeAddr string, addrTo string, from []string, memo string, assetID string, amount int, username string, password string) (string, string, error) {
	fmtp("-----------------== WalletSend()")
	JSONRPCMethod := "wallet.send"

	endpoint := "/ext/bc/" + chain + "/wallet"
	resField1 := "txID"
	resField2 := "changeAddr"
	params := Params{
		AssetID:    assetID,
		To:         chain + "-" + addrTo,
		From:       from,
		ChangeAddr: chain + "-" + changeAddr,
		Amount:     amount,
		Memo:       memo,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField1)
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField2)
	}
	return txID, changeAddrM, nil
}

//SendMultiple ...
func (c *Client) SendMultiple(chain string, changeAddr string, outputs []Output, from []string, memo string, username string, password string) (string, string, error) {
	fmtp("-----------------== SendMultiple()")
	JSONRPCMethod := "avm.sendMultiple"

	endpoint := "/ext/bc/" + chain
	resField1 := "txID"
	resField2 := "changeAddr"
	params := Params{
		Outputs:    outputs,
		From:       from,
		ChangeAddr: chain + "-" + changeAddr,
		Memo:       memo,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField1)
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField2)
	}
	return txID, changeAddrM, nil
}

//WalletSendMultiple ...
func (c *Client) WalletSendMultiple(chain string, changeAddr string, outputs []Output, from []string, memo string, username string, password string) (string, string, error) {
	fmtp("-----------------== WalletSendMultiple()")
	JSONRPCMethod := "wallet.sendMultiple"

	endpoint := "/ext/bc/" + chain + "/wallet"
	resField1 := "txID"
	resField2 := "changeAddr"
	params := Params{
		Outputs:    outputs,
		From:       from,
		ChangeAddr: chain + "-" + changeAddr,
		Memo:       memo,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField1)
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField2)
	}
	return txID, changeAddrM, nil
}

//ExportAVAX1 ...
func (c *Client) ExportAVAX1(chainFrom string, chainTo string, changeAddr string, addrTo string, assetID string, amount int, username string, password string) (string, error) {
	fmtp("-----------------== ExportAVAX1()")
	JSONRPCMethod := "avm.exportAVAX"

	endpoint := "/ext/bc/" + chainFrom
	resField1 := "txID"
	params := Params{
		To:               chainTo + "-" + addrTo,
		DestinationChain: chainTo,
		ChangeAddr:       chainFrom + "-" + changeAddr,
		Amount:           amount,
		Username:         username,
		Password:         password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	//Dump1("respOut.UTXOIDs:", respOut.UTXOIDs)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField1)
	}
	return txID, nil
}

//ExportAVAX2 ...
func (c *Client) ExportAVAX2(chainFrom string, chainTo string, addrFrom string, addrTo string, assetID string, amount int, username string, password string) (string, string, error) {
	fmtp("-----------------== ExportAVAX2()")
	var JSONRPCMethod string
	var endpoint string
	if chainFrom == "X" {
		JSONRPCMethod = "avm.exportAVAX"
		endpoint = "/ext/bc/X"
	} else if chainFrom == "P" {
		JSONRPCMethod = "platform.exportAVAX"
		endpoint = "/ext/P"
	} else if chainFrom == "C" {
		JSONRPCMethod = "avax.exportAVAX"
		endpoint = "/ext/bc/C/avax"
	} else {
		return "", "", errors.New("chainFrom is not valid")
	}
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		To:       chainTo + "-" + addrTo,
		From:     []string{chainFrom + "-" + addrFrom},
		Amount:   amount,
		Username: username,
		Password: password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField1)
	}
	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse " + resField2)
	}
	return txID, changeAddrM, nil
}

//ImportAVAX ...
func (c *Client) ImportAVAX(chainFrom string, chainTo string, addrTo string, assetID string, privateKey string, username string, password string) (string, string, string, error) {
	fmtp("-----------------== ImportAVAX()")
	resField1 := "txID"
	resField2 := "changeAddr"
	resField3 := "address"

	var JSONRPCMethod string
	var endpoint string
	if chainTo == "X" {
		JSONRPCMethod = "avm.importAVAX"
		endpoint = "/ext/bc/X"
	} else if chainTo == "P" {
		JSONRPCMethod = "platform.importAVAX"
		endpoint = "/ext/bc/P"
	} else if chainTo == "C" {
		JSONRPCMethod = "avax.importKey"
		endpoint = "/ext/bc/C/avax"
	} else {
		return "", "", "", errors.New("chainTo is not valid")
	}

	params := Params{
		To:          chainTo + "-" + addrTo,
		SourceChain: chainFrom,
		Username:    username,
		Password:    password,
		PrivateKey:  privateKey,
	} //pk only for importing to C chain

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", "", errors.New(respOut.Mesg)
	}

	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", "", errors.New("could not parse " + resField1)
	}
	changeAddr, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", "", errors.New("could not parse " + resField2)
	}
	address, ok := (respOut.Result[resField3]).(string)
	if !ok {
		return "", "", "", errors.New("could not parse " + resField3)
	}
	return txID, changeAddr, address, nil
}

//GetUTXOs ...
func (c *Client) GetUTXOs(chain string, addresses []string, limit int, encoding string) (string, []interface{}, string, string, error) {
	//@@@ https://docs.avax.network/v1.0/en/api/avm/#avmcreatefixedcapasset
	//numFetched, utxos, address, utxo, nil
	fmtp("-----------------== GetUTXOs()")
	JSONRPCMethod := "avm.getUTXOs"
	//chain := addresses[0][:1]
	endpoint := "/ext/bc/" + chain
	resField1 := "numFetched"
	resField2 := "utxos"
	resField3 := "endIndex"

	params := Params{
		Addresses: FixAddrSlice(chain, addresses),
		Limit:     limit,
		Encoding:  encoding,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", nil, "", "", errors.New(respOut.Mesg)
	}
	numFetched, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", nil, "", "", errors.New("cannot parse " + resField1)
	}
	Dump1("numFetched:", numFetched)

	utxos, ok := (respOut.Result[resField2]).([]interface{})
	if !ok {
		return "", nil, "", "", errors.New("utxos cannot be parsed")
	}
	Dump1("utxos:", utxos)

	endIndex, ok := (respOut.Result[resField3]).(map[string]interface{})
	if !ok {
		return "", nil, "", "", errors.New("endIndex cannot be parsed")
	}
	Dump1("endIndex:", endIndex)
	address, ok := endIndex["address"].(string)
	if !ok {
		return "", nil, "", "", errors.New("address cannot be parsed")
	}

	utxo, ok := endIndex["utxo"].(string)
	if !ok {
		return "", nil, "", "", errors.New("utxo cannot be parsed")
	}

	utxos0str, ok := utxos[0].(string)
	if !ok {
		return "", nil, "", "", errors.New("utxos0str cannot be parsed")
	}

	decoded, err := base58check.Decode(utxos0str)
	if err != nil {
		fmtp("err", err)
	}
	fmtp(decoded)

	return numFetched, utxos, address, utxo, nil
}

//Base58check ...
func (c *Client) Base58check() {
	fmtp("-----------------== Base58check()")
	// encoded, err := base58check.Encode("80", "44D00F6EB2E5491CD7AB7E7185D81B67A23C4980F62B2ED0914D32B7EB1C5581")
	// if err != nil {
	// 	fmtp(err)
	// }
	// fmtp(encoded)

	//encodedStr := "1mayif3H2JDC62S4N3rLNtBNRAiUUP99k"
	encodedStr := "116Q8K1wQ66wUnQWnZvR1duxJoJmNgvLAgCqomB5A2Gu3VYrG5jv1TLtLnVcgXCpH5s4FawVc7bnPUYbd76h6JLjPC2Qrv9kZi4g9eLWR1GhoySZCe67CdciXxkPvYEgPPUwDmbHsgveeyLEB7y9pXtStPA14rTu"
	decoded, err := base58check.Decode(encodedStr)
	if err != nil {
		fmtp("err:", err)
	}
	fmtp(decoded) // 00086eaa677895f92d4a6c5ef740c168932b5e3f44
}

// Base58 ...
func (c *Client) Base58() {
	//data := []byte("AVA Labs")
	data := []byte("Tesla Space X")
	encoded := base58.Encode(data)
	fmt.Println("Encoded:", encoded)

	//encoded := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	//encoded = "116VhGCxiSL4GrMPKHkk9Z92WCn2i4qk8qdN3gQkFz6FMEbHo82Lgg8nkMCPJcZgpVXZLQU6MfYuqRWfzHrojmcjKWbfwqzZoZZmvSjdD3KJFsW3PDs5oL3XpCHq4vkfFy3q1wxVY8qRc6VrTZaExfHKSQXX1KnC"
	//encoded := "116Q8K1wQ66wUnQWnZvR1duxJoJmNgvLAgCqomB5A2Gu3VYrG5jv1TLtLnVcgXCpH5s4FawVc7bnPUYbd76h6JLjPC2Qrv9kZi4g9eLWR1GhoySZCe67CdciXxkPvYEgPPUwDmbHsgveeyLEB7y9pXtStPA14rTu"
	decoded := base58.Decode(encoded) // return []byte
	fmt.Println("Decoded:", decoded)
	decodedStr := string(decoded)
	fmt.Println("DecodedStr:", decodedStr)

	src := decoded
	encodedHexStr := hex.EncodeToString(src)
	fmt.Printf("encodedHexStr: %s\n", encodedHexStr)
	//00000478f2398dd2163c34132ce7afa31f0ac503017f863bf4db87ea5553c52d7b57000000010478f2398dd2163c34132ce7afa31f0ac503017f863bf4db87ea5553c52d7b570000000a00000000000000000000000000000001000000013cb7d3842e8cee6a0ebd09f1fe884f6861e1b29c18b1efa5
	// https://www.better-converter.com/Encoders-Decoders/Base58Check-to-Hexadecimal-Decoder

	// decoded, version, err := base58.CheckDecode(encoded)
	// if err != nil {
	// 	fmtp(err)
	// 	return
	// }
	// // Show the decoded data.
	// fmt.Printf("Decoded data: %x\n", decoded)
	// fmtp("Version Byte:", version)

	// decoded2 := base58.Decode("2EWh72jYQvEJF9NLk")
	// fmt.Println("Decoded2:", decoded2)
	// decoded2Str := string(decoded2)
	// fmt.Println("Decoded2Str:", decoded2Str)
}

//CreateNFTAsset ...
func (c *Client) CreateNFTAsset(chain string, name string, symbol string, minterSets []MinterSet, username string, password string) (string, string, error) {
	fmtp("-----------------== CreateNFTAsset()")
	JSONRPCMethod := "avm.createNFTAsset"
	endpoint := "/ext/bc/" + chain
	resField1 := "assetID"
	resField2 := "changeAddr"
	params := Params{
		Name:       name,
		Symbol:     symbol,
		MinterSets: minterSets,
		Username:   username,
		Password:   password,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", "", errors.New(respOut.Mesg)
	}

	assetID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("could not parse assetID")
	}

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("could not parse changeAddr")
	}
	return assetID, changeAddrM, nil
}

// MintNFT ...
func (c *Client) MintNFT(chain string, assetID string, payloadStr string, addrTo string, username string, password string) (string, string, error) {
	fmtp("-----------------== MintNFT()")
	JSONRPCMethod := "avm.mintNFT"
	//chain := addrTo[:1]
	endpoint := "/ext/bc/" + chain
	resField1 := "txID"
	resField2 := "changeAddr"

	encoded := "2EWh72jYQvEJF9NLk"
	// data := []byte(payloadStr)
	// encoded := base58.CheckEncode(data, 0)
	fmt.Println("CB58 Encoded:", encoded)

	params := Params{
		AssetID:  assetID,
		Payload:  encoded,
		To:       chain + "-" + addrTo,
		Username: username,
		Password: password,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField1)
	}
	fmtp("txID:", txID)

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField2)
	}
	fmtp("changeAddrM:", changeAddrM)
	return txID, changeAddrM, nil
}

// SendNFT ...
func (c *Client) SendNFT(chain string, assetID string, addrTo string, from []string, changeAddr string, groupID int, username string, password string) (string, string, error) {
	fmtp("-----------------== MintNFT()")
	JSONRPCMethod := "avm.sendNFT"
	//chain := addrTo[:1]
	endpoint := "/ext/bc/" + chain
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		AssetID:    assetID,
		GroupID:    groupID,
		To:         chain + "-" + addrTo,
		From:       from,
		ChangeAddr: chain + "-" + changeAddr,
		Username:   username,
		Password:   password,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField1)
	}
	fmtp("txID:", txID)

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField2)
	}
	fmtp("changeAddrM:", changeAddrM)
	return txID, changeAddrM, nil
}

/*
([]interface {}) (len=12 cap=16) {
 (string) (len=160) "116Q8K1wQ66wUnQWnZvR1duxJoJmNgvLAgCqomB5A2Gu3VYrG5jv1TLtLnVcgXCpH5s4FawVc7bnPUYbd76h6JLjPC2Qrv9kZi4g9eLWR1GhoySZCe67CdciXxkPvYEgPPUwDmbHsgveeyLEB7y9pXtStPA14rTu",
 (string) (len=166) "112bfZVEWroUJxiRQ8aaBoqw2Jyhe2uwmEK3zYVG5HGJf22VfhiHJLnbAKdK7dMXv5CBgG5rzTe9SKQBQA1FWyfY8AwkPJg1irvpRTrXdRtDRMkFUV6cHsUm19GvWHxZ5LV9Y1ZAh863Jp3x675xY4UUY5KAtJ8HPHKqmQ",
 (string) (len=166) "11Hw2KkjD2EiczDXXMC1Kjc43YjKGrotq1AoXHk2W3GL5UAddEUrsecboKaDWcp32nENVggaQgjQbhA1LtKzL15jXW4FsAp7tv617xyA3YTMzFDX2QuUkd3J3Dm1zTgqhutnFUY9twpjvy9yLFAMnYHfi4FMzmUGQxR6yd",
 (string) (len=166) "115JvNKre4G1opQzCKCkgv7Z6jjZ9QX1NVvUxTdgCLkJEmgVnZinmLr7zKhSLN8yopwEEGcWDFeVpfjwXptyZybfH1ZGMnU46K9BjGfveiC6hYKhR56f9ryQL6vQFHPpEnrDbtCLYDnsAbv8ioCPMfuvSqFaVUMNYXis9G",
 (string) (len=155) "11MbcRUUzEdKAARuEG3UHcE9FCdDBLg4QgZWffMunkxABpYJaKSpcfvVXw8CwA6qZikPobXYUVFYSLKq3Y7EPkPG2TP7jygJghuzcsDsX4BTY552o3b4pCJAgRptZqhW2g9kH3nxzoPjVzrzxQZynCqm3ao",
 (string) (len=166) "11C4Xo3K7TDBbnNc8aLhQeqF9c3zgGMf6TasDTS3PwN3mzHVU5BWUUn6X9G5nuVfXubACufYKDLVLiB8cC7Ra9c46VwWkLKyef5rvYwUz44FnQTb8vdw9fnpGUWmwpKCVHhZQMysPPrFfFUyJgazUDsVuXmj86cSRwfHma",
 (string) (len=165) "11cKzDRBFHQKUiPHEZ7NGtrh4141Bg276mR27ggMqS1Mk2TWh1Zm3PZZWy3A6eMyNzHANZQntPjBCakL3LTecsAPbX83XMRw7WR4QMr62KQtCv4deseMaBvL3fY7oADcSM81cSmrhTe1BU6EWT7Rxd9ziWfpx4zTMzqrM",
 (string) (len=166) "11LC2AMsJnJk66VRqHzdstd7FSzF2QKruDYkwq1cGPHPB4Zgey6y21W9XPmaQbcYPaTSpp69WgA1qc263CAyuhsT6iAT9NSZ3ZDQJDuvapxkCXsE3GDwU5Ykg3P4sfanUPUK5cE5bYx8EVRWNJiGPGzHB5K9iizLcbaanN",
 (string) (len=166) "118YZDxocgpLca9uDSGJ9MMbPzdfxUfzaV8kwKEMD7LjvnaBTbwHBB7fLubA3QbdCREKGZVq4iYsG42z3z5PgDTE1X7S8zjPCFHPXLRSVwpSd2ZMUd92vPDwaViuSBsKzfvdR1wJpZjnegAWsP2mxjSz8kkk2MGDoQ2q4P",
 (string) (len=166) "112VwRNQdyU8hL9bh2FzNc51mPb6PErtU3zun1Vd4pgvxMQ6kRkWsuxT7RhXjLeD47uKoYqGWqwuQcHrX9obCsz4Ki3LemNK53DFFUt4ur5NgPUZgEr6kidxD9n2mkXfx6ao1LBepxx17pvqeP1dbzJaxZWkxfK9AjnfS6",
 (string) (len=166) "1186cKM3RSxgyfvvLouyBrMq3BKbcG5K9W9efeW2EW1iLe4xV6DFXoTW6cVo9fc2s9zwbahT6vkEFvCXdBngJ5D5jK4QMcU9wztDTAXWNvDwkjYWRTsys6apT3q3utv63i7kYHpZssHh5jkaXKc8QYf3HKYEw4mJ8YFMzV",
 (string) (len=166) "11ZQTgmmgQPTd57gzUafZZMqEQVp9CBMgWwq8htgoDKmkvERkVb4vjGP2YKbuxozAYYb6XrRDMDUwb2FH2tTy793VEtQy9TesAjNusRev5HATAxtA2JNnEx4SMVKvDPv9gksUjGDxFA5WjMSCR2sqHm5apwunbARK2SXzH"
*/

// BuildGenesis ...
func (c *Client) BuildGenesis(assetFixedCapName string,
	assetFixedCapSymbol string, assetVariableCapName string, assetVariableCapSymbol string, addr1 string, addr2 string, addr3 string) (string, error) {
	fmtp("-----------------== BuildGenesis()")
	JSONRPCMethod := "avm.buildGenesis"
	endpoint := "/ext/vm/avm"
	resField1 := "bytes"
	// assetFixedCapName := "FixedAssetName1"
	// assetFixedCapSymbol := "FIXA"
	// assetVariableCapName := "VariableAssetName1"
	// assetVariableCapSymbol := "VARA"

	//set assetname below
	type genesisData struct {
		Asset1 GenesisAsset `json:"asset1"`
		Asset2 GenesisAsset `json:"asset2"`
	}

	params := Params{
		GenesisData: genesisData{
			Asset1: GenesisAsset{
				Name:   assetFixedCapName,
				Symbol: assetFixedCapSymbol,
				InitialState: InitStateFixedCap{
					[]AddrBal{
						{Address: addr1, Amount: 1000},
						{Address: addr2, Amount: 3000},
						{Address: addr3, Amount: 5000},
					},
				},
			},
			Asset2: GenesisAsset{
				Name:   assetVariableCapName,
				Symbol: assetVariableCapSymbol,
				InitialState: InitStateVariableCap{
					[]MinterSet{
						{Minters: []string{addr1, addr2}, Threshold: 1},
						{Minters: []string{addr1, addr2, addr3}, Threshold: 2},
					},
				},
			},
		},
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", errors.New(respOut.Mesg)
	}
	bytes, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("bytes cannot be parsed as string")
	}
	fmtp("bytes:", bytes)
	return bytes, nil
}

// CreateBlockchain ...
func (c *Client) CreateBlockchain(subnetID string, vmID string, name string, genesisData string, username string, password string) (string, string, error) {
	fmtp("-----------------== CreateBlockchain()")
	JSONRPCMethod := "platform.createBlockchain"
	endpoint := "/ext/P"
	resField1 := "txID"
	resField2 := "changeAddr"

	params := Params{
		SubnetID:    subnetID,
		VMID:        vmID,
		Name:        name,
		GenesisData: genesisData,
		Username:    username,
		Password:    password,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", "", errors.New(respOut.Mesg)
	}
	txID, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField1)
	}
	fmtp("txID:", txID)

	changeAddrM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField2)
	}
	fmtp("changeAddrM:", changeAddrM)
	return txID, changeAddrM, nil
}

//GetBlockchains ...
func (c *Client) GetBlockchains() ([]Blockchain, error) {
	fmtp("-----------------== GetBlockchains()")
	JSONRPCMethod := "platform.getBlockchains"
	endpoint := "/ext/P"
	resField1 := "blockchains"

	params := Params{}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return nil, errors.New(respOut.Mesg)
	}
	blockchains, ok := (respOut.Result[resField1]).([]map[string]interface{})
	if !ok {
		return nil, errors.New("could not parse " + resField1)
	}

	blockchainsStr := make([]Blockchain, len(blockchains))
	for i := range blockchains {
		// item, ok := v.(map[string]interface{})
		// if !ok {
		// 	fmtp("err@ parsing utxoID @ index", i)
		// 	return -1, nil, "err@ parsing utxoID"
		// }

		id, ok := blockchains[i]["id"].(string)
		if !ok {
			fmtp("id cannot be parsed @ index", i)
			id = "-1"
		}

		name, ok := blockchains[i]["name"].(string)
		if !ok {
			fmtp("name cannot be parsed @ index", i)
			name = "-1"
		}

		subnetID, ok := blockchains[i]["subnetID"].(string)
		if !ok {
			fmtp("subnetID cannot be parsed @ index", i)
			subnetID = "-1"
		}

		vmID, ok := blockchains[i]["vmID"].(string)
		if !ok {
			fmtp("vmID cannot be parsed @ index", i)
			vmID = "-1"
		}
		fmtp("id:", id, ", name:", name, ", subnetID:", subnetID, ", vmID:", vmID)
		blockchainsStr[i] = Blockchain{id, name, subnetID, vmID}
	}

	return blockchainsStr, nil
}

// ExportKey ...
func (c *Client) ExportKey(chain string, address string, username string, password string) (string, error) {
	fmtp("-----------------== ExportKey()")
	var endpoint string
	var JSONRPCMethod string
	if chain == "X" {
		JSONRPCMethod = "avm.exportKey"
		endpoint = "/ext/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return "", errors.New("chain is not valid")
	}

	resField1 := "privateKey"

	params := Params{
		Address:  chain + "-" + address,
		Username: username,
		Password: password,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", errors.New(respOut.Mesg)
	}
	privateKey, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("cannot parse " + resField1)
	}
	fmtp(resField1+":", privateKey)
	return privateKey, nil
}

// GetAllBalances ...
func (c *Client) GetAllBalances(chain string, address string) ([]AssetBalance, error) {
	fmtp("-----------------== GetAllBalances()")
	var endpoint string
	var JSONRPCMethod string
	if chain == "X" {
		JSONRPCMethod = "avm.getAllBalances"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return nil, errors.New("chain is not valid")
	}
	resField1 := "balances"

	params := Params{
		Address: chain + "-" + address,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return nil, errors.New(respOut.Mesg)
	}

	balancesRaw, ok := (respOut.Result[resField1]).([]map[string]interface{})
	if !ok {
		return nil, errors.New("cannot parse " + resField1)
	}

	assetBalances := make([]AssetBalance, 0)
	for i := range balancesRaw {
		fmtp("peer of index", i, balancesRaw[i])
		asset, ok := (balancesRaw[i]["asset"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing asset")
		}

		balance, ok := (balancesRaw[i]["balance"]).(string)
		if !ok {
			return nil, errors.New("err@ parsing balance")
		}
		assetBalances = append(assetBalances, AssetBalance{asset, balance})
	}
	return assetBalances, nil
}

// GetAssetDescription ...
func (c *Client) GetAssetDescription(chain string, assetID string) (string, string, int, error) {
	fmtp("-----------------== GetAssetDescription()")
	var endpoint string
	var JSONRPCMethod string
	if chain == "X" {
		JSONRPCMethod = "avm.getAssetDescription"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return "", "", -1, errors.New("chain is not valid")
	}
	resField1 := "name"
	resField2 := "symbol"
	resField3 := "denomination"

	params := Params{
		AssetID: assetID,
	}
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", "", -1, errors.New(respOut.Mesg)
	}

	name, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", -1, errors.New("cannot parse " + resField1)
	}
	symbol, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", -1, errors.New("cannot parse " + resField2)
	}
	denomination, ok := (respOut.Result[resField3]).(int)
	if !ok {
		return "", "", -1, errors.New("cannot parse " + resField3)
	}
	return name, symbol, denomination, nil
}

// GetTx ...
func (c *Client) GetTx(chain string, txID string, encoding string) (string, string, error) {
	fmtp("-----------------== GetTx()")
	var endpoint string
	var JSONRPCMethod string
	if chain == "X" {
		JSONRPCMethod = "avm.getTx"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return "", "", errors.New("chain is not valid")
	}
	resField1 := "tx"
	resField2 := "encoding"

	params := Params{
		TxID:     txID,
		Encoding: encoding,
	} //chain + "-" +
	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		fmtp("err@ invokeRPC")
		return "", "", errors.New(respOut.Mesg)
	}

	tx, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField1)
	}
	encodingM, ok := (respOut.Result[resField2]).(string)
	if !ok {
		return "", "", errors.New("cannot parse " + resField2)
	}
	return tx, encodingM, nil
}

//IssueTx ... Send a signed transaction to the network. encoding specifies the format of the signed transaction. Can be either “cb58” or “hex”. Defaults to “cb58”.
func (c *Client) IssueTx(chain string, tx string, encoding string) (string, error) {
	fmtp("-----------------== IssueTx()")
	var JSONRPCMethod string
	var endpoint string
	if chain == "X" {
		JSONRPCMethod = "avm.issueTx"
		endpoint = "/ext/bc/X"
	} else if chain == "P" {
		JSONRPCMethod = "platform.createAddress"
		endpoint = "/ext/P"
	} else if chain == "C" {
		JSONRPCMethod = "contract.createAddress"
		endpoint = "/ext/C"
	} else {
		return "", errors.New("chain is not valid")
	}
	resField1 := "txID"
	params := Params{
		Tx:       tx,
		Encoding: encoding,
	}

	respOut := invokeRPC(JSONRPCMethod, params, c, endpoint)
	if respOut.Mesg != "ok" {
		return "", errors.New(respOut.Mesg)
	}
	val, ok := (respOut.Result[resField1]).(string)
	if !ok {
		return "", errors.New("could not parse " + resField1)
	}
	return val, nil
}

/*

 */
