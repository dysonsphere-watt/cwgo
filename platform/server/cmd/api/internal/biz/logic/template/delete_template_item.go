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

package template

import (
	"context"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/biz/model/template"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
)

const (
	successMsgDeleteTemplateItem = "" // TODO: to be filled...
)

type DeleteTemplateItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTemplateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTemplateItemLogic {
	return &DeleteTemplateItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTemplateItemLogic) DeleteTemplateItem(req *template.DeleteTemplateItemReq) (res *template.DeleteTemplateRes) {
	err := l.svcCtx.DaoManager.Template.DeleteTemplateItem(req.Ids)
	if err != nil {
		return &template.DeleteTemplateRes{
			Code: 400,
			Msg:  err.Error(),
		}
	}

	return &template.DeleteTemplateRes{
		Code: 0,
		Msg:  successMsgDeleteTemplateItem,
	}
}
