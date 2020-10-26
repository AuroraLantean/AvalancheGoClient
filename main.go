package main

import (
	"fmt"
	"main/avalanche"
	"os"
	"time"
)

func main() {
	args := []string{"0", "0", "0", "0", "0"}
	argsIn := os.Args[1:]
	for i := range argsIn {
		args[i] = argsIn[i]
	}
	fmtp("args", args)

	// rpcClient := jsonrpc.NewClient("http://127.0.0.1:9650/ext/info")
	// fmtp("check 0")
	// resp, err := rpcClient.Call("info.getNodeID")
	// fmtp("check 1", resp, err)

	avaClient := avalanche.Client{
		NodeURL: "127.0.0.1",
		Port:    "9650",
	}
	// username := "zulu2038"
	// password := "avagoflutter"
	//addr1 := "fuji1fyqtlkjhl0udzl2ldjyhgye4t8lz8qwusf0a0x"

	nodeID, err := avaClient.GetNodeID()
	fmtpf("GetNodeID()", err)
	fmtp("nodeID:", nodeID)

	isBootstrapped, err := avaClient.IsBootstrapped()
	fmtpf("IsBootstrapped()", err)
	fmtp("isBootstrapped:", isBootstrapped)

	if args[0] == "11" { // Make a new user
		fmtp("arg0 ==", args[0])
		username := "zulu2049"
		password := "avagoflutter"
		isSuccessful, err := avaClient.CreateUser(username, password)
		fmtpf("CreateUser()", err)
		fmtp("isSuccessful:", isSuccessful)
	}

	if args[0] == "12" { // Make a new address
		fmtp("arg0 ==", args[0])
		username := "zulu2049"
		password := "avagoflutter"
		chain := "X"
		address, err := avaClient.CreateAddress(chain, username, password)
		fmtpf("CreateAddress()", err)
		fmtp("address:", address)
		//fuji10spa9d8nt0s62rctd0g57qff3lz3094d6snp63
	}

	if args[0] == "13" { // List user addresses
		fmtp("arg0 ==", args[0])
		//username := "zulu2038"
		username := "zulu2048"
		password := "avagoflutter"
		chain := "X"
		addresses, err := avaClient.ListAddresses(username, password, chain)
		fmtpf("ListAddresses()", err)
		fmtp("addresses:", addresses)
	}

	if args[0] == "14" { // read AVAX balance ans UTXO IDs
		fmtp("arg0 ==", args[0])
		//username := "zulu2038"
		assetID := "AVAX"
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		//addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)

		balance, utxoIDs, err := avaClient.GetBalance(chain, addr1, assetID)
		fmtpf("GetBalance()", err)
		fmtp("balance:", balance, ", utxoIDs:", utxoIDs)
		/*balance1a: 4414997500 , utxoIDs: [{0 2wFJdDqLS4mcWoKgmeJW5SnTtvvKCkhSCTxgBnTN7M35LMRnG} {0 23yRyrDQRkxZ9p3c76GEid3DjdLzfvX6gWTjNzwqWAT538xJX1} {1 QyFxraVqzu99KqRGsR4eD9rF7TNwNJC3VcZaYMv2Zue9kjYFY} {1 5mERnUUNNKrFjxnPSMiWbThTuzC3zEs6dUsqFdyGJNsvg2mxB} {1 PYpNvLcxFE3LKNjzcNQBXKuJbqEsQi7xyiGTRSbAG32CVJsJX} {1 2mxfSyt3V3nYX1xkMJ23Q4x5KHyco3g9tGk4Y2KkqTTg2hnvKn}]


		 */
	} // @@@ What are these TxIDs and OutputIndexes???
	//3452997500 -> 4450997500

	if args[0] == "15" { // read AVAX Balance
		fmtp("arg0 ==", args[0])
		assetID := "AVAX"
		chain := "X"
		txID := "2uWowspvnxG1zxFjXockUxrag41SX1V3hJ3EnTY7GsMUwQWhxe"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
		fmtp("txStatus:", txStatus)
	} //3452997500 -> 4450997500 ???

	if args[0] == "16" { // send AVAX in X chain
		fmtp("arg0 ==", args[0])
		assetID := "AVAX"
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)

		amount := 1000000
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		memo := "hi there!"
		username := "zulu2038"
		password := "avagoflutter"

		txID, err := avaClient.SendAsset(chain, addr1, addr2, from, memo, assetID, amount, username, password)
		fmtpf("SendAsset()", err)
		fmtp("txID:", txID)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		balance1b, _, err1b := avaClient.GetBalance(chain, addr1, assetID)
		balance2b, _, err2b := avaClient.GetBalance(chain, addr2, assetID)

		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance1b:", balance1b, ", err1b:", err1b)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
		fmtp("balance2b:", balance2b, ", err2b:", err2b)
		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	} // @@@ balances here does not match with Explorer
	// @@@ What are these TxIDs and OutputIndexes??? assetIDs do not match!?
	// toAddr balance does not increase!!?? tx not sent?!!
	// id has to be changed??
	// from address???
	// what is changeAddr? what Changes? from what to what?
	/*
		balance1a: 4450997500
		balance1b: 4449997500
		balance2a: 1000000
		balance2b: 1000000
	*/

	if args[0] == "16w" { // walletsend in X chain
		fmtp("arg0 ==", args[0])
		assetID := "AVAX"
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)

		amount := 1000000
		changeAddr := addr1
		addrTo := addr2
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		memo := "hi there!"
		username := "zulu2038"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.WalletSend(chain, changeAddr, addrTo, from, memo, assetID, amount, username, password)
		fmtpf("SendAsset()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		balance1b, _, err1b := avaClient.GetBalance(chain, addr1, assetID)
		balance2b, _, err2b := avaClient.GetBalance(chain, addr2, assetID)

		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance1b:", balance1b, ", err1b:", err1b)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
		fmtp("balance2b:", balance2b, ", err2b:", err2b)
		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	}

	if args[0] == "17f" { //make fixed cap assets
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		fmtp("addr1:", addr1)

		name := "Stock OptionB"
		symbol := "OPTB"
		denomination := 0
		initialHolders := []avalanche.AddrBal{{
			Address: addr1,
			Amount:  1000000000,
		}}
		from := avalanche.FixAddrSlice(chain, []string{addr1})

		changeAddr := addr1
		username := "zulu2038"
		password := "avagoflutter"

		assetID, changeAddrM, err := avaClient.CreateFixedCapAsset(chain, name, symbol, denomination, initialHolders, from, changeAddr, username, password)
		fmtpf("CreateFixedCapAsset()", err)
		time.Sleep(9 * time.Second)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		fmtp("assetID:", assetID, ", changeAddrM:", changeAddrM)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
	}

	if args[0] == "17v" { //make variable cap assets
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		addr3 := "fuji16td6fmpkm3fquyqmsdtmx7rtt5vzjtfxd0t5u9"
		//addr4 := "fuji1zfqh37zg7zxzytm78e2x6j7zr7evmrjjt8sxzp"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
		fmtp("addr3:", addr3)

		name := "Stock OptionV"
		symbol := "OPTV"
		denomination := 0
		minters1 := avalanche.FixAddrSlice(chain, []string{addr1})
		minters123 := avalanche.FixAddrSlice(chain, []string{addr1, addr2, addr3})

		minterSets := []avalanche.MinterSet{
			{
				Minters:   minters1,
				Threshold: 1,
			},
			{
				Minters:   minters123,
				Threshold: 2,
			},
		}
		from := []string{addr1}
		changeAddr := addr1
		username := "zulu2038"
		password := "avagoflutter"

		assetID, changeAddrM, err := avaClient.CreateVariableCapAsset(chain, name, symbol, denomination, minterSets, from, changeAddr, username, password)
		fmtpf("CreateFixedCapAsset()", err)
		time.Sleep(9 * time.Second)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		fmtp("assetID:", assetID, ", changeAddrM:", changeAddrM)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
	}

	if args[0] == "17vm" { //Mint variable cap assets
		fmtp("arg0 ==", args[0])
		chain := "X"
		assetID := "RvawJiqLeNfk44VPBRvT6KAxi18AD4mYXn1wTHhPWosC7WQ3K"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		fmtp("addr1:", addr1)

		amount := 1000
		addrTo := chain + "-" + addr1
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		changeAddr := chain + "-" + addr1
		username := "zulu2038"
		password := "avagoflutter"

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)

		txID, changeAddrM, err := avaClient.MintAsset(chain, amount, assetID, addrTo, from, changeAddr, username, password)
		fmtpf("MintAsset()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		balance1b, _, err1b := avaClient.GetBalance(chain, addr1, assetID)

		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance1b:", balance1b, ", err1b:", err1b)
		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	}

	if args[0] == "18" { //send fix cap assets
		fmtp("arg0 ==", args[0])
		chain := "X"
		assetID := "RvawJiqLeNfk44VPBRvT6KAxi18AD4mYXn1wTHhPWosC7WQ3K"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)

		amount := 1000
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		memo := "hi there!"
		username := "zulu2038"
		password := "avagoflutter"

		txID, err := avaClient.SendAsset(chain, addr2, addr2, from, memo, assetID, amount, username, password)
		fmtpf("SendAsset()", err)
		fmtp("txID:", txID)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		balance1b, _, err1b := avaClient.GetBalance(chain, addr1, assetID)
		balance2b, _, err2b := avaClient.GetBalance(chain, addr2, assetID)

		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance1b:", balance1b, ", err1b:", err1b)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
		fmtp("balance2b:", balance2b, ", err2b:", err2b)
		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	} /*
		# OPTA: fGXtb4N9tu624iNxgHECqkBUY7WD4rAzd4s8grWGpF9zwCdFj
		# OPTB:
		RvawJiqLeNfk44VPBRvT6KAxi18AD4mYXn1wTHhPWosC7WQ3K

	*/

	if args[0] == "18m" { //SendMultiple
		fmtp("arg0 ==", args[0])
		chain := "X"
		assetID := "RvawJiqLeNfk44VPBRvT6KAxi18AD4mYXn1wTHhPWosC7WQ3K"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		addr3 := "fuji16td6fmpkm3fquyqmsdtmx7rtt5vzjtfxd0t5u9"
		//addr4 := "fuji1zfqh37zg7zxzytm78e2x6j7zr7evmrjjt8sxzp"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
		fmtp("addr3:", addr3)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)

		changeAddr := addr2
		outputs := []avalanche.Output{
			{AssetID: "AVAX", To: chain + "-" + addr2, Amount: 1000},
			{AssetID: assetID, To: chain + "-" + addr3, Amount: 10},
		}
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		memo := "hi there!"
		username := "zulu2038"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.SendMultiple(chain, changeAddr, outputs, from, memo, username, password)
		fmtpf("SendAsset()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		balance1b, _, err1b := avaClient.GetBalance(chain, addr1, assetID)
		balance2b, _, err2b := avaClient.GetBalance(chain, addr2, assetID)

		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance1b:", balance1b, ", err1b:", err1b)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
		fmtp("balance2b:", balance2b, ", err2b:", err2b)
		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	}

	if args[0] == "18wm" { //Send Wallet Multiple
		fmtp("arg0 ==", args[0])
		chain := "X"
		assetID := "RvawJiqLeNfk44VPBRvT6KAxi18AD4mYXn1wTHhPWosC7WQ3K"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		addr3 := "fuji16td6fmpkm3fquyqmsdtmx7rtt5vzjtfxd0t5u9"
		//addr4 := "fuji1zfqh37zg7zxzytm78e2x6j7zr7evmrjjt8sxzp"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
		fmtp("addr3:", addr3)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)

		changeAddr := addr2
		outputs := []avalanche.Output{
			{AssetID: "AVAX", To: chain + "-" + addr2, Amount: 1000},
			{AssetID: assetID, To: chain + "-" + addr3, Amount: 10},
		}
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		memo := "hi there!"
		username := "zulu2038"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.WalletSendMultiple(chain, changeAddr, outputs, from, memo, username, password)
		fmtpf("SendAsset()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		balance1b, _, err1b := avaClient.GetBalance(chain, addr1, assetID)
		balance2b, _, err2b := avaClient.GetBalance(chain, addr2, assetID)

		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance1b:", balance1b, ", err1b:", err1b)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
		fmtp("balance2b:", balance2b, ", err2b:", err2b)
		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	}

	if args[0] == "19" { // read asset balance
		fmtp("arg0 ==", args[0])
		chain := "X"
		assetID := "RvawJiqLeNfk44VPBRvT6KAxi18AD4mYXn1wTHhPWosC7WQ3K"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		balance1a, _, err1a := avaClient.GetBalance(chain, addr1, assetID)
		balance2a, _, err2a := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("balance1a:", balance1a, ", err1a:", err1a)
		fmtp("balance2a:", balance2a, ", err2a:", err2a)
	}

	if args[0] == "20" { //create NFT Assets
		fmtp("arg0 ==", args[0])
		symbols := []string{"ABCD", "SPHX", "MUMY", "PYRD"}
		symbolIndex := 0
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		addr3 := "fuji16td6fmpkm3fquyqmsdtmx7rtt5vzjtfxd0t5u9"
		//addr4 := "fuji1zfqh37zg7zxzytm78e2x6j7zr7evmrjjt8sxzp"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
		fmtp("addr3:", addr3)

		name := "Egyptian Artifacts"
		symbol := symbols[symbolIndex]

		minters123 := avalanche.FixAddrSlice(chain, []string{addr1, addr2, addr3})

		minterSets := []avalanche.MinterSet{
			{
				Minters:   minters123,
				Threshold: 2,
			},
		}

		username := "zulu2038"
		password := "avagoflutter"

		assetID, changeAddrM, err := avaClient.CreateNFTAsset(chain, name, symbol, minterSets, username, password)
		//time.Sleep(9 * time.Second)
		fmtpf("CreateNFTAsset()", err)
		fmtp("assetID:", assetID, ", changeAddrM:", changeAddrM)
	} /*
		# TWFC fixed: oCoedYUYGeLk5bbAkJ5fqyKWBgoz7j1aeRUTK9LJTLTnZeDu5
		# TESL varied: qKD4sHi9AfwjrGZT14Vko2QZsLWMEsKsvVR27Ci6oQCGAPWFN
		# ABCD NFT varied: 2wFJdDqLS4mcWoKgmeJW5SnTtvvKCkhSCTxgBnTN7M35LMRnG
		# ABCD NFT varied: 2SktzhDaGVAmtRTedG7M88fk6rgrES5nGHdUxX3KbyuNmCEvVp
		# ABCD NFT varied:
		2QS7HSXDhnxBMHkXcGwJVtEAHL1snTH6U21AyMU7P3LevvBUyd
	*/

	if args[0] == "23" { //Mint NFT Assets
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		fmtp("addr1:", addr1)
		assetID := "2QS7HSXDhnxBMHkXcGwJVtEAHL1snTH6U21AyMU7P3LevvBUyd"
		payload := "Tesla Motor and Space X"
		addrTo := addr1
		username := "zulu2038"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.MintNFT(chain, assetID, payload, addrTo, username, password)
		//time.Sleep(9 * time.Second)
		fmtpf("MintNFT()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
	}
	/*
		mdXQwPCCyUSYvDeWh5QcdWx8hbywwXteQXSPLRbscbfXPKVpH
	*/

	if args[0] == "21" {
		fmtp("arg0 ==", args[0])
		//avaClient.Base58check()
		avaClient.Base58()
	}

	if args[0] == "22" { // Get UTXOs
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		fmtp("addr1:", addr1)
		addresses := []string{addr1}
		limit := 5
		encoding := "cb58"

		numFetched, utxos, address, utxo, err := avaClient.GetUTXOs(chain, addresses, limit, encoding)
		fmtpf("GetUTXOs()", err)
		fmtp("numFetched:", numFetched)
		fmtp("address:", address)
		fmtp("utxo:", utxo)
		avalanche.Dump1("utxos:", utxos)
	}

	if args[0] == "24" { //Mint NFT Assets
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		assetID := "2QS7HSXDhnxBMHkXcGwJVtEAHL1snTH6U21AyMU7P3LevvBUyd"
		addrTo := addr1
		groupID := 0
		username := "zulu2038"
		password := "avagoflutter"
		addresses := []string{addrTo}
		limit := 5
		encoding := "cb58"
		from := avalanche.FixAddrSlice(chain, []string{addr1})
		changeAddr := addr1

		txID, changeAddrM, err := avaClient.SendNFT(chain, assetID, addrTo, from, changeAddr, groupID, username, password)
		//time.Sleep(9 * time.Second)
		fmtpf("SendNFT()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(8 * time.Second)

		numFetched, utxos, address, utxo, err := avaClient.GetUTXOs(chain, addresses, limit, encoding)
		fmtpf("GetUTXOs()", err)
		fmtp("numFetched:", numFetched)
		fmtp("address:", address)
		fmtp("utxo:", utxo)
		avalanche.Dump1("utxos:", utxos)
	}

	if args[0] == "25" { // Generate control keys + Subnet
		fmtp("arg0 ==", args[0])
		chain := "P"
		username := "zulu2049"
		password := "avagoflutter"
		ctrlKey1, err := avaClient.CreateAddress(chain, username, password)
		fmtpf("CreateAddress()", err)
		fmtp("ctrlKey1:", ctrlKey1)

		ctrlKey2, err := avaClient.CreateAddress(chain, username, password)
		fmtpf("CreateAddress()", err)
		fmtp("ctrlKey2:", ctrlKey2)

		controlKeys := []string{ctrlKey1, ctrlKey2}
		threshold := 2
		txID, changeAddr, err := avaClient.CreateSubnet(controlKeys, threshold, username, password)
		fmtpf("CreateSubnet()", err)
		fmtp("txID:", txID, ", changeAddr:", changeAddr)
		time.Sleep(9 * time.Second)

		id, ctrlKeys, thresholdM, err := avaClient.GetSubnets(username, password)
		fmtpf("GetSubnets()", err)
		fmtp("id:", id, ", ctrlKeys:", ctrlKeys, ", thresholdM:", thresholdM)
	}

	if args[0] == "26" {
		// Add Validator to mainnet and a subnet
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		nodeID := ""
		startTime := "" //$(date --date="10 minutes" +%s)
		endTime := ""   //$(date --date="30 days" +%s)
		stakeAmount := 2000000000000
		rewardAddress := addr1
		changeAddr := addr1
		delegationFeeRate := 10
		username := "zulu2049"
		password := "avagoflutter"

		txID, changeAddr, err := avaClient.AddValidator(chain, nodeID, startTime, endTime, stakeAmount, rewardAddress, changeAddr, delegationFeeRate, username, password)
		fmtpf("AddValidator()", err)
		fmtp("txID:", txID, ", changeAddr:", changeAddr)
		time.Sleep(9 * time.Second)

		chain = "P"
		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)
		//txStatus should be Committed
		nodeIDM, startTimeM, endTimeM, stakeAmountM, _, err := avaClient.GetPendingValidators("")
		fmtpf("GetPendingValidators()", err)
		fmtp("nodeIDM:", nodeIDM, ", startTimeM:", startTimeM, ", endTimeM:", endTimeM, ", stakeAmountM:", stakeAmountM)
		time.Sleep(3 * time.Second)

		//-----------== Subnet
		subnetID := ""
		weight := 1
		txID2, changeAddr, err := avaClient.AddSubnetValidator(nodeID, subnetID, startTime, endTime, weight, changeAddr, username, password)
		fmtpf("AddSubnetValidator()", err)
		fmtp("txID2:", txID2, ", changeAddr:", changeAddr)
		time.Sleep(9 * time.Second)

		txStatus2, err := avaClient.GetTxStatus(txID2, chain)
		fmtp("txStatus2:", txStatus2)
		//txStatus should be Committed
		fmtpf("GetTxStatus()", err)

		nodeIDM2, startTimeM2, endTimeM2, _, weight3M, err := avaClient.GetPendingValidators(subnetID)
		fmtpf("GetPendingValidators()", err)
		fmtp("nodeIDM2:", nodeIDM2, ", startTimeM2:", startTimeM2, ", endTimeM2:", endTimeM2, ", weight3M:", weight3M)
	}

	if args[0] == "27" { // BuildGenesis
		fmtp("arg0 ==", args[0])
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		addr3 := "fuji16td6fmpkm3fquyqmsdtmx7rtt5vzjtfxd0t5u9"
		//addr4 := "fuji1zfqh37zg7zxzytm78e2x6j7zr7evmrjjt8sxzp"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
		fmtp("addr3:", addr3)

		assetFixedCapName := "FixedAssetName1"
		assetFixedCapSymbol := "FIXA"
		assetVariableCapName := "VariableAssetName1"
		assetVariableCapSymbol := "VARA"

		bytesStr, err := avaClient.BuildGenesis(assetFixedCapName, assetFixedCapSymbol, assetVariableCapName, assetVariableCapSymbol, addr1, addr2, addr3)
		fmtpf("BuildGenesis()", err)
		fmtp("bytesStr:", bytesStr)
		//time.Sleep(9 * time.Second)
	}

	if args[0] == "28" {
		// CreateBlockchain -> GetBlockchains
		fmtp("arg0 ==", args[0])
		subnetID := ""
		vmID := "avm"
		name := "AVM1"
		genesisData := ""
		username := "zulu2049"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.CreateBlockchain(subnetID, vmID, name, genesisData, username, password)
		fmtpf("CreateBlockchain()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(9 * time.Second)

		blockchainsStr, err := avaClient.GetBlockchains()
		fmtpf("GetBlockchains()", err)
		avalanche.Dump1("blockchainsStr:", blockchainsStr)
	}

	if args[0] == "29" { // GetBlockchains + GetBalance
		fmtp("arg0 ==", args[0])
		chain := ""
		assetID1 := "asset1"
		assetID2 := "asset2"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		addr3 := "fuji16td6fmpkm3fquyqmsdtmx7rtt5vzjtfxd0t5u9"
		//addr4 := "fuji1zfqh37zg7zxzytm78e2x6j7zr7evmrjjt8sxzp"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
		fmtp("addr3:", addr3)

		blockchainsStr, err := avaClient.GetBlockchains()
		fmtpf("GetBlockchains()", err)
		avalanche.Dump1("blockchainsStr:", blockchainsStr)

		balance1, _, err1 := avaClient.GetBalance(chain, addr1, assetID1)
		balance2, _, err2 := avaClient.GetBalance(chain, addr2, assetID1)
		balance3, _, err3 := avaClient.GetBalance(chain, addr3, assetID1)
		fmtp("---------== Asset1 Fixed Cap")
		fmtp("balance1:", balance1, ", err1:", err1)
		fmtp("balance2:", balance2, ", err2:", err2)
		fmtp("balance3:", balance3, ", err3:", err3)

		balance1, _, err1 = avaClient.GetBalance(chain, addr1, assetID2)
		balance2, _, err2 = avaClient.GetBalance(chain, addr2, assetID2)
		balance3, _, err3 = avaClient.GetBalance(chain, addr3, assetID2)
		fmtp("---------== Asset2 Variable Cap")
		fmtp("balance1:", balance1, ", err1:", err1)
		fmtp("balance2:", balance2, ", err2:", err2)
		fmtp("balance3:", balance3, ", err3:", err3)

	}

	if args[0] == "30" { // ExportAVAX2 AVAX X to P
		fmtp("arg0 ==", args[0])
		assetID := "AVAX"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		chain := "X"
		balance1XB4, _, err1XB4 := avaClient.GetBalance(chain, addr1, assetID)
		balance2XB4, _, err2XB4 := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("------------== balances in X chain")
		fmtp("balance1XB4:", balance1XB4, ", err1XB4:", err1XB4)
		fmtp("balance2XB4:", balance2XB4, ", err2XB4:", err2XB4)

		chain = "P"
		balance1PB4, _, err1PB4 := avaClient.GetBalance(chain, addr1, assetID)
		balance2PB4, _, err2PB4 := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("------------== balances in X chain")
		fmtp("balance1PB4:", balance1PB4, ", err1PB4:", err1PB4)
		fmtp("balance2PB4:", balance2PB4, ", err2PB4:", err2PB4)

		chainFrom := "X"
		chainTo := "P"
		addrFrom := addr1
		addrTo := addr2
		amount := 1000000
		username := "zulu2038"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.ExportAVAX2(chainFrom, chainTo, addrFrom, addrTo, assetID, amount, username, password)
		fmtpf("ExportAVAX2()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		time.Sleep(9 * time.Second)

		txStatus, err := avaClient.GetTxStatus(chain, txID)
		fmtpf("GetTxStatus()", err)
		fmtp("txStatus:", txStatus)

		chain = "X"
		balance1XAF, _, err1XAF := avaClient.GetBalance(chain, addr1, assetID)
		balance2XAF, _, err2XAF := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("------------== balances in X chain")
		fmtp("balance1XAF:", balance1XAF, ", err1XAF:", err1XAF)
		fmtp("balance2XAF:", balance2XAF, ", err2XAF:", err2XAF)

		chain = "P"
		balance1PAF, _, err1PAF := avaClient.GetBalance(chain, addr1, assetID)
		balance2PAF, _, err2PAF := avaClient.GetBalance(chain, addr2, assetID)
		fmtp("------------== balances in X chain")
		fmtp("balance1PAF:", balance1PAF, ", err1PAF:", err1PAF)
		fmtp("balance2PAF:", balance2PAF, ", err2PAF:", err2PAF)

		fmtp("txID:", txID)
		fmtp("txStatus:", txStatus)
	} // @@@ balances here does not match with Explorer

	if args[0] == "31" {
		fmtp("arg0 ==", args[0])
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		chainFrom := ""
		chainTo := ""
		addrTo := "AVM1"
		assetID := ""
		privateKey := ""
		username := "zulu2049"
		password := "avagoflutter"

		txID, changeAddrM, address, err := avaClient.ImportAVAX(chainFrom, chainTo, addrTo, assetID, privateKey, username, password)
		fmtpf("ImportAVAX()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM, ", address:", address)
		//time.Sleep(9 * time.Second)
	}

	if args[0] == "32" {
		fmtp("arg0 ==", args[0])
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)

		chainFrom := ""
		chainTo := ""
		addrFrom := "AVM1"
		addrTo := "AVM1"
		assetID := ""
		amount := 1000
		username := "zulu2049"
		password := "avagoflutter"

		txID, changeAddrM, err := avaClient.ExportAVAX2(chainFrom, chainTo, addrFrom, addrTo, assetID, amount, username, password)
		fmtpf("ExportAVAX2()", err)
		fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
		//time.Sleep(9 * time.Second)
	}

	if args[0] == "33" {
		fmtp("arg0 ==", args[0])
		peers, err := avaClient.GetPeers()
		fmtpf("GetPeers()", err)
		avalanche.Dump1("peers:", peers)
	}

	if args[0] == "34" {
		fmtp("arg0 ==", args[0])
		username := "zulu2049"
		password := "avagoflutter"

		err := avaClient.ImportAVAXlocalNework(username, password)
		fmtpf("ImportAVAXlocalNework()", err)
		//fmtp("txID:", txID, ", changeAddrM:", changeAddrM)
	}

	if args[0] == "35" {
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		address := addr1
		username := "zulu2049"
		password := "avagoflutter"

		privateKey, err := avaClient.ExportKey(chain, address, username, password)
		fmtpf("ExportKey()", err)
		fmtp("privateKey:", privateKey)
	}
	/*
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		addr2 := "fuji1wh9l2njkq0wf7xl05xwz0jhv2m9pt7k7g0p6zg"
		fmtp("addr1:", addr1)
		fmtp("addr2:", addr2)
	*/
	if args[0] == "36" {
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		address := addr1

		assetBalances, err := avaClient.GetAllBalances(chain, address)
		fmtpf("GetAllBalances()", err)
		fmtp("assetBalances:", assetBalances)
	}

	if args[0] == "37" {
		fmtp("arg0 ==", args[0])
		chain := "X"
		addr1 := "fuji157fqpkufy7cz93e6rnua7t34rmhttv07wdrvdu"
		assetID := addr1

		name, symbol, denomination, err := avaClient.GetAssetDescription(chain, assetID)
		fmtpf("GetAssetDescription()", err)
		fmtp("name:", name, ", symbol:", symbol, "denomination:", denomination)
	}

	if args[0] == "38" {
		fmtp("arg0 ==", args[0])
		chain := "X"
		txID := ""
		encoding := "cb58" // cb58 or hex
		tx, encodingM, err := avaClient.GetTx(chain, txID, encoding)
		fmtpf("GetTx()", err)
		fmtp("tx:", tx, ", encodingM:", encodingM)
	}

	if args[0] == "39" {
		fmtp("arg0 ==", args[0])
		chain := "X"
		privateKey := ""
		username := "zulu2049"
		password := "avagoflutter"

		address, err := avaClient.ImportKey(chain, privateKey, username, password)
		fmtpf("ImportKey()", err)
		fmtp("address:", address)
	}

	if args[0] == "40" {
		fmtp("arg0 ==", args[0])
		chain := "X"
		tx := ""
		encoding := "cb58"

		address, err := avaClient.IssueTx(chain, tx, encoding)
		fmtpf("IssueTx()", err)
		fmtp("address:", address)
	}

	if args[0] == "41" {
		fmtp("arg0 ==", args[0])
	}

}

/* TODO
https://www.avalabs.org/avalanche-x/apply-for-general-grants

https://medium.com/avalabs/deploy-a-smart-contract-on-ava-using-remix-and-metamask-98933a93f436
*/
func fmtp(a ...interface{}) {
	fmt.Println(a...)
}

func fmtpf(descriptn string, err error) {
	if err != nil {
		fmt.Println("err@ "+descriptn+": ", err)
		os.Exit(1)
	}
}
