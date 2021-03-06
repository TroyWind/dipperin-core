// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package commands

import (
	"errors"
	"github.com/dipperin/dipperin-core/core/model"
	"github.com/dipperin/dipperin-core/core/rpc-interface"
	"testing"

	"github.com/dipperin/dipperin-core/common"
	"github.com/dipperin/dipperin-core/core/economy-model"
	"github.com/golang/mock/gomock"
	"github.com/urfave/cli"
)

func Test_rpcCaller_GetBlockDiffVerifierInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "m", Usage: "operation"},
		cli.StringFlag{Name: "p", Usage: "parameters"},
	}

	app.Action = func(c *cli.Context) {
		client = NewMockRpcClient(ctrl)
		caller := &rpcCaller{}

		InnerRpcForbid = true
		caller.GetBlockDiffVerifierInfo(c)

		InnerRpcForbid = false
		SyncStatus.Store(false)
		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), gomock.Any()).Return(errors.New("test"))
		caller.GetBlockDiffVerifierInfo(c)

		SyncStatus.Store(true)
		caller.GetBlockDiffVerifierInfo(c)

		c.Set("m", "test")
		c.Set("p", "")
		caller.GetBlockDiffVerifierInfo(c)

		c.Set("p", "test")
		caller.GetBlockDiffVerifierInfo(c)

		c.Set("p", "1")
		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test"))
		caller.GetBlockDiffVerifierInfo(c)

		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(result interface{}, method string, args ...interface{}) error {
			*result.(*map[economy_model.VerifierType][]common.Address) = map[economy_model.VerifierType][]common.Address{}
			return nil
		})
		caller.GetBlockDiffVerifierInfo(c)

	}

	app.Run([]string{"xxx"})
	client = nil
}

func Test_printAddress(t *testing.T) {
	printAddress([]common.Address{common.HexToAddress("0x1234")})
}

func Test_rpcCaller_CheckVerifierType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "m", Usage: "operation"},
		cli.StringFlag{Name: "p", Usage: "parameters"},
	}

	app.Action = func(c *cli.Context) {
		client = NewMockRpcClient(ctrl)
		caller := &rpcCaller{}

		InnerRpcForbid = true
		caller.CheckVerifierType(c)

		InnerRpcForbid = false
		SyncStatus.Store(false)
		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), gomock.Any()).Return(errors.New("test"))
		caller.CheckVerifierType(c)

		SyncStatus.Store(true)
		caller.CheckVerifierType(c)

		c.Set("m", "test")
		c.Set("p", "")
		caller.CheckVerifierType(c)

		c.Set("p", "test,test")
		caller.CheckVerifierType(c)

		c.Set("p", "1,test")
		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), gomock.Any()).Return(errors.New("test"))
		caller.CheckVerifierType(c)

		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), getDipperinRpcMethodByName("CurrentBlock")).DoAndReturn(func(result interface{}, method string, args ...interface{}) error {
			*result.(*rpc_interface.BlockResp) = rpc_interface.BlockResp{
				Header: model.Header{
					Number: 2,
				},
			}
			return nil
		}).Times(4)
		caller.CheckVerifierType(c)

		c.Set("p", "0,test")
		caller.CheckVerifierType(c)

		c.Set("p", "0,"+common.HexToAddress("0x1234").Hex())
		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), getDipperinRpcMethodByName("GetBlockDiffVerifierInfo"), gomock.Any()).Return(errors.New("test"))
		caller.CheckVerifierType(c)

		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), getDipperinRpcMethodByName("GetBlockDiffVerifierInfo"), gomock.Any()).DoAndReturn(func(result interface{}, method string, args ...interface{}) error {
			tmp := map[economy_model.VerifierType][]common.Address{}
			tmp[economy_model.MasterVerifier] = []common.Address{
				common.HexToAddress("0x1234"),
			}
			*result.(*map[economy_model.VerifierType][]common.Address) = tmp
			return nil
		}).Times(2)

		client.(*MockRpcClient).EXPECT().Call(gomock.Any(), getDipperinRpcMethodByName("GetBlockDiffVerifierInfo"), gomock.Any()).Return(nil).Times(1)
		caller.CheckVerifierType(c)
	}

	app.Run([]string{"xxx"})
	client = nil
}
