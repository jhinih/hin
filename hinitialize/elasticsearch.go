package hinitialize

import (
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
)

func InitElasticsearch() {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + hglobal.Config.Elasticsearch.Host + ":" + hglobal.Config.Elasticsearch.Port,
		},
		Username: hglobal.Config.Elasticsearch.UserName,
		Password: hglobal.Config.Elasticsearch.Password,
	}

	// 创建客户端连接
	var err error
	hglobal.ESClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		zlog.Errorf("ElasticSearch 初始化失败: %v", err)
		return
	}
	zlog.Infof("ElasticSearch 初始化成功")

	//err = UpdateElasticsearch()
	if err != nil {
		zlog.Errorf("同步ElasticSearch数据失败: %v", err)
		return
	}
}

//func UpdateElasticsearch() (err error) {
//	// 获取最近更新时间
//	var updateTime int64
//	updateTime = 0
//	m, err := elasticSearchUtils.Get(hglobal.ESClient, "update_time", "post")
//	if err != nil {
//		zlog.Errorf("获取更新时间失败: %v", err)
//		return
//	}
//	if m == nil {
//		updateTime = 0
//	} else {
//		updateTime = int64(m["timestamp"].(float64))
//	}
//	zlog.Infof("ElasticSearch 上次同步时间为: %s", time.UnixMilli(updateTime).Format("2006-01-02 15:04:05"))
//	newUpdateTime := time.Now().UnixMilli()
//
//	// 同步数据
//	posts, err := repository.NewPostrequest(hglobal.DB).GetPostAfterUpdateTime(updateTime)
//	if err != nil {
//		zlog.Errorf("获取文章失败: %v", err)
//		return
//	}
//	zlog.Debugf("需要同步的文章数量为: %v", len(posts))
//
//	for _, post := range posts {
//		// 获取作者名称
//		var userProfile model.User
//		userProfile, err = repository.NewUserrequest(hglobal.DB).GetUserProfileByID(post.UserID)
//		if err != nil {
//			zlog.Errorf("获取作者信息失败: %v", err)
//			return
//		}
//		// 类型名词
//		typeName := "未知"
//		if post.Type == "diary" {
//			typeName = "周记"
//		} else if post.Type == "tutorial" {
//			typeName = "教程"
//		} else if post.Type == "solution" {
//			typeName = "题解"
//		} else if post.Type == "contest" {
//			typeName = "比赛"
//		} else if post.Type == "fun" {
//			typeName = "闲聊"
//		}
//
//		err = elasticSearchUtils.Update(hglobal.ESClient, "post", strconv.FormatInt(post.ID, 10), map[string]interface{}{
//			"author_name":  userProfile.Username,
//			"id":           strconv.FormatInt(post.ID, 10),
//			"title":        post.Title,
//			"content":      post.Content,
//			"type":         post.Type,
//			"type_name":    typeName,
//			"source":       post.Source,
//			"created_time": post.CreatedTime,
//		})
//		if err != nil {
//			zlog.Errorf("同步文章失败: %v", err)
//			return
//		}
//		zlog.Debugf("同步文章成功: %v", post.Title)
//	}
//
//	// 更新更新时间
//	err = elasticSearchUtils.Update(hglobal.ESClient, "update_time", "post", map[string]interface{}{
//		"timestamp": newUpdateTime,
//	})
//	if err != nil {
//		zlog.Errorf("更新更新时间失败: %v", err)
//		return
//	}
//	return
//}
