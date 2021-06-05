package service

import (
	_stringUtils "github.com/utlai/utl/internal/pkg/libs/string"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluIntentService struct {
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
}

func NewNluIntentService() *NluIntentService {
	return &NluIntentService{}
}

func (s *NluIntentService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluIntent, total int64) {
	pos, total = s.NluIntentRepo.Query(keywords, status, pageNo, pageSize)
	return
}
func (s *NluIntentService) ListByTask(taskId uint) (pos []model.NluIntent) {
	pos = s.NluIntentRepo.ListByTaskId(taskId)
	return
}

func (s *NluIntentService) Get(id uint) (po model.NluIntent) {
	po = s.NluIntentRepo.Get(id)

	sents := s.NluSentRepo.ListByIntent(id)
	po.Sents = sents

	return
}

func (s *NluIntentService) Create(taskIdStr, targetIdStr, mode, name string) (po model.NluIntent, err error) {
	po.Name = name
	po.TaskId = _stringUtils.ParseUint(taskIdStr)
	targetId := _stringUtils.ParseUint(targetIdStr)

	ordr := 0
	if mode == "child" {
		ordr = s.NluIntentRepo.GetMaxOrder(po.TaskId)
	} else if mode == "neighbor" {
		target := s.NluIntentRepo.Get(targetId)
		ordr = target.Ordr + 1

		s.NluIntentRepo.AddOrderForNext(0, target.Ordr, po.TaskId)
	} else {
		return
	}

	po.Ordr = ordr
	err = s.NluIntentRepo.Save(&po)

	return
}

func (s *NluIntentService) Update(po *model.NluIntent) (err error) {
	err = s.NluIntentRepo.Update(po)

	return
}

func (s *NluIntentService) SetDefault(id uint) (err error) {
	err = s.NluIntentRepo.SetDefault(id)

	return
}

func (s *NluIntentService) Disable(id uint) (err error) {
	err = s.NluIntentRepo.Disable(id)

	return
}

func (s *NluIntentService) Delete(id uint) (err error) {
	err = s.NluIntentRepo.Delete(id)

	return
}

func (s *NluIntentService) BatchDelete(ids []int) (err error) {
	err = s.NluIntentRepo.BatchDelete(ids)

	return
}
