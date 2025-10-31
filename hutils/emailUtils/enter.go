package email

import (
	"crypto/tls"
	"fmt"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
	"gopkg.in/gomail.v2"
)

// Send 发送邮件
func Send(to []string, subject string, message string) error {
	// 1. 连接SMTP服务器
	host := hglobal.Config.Email.Host
	port := hglobal.Config.Email.Port
	userName := hglobal.Config.Email.UserName
	password := hglobal.Config.Email.Password

	// 2. 构建邮件对象
	m := gomail.NewMessage()
	m.SetHeader("From", userName)   // 发件人
	m.SetHeader("To", to...)        // 收件人
	m.SetHeader("Subject", subject) // 主题
	m.SetBody("text/html", message) // 正文

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		zlog.Errorf("邮件发送失败：%v", err)
		return err
	}
	return nil
}

// SendCode 发送验证码
func SendCode(to string, code int64) error {
	message := `
	<p style="text-indent:2em;">你的邮箱验证码为: %06d </p> 
	<p style="text-indent:2em;">此验证码的有效期为5分钟，请尽快使用。</p>
	`
	return Send([]string{to}, "[你好] [邮箱验证码]", fmt.Sprintf(message, code))
}
func SendResume(to string) error {
	message := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <title>王鑫宇 - 简历</title>
  <style>
    body {
      font-family: "PingFang SC", "Helvetica Neue", Arial, sans-serif;
      margin: 0;
      padding: 40px;
      background-color: #fafafa;
      color: #333;
    }
    .container {
      max-width: 800px;
      margin: auto;
      background: #fff;
      padding: 40px;
      border-radius: 8px;
      box-shadow: 0 2px 8px rgba(0,0,0,0.05);
    }
    h1 {
      font-size: 32px;
      margin-bottom: 10px;
    }
    .contact {
      margin-bottom: 30px;
      font-size: 15px;
      color: #555;
    }
    .contact span {
      margin-right: 20px;
    }
    h2 {
      font-size: 20px;
      margin-top: 40px;
      margin-bottom: 15px;
      border-left: 4px solid #0078ff;
      padding-left: 10px;
    }
    ul {
      margin: 0;
      padding-left: 20px;
    }
    li {
      margin-bottom: 10px;
      line-height: 1.6;
    }
    .section {
      margin-bottom: 30px;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>王鑫宇</h1>
    <div class="contact">
      <span>微信：wxy13148880120</span>
      <span>简历下载：<a href="https://jhinih-tiktok.oss-cn-hongkong.aliyuncs.com/fa11dfd3b64fa951f221d51c7354d07795a78dbf0f21f571e57ecd2dc29e3f13.pdf" target="_blank" rel="noopener">点击下载 PDF 简历</a>
    </div>

    <div class="section">
      <h2>🏆 获奖情况</h2>
      <ul>
        <li>2025年 ACM/ICPC 国际大学生程序设计竞赛邀请赛（武汉）银奖</li>
        <li>2025年 ACM/GDCPC 广东省大学生程序设计竞赛 铜奖</li>
      </ul>
    </div>

    <div class="section">
      <h2>🛠 项目经历</h2>
      <h3>仿真抖音-AI | 2025.07 - 至今</h3>
	  <span><a href="http//jhinih.com" target="_blank" rel="noopener">项目体验</a>
      <p><strong>项目简介：</strong>基于 Gin 开发，部分业务已迁移至 Go-Zero 框架，使用 gRPC 实现微服务通信。基于 Eino 框架设计 Agent，提供可定制化 AI 管家，并扩展多种 Agent。项目提供登录注册、短视频、好友群聊、AI 等服务。</p>
      <ul>
        <li><strong>用户认证：</strong>设计并实现 JWT 双 Token（AccessToken + RefreshToken）无感刷新机制，将用户登录状态从 24 小时延长至 30 天。</li>
        <li><strong>短视频模块：</strong>引入 RabbitMQ 解耦，异步实现文件上传功能；池化 channel，避免重复建/毁 Channel 的开销；设计 Redis 缓存策略，减少频繁查询数据库带来的磁盘 IO；升级为 SSE 请求，避免轮询上传结果，减轻后端压力。</li>
        <li><strong>IM 系统：</strong>基于 WebSocket 实现实时聊天社交系统；引入优化后的 AI，提供更专业的问答服务。</li>
        <li><strong>AI 交互模块：</strong>独立开发，基于 Go-Zero 与 Eino 框架，使用 gRPC 实现微服务通信；将 API 封装为 tool，实现 AI 对项目的全方位调用；支持用户定制个性化 AI。</li>
        <li><strong>Agent 开发：</strong>基于 Eino 搭建工作流与智能体开发功能；使用 Doubao-1.5-vision-pro 模型，针对 AI 控制下的视频获取、搜索、添加好友等场景进行 Prompt 调优。</li>
        <li><strong>RAG 向量数据库：</strong>建立知识库，实现 AI 的定制化配置，提升 API 调用质量与准确性。</li>
        <li><strong>系统稳定性保障：</strong>设计并实现自定义频率限流中间件（基于令牌桶）；严格校验数据合法性，提升系统安全性与稳定性。</li>
      </ul>
    </div>

    <div class="section">
      <h2>🧠 专业技能</h2>
      <ul>
        <li>熟练使用 Go 语言进行项目开发，了解内存逃逸、GMP 模型、GC、三色标记法、混合写屏障。</li>
        <li>熟练使用 Gin、Go-Zero、Eino 等框架，具备良好的代码开发习惯，能针对业务进行二次开发。</li>
        <li>熟悉 MySQL 增删改查操作，掌握存储引擎、事务隔离级别、锁、索引。</li>
        <li>熟悉 Redis 常用数据结构、持久化、穿透、击穿、雪崩、淘汰策略。</li>
        <li>能使用 RabbitMQ 消息队列进行生产与消费，掌握消息丢失、积压、顺序消息、重复消息的解决方案。</li>
        <li>熟悉分布式系统，了解 CAP 和 BASE 理论、分布式 ID、分布式锁、一致性算法（如 Raft）。</li>
        <li>了解微服务架构，能使用 Go-Zero 进行业务开发。</li>
        <li>具备大模型基础知识，能熟练使用 Coze 开发工作流与 Agent，掌握 Prompt 调优，了解 RAG 与 MCP。</li>
      </ul>
    </div>

    <div class="section">
      <h2>🎓 教育经历</h2>
      <p><strong>东莞理工学院（一本）</strong> | 2024.09 - 2028.06<br>软件工程 本科</p>
      <ul>
        <li>主修课程：算法与数据结构、计算机网络、计算机操作系统、计算机组成原理</li>
        <li>学校 AcKing 实验室成员</li>
        <li>竞赛中心负责人</li>
      </ul>
    </div>
  </div>
</body>
</html>
	`
	return Send([]string{to}, "[你好] [jhinih的简历已送达]", fmt.Sprintf(message))
}
