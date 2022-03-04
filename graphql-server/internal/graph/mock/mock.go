package mock

import (
	"graphql-server/internal/graph/model"
	"time"
)

var Events = []*model.Event{
	{
		ID:       "1",
		Name:     "週末夜喜劇秀",
		Briefly:  "每週五、六推出固定的\n「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！",
		Category: "standup",
		Host: &model.Host{
			ID:       "1",
			Name:     "卡米地喜劇基地Comedy Base",
			Location: "台北市中山區八德路二段325號3樓",
		},
		Shows: []*model.Show{
			{
				ID:       "1",
				Time:     time.Date(2022, 2, 11, 20, 00, 0, 0, time.Local),
				Location: "台北市中山區八德路二段325號3樓",
				Price:    []int{300},
				Actors:   []*model.Actor{},
				Link:     "https://comedyclub.kktix.cc/events/livecomedy",
			},
			{
				ID:       "2",
				Time:     time.Date(2022, 2, 12, 20, 00, 0, 0, time.Local),
				Location: "台北市中山區八德路二段325號3樓",
				Price:    []int{300},
				Actors:   []*model.Actor{},
				Link:     "https://comedyclub.kktix.cc/events/livecomedy",
			},
		}},
	{
		ID:       "2",
		Name:     "炎上艾莉莎莎",
		Briefly:  "不被審查也不被停播的吐槽大會",
		Category: "standup",
		Host: &model.Host{
			ID:       "2",
			Name:     "23喜劇",
			Location: "台北市中山區林森北路286號B1",
		},
		Shows: []*model.Show{
			{
				ID:       "3",
				Time:     time.Date(2022, 2, 13, 20, 00, 0, 0, time.Local),
				Location: "台北流行音樂中心 台北市南港區市民大道八段99號​",
				Price:    []int{1500, 1800, 2200, 2600, 4000},
				Actors:   []*model.Actor{},
				Link:     "https://strnetwork.kktix.cc/events/3burn0414",
			},
		},
	},
	{
		ID:       "3",
		Name:     "造音少年3！音樂脫口秀 Music Talk Show",
		Briefly:  "脫口秀 x 嘻哈 x DJ x 街舞",
		Category: "standup",
		Host: &model.Host{
			ID:       "3",
			Name:     "超級圓頂娛樂製作有限公司",
			Location: "台北市中正區八德路1段23號3樓",
		},
		Shows: []*model.Show{
			{
				ID:       "4",
				Time:     time.Date(2022, 2, 24, 20, 00, 0, 0, time.Local),
				Location: "Pipe Live Music 台北市中正區思源街1號",
				Price:    []int{75, 120, 150},
				Actors:   []*model.Actor{},
				Link:     "https://kktix.com/virtual_venue/music-talk-show",
			},
		},
	},
}
