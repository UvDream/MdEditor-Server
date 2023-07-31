package code

const (
	ErrLedgerExist                              = 11001 //	账本已存在
	ErrCreateLedger                             = 11002 //	创建账本失败
	ErrGetLedgerList                            = 11003 //	获取账本列表失败
	ErrCreateLedgerCategory                     = 11004 //	创建账本分类失败
	ErrCreateLedgerCategoryRelation             = 11005 //	创建账本分类关联关系失败
	ErrorDeleteLedger                           = 11006 //	删除账本失败
	ErrorLedgerNotExist                         = 11007 //	账本不存在
	ErrorGetLedgerCategoryList                  = 11008 //	查询账本分类失败
	ErrorGetLedgerCategoryRelationList          = 11009 //	查询账本分类关联关系失败
	ErrorDeleteLedgerCategoryRelation           = 11010 //	删除账本分类关联关系失败
	ErrorDeleteLedgerCategory                   = 11011 //	删除账本分类失败
	ErrorUpdateLedger                           = 11012 //	更新账本失败
	ErrorGetLedger                              = 11013 //	查询账本失败
	ErrorCreateCategory                         = 11014 //   创建分类失败
	ErrCategoryExist                            = 11015 //	分类已存在
	ErrorCategoryNotExist                       = 11016 //	分类不存在
	ErrorUpdateCategory                         = 11017 //	更新分类失败
	ErrorDeleteCategory                         = 11018 //	删除分类失败
	ErrCreateBill                               = 11019 //	创建账单失败
	ErrorDeleteBill                             = 11020 //	删除账单失败
	ErrorUpdateBill                             = 11021 //	更新账单失败
	ErrorGetBill                                = 11022 //	查询账单失败
	ErrorMissingLedgerId                        = 11023 //	缺少账本id
	ErrCreateBudget                             = 11024 //	预算添加失败
	ErrDeleteBudget                             = 11025 //	预算删除失败
	ErrGetBudget                                = 11026 //	预算查找失败
	ErrUpdateBudget                             = 11027 //	预算更新失败
	ErrBudgetNotExist                           = 11028 //	预算不存在
	ErrBudgetExist                              = 11029 //	预算已存在
	ErrorGetCategory                            = 11030 //	查询分类失败
	ErrorGetCategoryStatistics                  = 11031 //	查询分类统计失败
	ErrorGetIncomeExpenditureStatistics         = 11032 //	查询收支统计失败
	ErrorNotLedgerCreator                       = 11033 //	不是账本创建者
	ErrorShareLedger                            = 11034 //	账本分享失败
	ErrShareCodeExpired                         = 11035 //共享码过期
	ErrShareCodeNotExist                        = 11036 //共享码不存在
	ErrorJoinLedger                             = 11037 //	加入账本失败
	ErrorUserAlreadyJoined                      = 11039 //	用户已加入账本
	ErrorGetMemberStatistics                    = 11040 //	查询成员统计失败
	ErrorGetIncomeExpenditureStatisticsByMember = 11041 //	查询成员收支统计失败
	ErrorGetPersonalStatistics                  = 11042 //	查询个人统计失败
	ErrorGetBudget                              = 11043 //	查询预算统计失败
	ErrCreateLedgerNotMember                    = 11044 //不是会员
	ErrorCreateLedgerNotMember                  = 11045 //	非会员只能创建一个账本
	ErrorCreateLedgerMember                     = 11046 //	会员可以创建五个账本
	ErrorGetCategoryStatisticsDetail            = 11047 //	查询分类详情统计失败
	ErrorGetTotalAmount                         = 11048 //	查询总金额失败
	ErrorMissingSearchParam                     = 11049 //	缺少搜索参数

)
