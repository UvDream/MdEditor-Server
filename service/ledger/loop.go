package ledger

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"server/code"
	"server/global"
	"server/models/ledger"
	"time"
)

// LoopAccountService 新增
func (*LedgersService) LoopAccountService(loopAccount ledger.LoopAccount) (cd int, err error) {
	if err := global.DB.Create(&loopAccount).Error; err != nil {
		return code.ErrorLoopAccount, err
	}
	var service LedgersService
	service.CornService()
	return code.SUCCESS, err
}

// UpdateLoopAccountService 更新
func (*LedgersService) UpdateLoopAccountService(loopAccount ledger.LoopAccount) (cd int, err error) {
	if err := global.DB.Model(&loopAccount).Updates(&loopAccount).Error; err != nil {
		return code.ErrorUpdateLoopAccount, err
	}
	var service LedgersService
	service.CornService()
	return code.SUCCESS, err
}

// DeleteLoopAccountService 删除
func (*LedgersService) DeleteLoopAccountService(id string) (cd int, err error) {
	//查询是否存在
	var loopAccount ledger.LoopAccount
	if err := global.DB.Where("id = ?", id).First(&loopAccount).Error; err != nil {
		return code.ErrorDeleteLoopAccount, err
	}
	if err := global.DB.Where("id = ?", id).Unscoped().Delete(&ledger.LoopAccount{}).Error; err != nil {
		return code.ErrorDeleteLoopAccount, err
	}
	var service LedgersService
	service.CornService()
	return code.SUCCESS, err
}
func (*LedgersService) ChangeLoopAccountStatusService(id string) (cd int, err error) {
	//先查询是否存在
	db := global.DB
	var loopAccount ledger.LoopAccount
	if err := db.Where("id = ?", id).Find(&loopAccount).Error; err != nil {
		return code.ErrorLoopAccountNotExist, err
	}
	if loopAccount.Status == "0" {
		loopAccount.Status = "1"
	} else {
		loopAccount.Status = "0"
	}
	//更新数据
	if err := db.Model(&loopAccount).Updates(&loopAccount).Error; err != nil {
		return code.ErrorUpdateLoopAccount, err
	}
	var service LedgersService
	service.CornService()
	return code.SUCCESS, err
}

// CornService 定时任务
func (*LedgersService) CornService() {
	db := global.DB
	c := cron.New(cron.WithSeconds())
	c.Stop()
	var loopAccounts []ledger.LoopAccount
	if err := db.Find(&loopAccounts).Error; err != nil {
		global.Log.Error("查询定时任务失败")
	}
	var ledgerService LedgersService
	if len(loopAccounts) > 0 {
		for _, loopAccount := range loopAccounts {
			if loopAccount.Status == "1" {
				id, er := c.AddFunc(loopAccount.Corn, func() {
					fmt.Println("执行定时任务------定时任务ID:" + loopAccount.ID + "时间:" + time.Now().Format("2006-01-02 15:04:05"))
					var bill ledger.Bill
					bill.Amount = loopAccount.Amount
					bill.CategoryID = loopAccount.CategoryID
					bill.CreatorID = loopAccount.CreatorID
					bill.LedgerID = loopAccount.LedgerID
					bill.NotBudget = loopAccount.NotBudget
					bill.Remark = loopAccount.Remark
					bill.Type = loopAccount.Type
					bill.CreateTime = time.Now()
					_, _, err := ledgerService.AddBillService(bill)
					if err != nil {
						global.Log.Error("定时任务执行失败,定时任务ID" + loopAccount.ID)
					}
				})
				if er != nil {
					global.Log.Error("添加定时任务失败,定时任务ID" + loopAccount.ID)
				}
				fmt.Println("定时任务ID", id)
			}

		}
	}
	c.Start()

}
