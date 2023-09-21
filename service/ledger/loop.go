package ledger

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"server/code"
	"server/global"
	"server/models/ledger"
	"time"
)

func (*LedgersService) LoopAccountService(loopAccount ledger.LoopAccount) (cd int, err error) {
	if err := global.DB.Create(&loopAccount).Error; err != nil {
		return code.ErrorLoopAccount, err
	}
	var service LedgersService
	service.CornService()
	return code.SUCCESS, err
}

func (*LedgersService) CornService() {
	db := global.DB
	c := cron.New()
	var loopAccounts []ledger.LoopAccount
	if err := db.Find(&loopAccounts).Error; err != nil {
		global.Log.Error("查询定时任务失败")
	}
	var ledgerService LedgersService
	for _, loopAccount := range loopAccounts {
		_, er := c.AddFunc(loopAccount.Corn, func() {
			var bill ledger.Bill
			bill.CreatorID = loopAccount.CreatorID
			bill.Amount = loopAccount.Amount
			bill.Type = loopAccount.Type
			bill.Remark = loopAccount.Remark
			bill.CategoryID = loopAccount.CategoryID
			bill.LedgerID = loopAccount.LedgerID
			bill.NotBudget = loopAccount.NotBudget
			bill.CreateTime = time.Now()
			_, _, errs := ledgerService.AddBillService(bill)
			fmt.Println("入库")
			if errs != nil {
				//记录下定时任务的id转换为string
				global.Log.Error("入库失败,定时任务ID:  " + loopAccount.ID)
			}
		})
		if er != nil {
			global.Log.Error("添加定时任务失败,定时任务ID:   " + loopAccount.ID)
		}
	}
	c.Start()

}
