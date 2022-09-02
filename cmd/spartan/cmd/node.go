package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/p2p"

	"github.com/cosmos/cosmos-sdk/client"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
)

type SignInfo struct {
	NodeID    string `json:"node_id"`
	PubKey    string `json:"pub_key"`
	Signature []byte `json:"signature"`
}

// NodeCmd creates a main CLI command
func NodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Sign and Verify node information",
		RunE:  client.ValidateCmd,
	}

	cmd.AddCommand(SignCmd())
	cmd.AddCommand(VerifyCmd())
	return cmd
}

func SignCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sign-info",
		Short: "sign with nodeID and export signature",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			cfg := serverCtx.Config

			nodeKey, err := p2p.LoadNodeKey(cfg.NodeKeyFile())
			if err != nil {
				return err
			}

			pubKey := nodeKey.PubKey()
			nodeID := string(p2p.PubKeyToID(pubKey))

			sdkPK, err := cryptocodec.FromTmPubKeyInterface(pubKey)
			if err != nil {
				return err
			}

			sigByte, err := nodeKey.PrivKey.Sign([]byte(nodeID))
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			pkByte, err := clientCtx.Codec.MarshalInterfaceJSON(sdkPK)
			if err != nil {
				return err
			}

			nodeInfo := SignInfo{
				NodeID:    nodeID,
				PubKey:    string(pkByte),
				Signature: sigByte,
			}

			bz, err := json.Marshal(nodeInfo)
			if err != nil {
				return err
			}

			fmt.Println(string(bz))
			return nil
		},
	}
}

func VerifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "verify [sign-info]",
		Short: "Verify a node sign-info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var signInfo SignInfo
			if err := json.Unmarshal([]byte(args[0]), &signInfo); err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			var sdkPK cryptotypes.PubKey
			if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(signInfo.PubKey), &sdkPK); err != nil {
				return err
			}

			pubKey, err := cryptocodec.ToTmPubKeyInterface(sdkPK)
			if err != nil {
				return err
			}

			nodeID := string(p2p.PubKeyToID(pubKey))
			if nodeID != signInfo.NodeID {
				return fmt.Errorf("invalid sign-info,except node_id : %s, but got :%s", nodeID, signInfo.NodeID)
			}

			if !pubKey.VerifySignature([]byte(nodeID), signInfo.Signature) {
				return fmt.Errorf("invalid signature")
			}
			fmt.Println("verify success")
			return nil
		},
	}
}
