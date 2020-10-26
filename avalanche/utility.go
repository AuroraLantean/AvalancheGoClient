package avalanche

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
	"github.com/ybbus/jsonrpc"
)

// RoutineInputs ...
type RoutineInputs struct {
	RoutineName   string `json:"routineName"`
	JSONRPC       string `json:"jsonrpc"`
	JSONRPCMethod string `json:"jsonrpcMethod"`
	Params        Params `json:"params"`
	ID            int    `json:"id"`
	URL           string `json:"url"`
	Timeout       int    `json:"timeout"`
}

// AddrBal ...
type AddrBal struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

// GenesisAsset ...
type GenesisAsset struct {
	Name         string      `json:"name"`
	Symbol       string      `json:"symbol"`
	InitialState interface{} `json:"initialState"`
}

// InitStateFixedCap ...
type InitStateFixedCap struct {
	FixedCap []AddrBal `json:"fixedCap"`
}

// InitStateVariableCap ...
type InitStateVariableCap struct {
	VariableCap []MinterSet `json:"variableCap"`
}

// OutputIdxTxID ...
type OutputIdxTxID struct {
	OutputIndex int64  `json:"outputIndex"`
	TxID        string `json:"txID"`
}

// EndIndex1 ...
type EndIndex1 struct {
	Address string `json:"address"`
	UTXO    string `json:"utxo"`
}

// MinterSet ...
type MinterSet struct {
	Minters   []string `json:"minters"`
	Threshold int      `json:"threshold"`
}

// Blockchain ...
type Blockchain struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	SubnetID string `json:"subnetID"`
	VMID     string `json:"vmID"`
}

