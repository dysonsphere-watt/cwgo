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
	successMsgSyncIDLs = "" // TODO: to be filled...
)

type SyncIDLsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncIDLsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncIDLsLogic {
	return &SyncIDLsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncIDLsLogic) SyncIDLs(req *idl.SyncIDLsByIdReq) (res *idl.SyncIDLsByIdRes) {
	//TODO: 调用agent

	return &idl.SyncIDLsByIdRes{
		Code: 0,
		Msg:  successMsgSyncIDLs,
	}
}
