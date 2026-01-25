package db

import (
	"fmt"

	"github.com/feilongjump/jigsaw-api/domain/entity"
)

func Seed() error {
	return seedSystemLedgerCategories()
}

func seedSystemLedgerCategories() error {
	type seedChild struct {
		name string
		sort int
	}
	type seedParent struct {
		name     string
		typ      uint8
		sort     int
		children []seedChild
	}

	seeds := []seedParent{
		// 支出 - 餐饮美食
		{
			name: "餐饮美食", typ: entity.LedgerCategoryTypeExpense, sort: 100,
			children: []seedChild{
				{name: "早餐", sort: 100}, {name: "午餐", sort: 90}, {name: "晚餐", sort: 80},
				{name: "宵夜", sort: 70}, {name: "买菜", sort: 60},
			},
		},
		// 支出 - 居住物业
		{
			name: "居住物业", typ: entity.LedgerCategoryTypeExpense, sort: 90,
			children: []seedChild{
				{name: "房租", sort: 100}, {name: "水费", sort: 90}, {name: "电费", sort: 80},
				{name: "管理费", sort: 70}, {name: "维修保养", sort: 60}, {name: "宽带费", sort: 50},
			},
		},
		// 支出 - 交通出行
		{
			name: "交通出行", typ: entity.LedgerCategoryTypeExpense, sort: 80,
			children: []seedChild{
				{name: "地铁公交", sort: 100}, {name: "打车", sort: 90}, {name: "油费", sort: 80},
				{name: "停车费", sort: 70}, {name: "车辆保养", sort: 60}, {name: "车险", sort: 50},
				{name: "违章罚款", sort: 40},
			},
		},
		// 支出 - 购物消费
		{
			name: "购物消费", typ: entity.LedgerCategoryTypeExpense, sort: 70,
			children: []seedChild{
				{name: "日常用品", sort: 100}, {name: "服饰鞋帽", sort: 90}, {name: "数码电器", sort: 80},
				{name: "美容护肤", sort: 70},
			},
		},
		// 支出 - 休闲娱乐
		{
			name: "休闲娱乐", typ: entity.LedgerCategoryTypeExpense, sort: 60,
			children: []seedChild{
				{name: "电影演出", sort: 100}, {name: "游戏会员", sort: 90}, {name: "旅游度假", sort: 80},
				{name: "运动健身", sort: 70},
			},
		},
		// 支出 - 医疗健康
		{
			name: "医疗健康", typ: entity.LedgerCategoryTypeExpense, sort: 50,
			children: []seedChild{
				{name: "药品", sort: 100}, {name: "诊疗费", sort: 90}, {name: "保险费", sort: 80},
			},
		},
		// 支出 - 人情社交
		{
			name: "人情社交", typ: entity.LedgerCategoryTypeExpense, sort: 40,
			children: []seedChild{
				{name: "请客吃饭", sort: 100}, {name: "送礼红包", sort: 90}, {name: "孝敬长辈", sort: 80},
			},
		},
		// 支出 - 教育进修
		{
			name: "教育进修", typ: entity.LedgerCategoryTypeExpense, sort: 30,
			children: []seedChild{
				{name: "书籍资料", sort: 100}, {name: "课程培训", sort: 90},
			},
		},
		// 支出 - 生活服务
		{
			name: "生活服务", typ: entity.LedgerCategoryTypeExpense, sort: 20,
			children: []seedChild{
				{name: "手机话费", sort: 100}, {name: "寄快递", sort: 90}, {name: "家政保洁", sort: 80},
			},
		},

		// 收入 - 职业收入
		{
			name: "职业收入", typ: entity.LedgerCategoryTypeIncome, sort: 100,
			children: []seedChild{
				{name: "工资收入", sort: 100}, {name: "奖金收入", sort: 90}, {name: "加班费", sort: 80},
				{name: "兼职收入", sort: 70},
			},
		},
		// 收入 - 理财收入
		{
			name: "理财收入", typ: entity.LedgerCategoryTypeIncome, sort: 90,
			children: []seedChild{
				{name: "投资收入", sort: 100}, {name: "利息收入", sort: 90}, {name: "租金收入", sort: 80},
			},
		},
		// 收入 - 其他收入
		{
			name: "其他收入", typ: entity.LedgerCategoryTypeIncome, sort: 80,
			children: []seedChild{
				{name: "人情红包", sort: 100}, {name: "退款赔付", sort: 90}, {name: "意外所得", sort: 80},
			},
		},
	}

	for _, p := range seeds {
		// 1. 处理父级
		var parent entity.LedgerCategory
		var count int64
		// 查找父级是否存在
		err := globalDB.Model(&entity.LedgerCategory{}).
			Where("user_id = ? AND parent_id = ? AND type = ? AND name = ?", 0, 0, p.typ, p.name).
			Count(&count).Error
		if err != nil {
			return err
		}

		if count == 0 {
			// 创建父级
			parent = entity.LedgerCategory{
				UserID:   0,
				ParentID: 0,
				Type:     p.typ,
				Name:     p.name,
				Path:     "0-",
				Sort:     p.sort,
			}
			if err := globalDB.Create(&parent).Error; err != nil {
				return err
			}
		} else {
			// 如果已存在，查出来以便获取ID给子级用
			if err := globalDB.Where("user_id = ? AND parent_id = ? AND type = ? AND name = ?", 0, 0, p.typ, p.name).First(&parent).Error; err != nil {
				return err
			}
		}

		// 2. 处理子级
		for _, c := range p.children {
			var childCount int64
			err := globalDB.Model(&entity.LedgerCategory{}).
				Where("user_id = ? AND parent_id = ? AND name = ?", 0, parent.ID, c.name).
				Count(&childCount).Error
			if err != nil {
				return err
			}

			if childCount == 0 {
				child := entity.LedgerCategory{
					UserID:   0,
					ParentID: parent.ID,
					Type:     parent.Type,
					Name:     c.name,
					Path:     fmt.Sprintf("0-%d-", parent.ID),
					Sort:     c.sort,
				}
				if err := globalDB.Create(&child).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