// Params ...
type Params struct {
	Chain             string      `json:"chain"`
	Username          string      `json:"username"`
	Password          string      `json:"password"`
	Address           string      `json:"address"`
	Addresses         []string    `json:"addresses"`
	AssetID           string      `json:"assetID"`
	To                string      `json:"to"`
	ChangeAddr        string      `json:"changeAddr"`
	DestinationChain  string      `json:"destinationChain"`
	SourceChain       string      `json:"sourceChain"`
	PrivateKey        string      `json:"privateKey"`
	Amount            int         `json:"amount"`
	TxID              string      `json:"txID"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Denomination      int         `json:"denomination"`
	GroupID           int         `json:"groupID"`
	InitialHolders    []AddrBal   `json:"initialHolders"`
	From              []string    `json:"from"`
	MinterSets        []MinterSet `json:"minterSets"`
	Payload           string      `json:"payload"`
	ControlKeys       []string    `json:"controlKeys"`
	Threshold         int         `json:"threshold"`
	NodeID            string      `json:"nodeID"`
	SubnetID          string      `json:"subnetID"`
	Weight            int         `json:"weight"`
	StartTime         string      `json:"startTime"`
	EndTime           string      `json:"endTime"`
	StakeAmount       int         `json:"stakeAmount"`
	RewardAddress     string      `json:"rewardAddress"`
	DelegationFeeRate int         `json:"delegationFeeRate"`
	GenesisData       interface{} `json:"genesisData"`
	VMID              string      `json:"vmID"`
	Encoding          string      `json:"encoding"`
	Tx                string      `json:"tx"`
	Memo              string      `json:"memo"`
	Limit             int         `json:"limit"`
	Outputs           []Output    `json:"outputs"`
}

// RespOut ...
type RespOut struct {
	Mesg   string                 `json:"mesg"`
	Result map[string]interface{} `json:"result"`
}

// AssetBalance ...
type AssetBalance struct {
	Asset   string `json:"asset"`
	Balance string `json:"balance"`
}

// Peer ...
type Peer struct {
	IP           string `json:"ip"`
	PublicIP     string `json:"publicIP"`
	ID           string `json:"id"`
	Version      string `json:"version"`
	LastSent     string `json:"lastSent"`
	LastReceived string `json:"lastReceived"`
}

// Output ...
type Output struct {
	AssetID string `json:"assetID"`
	To      string `json:"to"`
	Amount  int    `json:"amount"`
}

// AvaOut ...
// type AvaOut struct {
// 	UTXOs          []interface{}   `json:"utxos"`
// 	Mesg           string          `json:"mesg"`
// 	TxID           string          `json:"txID"`
// 	Status         string          `json:"status"`
// 	Balance        int             `json:"balance"`
// 	AssetID        string          `json:"assetID"`
// 	UTXOIDs        []OutputIdxTxID `json:"utxoIDs"`
// 	Address        string          `json:"address"`
// 	Addresses      []string        `json:"addresses"`
// 	Success        bool            `json:"success"`
// 	ChangeAddr     string          `json:"changeAddr"`
// 	NodeID         string          `json:"nodeID"`
// 	IsBootstrapped bool            `json:"isBootstrapped"`
// 	NumFetched     string          `json:"numFetched"`
// 	UTXO           string          `json:"utxo"`
// }

/*
curl -X POST --data '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"avm.getTxStatus",
    "params" :{
        "txID":"2EAgR1Y.....heewnxSR"
    }
}' -H 'content-type:application/json;' 127.0.0.1:9650/ext/bc/X
*/
func invokeRPC(JSONRPCMethod string, params Params, c *Client, endpoint string) RespOut {
	fmtp("-----------------== invokeRPC()")
	URL := "http://" + c.NodeURL + ":" + c.Port + endpoint
	Dump1("params:", params)

	//https://pkg.go.dev/github.com/ybbus/jsonrpc@v2.1.2+incompatible?tab=overview
	rpcClient := jsonrpc.NewClient(URL)
	// rpcClient := jsonrpc.NewClientWithOpts(URL, &jsonrpc.RPCClientOpts{
	// 	CustomHeaders: map[string]string{
	// 		"content-type": "application/json",
	// 	},
	// })
	fmtp("check 1")

	resp, err := rpcClient.Call(JSONRPCMethod, params)
	Dump1("resp:", resp)
	fmtp("check 2")
	if err != nil || resp == nil {
		Dump1("err@rpcClient.Call:", err)
		// switch e := err.(type) {
		//   case nil: // if error is nil, do nothing
		//   case *HTTPError:
		//     // use e.Code here
		//     return
		//   default:
		//     // any other error
		//     return
		// }
		return RespOut{Mesg: "err@rpcClient.Call"}
	}
	if resp.Error != nil {
		fmtp("rpc-json protocol error")
		// check response.Error.Code, response.Error.Message and optional response.Error.Data
		return RespOut{Mesg: "rpc-json protocol error"}
	}

	//var rpcResp *jsonrpc.RPCResponse
	//err = resp.GetObject(&rpcResp) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	//Dump1("rpcResp:", rpcResp)

	fmtp("check 3")
	Dump1("resp.Result:", resp.Result)
	if (*resp).Result != nil {
		result, ok := ((*resp).Result).(map[string]interface{})
		if !ok {
			return RespOut{Mesg: "cannot parse (*resp).Result"}
		}
		return RespOut{Mesg: "ok", Result: result}
	}
	return RespOut{Mesg: "(*resp).Result is nil"}
}

// Dump1 ... to print structs
var Dump1 = spew.Dump

func fmtp(a ...interface{}) {
	fmt.Println(a...)
}

// to convert int to string
func toStr(i int) string {
	return strconv.Itoa(i)
}

// to convert int to string
func toStr64(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// to convert string to int
func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmtp("err@converting string to int. s:", s, ", err:", err)
		return -371
	}
	return i
}

// to convert string to float
func toFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmtp("err@converting string to a float. s:", s, ", err:", err)
		return -371.00
		//fmtp(f) // bitSize is 32 for float32 convertible,
		// 64 for float64
	}
	return f
}

// to check input for minimum length
func checkStrFixLength(s string, fixedLen int, inputName string) (bool, error) {
	if s == "" {
		fmtp(inputName + " is empty")
		return false, errors.New(inputName + " is empty:")
	}
	strlen := utf8.RuneCountInString(s)
	if strlen != fixedLen {
		fmtp(inputName, "of length", strlen, "should be of", toStr(fixedLen), "characters in length")
		return false, errors.New(inputName + " should be of " + toStr(fixedLen) + " characters in length")
	}
	fmtp(inputName + " is valid via checkStrFixLength")
	return true, nil
}

// to check input for minimum length
func checkInput(s string, minLen int, inputName string) (bool, error) {
	if s == "" {
		fmtp(inputName + " is empty")
		return false, errors.New(inputName + " is empty:")
	}
	if utf8.RuneCountInString(s) < minLen {
		fmtp(inputName + " should be at least " + toStr(minLen) + " characters in length")
		return false, errors.New(inputName + " should be at least " + toStr(minLen) + " characters in length")
	}
	fmtp(inputName + " is valid via checkInput")
	return true, nil
}

func checkCharLength(s string) int {
	return utf8.RuneCountInString(s)
}

func getErr(errs []error) (int, error) {
	for idx, err := range errs {
		fmtp(err)
		if err != nil {
			return idx, err
		}
	}
	return -1, nil
}

// to log fatal error and stop execution
// func logFatalErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

/*
// MakeJSONRPC ...
func MakeJSONRPC(ch1 chan *RoutineOut,
	requestURL string, jsonrpcMethod string) {
	fmtp("----------== MakeJSONRPC")
	//Dump1(requestURL)
	client := &http.Client{}
	req, err := http.NewRequest(jsonrpcMethod, requestURL, nil)
	//resp, err := http.Get(requestURL)
	if err != nil {
		fmtp("sending SMS error@http.NewRequest():", err)
		ch1 <- &RoutineOut{"sending SMS error", "NA"}
	}
	resp, err := client.Do(req)
	if resp == nil || resp.Body == nil {
		fmtp("resp or resp.Body is niil:", resp)
		ch1 <- &RoutineOut{"HTTP response is nil or its body is nil", "NA"}
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmtp("reading response error@ioutil.ReadAll:", err)
		ch1 <- &RoutineOut{"reading response error", "NA"}
	}
	respStr := string(respBody)
	Dump1("respStr:", respStr)

	items := strings.Split(respStr, ",")
	if len(items) < 1 {
		fmtp("response length not valid")
		ch1 <- &RoutineOut{"response length not valid", respStr}
	}
	balance := toFloat(items[0])
	if balance < 0 {
		fmtp("SMS delivery failed")
		ch1 <- &RoutineOut{"SMS delivery failed", respStr}
	}
	err = resp.Body.Close()
	if err != nil {
		fmtp("response close error@ resp.Body.Close():", err)
		ch1 <- &RoutineOut{"response close error",
			respStr}
	}
	fmtp("SMS sending is successful")
	ch1 <- &RoutineOut{"OK", respStr}
}

// ExecuteRoutine ...
func ExecuteRoutine(routineInputs RoutineInputs) (*RoutineOut, error) {
	fmtp("---------------== ExecuteRoutine")
	Dump1("routineInputs:", routineInputs)
	routineName := routineInputs.RoutineName
	jsonrpcMethod := routineInputs.JSONRPCMethod
	URL := routineInputs.URL
	timeout := routineInputs.Timeout

	ch1 := make(chan *RoutineOut)
	switch {
	// case routineName == "MakeGetRequest":
	// 	go MakeGetRequest(ch1, routineAddr)
	case routineName == "MakeJSONRPC":
		go MakeJSONRPC(ch1, URL, jsonrpcMethod)
	default:
		fmtp("routineName has no match!")
		return &RoutineOut{"function input not valid",
			"NA"}, nil
	}
	var RoutineOutPtr *RoutineOut
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		fmtp("routine takes too long. timeout has been reached")
		RoutineOutPtr = &RoutineOut{
			"routine takes too long: " + toStr(timeout) + " seconds", "NA"}
	case RoutineOutPtr = <-ch1:
		fmtp("Success@ CallGoroutine():" +
			"channel value has been returned")
	}
	return RoutineOutPtr, nil
}*/
