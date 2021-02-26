package dto

import (
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type WfProcessClassifySearch struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name" search:"type:contains;column:name;table:wf_process_classify" comment:"名称"`

    
}

func (m *WfProcessClassifySearch) GetNeedSearch() interface{} {
	return *m
}

func (m *WfProcessClassifySearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

type WfProcessClassifyControl struct {
    
    Id int `uri:"id" comment:"编码"` // 编码

    Name string `json:"name" comment:"名称"`
    
}

func (s *WfProcessClassifyControl) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBindUri(s)
    if err != nil {
        log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
        return err
    }
    err = ctx.ShouldBind(s)
    if err != nil {
        log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
    }
    return err
}

func (s *WfProcessClassifyControl) Generate() (*models.WfProcessClassify, error) {
	return &models.WfProcessClassify{
	
        Model:     common.Model{ Id: s.Id },
        Name:  s.Name,
	}, nil
}

func (s *WfProcessClassifyControl) GetId() interface{} {
	return s.Id
}

type WfProcessClassifyById struct {
	Id  int   `uri:"id"`
	Ids []int `json:"ids"`
}

func (s *WfProcessClassifyById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *WfProcessClassifyById) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
	}
	return err
}

func (s *WfProcessClassifyById) SetUpdateBy(id int) {

}