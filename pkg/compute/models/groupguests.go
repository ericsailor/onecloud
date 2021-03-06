package models

import (
	"context"

	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/mcclient"
)

type SGroupguestManager struct {
	SGroupJointsManager
}

var GroupguestManager *SGroupguestManager

func init() {
	db.InitManager(func() {
		GroupguestManager = &SGroupguestManager{
			SGroupJointsManager: NewGroupJointsManager(
				SGroupguest{},
				"guestgroups_tbl",
				"groupguest",
				"groupguests",
				GuestManager,
			),
		}
	})
}

type SGroupguest struct {
	SGroupJointsBase

	Tag     string `width:"256" charset:"ascii" nullable:"true" list:"user" update:"user" create:"optional"` // Column(VARCHAR(256, charset='ascii'), nullable=True)
	GuestId string `width:"36" charset:"ascii" nullable:"false" list:"user" create:"required"`               // Column(VARCHAR(36, charset='ascii'), nullable=False)
}

func (joint *SGroupguest) Master() db.IStandaloneModel {
	return db.JointMaster(joint)
}

func (joint *SGroupguest) Slave() db.IStandaloneModel {
	return db.JointSlave(joint)
}

func (self *SGroupguest) GetCustomizeColumns(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject) *jsonutils.JSONDict {
	extra := self.SGroupJointsBase.GetCustomizeColumns(ctx, userCred, query)
	return db.JointModelExtra(self, extra)
}

func (self *SGroupguest) GetExtraDetails(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject) (*jsonutils.JSONDict, error) {
	extra, err := self.SGroupJointsBase.GetExtraDetails(ctx, userCred, query)
	if err != nil {
		return nil, err
	}
	return db.JointModelExtra(self, extra), nil
}

func (self *SGroupguest) GetGuest() *SGuest {
	guest, _ := GuestManager.FetchById(self.GuestId)
	if guest != nil {
		return guest.(*SGuest)
	}
	return nil
}

func (self *SGroupguest) Delete(ctx context.Context, userCred mcclient.TokenCredential) error {
	return db.DeleteModel(ctx, userCred, self)
}

func (self *SGroupguest) Detach(ctx context.Context, userCred mcclient.TokenCredential) error {
	return db.DetachJoint(ctx, userCred, self)
}
