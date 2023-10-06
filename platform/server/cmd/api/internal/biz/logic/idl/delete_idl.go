/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package idl

import (
	"context"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/biz/model/idl"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
)

const (
	successMsgDeleteIDL = "" // TODO: to be filled...
)

type DeleteIDLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIDLLogic {
	return &DeleteIDLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteIDLLogic) DeleteIDL(req *idl.DeleteIDLsReq) (res *idl.DeleteIDLsRes) {
	err := l.svcCtx.DaoManager.Idl.DeleteIDLs(req.Ids)
	if err != nil {
		return &idl.DeleteIDLsRes{
			Code: 400,
			Msg:  err.Error(),
		}
	}

	return &idl.DeleteIDLsRes{
		Code: 0,
		Msg:  successMsgDeleteIDL,
	}
}
